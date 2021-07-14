package models

// Permission have id and name , it is define the Permission of role
type Permission struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
