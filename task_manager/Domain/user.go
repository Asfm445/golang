package domain

type User struct {
	ID       string `bson:"_id,omitempty"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Role     string `bson:"role"`
}

type UserRepository interface {
	Register(user User) error
	Login(email, password string) (string, error)
	Promote(email string) error
	FindByEmail(email string) (User, error)
}
