package data

import (
	"encoding/json"
	"io"
	"time"
)

type Repuesto struct {
	ID        int     `json:"id"`
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	Moto      uint    `json:"moto"`
	Price     float64 `json:"price"`
	Anaquel   uint    `json:"anaquel"`
	Seccion   string  `json:"seccion"`
	Caja      uint    `json:"caja"`
	Cantidad  uint    `json:"cantidad"`
	CreatedAt string  `json:"-"`
	UpdatedAt string  `json:"-"`
	DeletedAt string  `json:"-"`
}

var repuestosList = []Repuesto{
	{
		ID:        1,
		Name:      "parrilla",
		Code:      "52100K320001",
		Moto:      1,
		Price:     10.00,
		Anaquel:   1,
		Seccion:   "A-A",
		Caja:      1,
		Cantidad:  4,
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
	{
		ID:        2,
		Name:      "parrilla",
		Code:      "52100K320002",
		Moto:      2,
		Price:     10.00,
		Anaquel:   1,
		Seccion:   "A-A",
		Cantidad:  5,
		Caja:      1,
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
	{
		ID:        3,
		Name:      "parrilla",
		Code:      "52100K320003",
		Moto:      3,
		Price:     10.00,
		Anaquel:   1,
		Seccion:   "A-A",
		Caja:      1,
		Cantidad:  6,
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
	{
		ID:        4,
		Name:      "parrilla",
		Code:      "52100K320004",
		Moto:      4,
		Price:     10.00,
		Anaquel:   1,
		Seccion:   "A-A",
		Cantidad:  2,
		Caja:      1,
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
}

type Repuestos []Repuesto

func (re *Repuestos) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(re)
}

func GetRespuestos() Repuestos {
	return repuestosList
}
