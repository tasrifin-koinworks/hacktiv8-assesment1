package models

import "fmt"

type People struct {
	No     int
	Name   string
	Job    string
	Reason string
}

func (p *People) Print(numb int) {
	if numb == p.No {
		fmt.Println("No \t\t : ", p.No)
		fmt.Println("Name \t\t : ", p.Name)
		fmt.Println("Job Position \t : ", p.Job)
		fmt.Println("Reason \t\t : ", p.Reason)
	}
}
