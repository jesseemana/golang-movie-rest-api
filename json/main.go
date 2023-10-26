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
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	employees := []employee{
		{"John Malurky", "72bf0!sd2", "johnmarljohn@gmail.com", "Lead Developer", []string{"Admin", "Staff"}},
		{"James Garner", "12bkj91", "garnerjames@acme.dev", "Sr. Engineer", []string{"Admin", "Staff", "Editor"}},
		{"Tess Thompson", "901nc92ek", "tessthompson@acme.dev", "Product Designer", []string{"Staff", "Editor"}},
		{"Harry Pistoli", "@n890dk", "pistoriharry@gmail.com", "Intern Developer", nil},
	}

	finaljson, err := json.MarshalIndent(employees, "", "\t")
	if err != nil {
		panic(err)
	}	

	fmt.Printf("%s\n", finaljson)
}

func DecodeJson() {
	json_data := []byte(`
		{
            "employee": "James Garner",
            "email": "garnerjames@acme.dev",
            "position": "Sr. Engineer",
            "role": ["Admin","Staff","Editor"]
        }
	`)

	var employee_data employee

	if json.Valid(json_data) {
		fmt.Println("Valid JSON received")
		json.Unmarshal(json_data, &employee_data)
		fmt.Printf("%#v\n", employee_data)
	} else {
		fmt.Println("JSON WAS NOT VALID!")
	}

	// STORE AS KEY/VALUE PAIRS 
	var webdata map[string]interface{}
	json.Unmarshal(json_data, &webdata)
	fmt.Printf("%#v\n", webdata)
}	