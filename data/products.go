package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Repuesto struct {
	ID        uint    `json:"id"`
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

func (re *Repuesto) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(re)

}

func GetRespuestos() Repuestos {
	return repuestosList
}

func AddRepuestos(r *Repuesto) {
	r.ID = GetlastID()
	repuestosList = append(repuestosList, *r)

}

func UpdateRepuesto(id uint, r Repuesto) error {
	_, pos, err := findrepuesto(id)
	if err != nil {
		return err
	}
	r.ID = id

	if r.Name != "" {
		repuestosList[pos].Name = r.Name
	}
	if r.Code != "" {
		repuestosList[pos].Code = r.Code
	}
	if r.Moto != 0 {
		repuestosList[pos].Moto = r.Moto
	}
	if r.Price != 0.00 {
		repuestosList[pos].Price = r.Price
	}
	if r.Anaquel != 0 {
		repuestosList[pos].Anaquel = r.Anaquel
	}
	if r.Seccion != "" {
		repuestosList[pos].Seccion = r.Seccion
	}
	if r.Caja != 0 {
		repuestosList[pos].Caja = r.Caja
	}
	if r.Cantidad != 0 {
		repuestosList[pos].Cantidad = r.Cantidad
	}

	repuestosList[pos].UpdatedAt = time.Now().UTC().String()

	return nil
}

var ErrProductNotFound = fmt.Errorf("Repuesto not found")

func findrepuesto(id uint) (*Repuesto, int, error) {
	for i, r := range repuestosList {
		if r.ID == id {
			return &r, i, nil
		}
	}

	return nil, 0, ErrProductNotFound
}

func GetlastID() uint {
	lastR := repuestosList[len(repuestosList)-1]
	return lastR.ID + 1
}
