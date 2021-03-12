package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/hitolv4/fm-api/data"
)

type Repuestos struct {
	l *log.Logger
}

func NewRepuestos(l *log.Logger) *Repuestos {
	return &Repuestos{l}
}

func (re *Repuestos) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case "GET":
		re.getRepuestos(w, r)
		return
	case "POST":
		re.addRepuestos(w, r)
	case "PUT":
		re.UpdateRepuestos(w, r)
	case "DELETE":
		re.DeleteRepuestos(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (re *Repuestos) getRepuestos(w http.ResponseWriter, r *http.Request) {
	re.l.Println("Handle GET Repuestos")
	listR := data.GetRespuestos()
	err := listR.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal Json", http.StatusInternalServerError)
	}
}

func (re *Repuestos) addRepuestos(w http.ResponseWriter, r *http.Request) {
	re.l.Println("Handle POST Repuestos")
	newR := data.Repuesto{}
	err := newR.FromJson(r.Body)
	if err != nil {
		http.Error(w, "Can't save new data", http.StatusBadRequest)

	}
	if newR.Name == "" {
		http.Error(w, "Name is empty", http.StatusBadRequest)
		return
	}
	if newR.Code == "" {
		http.Error(w, "Code is empty", http.StatusBadRequest)
		return
	}
	if newR.Moto == 0 {
		http.Error(w, "Moto is empty", http.StatusBadRequest)
		return
	}
	if newR.Price == 0 {
		http.Error(w, "Price is empty", http.StatusBadRequest)
		return
	}
	if newR.Anaquel == 0 {
		http.Error(w, "Anaquel is empty", http.StatusBadRequest)
		return
	}
	if newR.Seccion == "" {
		http.Error(w, "Seccion is empty", http.StatusBadRequest)
		return
	}
	if newR.Caja == 0 {
		http.Error(w, "Caja is empty", http.StatusBadRequest)
		return
	}
	if newR.Cantidad == 0 {
		http.Error(w, "Cantidad is empty", http.StatusBadRequest)
		return
	}
	re.l.Printf("Respuesto %#v", newR)
	data.AddRepuestos(&newR)
	fmt.Fprintf(w, "data saved \n %#v  ", newR)
}

// must refactor later
func (re *Repuestos) UpdateRepuestos(w http.ResponseWriter, r *http.Request) {
	re.l.Println("Handle PUT Repuestos", r.URL.Path)

	id, err := IdFromURL(r)
	if err != nil {
		re.l.Println(err)
		return
	}
	UpdateR := data.Repuesto{}
	err = UpdateR.FromJson(r.Body)
	if err != nil {
		http.Error(w, "Can't Update Respuesto", http.StatusBadRequest)
	}
	err = data.UpdateRepuesto(uint(id), UpdateR)
	if err == data.ErrProductNotFound {
		http.Error(w, "Prduct not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
	}
	fmt.Fprintf(w, "data Updated\n %#v  ", id)
}
func (re *Repuestos) DeleteRepuestos(w http.ResponseWriter, r *http.Request) {
	re.l.Println("Handle DELETE Repuestos")
}

func IdFromURL(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		return 0, errors.New("not found")
	}
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return 0, errors.New("not id")
	}
	return id, nil
}
