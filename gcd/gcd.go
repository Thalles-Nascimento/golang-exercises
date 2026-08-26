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