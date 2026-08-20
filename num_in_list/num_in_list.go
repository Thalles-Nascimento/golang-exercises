package main

import "fmt"

// NumInList retornará verdadeiro se o número estiver na
// Slice list, e falso caso contrário.

// Vou utilizar uma pesquisa simples e não uma pesquisa binária, pois serão testadas as seguintes listas:
// NumInList([]int{1,2,3,4,5}, 5) // true
// NumInList([]int{3,3,3,3,3}, 5) // false
// NumInList([]int{3,5,3,5,3}, 5) // true
// NumInList([]int{4,2,22,-10,8}, -10) // true
// empty lists!
// NumInList(nil, 5) // false
// NumInList([]int{}, 5) // false
func NumInList(list []int, num int) bool {
	// Ao usar a variável numList dentro do loop o Go faz uma cópia dos números da lista a cada iteração e insere na variável. Com tipos primitivos, não há perda de perfomance
	// e não é muito relevante. Mas para tipos structs pode causar uma perda de perfomance dependo do tamanho do struct.
	// Para deixar esse loop enxuto e otimizado, podemos rodar esse comando que aceita qualquer tipo: slices.Contains(list, num), ou criar um loop assim:
		// for i := range list { <-- o slices.Contains(list, num) funciona exatamente assim. Ele chama um função Index passando como parâmetro o slice e o valor, a função Index retorna um inteiro (o índice do valor encontrado ou -1 para não encontrado) e a função Contains do slice retorna o bool de acordo com o retorno de Index
		// 	if list[i] == num { 
		// 		return true
		// 	} 		
		// }
	// Por se tratar de números inteiros, a cópia é quase instantânea para a variável e não há perda de perfomance ou seria otimizado se não fizesse a cópia, por isso a escolha
	// pelo loop abaixo
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

	// Empty teste
	fmt.Println("Nil Teste: ", NumInList(nil, 5))
	fmt.Println("Vazia Teste: ", NumInList([]int{}, 5))

}
