package main

import (
	"encoding/json"
)

type Location struct {
	Long string `json:"long"`
	Lat  string `json:"lat"`
}

func (l Location) toJSON() ([]byte, error) {
	b, err := json.Marshal(l)
	if err != nil {
		return b, nil
	}

	return b, nil
}
