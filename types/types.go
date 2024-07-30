package types

type User struct {
	ID        int    `json: "id"`
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
	Email     string `json: "email"`
	Password  string `json: "password"`
	CreatedAt string `json: "createdAt"`
}

type UserRepository interface {
	// FindByEmail(email string) (*User, error)
	// FindById(id int) (*User, error)
	Create(User) error
}

type RegisterUserPayload struct {
	FirstName string `json: "firstName" validate: "required"`
	LastName  string `json: "lastName" validate: "required"`
	Email     string `json: "email" validate: "required"`
	Password  string `json: "password" validate: "required,min=3,max=130"`
}
