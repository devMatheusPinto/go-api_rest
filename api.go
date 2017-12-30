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
	Fornecedor{ID: 1, Nome: "Americanas", CNPJ: "33.014.556/0001-96"},
	Fornecedor{ID: 2, Nome: "B2W - Companhia Digital", CNPJ: "00.776.574/0006-60"},
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
