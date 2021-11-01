package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manishdangi98/banking/domain"
	"github.com/manishdangi98/banking/service"
)

func Start() {

	router := mux.NewRouter()
	//wiring
	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
