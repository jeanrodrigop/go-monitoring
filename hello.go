package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
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
			lerLogs()
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

	fmt.Print("Monitorando ", siteLido, "...\n")
	for {
		resp, err := http.Get(site)
		if err != nil {
			fmt.Printf("\rErro ao acessar o site: %v", err)
			continue
		}

		if resp.StatusCode == 200 {
			fmt.Printf("\rSite: %s foi carregado com sucesso!", siteLido)
			escreveLog(siteLido, true)
		} else {
			fmt.Printf("\rSite: %s está com problemas. Status Code: %d", siteLido, resp.StatusCode)
			escreveLog(siteLido, false)
		}

		time.Sleep(delay * time.Second)
	}
}

func escreveLog(siteLido string, status bool) {
	arquivo, err := os.OpenFile("monitoramento.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + siteLido + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func lerLogs() {

	arquivo, err := os.ReadFile("monitoramento.log")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
