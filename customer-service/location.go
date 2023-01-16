package main

import (
	"encoding/json"
)

type Location struct {
	Long string `json:"long"`
	Lat  string `json:"lat"`
}

func (l *Location) fromJSON(data []byte) error {
	if err := json.Unmarshal(data, l); err != nil {
		return err
	}

	return nil
}
