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
	URLSOMA := "http://localhost:6969/soma"
	URLSUB := "http://localhost:6969/sub"
	URLDIV := "http://localhost:6969/div"
	URLMULT := "http://localhost:6969/mult"

	cStr := "$ Consultar Preços de combustível"
	mStr := "📐 Utilizar a API de calculadora"
	subStr := "➖ Subtração"
	somaStr := "➕ Soma"
	divStr := "➗ Divisão"
	multStr := "✖ Multiplicação"

outer:
	for {
		serviçoPrompt := promptui.Select{
			Label: "Selecione o serviço desejado",
			Items: []string{mStr, cStr},
		}
		_, serviço, _ := serviçoPrompt.Run()
		// copilot me da o emoji que tem a ver com matemática

		// ─────────────────────────
		// Escolha Combustível
		// ─────────────────────────
		if serviço == cStr {
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

			sairPrompt := promptui.Select{
				Label: "Escolha uma opção",
				Items: []string{"↩ Voltar ao menu", "⛔ Encerrar"},
			}
			_, sair, _ := sairPrompt.Run()

			if sair == "↩ Voltar ao menu" {
				continue
			}
			break outer
		} else {
			promptOperacao := promptui.Select{
				Label: "📐 Selecione a operação matemática",
				Items: []string{"➕ Soma", "➖ Subtração", "✖ Multiplicação", "➗ Divisão"},
			}
			_, operacao, _ := promptOperacao.Run()
			// copilot me ensina a pegar a input do usuário para fazer operações matemáticas usando a API

			number1Prompt := promptui.Prompt{
				Label: "Digite o primeiro número",
			}
			num1, _ := number1Prompt.Run()

			number2Prompt := promptui.Prompt{
				Label: "Digite o segundo número",
			}
			num2, _ := number2Prompt.Run()
			switch operacao {
			case somaStr:
				api.RequestMath(URLSOMA, num1, num2)
			case subStr:
				api.RequestMath(URLSUB, num1, num2)
			case divStr:
				api.RequestMath(URLDIV, num1, num2)
			case multStr:
				api.RequestMath(URLMULT, num1, num2)
			}

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
}

func UIHeader() {
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("      ⛽ CONSULTA DE PREÇOS – CLI       ")
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("   Busque valores de Diesel e Gasolina ")
	fmt.Println("═══════════════════════════════════════\n")
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("     📐  CALCULADORA SIMPLES – CLI     ")
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("   Faça operações simples com uma API  ")
	fmt.Println("═══════════════════════════════════════\n")
}
