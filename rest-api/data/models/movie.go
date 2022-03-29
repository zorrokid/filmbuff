package models

import (
	"encoding/json"
	"io"

	"gorm.io/gorm"
)

type Movies []*Movie

func (m *Movies) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

type Movie struct {
	gorm.Model
	Name string `json:"name"`
}

func (m *Movie) FromJson(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(m)
}
