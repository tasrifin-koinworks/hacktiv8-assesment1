package repositories

import (
	"hacktiv8-assesment1/models"
	"hacktiv8-assesment1/params"
)

func SetPeople(req *params.SetPeople) models.People {
	var people models.People

	people.No = req.No
	people.Name = req.Name
	people.Job = req.Job
	people.Reason = req.Reason

	return people

}
