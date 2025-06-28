package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Agendamento struct {
	Nome string `json:"nome"`
	Data string `json:"data"`
	Hora string `json:"hora"`
}

func coletarAgendamentosExistentes() []Agendamento {
	var agendamentos []Agendamento
	arquivo, err := os.Open("agendamentos.json")
	if err == nil {
		defer arquivo.Close()
		decoder := json.NewDecoder(arquivo)
		decoder.Decode(&agendamentos)
	}
	return agendamentos
}

func salvarAgendamentos(agendamentos []Agendamento) {
	arquivo, err := os.Create("agendamentos.json")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer arquivo.Close()

	encoder := json.NewEncoder(arquivo)
	err = encoder.Encode(agendamentos)
	if err != nil {
		fmt.Println("Erro ao salvar dados:", err)
		return
	}
}

func mostrarAgendamentos(agendamentos []Agendamento) {
	if len(agendamentos) == 0 {
		fmt.Println("Nenhum agendamento cadastrado.")
		return
	}
	fmt.Println("\nAgendamentos cadastrados:")
	for _, ag := range agendamentos {
		fmt.Printf("%s - %s - %s\n", ag.Nome, ag.Data, ag.Hora)
	}
}

func adicionarAgendamento(agendamentos []Agendamento) []Agendamento {
	var nome, data, hora string
	fmt.Print("Digite o nome do cliente: ")
	fmt.Scanln(&nome)
	fmt.Print("Digite a data do agendamento (ex: 10/10/2025): ")
	fmt.Scanln(&data)
	fmt.Print("Digite o horário do agendamento: ")
	fmt.Scanln(&hora)

	novo := Agendamento{Nome: nome, Data: data, Hora: hora}
	agendamentos = append(agendamentos, novo)
	fmt.Println("Agendamento adicionado com sucesso!")
	return agendamentos
}

func removerAgendamento(agendamentos []Agendamento) []Agendamento {
	if len(agendamentos) == 0 {
		fmt.Println("Nenhum agendamento para remover.")
		return agendamentos
	}

	var nome string
	fmt.Print("Digite o nome do cliente que deseja remover o agendamento: ")
	fmt.Scanln(&nome)

	var novoSlice []Agendamento
	var removido bool
	for _, ag := range agendamentos {
		if ag.Nome != nome {
			novoSlice = append(novoSlice, ag)
		} else {
			removido = true
		}
	}

	if removido {
		fmt.Println("Agendamento removido com sucesso!")
	} else {
		fmt.Println("Nenhum agendamento encontrado com esse nome.")
	}

	return novoSlice
}

func main() {
	agendamentos := coletarAgendamentosExistentes()

	for {
		fmt.Println("\nEscolha uma opção:")
		fmt.Println("1 - Adicionar agendamento")
		fmt.Println("2 - Remover agendamento")
		fmt.Println("3 - Mostrar agendamentos")
		fmt.Println("4 - Sair")
		fmt.Print("Opção: ")

		var opcao int
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			agendamentos = adicionarAgendamento(agendamentos)
			salvarAgendamentos(agendamentos)
		case 2:
			agendamentos = removerAgendamento(agendamentos)
			salvarAgendamentos(agendamentos)
		case 3:
			mostrarAgendamentos(agendamentos)
		case 4:
			fmt.Println("Saindo... Até mais!")
			return
		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}
