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