package person_create

type Request struct {
	LastName		string		`json:"lastName"`
	FirstName		string		`json:"firstName"`
	Gender			string		`json:"gender"`
}