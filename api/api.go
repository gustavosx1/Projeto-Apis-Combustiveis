// Package api
package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func Request(URL string, estado string, combustivel string) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	var dados PrecoPosto
	err = json.Unmarshal(body, &dados)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nStatus:\n %v", resp.Status)
	fmt.Println("\nData de Coleta dos Dados:\n", dados.Data)
	switch combustivel {
	case "diesel":
		// swtich estado
		fmt.Println("\nPreços:\n", dados.Precos.Diesel)
	case "gasolina":
		// switch estado
		fmt.Println("\nPreços:\n", dados.Precos.Gasolina)
	}
}
