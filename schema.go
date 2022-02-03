package main

type Person struct {
	_id         string `json:"_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Designation string `json:"designation,omitempty"`
}
