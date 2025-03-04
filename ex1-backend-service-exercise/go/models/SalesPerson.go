package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Sales struct {
	Month  int `json:"month"`
	Amount int `json:"amount"`
}

type SalesPerson struct {
	Name  string  `json:"name"`
	Sales []Sales `json:"sales"`
}

func (s *SalesPerson) FindAll() ([]SalesPerson, error) {
	jsonFile, err := os.Open("models/records.json")
	if err != nil {
		return []SalesPerson{}, fmt.Errorf("error fetching SalesPerson data %s", err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var salesPeople []SalesPerson
	err = json.Unmarshal(byteValue, &salesPeople)
	if err != nil {
		return []SalesPerson{}, fmt.Errorf("error fetching SalesPerson data %s", err)
	}

	return salesPeople, nil
}
