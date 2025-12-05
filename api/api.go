// Package api
package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type PrecoPosto struct {
	Precos Combustivel `json:"precos"`
	Data   string      `json:"data_coleta"`
}

type Combustivel struct {
	Diesel   Estados `json:"diesel"`
	Gasolina Estados `json:"gasolina"`
}

type Estados struct {
	Br string `json:"br"`
	Al string `json:"al"`
	Am string `json:"am"`
	Ce string `json:"ce"`
	Df string `json:"df"`
	Es string `json:"es"`
	Go string `json:"go"`
	Ma string `json:"ma"`
	Mt string `json:"mt"`
	Pg string `json:"pg"`
	Mg string `json:"mg"`
	Pr string `json:"pr"`
	Pa string `json:"pa"`
	Pe string `json:"pe"`
	Rs string `json:"rs"`
	Sp string `json:"sp"`
	Sc string `json:"sc"`
	Rj string `json:"rj"`
}

func RequestMath(URL string, num1 string, num2 string) {
	client := http.Client{Timeout: 10 * time.Second}

	req, _ := http.NewRequest("POST", URL, nil)
	req.Header.Set("Request", "Posto")
	req.Header.Set("Number1", num1)
	req.Header.Set("Number2", num2)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao conectar:", err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		println("Erro ao ler o corpo da resposta:", err)
		return
	}
	fmt.Println(string(body))
}

func Request(URL string, estado string, combustivel string) {
	client := http.Client{Timeout: 10 * time.Second}

	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("Request", "Postos")
	req.Header.Set("User-Agent", "Mozilla")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao conectar:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var dados PrecoPosto
	json.Unmarshal(body, &dados)

	fmt.Printf("\n📅 Data da coleta: %s\n", dados.Data)

	var precoEstados Estados

	if combustivel == "Diesel" {
		precoEstados = dados.Precos.Diesel
	} else {
		precoEstados = dados.Precos.Gasolina
	}

	if estado == "Todos" {
		PrintAll(precoEstados)
	} else {
		PrintSingle(estado, precoEstados)
	}
}

func PrintSingle(estado string, e Estados) {
	valor := getField(e, strings.ToLower(estado))
	fmt.Printf("\n💰 Preço em %s: R$ %s\n", estado, valor)
}

func PrintAll(e Estados) {
	fmt.Println("\n💰 Preços em todos os Estados:")
	fmt.Println("───────────────────────────────")

	est := map[string]string{
		"Br": e.Br, "Al": e.Al, "Am": e.Am, "Ce": e.Ce, "Df": e.Df,
		"Es": e.Es, "Go": e.Go, "Ma": e.Ma, "Mt": e.Mt, "Pg": e.Pg,
		"Mg": e.Mg, "Pr": e.Pr, "Pa": e.Pa, "Pe": e.Pe, "Rs": e.Rs,
		"Sp": e.Sp, "Sc": e.Sc, "Rj": e.Rj,
	}

	for uf, preco := range est {
		fmt.Printf(" • %s → R$ %s\n", uf, preco)
	}
}

// pega o valor do campo pelo nome
func getField(e Estados, nome string) string {
	v := map[string]string{
		"br": e.Br, "al": e.Al, "am": e.Am, "ce": e.Ce, "df": e.Df,
		"es": e.Es, "go": e.Go, "ma": e.Ma, "mt": e.Mt, "pg": e.Pg,
		"mg": e.Mg, "pr": e.Pr, "pa": e.Pa, "pe": e.Pe, "rs": e.Rs,
		"sp": e.Sp, "sc": e.Sc, "rj": e.Rj,
	}
	return v[nome]
}

func printRaw() {}
