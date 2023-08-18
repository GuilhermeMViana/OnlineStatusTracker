package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const aux = 1
const delay = 3

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
			readFile()
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
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("ERRO:", err)
	}

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
	_, err := fmt.Scan(&command)
	if err != nil {
		fmt.Println("ERRO:", err)
	}

	return command
}

func handleVerifySiteStatus() {

	fmt.Println("Monitorando...")
	sites := readFile()

	for i := 0; i < aux; i++ {
		//for i := 0 ; i < len(sites) ; i++
		for i, site := range sites {
			fmt.Println("Testando", i+1, "º site:", site)
			testSite(site)

			time.Sleep(delay * time.Second)
		}
	}

}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("ERRO na conexão:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O Site", site, "foi carregado com sucesso!", resp.StatusCode)
	} else {
		fmt.Println("O Site", site, "encontrou o problema: ", resp.StatusCode)
	}
}

func readFile() []string {

	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("ERRO na leitura do arquivo:", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		sites = append(sites, line)
	}

	erro := file.Close()
	if erro != nil {
		return nil
	}

	return sites
}
