package main

import (
	"fmt"

	"go-cli/api"

	"github.com/manifoldco/promptui"
)

var Estados = map[string][]string{
	"Diesel": {
		"Todos", "Br", "Al", "Am", "Ce", "Df", "Es", "Go",
		"Ma", "Mt", "Pg", "Mg", "Pr", "Pa", "Pe",
		"Rs", "Sp", "Sc", "Rj",
	},
	"Gasolina": {
		"Todos", "Br", "Al", "Am", "Ce", "Df", "Es", "Go",
		"Ma", "Mt", "Pg", "Mg", "Pr", "Pa", "Pe",
		"Rs", "Sp", "Sc", "Rj",
	},
}

func main() {
	UIHeader()

	URL := "https://combustivelapi.com.br/api/precos/"

outer:
	for {

		// ─────────────────────────
		// Escolha Combustível
		// ─────────────────────────
		combustivelPrompt := promptui.Select{
			Label: "⛽ Selecione o tipo de combustível",
			Items: []string{"Diesel", "Gasolina"},
		}
		_, combustivel, _ := combustivelPrompt.Run()

		// ─────────────────────────
		// Escolha Estado
		// ─────────────────────────
		estadoPrompt := promptui.Select{
			Label: "📍 Selecione o estado",
			Items: Estados[combustivel],
		}
		_, estado, _ := estadoPrompt.Run()

		fmt.Println("\n-----------------------------")
		fmt.Printf(" 🔎 Consultando preços para %s em %s...\n", combustivel, estado)
		fmt.Println("-----------------------------")

		api.Request(URL, estado, combustivel)

		// ─────────────────────────
		// Voltar ou Encerrar
		// ─────────────────────────
		sairPrompt := promptui.Select{
			Label: "Escolha uma opção",
			Items: []string{"↩ Voltar ao menu", "⛔ Encerrar"},
		}
		_, sair, _ := sairPrompt.Run()

		if sair == "↩ Voltar ao menu" {
			continue
		}
		break outer
	}
}

func UIHeader() {
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("      ⛽ CONSULTA DE PREÇOS – CLI       ")
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("   Busque valores de Diesel e Gasolina ")
	fmt.Println("═══════════════════════════════════════\n")
}
