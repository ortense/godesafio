package model

type User struct {
	Name     string  `json:"nome"`
	Email    string  `json:"email"`
	Password string  `json:"senha"`
	Phones   []Phone `json:"telefones"`
}
