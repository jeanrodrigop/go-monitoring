package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const delay = 5

func main() {

	introducao()

	for {
		exibeMenu()

		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa.")
			os.Exit(0)
		default:
			fmt.Println("Comando não existe")
			os.Exit(-1)
		}
	}

}

func introducao() {
	nome := "Jean Rodrigo"
	versao := 1.1

	fmt.Println("Olá,", nome)
	fmt.Println("Este programa esta na versão", versao)
}

func exibeMenu() {
	fmt.Println("1-Iniciar o monitoramento")
	fmt.Println("2-Exibir logs")
	fmt.Println("0-Sair do programa")
}

func lerComando() int {
	var comandoLido int

	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	var siteLido string

	fmt.Println("Qual site deseja monitorar?")
	fmt.Scan(&siteLido)

	site := "https://" + siteLido

	fmt.Print("Monitorando, pressione \"x\" para encerrar...\n")
	for {
		resp, err := http.Get(site)
		if err != nil {
			fmt.Printf("\rErro ao acessar o site: %v", err)
			continue
		}

		if resp.StatusCode == 200 {
			fmt.Printf("\rSite: %s foi carregado com sucesso!", siteLido)
		} else {
			fmt.Printf("\rSite: %s está com problemas. Status Code: %d", siteLido, resp.StatusCode)
		}

		time.Sleep(delay * time.Second)

		var parar string
		fmt.Scan(&parar)

		if parar == "x" {
			fmt.Println("\rMonitoramento encerrado.")

			break
		}
	}
}
