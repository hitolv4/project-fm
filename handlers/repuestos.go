package handlers

import (
	"log"
	"net/http"

	"github.com/hitolv4/fm-api/data"
)

type Repuestos struct {
	l *log.Logger
}

func NewRepuestos(l *log.Logger) *Repuestos {
	return &Repuestos{}
}

func (re *Repuestos) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		re.getRepuestos(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (re *Repuestos) getRepuestos(w http.ResponseWriter, r *http.Request) {
	lp := data.GetRespuestos()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal Json", http.StatusInternalServerError)
	}
}
