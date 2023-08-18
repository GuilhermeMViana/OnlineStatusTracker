package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

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
			readLog()
			for i := 3; i > 0; i-- {
				fmt.Println("Voltando para o menu em", i, "segundos.")
				time.Sleep(1 * time.Second)
			}
		case 3:
			addSite()
		case 0:
			fmt.Println("Programa finalizado")
			os.Exit(0)
		default:
			fmt.Println("Verifique seu comando")
			for i := 3; i > 0; i-- {
				fmt.Println("Voltando para o menu em", i, "segundos.")
				time.Sleep(1 * time.Second)
			}

			showMenu()
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
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Adicionar site")
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

	//for i := 0 ; i < len(sites) ; i++
	for i, site := range sites {
		fmt.Println("Testando", i+1, "º site:", site)
		testSite(site)

		time.Sleep(delay * time.Second)
	}

}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("ERRO na conexão:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O Site", site, "foi carregado com sucesso!", resp.StatusCode)
		logRegister(site, true)
	} else {
		fmt.Println("O Site", site, "encontrou o problema: ", resp.StatusCode)
		logRegister(site, false)
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
		sites = append(sites, line)
	}

	closeErr := file.Close()
	if closeErr != nil {
		return nil
	}

	return sites
}

func logRegister(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("ERRO:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04") + " - " + site + "- Online:" + strconv.FormatBool(status) + "\n")

	closeErr := file.Close()

	if closeErr != nil {
		fmt.Println("ERRO:", closeErr)
	}
	fmt.Println("Log Gerado!")
}

func readLog() {
	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("ERRO:", err)
	}

	fmt.Println(string(file))

}

func addSite() {
	var command int
	fmt.Println("Deseja adicionar um novo site?")
	fmt.Println("1 - Adicionar um novo site")
	fmt.Println("0 - Voltar para o menu")
	_, err := fmt.Scan(&command)
	if err != nil {
		fmt.Println("Erro ao adicionar novo site:", err)
	}
	switch command {
	case 1:
		fmt.Println("Insira a url do site")
		var link string
		_, err := fmt.Scan(&link)

		if err != nil {
			fmt.Println("Problema com o link:", err)
		}

		file, _ := os.OpenFile("sites.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

		if file != nil {
			_, err := file.WriteString(link + "\n")

			if err != nil {
				fmt.Println("Erro ao adicionar linha:", err)
			}
		}

		file.Close()
		showMenu()
	case 0:
		showMenu()
	default:
		showMenu()

	}

}
