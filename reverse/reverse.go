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
// 	 Reverse("ca t") => t ac
//
func Reverse(word string) string {
	
	// Poderia aplicar a conversão de tipo de string para um slice de bytes. 
	// No entanto, há um problema em inverter bytes: strings com caracteres especiais irão quebrar, pois são representados por 2 ou mais bytes 
	// (até 4 bytes). Quis testar isso na prática e por isso realizei um teste com a palavra "Goiás"

	// Para resolver o exercício, terei que transformar a strings em slices de runes, que são representações Code Points (4 bytes). 
	// É um álias para int32 - É uma representação dos caracteres Unicode.

	// O type rune usa 4 bytes para representar um caracter e portanto são mais pesados. Porém para realizar operações de inversão, sai mais barato e rápido realizar
	// a conversão em rune, inverter e depois voltar para byte.
	word_rune := []rune(word)

	slices.Reverse(word_rune)

	var word_reversed strings.Builder
	// Passando o Grow para o Builder o tamanho exato da nova string. Importante para ele saber o tamanho exato em bytes
	word_reversed.Grow(len(word))

	for _, s := range word_rune {
		word_reversed.WriteRune(s)
	}

	return word_reversed.String()
}


func main(){
	// "cat" => 3 letras -- 3 bytes -- Rune = 12 bytes
	fmt.Println(Reverse("cat"))

	// "ca t" => 3 letras + um espaço -- 4 bytes -- Rune = 16 bytes
	fmt.Println(Reverse("ca t"))

	// "alphabet" => 8 letras -- 8 bytes -- Rune = 32 bytes
	fmt.Println(Reverse("alphabet"))

	// Usando esse teste para validar a inversão com caracter especial
	// "Goiás" => 5 letras -- 6 bytes -- Rune = 20 bytes
	fmt.Println(Reverse("Goiás")) // "á" é representado por 2 bytes em UTF-8.


}
