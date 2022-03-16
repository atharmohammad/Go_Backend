package main

type Person struct {
	_id         string   `json:"_id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Designation string   `json:"designation,omitempty"`
	Assignments []string `json:"assignments,omitempty"`
}

type Assignment struct {
	_id    string `json:"_id,omitempty"`
	Title  string `json:"title,omitempty"`
	Tasks  string `json:"tasks,omitempty"`
	Person string `json:"person,omitempty"`
}
