package domain

type User struct {
	id        int64
	username  string
	firstName string
	lastName  string
	email     string
	password  string
	phone     string
}

func (u User) ID() int64 {
	return u.id
}

func (u User) Username() string {
	return u.username
}

func (u User) FirstName() string {
	return u.firstName
}

func (u User) LastName() string {
	return u.lastName
}

func (u User) Email() string {
	return u.email
}

func (u User) Password() string {
	return u.password
}

func (u User) Phone() string {
	return u.phone
}

func NewUser(id int64, firstName, lastName, email, password string) (User, error) {
	return User{id: id, firstName: firstName, lastName: lastName, email: email, password: password}, nil
}
