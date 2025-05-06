package handlers

import (
	"e5/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, er := strconv.Atoi(vars["id"])
	if er != nil {
		http.Error(w, "Unable to find ID", http.StatusBadRequest)
	}
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to unmarshal Json", http.StatusBadRequest)
	}
	erp := data.UpdateProduct(id, prod)

	if erp == data.ErrProductNotFound || erp != nil {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}

}

func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Product")
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
