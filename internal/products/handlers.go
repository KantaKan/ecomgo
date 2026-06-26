package products

import (
	"net/http"

	"github.com/KantaKan/ecomgo/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	//call the service  > list products
	//2 return JSON in http response
	products :=  struct {
		Products []string `json:"products"`
	}{}

	json.Write(w, http.StatusOK,products)
	
}