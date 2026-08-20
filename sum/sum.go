package main

import "fmt"

// Você deverá somar todos os números passados ​​como entrada e retornar o resultado.

func Sum(numbers []int) int {
	var acc int // ZeroValue - Significa que a variável do tipo int inicializará com valor padrão = 0

	for _, num := range numbers {
		acc += num
	}

	return acc
}

func main(){
	fmt.Println("Primeiro Teste: ", Sum([]int{1,2,3,4,5})) // 15
	fmt.Println("Segundo Teste: ", Sum([]int{3, 3})) // 6
	fmt.Println("Terceiro Teste: ", Sum([]int{3,5,3,5,3})) // 19
	fmt.Println("Quarto Teste: ", Sum([]int{4,2,22,-10,8})) // 26

	// Nil ou vazia
	fmt.Println("Nil: ", Sum(nil))
	fmt.Println("Vazia: ", Sum([]int{}))

}
