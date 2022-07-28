package people

type Getter interface {
	GetAll() []Person
}

type Adder interface {
	Add(person Person)
}

type Person struct {
	Id        int    `json:"id"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Gender    string `json:"gender"`
}

type Repo struct {
	People []Person
}

func Init() *Repo {
	return &Repo{
		People: []Person{},
	}
}

func (repo *Repo) Add(person Person) {
	person.Id = len(repo.People) + 1
	repo.People = append(repo.People, person)
}

func (repo *Repo) Find(id int) Person {
	for index := range repo.People {
		if repo.People[index].Id == id {
			return repo.People[index]
		}
	}
	return Person{}
}

func (repo *Repo) GetAll() []Person {
	return repo.People
}
