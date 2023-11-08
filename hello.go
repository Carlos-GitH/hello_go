package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	intro()
	for {
		comando := lerComando()
		// if comando == 1 {
		// 	println("Iniciando Monitoramento")
		// } else if comando == 2 {
		// 	println("Exibindo Logs")
		// } else if comando == 0 {
		// 	println("Saindo do Programa")
		// } else {
		// 	println("Comando inválido")
		// }
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			println("Exibindo Logs")
			lerLogs()
		case 0:
			println("Saindo do Programa")
			os.Exit(0)
		default:
			println("Comando inválido")
			os.Exit(-1)
		}
	}
}

func intro() {
	versao := "1.0"
	println("Olá, versão do programa", versao)

	println("1- Iniciar Monitoramento")
	println("2- Exibir Logs")
	println("0- Sair do Programa")
}

func lerComando() int {
	var comando int
	fmt.Scanf("%d", &comando)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// site := "https://www.google.com/"

	// sites := []string{
	// 	"https://www.google.com/",
	// 	"https://www.youtube.com/",
	// 	"https://www.facebook.com/",
	// 	"https://www.x.com/",
	// }

	sites := leSitesDoArquivo()

	for site := range sites {
		testeSites(sites[site])
	}
}

func testeSites(site string) {
	response, _ := http.Get(site)
	println(response.Status)

	if response.StatusCode == 200 {
		fmt.Println("Consegui acessar o site", site, "Status Code:", response.StatusCode)
		registraLog(site, true)
	} else {
		fmt.Println("Não consegui acessar o site", site, "Status Code:", response.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		println("Ocorreu um erro:", err)
	} else {
		fmt.Println("Arquivo aberto com sucesso")
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}

	}
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		println("Ocorreu um erro:", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func lerLogs() {
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		println("Ocorreu um erro:", err)
	}
	println(string(arquivo))
}
