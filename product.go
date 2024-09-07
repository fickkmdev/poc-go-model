package model

import "fmt"

type Product struct {
	Standard
	Unit  Standard `json:"unit"`
	Price float64  `json:"price"`
}

func (p Product) NameWithCode() string {
	return fmt.Sprintf("code: %s, name: %s", p.Code, p.Name)
}
