package handlers

import (
	"github.com/gorilla/mux"
	"github.com/thitemalhar/GoLang/Projects/Product-API/data"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products  {
	return &Products{l}
}

//func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
//	if r.Method == http.MethodGet {
//		p.getProducts(rw, r)
//		return
//	}
//	if r.Method == http.MethodPost {
//		p.addProducts(rw, r)
//		return
//	}
//
//	rw.WriteHeader(http.StatusMethodNotAllowed)
//}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "Json error, unable to marshal", http.StatusInternalServerError)
	}
	//d, err := json.Marshal(lp)
	//if err != nil {
	//	http.Error(rw, "Json error, unable to marshal", http.StatusInternalServerError)
	//}
	p.l.Println("Json is returned")
	//rw.Write(d)
}

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to Unmarshal json", http.StatusBadRequest)
	}

	//p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unble to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", id)

	prod := &data.Product{}

	err = prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal the json", http.StatusBadRequest)
	}

	err = data.UpdateProducts(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}
}

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unble to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", id)

	prod := &data.Product{}
	err = data.DeleteProducts(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}
}