package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"argument"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, argument := range os.Args[1:] {
		url := "https://viacep.com.br/ws/" + argument + "/json/"
		fmt.Printf("Buscando CEP %s", argument)
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v", err)
		}
		// fmt.Printf("%s", res)
		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse do JSON: %v", err)
		}
		fmt.Println(data)
		file, err := os.Create("resultado.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v", err)
		}
		defer file.Close()
		file.WriteString(fmt.Sprintf("%s", data))
		if argument == "delete" {
			err := os.Remove("resultado.txt")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao deletar arquivo: %v", err)
			}
		}
	}
}
