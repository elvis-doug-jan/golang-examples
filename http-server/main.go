package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
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
	http.HandleFunc("/cep", BuscaCEPHandler)
	http.ListenAndServe(":8080", nil)
}

type ResponseTeste struct {
	Message ViaCEP `json:"message"`
	Status  bool   `json:"status"`
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/cep" && r.URL.Path == "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("value")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	responseCEP, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var response ResponseTeste
	response.Message = *responseCEP
	response.Status = true
	json.NewEncoder(w).Encode(response)
	// bodyResponse, err := json.Marshal(&response)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(bodyResponse)
}

func BuscaCep(cep string) (*ViaCEP, error) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var data ViaCEP
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
