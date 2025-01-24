package domain

type User struct {
	id       int64
	username string
}

func (u User) ID() int64 {
	return u.id
}

func (u User) Username() string {
	return u.username
}

func NewUser(id int64, username string) (User, error) {
	return User{id: id, username: username}, nil
}
