package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Name     string   `json:"employee"`
	Password string   `json:"-"`
	Email    string   `json:"email"`
	Position string   `json:"position"`
	Role     []string `json:"role,omitempty"`
}

func main() {
	fmt.Println("Working with JSON data")
	EncodeJson()
}

func EncodeJson() {
	employees := []employee{
		{"John Malurky", "72bf0!sd2", "johnmarljohn@gmail.com", "Lead Developer", []string{"Admin", "Staff"}},
		{"James Garner", "12bkj91", "garnerjames@acme.dev", "Sr. Engineer", []string{"Admin", "Staff", "Editor"}},
		{"Tess Thompson", "901nc92ek", "tessthompson@acme.dev", "Product Designer", []string{"Staff", "Editor"}},
		{"Harry Pistoli", "@n890dk", "pistoriharry@gmail.com", "Intern Developer", nil},
	}

	finaljson, error := json.MarshalIndent(employees, "", "\t")
	if error != nil {
		panic(error)
	}

	fmt.Printf("%s\n", finaljson)
}
