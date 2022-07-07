package main

import (
	"fmt"
	"time"
)

func update(m *map[string]int, word string){
	for _,r := range word {
		(*m)[string(r)]++
	}
}
func main() {
	words := []string{"quick","brown","fox","lazy","dog"}
	var AlphabetMap =make(map[string]int)
	for _,word := range words{
		go update(&AlphabetMap, word)
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println(AlphabetMap)
}
