package domain

type User struct {
	Id        int
	Nickname  string
	Age       int
	Gender    int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UserInput struct {
	Nickname  string `json:"nickname"`
	Age       int    `json:"age"`
	Gender    int    `json:"gender"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
