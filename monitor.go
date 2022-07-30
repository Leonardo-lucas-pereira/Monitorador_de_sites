package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// Constantes
const monitoramentos = 3
const delay = 5

// Função principal
func main() {

	exibeIntroducao()

	for {

		exibeMenu()

		comando := leComando()
		fmt.Println()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

// Funções secundarias
func exibeIntroducao() {
	fmt.Println("######################################")
	nome := "Leonardo"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println()
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println("######################################")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	fmt.Println()
	sites := []string{"https://www.alura.com.br",
		"https://random-status-code.herokuapp.com",
		"https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		fmt.Println("Sleeping...")
		time.Sleep(delay * time.Second)
		fmt.Println()

	}

	fmt.Println()
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println(" [+] -> Site:", site, "foi carregado com sucesso!")
		fmt.Println()
	} else {
		fmt.Println(" [-] ->Site:", site, "esta com problemas. Staus Code:", resp.StatusCode)
		fmt.Println()
	}
}
