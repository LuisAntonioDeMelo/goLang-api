package main

import "fmt"
import "os"
import "net/http"
import "time"
import "io"
import "bufio"

func main()  {
	
	exibeIntroducao()
	fmt.Println("")
	for {
		exibeMenu()
		comando := lerComando()
		
	
		switch comando {
		case 1:
			iniciarMonitoramento()
			fmt.Println("")
		case 2:
			fmt.Println("Exibindo Logs ...")
		case 0:
			fmt.Println("Saindo do programa ...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}


func exibeIntroducao()  {
	nome := ""
	versao := 1.0

	fmt.Println("MApS - monitoramento ::::::")
	fmt.Println("Acesso", nome," @versao : ", versao)
	fmt.Println("")
}

func exibeMenu(){
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
}

func lerComando() int {
	var comando int 
	//fmt.Scanf("%d", &comando) // scanf
	fmt.Scan(&comando)
	fmt.Println("Comando :", comando, "foi selecionado!") 
	return comando
}

func iniciarMonitoramento(){
	fmt.Println("Monitorando ...")
	
	// sites := []string {
	// 	"https://random-status-code.herokuapp.com/",
	// 	"https://www.alura.com.br",
	// 	"https://www.caelum.com.br",
	// }

   sites := leSitesDoArquivo()
	//for i:= 0; i < len(sites); i++ {
   for i:= 0; i < 5; i++ {
	for i, site := range sites {
		fmt.Println("Testando site ",i, ":", site)
		testeSite(site)
	}
	time.Sleep(5 * time.Second )
   }
}

func testeSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:" , err)
	}else {

		if resp.StatusCode == 200 {
			fmt.Println("Site:", site, "foi carregado com sucesso!")
		}else {
			fmt.Println("Site:", site, "está com problemas.")
		}	
	}

}


func leSitesDoArquivo() []string {

    var sites []string

    arquivo, err := os.Open("sites.txt")
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    for {
        linha, err := leitor.ReadString('\n')
        fmt.Println(linha)
        if err == io.EOF {
            break
        }
    }

    return sites
}