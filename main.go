package main

import (
	"fmt"
	"hacktiv8-assesment1/models"
	"hacktiv8-assesment1/params"
	"hacktiv8-assesment1/repositories"
	"os"
	"strconv"
)

func main() {
	var peoples []models.People

	//Check Is os.Args is not null
	if len(os.Args) > 1 {
		//Set os.Args to variable and make sure it set with integer
		RequestNumber, err := strconv.Atoi(os.Args[1])

		if err != nil {
			fmt.Println("Please make sure you set the number not with string!")
			return
		}

		p1 := params.SetNewPeople(1, "Tasrifin", "Back-End Engineer", "Improving skill and find new oppotunity")
		p2 := params.SetNewPeople(2, "Paijo", "Front-End Engineer", "Improving skill and find new oppotunity")
		p3 := params.SetNewPeople(3, "Kusumo", "Android Engineer", "Improving skill and find new oppotunity")

		peoples = append(peoples, repositories.SetPeople(p1))
		peoples = append(peoples, repositories.SetPeople(p2))
		peoples = append(peoples, repositories.SetPeople(p3))

		for _, people := range peoples {
			people.Print(RequestNumber)
		}
	} else {
		fmt.Println("Please set the number!")
	}
}
