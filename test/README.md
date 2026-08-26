# Testes de Benchmark

Esta pasta é utilizada para realizar os testes de benchmark dos scripts desenvolvidos nesse repositório.

## Sobre os testes
- Cada teste será realizado com duas abordagens: 1º Considerada a melhor para aquele cenário; 2º A pior para aquele cenário
- Existe um arquivo em cada pasta de testes que é o `resultado.txt`, onde estarão todos os resultados dos testes daquele script espercífico
- Para executar os testes e visualizar em um arquivo:
    - 1º Crie um arquivo `<nome>.txt`
    - 2º Execute o teste específico `go test -bench=. -benchmem ./test/<pasta de teste>/<arquivo de teste>_test.go -count=10 > ./test/<pasta do teste>/<nome do arquivo de resultado>.txt`
    - 3º Caso possua instalado o `benchstat` Execute a visualização dos resultados no terminal `benchstat ./test/<pasta de teste>/<nome do arquivo de resultado>.txt`

## Índice de testes realizados

- [MDC_test](#mdc_test)

---

## MDC_test

**Pasta:** [test/gcd/gcd_test.go](/test/gcd/gcd_test.go)

### Proposta

Dados dois números, o MDC calcula o maior divisor pelo qual ambos os números podem ser divididos sem deixar resto.
Calculei usando o método da fatoração - ruim - e o algoritmo de Euclides - bom.

## Resultados
### Entrada de dados testados para ambos os casos
``` go
GCD(270, 192)
GCD(20, 10)
```

### Euclides
``` go
- Memória: 0 allocs/op e 0 B/op
- Tempo de processamento médio: 25.89 ns/op
- Número de iterações: 45 milhões vezes
```
### Fatoração
``` go
- Memória: 0 allocs/op e 0 B/op
- Tempo de processamento médio: 1013,68 ns/op ~ 1,01 microssegundos
- Número de iterações: 1,1 milhões vezes
```
### Resultados por entrada
### **GCD(270, 192)**
```
BenchmarkGCD/impl=euclides-4         	43245598	        26.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	44645617	        26.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	45585439	        25.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	45412999	        26.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	45610796	        25.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	44620116	        25.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	45655174	        26.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	43581458	        25.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	41300751	        27.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	43935628	        26.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 1000000	      1025 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 1000000	      1005 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 1000000	      1004 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 1000000	      1012 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 1203979	      1039 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 1183117	       997.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 1186344	      1011 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 1152294	      1003 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 1000000	      1001 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 1000000	      1069 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	command-line-arguments	26.684s
```
### **GCD(20, 10)**
```
BenchmarkGCD/impl=euclides-4         	143150198	         7.779 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	153676141	         8.003 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	149667937	         7.925 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	144178099	         8.792 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	153257985	         7.941 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	152452230	         7.949 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	151223755	         7.777 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	153545863	  package main

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
}       7.890 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	153141804	         7.861 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=euclides-4         	150398152	         8.114 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 2072995	       572.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 2018956	       572.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 2030404	       569.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 2013178	       570.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 2044306	       570.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 1857270	       588.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 2003683	       570.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 2064066	       571.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 2023126	       574.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGCD/impl=fatoracao-4        	 1998865	       567.8 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	command-line-arguments	37.843s
```

## Insights sobre os testes
- O número de iterações é maior no algoritmo de Euclides por conta do mecanismo do Benchmark em Golang que estipula 1 segundo de tempo para rodar cada linha de teste. Por ser mais leve de executar, o algoritmo de Euclides consegue rodar 45 milhões de vezes até bater o tempo estipulado por Go para se chegar a uma média matemática justa. Essa iterações acontecem no Loop, onde existe a variável interna b.N.
- Memória alocada foi igual a 0 em ambos os casos por causa do compilador de Go que utilizou Escape Analysis decidindo pela alocação das variáveis na Stack, e não na Heap. [Referência](https://tip.golang.org/doc/gc-guide#Escape_analysis).
- A diferença de tempo de processamento é de 39,15 vezes.

### Código
``` go
package main

import (
	"testing"
)

// Implementação A: Algoritmo de Euclides com Recusividade
func GCDRecursivo(a, b int) int {
	if b == 0 {return a} else {
		return GCDRecursivo(b,a%b)
	}

}


// Implementação B: Fatorar
func Fatorar(n int) map[int]int {
	fatores := make(map[int]int)
	
	// Trata o fator 2 separadamente para permitir passos maiores no loop seguinte
	for n%2 == 0 {
		fatores[2]++
		n /= 2
	}

	// Testa os números ímpares a partir de 3 até a raiz quadrada de n
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			fatores[i]++
			n /= i
		}
	}

	// Se sobrou algum número maior que 1, ele mesmo é um número primo
	if n > 1 {
		fatores[n]++
	}

	return fatores
}

// GCDFactorization calcula o MDC comparando os fatores primos comuns
func GCDFactorization(a, b int) int {
	if a == 0 { return b }
	if b == 0 { return a }

	fatoresA := Fatorar(a)
	fatoresB := Fatorar(b)
	mdc := 1

	// Compara os fatores comuns e multiplica usando o menor expoente
	for primo, expA := range fatoresA {
		if expB, existe := fatoresB[primo]; existe {
			// Encontra o menor expoente entre as duas fatorações
			menorExp := expA
			if expB < expA {
				menorExp = expB
			}
			
			// Multiplica o MDC pelo primo elevado ao menor expoente
			for i := 0; i < menorExp; i++ {
				mdc *= primo
			}
		}
	}

	return mdc
}

// Único Benchmark que testa ambas as abordagens
func BenchmarkGCD(b *testing.B) {

	b.Run("impl=euclides", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = GCDRecursivo(20, 10)
		}
	})

	b.Run("impl=fatoracao", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = GCDFactorization(20, 10)
		}
	})
}
```

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![linkedin](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)

