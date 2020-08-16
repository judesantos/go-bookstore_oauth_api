package users

type User struct {
	Id        int64
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string
}

type UserLogin struct {
	Email    string
	Password string
}
