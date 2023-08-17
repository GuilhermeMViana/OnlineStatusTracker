package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	showIntroduction()
	for {
		showMenu()

		command := receiveCommand()

		switch command {
		case 1:
			handleVerifySiteStatus()
		case 2:
			fmt.Println("Gerando Logs...")
		case 0:
			fmt.Println("Programa finalizado")
			os.Exit(0)
		default:
			fmt.Println("Verifique seu comando")
			os.Exit(-1)
		}
	}

}

func showIntroduction() {
	var name string
	fmt.Println("Qual é o seu nome:")
	fmt.Scanf("%s", &name)
	version := 1.1
	fmt.Println("Olá, Sr(a).", name)
	fmt.Println("Versão atual do programa: ", version)
}

func showMenu() {
	fmt.Println("---------------------------")
	fmt.Println("1 - Iniciar Monitoranto")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
	fmt.Println("---------------------------")
}

func receiveCommand() int {
	var command int

	fmt.Println("Selecione uma opção: ")
	fmt.Scan(&command)

	return command
}

func handleVerifySiteStatus() {

	fmt.Println("Monitorando...")
	sites := []string{"https://www.alura.com.br", "https://go.dev"}
	resp, _ := http.Get(sites[1])
	fmt.Println(resp.StatusCode)

	if resp.StatusCode == 200 {
		fmt.Println("O Site", sites[1], "foi carregado com sucesso!")
	} else {
		fmt.Println("O Site", sites[1], "encontrou o problema: ", resp.StatusCode)
	}
}
