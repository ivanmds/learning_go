package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Birthdate string
type AddressType int
type ContactType int

const (
	AddressHome AddressType = iota
	AddressWork
)

const (
	ContactPhone ContactType = iota
	ContactEmail
)

type Address struct {
	Street       string      `json:"street"`
	Number       string      `json:"number"`
	Neighborhood string      `json:"neighborhood"`
	City         string      `json:"city"`
	State        string      `json:"state"`
	Country      string      `json:"country"`
	Complement   string      `json:"complement"`
	Type         AddressType `json:"type"`
}

type Contact struct {
	Value string      `json:"value"`
	Type  ContactType `json:"type"`
}

type Customer struct {
	ID         primitive.ObjectID `bson:"_id" json:"id",omitempty`
	Name       string             `json:"name"`
	Birthdate  Birthdate          `json:"birthdate"`
	MotherName string             `json:"motherName"`
	Contacts   []Contact          `json:"contacts"`
	Addresses  []Address          `json:"addresses"`
	Active     bool               `bson:"active" json:"-"`
}
