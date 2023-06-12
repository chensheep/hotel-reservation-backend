package types

import (
	"fmt"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	minFirstNameLength = 2
	minLastNameLength  = 2
	minPasswordLength  = 7
)

type CreateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() map[string]string {

	errors := map[string]string{}

	if len(params.FirstName) < minFirstNameLength {
		errors["first_name"] = fmt.Sprintf("first name must be at least %d characters long", minFirstNameLength)
	}

	if len(params.LastName) < minLastNameLength {
		errors["last_name"] = fmt.Sprintf("last name must be at least %d characters long", minLastNameLength)
	}

	if len(params.Password) < minPasswordLength {
		errors["password"] = fmt.Sprintf("password must be at least %d characters long", minPasswordLength)
	}

	//validate email
	if _, err := mail.ParseAddress(params.Email); err != nil {
		errors["email"] = "invalid email address"
	}

	return errors
}

type UpdateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (p UpdateUserParams) ToBSONM() bson.M {
	m := bson.M{}

	if p.FirstName != "" {
		m["first_name"] = p.FirstName
	}

	if p.LastName != "" {
		m["last_name"] = p.LastName
	}

	return m
}

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName      string             `bson:"first_name" json:"first_name"`
	LastName       string             `bson:"last_name" json:"last_name"`
	Email          string             `bson:"email" json:"email"`
	HashedPassword string             `bson:"hashed_password" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:      params.FirstName,
		LastName:       params.LastName,
		Email:          params.Email,
		HashedPassword: string(hashedPassword),
	}, nil
}
