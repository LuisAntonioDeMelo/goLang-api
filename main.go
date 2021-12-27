package main

import "net/http"
import "fmt"
import ("text/template")


type Produto struct {
	nome string
	descricao string
	preco float64
	quantidade int 
} 

var temp =  template.Must(template.ParseGlob("templates/*.html"))

func main()  {
	fmt.Println("start :: 🔥🔥 server >> port :8000")
	
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000",nil)
}

func index(w http.ResponseWriter, r *http.Request){
	produtos := []Produto {
		{nome: "Camisete", descricao: "Azul, bem bonita", preco:39, quantidade:5},
		{"Tenis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Produto n", "teste novo", 10,1},
	}
	fmt.Println(produtos)
	temp.ExecuteTemplate(w,"index", produtos)
}