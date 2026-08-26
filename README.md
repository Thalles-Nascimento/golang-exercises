# Exercícios de Go

Repositório de estudos para praticar sintaxe e conceitos da linguagem Go. Cada exercício vem de um curso e é resolvido individualmente; este README é atualizado incrementalmente conforme novos exercícios são resolvidos, registrando a proposta, as decisões de implementação e o código de cada um.

## Sobre o repositório

- Cada exercício tem sua própria pasta em `exercises/<nome>/<nome>.go`.
- Para rodar, basta ter o Golang instalado na máquina e roda o comando `go run exercises/<nome>/<nome>.go` => Exemplo: `go run reverse/reverse.go`
- Exercícios **ainda não resolvidos** ficam como o stub original do curso: `package module01` com a função declarada mas com corpo vazio (ex: `fizz_buzz/fizz_buzz.go`).
- Exercícios **já resolvidos** viram `package main`, com a lógica implementada e uma função `main()` demonstrando chamadas de teste via `fmt.Println`.

## Índice de exercícios resolvidos

- [Sum](#sum)
- [Reverse](#reverse)
- [NumInList](#numinlist)
- [MDC](#mdc)

---

## Sum

**Pasta:** [sum/sum.go](sum/sum.go)

### Proposta

Somar todos os números passados como entrada (uma slice de `int`) e retornar o resultado.

### Minhas decisões

- Usei o *zero value* do `int` (`var acc int`) para iniciar o acumulador em `0`, em vez de atribuir manualmente — aproveitando o valor padrão que o Go já garante para o tipo.
- Percorri a slice com `for _, num := range numbers`, somando cada elemento ao acumulador. Para esse problema em que eu sei que serão listas de inteiros e que serão pequenas, utilizei a abordagem de blank identifier/underline para o índice e usei a variável `num` para receber o itens.
- Outra forma de fazer seria capturando apenas o índice, o que melhoria a perfomance e evitaria overhead, pois, assim Go não cria cópias para iterar. No entanto, esse caso se encaixa quando estamos falando de structs ou listas enormes.
- Testei explicitamente os casos de slice `nil` e slice vazia (`[]int{}`) para confirmar que `Sum` retorna `0` em ambos, sem tratamento especial extra — o `range` sobre uma slice `nil` ou vazia simplesmente não itera.


### Código

```go
package main

import "fmt"


func Sum(numbers []int) int {
	var acc int

	for _, num := range numbers {
		acc += num
	}

	return acc
}

func main(){
	fmt.Println("Primeiro Teste: ", Sum([]int{1,2,3,4,5}))
	fmt.Println("Segundo Teste: ", Sum([]int{3, 3}))
	fmt.Println("Terceiro Teste: ", Sum([]int{3,5,3,5,3}))
	fmt.Println("Quarto Teste: ", Sum([]int{4,2,22,-10,8}))

	fmt.Println("Nil: ", Sum(nil))
	fmt.Println("Vazia: ", Sum([]int{}))
}
```

---

## Reverse

**Pasta:** [reverse/reverse.go](reverse//reverse.go)

### Proposta

Receber uma palavra (`string`) e retornar essa palavra na ordem inversa.

### Minhas decisões

- Converti a string em um slice `rune` (`int32`) para rodar o algoritmo de reversão. Essa conversão é extremamente rápida nesse cenário e quis testar com caracteres especiais.
- Poderia quebrar a string em uma slice de substrings (uma por caractere) com `strings.Split(word, "")`, em vez de converter para `[]byte` ou `[]rune` manualmente. No entanto, eu criaria um overhead gerando múltiplas alocações em memória dependendo do tamanho da string (ex: um texto grande). O rune apenas decodifica para Code points os bytes da string, gerando uma única alocação em memória.
- Poderia também utilizar a conversão para `byte`, no entanto gerará um problema quando houver caracteres especiais - Unicode - onde são usado 2 bytes para representar esse caracter, ao inverter causaria a quebra dele.
Quis testar esse cenário de caracter especial, por isso inseri um teste com a palavra Goiás - onde `á` é armazenada com 2 bytes.
- Usei `slices.Reverse` para inverter a slice resultante no lugar, aproveitando a função pronta do pacote `slices` em vez de escrever um loop de troca de índices.
- Usei novamente o blank identifier para o índice no loop para rescrever a string.
- Reconstruí a string com `strings.Builder`, escrevendo cada `rune` com `WriteRune` — abordagem mais eficiente que concatenação simples de strings em loop, já que `strings.Builder` evita realocações repetidas e faz a conversão para slices de bytes. Além disso, utilizei a função `Grow()` para estabelecer/fixar o quanto de buffer seria reservado para a nova string. References: [Builder.Grow()](https://cs.opensource.google/go/go/+/refs/tags/go1.27.0:src/strings/builder.go;l=75)


### Código

```go
package main

import (
	"fmt"
	"slices"
	"strings"
)

// Reverse retornará a palavra fornecida na ordem inversa
// Ex:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	word_rune := []rune(word)
	slices.Reverse(word_rune)

	var word_reversed strings.Builder
	word_reversed.Grow(len(word))

	for _, s := range word_rune {
		word_reversed.WriteRune(s)
	}
	return word_reversed.String()
}

func main(){
	fmt.Println(Reverse("cat"))
	fmt.Println(Reverse("ca t"))
	fmt.Println(Reverse("alphabet"))
	fmt.Println(Reverse("Goiás"))

}
```

---

## NumInList

**Pasta:** [num_in_list/num_in_list.go](num_in_list/num_in_list.go)

### Proposta

Retornar `true` se um número estiver presente em uma slice de `int`, e `false` caso contrário — incluindo os casos de slice `nil` e slice vazia.

### Minhas decisões

- Optei por uma busca linear simples (`for _, numList := range list`) em vez de busca binária, já que as listas de teste não estão ordenadas e o volume de dados não justifica a complexidade extra. Dependendo do tamanho da lista, a busca binária se sairia melhor. O(n) x O(log n);
- Considerei usar `slices.Contains(list, num)`, que internamente já faz esse mesmo loop (via `slices.Index`), mas mantive o loop explícito para fins de estudo da sintaxe.
- Documentei no próprio código o trade-off de usar a variável de iteração (`numList`) dentro do `range`: para tipos primitivos como `int`, o Go copia o valor a cada iteração sem custo relevante de performance; já para `structs` grandes essa cópia poderia ser custosa, sendo um ponto de atenção para otimizar futuramente se o tipo mudar.

### Código

```go
package main

import "fmt"

// NumInList retornará verdadeiro se o número estiver na
// Slice list, e falso caso contrário.
func NumInList(list []int, num int) bool {

	for _, numList := range list {
		if numList == num {
			return true
		} 		
	}

	return false
}

func main(){
	fmt.Println("Primeiro Teste: ", NumInList([]int{1,2,3,4,5}, 5))
	fmt.Println("Segundo Teste: ", NumInList([]int{3,3,3,3,3}, 5))
	fmt.Println("Terceiro Teste: ", NumInList([]int{3,5,3,5,3}, 5))
	fmt.Println("Quarto Teste: ", NumInList([]int{4,2,22,-10,8}, -10))

	fmt.Println("Nil Teste: ", NumInList(nil, 5))
	fmt.Println("Vazia Teste: ", NumInList([]int{}, 5))
}
```

## MDC

**Pasta:** [gcd/gcd.go](gcd/gcd.go)

### Proposta

MDC/GCD significa máximo divisor comum.
Dados dois números, o MDC calcula o maior divisor pelo qual ambos os números podem ser divididos sem deixar resto.

### Minhas decisões

- Escolhi o Algoritmo de Euclides (Algoritmo Euclidiano) por ser a implementação mais simples do cálculo de MDC e mais perfomática. Ele reduz o problema maior a um menor e mais fácil de resolver.
- Existe também o método de Fatoração para encontrar o MDC de dois números. No entanto, este é mais complexo de implementar e menos perfomático. A complexidade de tempo do Algoritmo Euclidiano é O(Log n), enquanto a Fatoração é de O(Raiz quadrada de N). Portanto, a escolha pelo Algoritmo Euclidiano é assegurada por esses fatores.
- Utilizei apenas fmt como stdlib para printar no console o MDC.
- O principal trade-off é a implementação do algoritmo, que pode se dá por meio da Recursão ou Loop For. Optei pela recursão por ser mais prático de implementar. 


### Código
``` go
package main

import "fmt"

// MDC/GCD significa máximo divisor comum.
// Dados dois números, o MDC calcula o maior número pelo qual ambos os números podem ser divididos sem deixar resto.
func GCD(a, b int) int {
	// Usarei o algoritmo Euclidiano para solucionar esse problema, reduzindo-o a um problema menor.
	// Usarei também a função recursiva para a solução.
	// O algoritmo Euclidiano fala em: Dado um MDC(A,B), onde A != 0 e B != 0 => dividi-se A/B, resultando em um novo MDC(B,R), onde R = resto da divisão anterior,
	// reduzindo o MDC até que um dos inteiros seja 0.
	// As 3 Propriedades do Algoritmo são:
		// - MDC(A,0) = A
		// - MDC(0,B) = B
		// - Se A = B⋅Q + R e B≠0, então MDC(A,B) = MDC(B,R) sendo Q um inteiro, e R um inteiro entre 0 e B-1 
	fmt.Printf("MDC(%v, %v)\n", a, b)

	if b == 0 {return a} else {
		return GCD(b,a%b)
	}


}


func main(){
	a, b := 270,192
	fmt.Printf("O máximo divisor comum de %v e %v é %v\n",a , b, GCD(a, b))
}
```

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![linkedin](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)

