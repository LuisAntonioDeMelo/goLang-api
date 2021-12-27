package main

import "fmt"
import "reflect"


func main() {
	nome  := "Luis"
	idade := 20
	versao  := 1.1
	fmt.Println("Olá sr." , reflect.TypeOf(nome))
	fmt.Println("idade ", idade)
	fmt.Println("Este programa está na versão", versao)
}