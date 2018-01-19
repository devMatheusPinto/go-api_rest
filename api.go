package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	rotas := mux.NewRouter().StrictSlash(true)

	rotas.HandleFunc("/clientes", getClientes).Methods("GET")
	rotas.HandleFunc("/clientes/{idNumber}", getOneClientes).Methods("GET")
	rotas.HandleFunc("/clientes", postClientes).Methods("POST")

	rotas.HandleFunc("/fornecedores", getFornecedores).Methods("GET")
	rotas.HandleFunc("/fornecedores/{idNumber}", getOneFornecedores).Methods("GET")
	rotas.HandleFunc("/fornecedores", getFornecedores).Methods("POST")

	rotas.HandleFunc("/produtos", getProdutos).Methods("GET")
	rotas.HandleFunc("/produtos/{idNumber}", getOneProdutos).Methods("GET")
	rotas.HandleFunc("/produtos", getProdutos).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", rotas))
}

type Cliente struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

type Fornecedor struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
	CNPJ string `json:"cnpj"`
}

type Produto struct {
	ID            int    `json:"id"`
	ID_Fornecedor int    `json:"id_fornecedor"`
	Nome          string `json:"nome"`
	Quantidade    int    `json:"quantidade"`
}

var clientes = []Cliente{
	Cliente{ID: 1, Nome: "Matheus Pinto", Email: "matheus@teste.com"},
	Cliente{ID: 2, Nome: "Rafaele Pinto", Email: "rafaele@teste.com"},
}

func getClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(clientes)
}

func getOneClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vars = mux.Vars(r)
	idNumber, _ := strconv.Atoi(vars["idNumber"])

	for _, values := range clientes {
		if values.ID == idNumber {
			json.NewEncoder(w).Encode(values)
		}
	}
}

func postClientes(w http.ResponseWriter, r *http.Request) {
	var t Cliente

	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &t)

	clientes = append(clientes, t)
}

var fornecedores = []Fornecedor{
	Fornecedor{ID: 1, Nome: "Lojinha", CNPJ: "00.000.001/0001-01"},
	Fornecedor{ID: 2, Nome: "Loja", CNPJ: "00.100.110/0001-00"},
}

func getFornecedores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(fornecedores)
}

func getOneFornecedores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vars = mux.Vars(r)
	idNumber, _ := strconv.Atoi(vars["idNumber"])

	for _, values := range fornecedores {
		if values.ID == idNumber {
			json.NewEncoder(w).Encode(values)
		}
	}
}

func postFornecedores(w http.ResponseWriter, r *http.Request) {
	var t Fornecedor

	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &t)

	fornecedores = append(fornecedores, t)
}

var produtos = []Produto{
	Produto{ID: 1, ID_Fornecedor: 2, Nome: "Xbox One", Quantidade: 103},
	Produto{ID: 2, ID_Fornecedor: 1, Nome: "PS4", Quantidade: 86},
}

func getProdutos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(produtos)
}

func getOneProdutos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vars = mux.Vars(r)
	idNumber, _ := strconv.Atoi(vars["idNumber"])

	for _, values := range produtos {
		if values.ID == idNumber {
			json.NewEncoder(w).Encode(values)
		}
	}
}

func postProdutos(w http.ResponseWriter, r *http.Request) {
	var t Produto

	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &t)

	produtos = append(produtos, t)
}
