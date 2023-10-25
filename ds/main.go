package main

import (
	"fmt"
	"sort"
)

func main() {
	maps()

	jesse := User{
		"Jesse Emana",
		"Man United till I die",
		"Liwonde, Machinga",
		"0881815626",
		"jesseemana@gmail.com",
		true,
	}
	
	fmt.Printf("Users' details are: %+v\n", jesse)
	jesse.getStatus()
}

type User struct {
	Full_name string
	Bio       string
	Location  string
	Contact   string
	Email     string
	Active    bool
}

func (u User) getStatus() {
	fmt.Println("Is user logged in?", u.Active)
}

func arrays() {
	var my_fpl = [11]string{
		"Becker", "Walker", "Saliba", "Andersen",
		"MoSalah", "Son", "Saka", "Gordon/Diaz",
		"Mbeumo", "Watkins", "Alvarez",
	}

	fmt.Println("My starting team:", my_fpl)
}

func slices() {
	var my_team = []string{
		"Becker", "Walker", "Saliba",
		"Andersen", "MoSalah", "Son", "Saka",
		"Gordon/Diaz", "Mbeumo", "Watkins", "Alvarez",
	}

	my_team = append(my_team, "Areola", "Aurier", "Udogie")
	fmt.Println("My full fpl team:", my_team)

	// slicing > [index:] [:index] [index:index]
	my_team = append(my_team[:11])
	fmt.Println("My starting team:", my_team)

	// Allocating new memory with make()
	high_scores := make([]int, 4)
	high_scores[0] = 321
	high_scores[1] = 213
	high_scores[2] = 194
	high_scores[3] = 487

	high_scores = append(high_scores, 3802, 982, 142)

	sort.Ints(high_scores)
	fmt.Println(high_scores)
}

func maps() {
	languages := make(map[string]string)

	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	languages["JS"] = "JavaScript"
	languages["TS"] = "TypeScript"

	fmt.Println("My languages:", languages)
	fmt.Println("JS stands for:", languages["JS"])

	delete(languages, "PY")
	fmt.Println("Updated languages:", languages)

	for index, value := range languages {
		fmt.Printf("%v: %v\n", index, value)
	}

}
