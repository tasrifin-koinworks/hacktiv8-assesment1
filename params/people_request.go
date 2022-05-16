package params

type SetPeople struct {
	No     int
	Name   string
	Job    string
	Reason string
}

func SetNewPeople(no int, name string, job string, reason string) *SetPeople {
	return &SetPeople{
		No:     no,
		Name:   name,
		Job:    job,
		Reason: reason,
	}
}
