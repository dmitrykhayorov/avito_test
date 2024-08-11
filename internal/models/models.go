package models

import "time"

type UserRole string
type Status string

const (
	// Possible users roles
	Moderator UserRole = "moderator"
	Client    UserRole = "client"
	// Moderation statuses of flats
	StatusCreated      Status = "created"
	StatusApproved     Status = "approved"
	StatusDeclined     Status = "declined"
	StatusOnModeration Status = "on moderation"
)

type Response500 struct {
	Message   string `json:"message"`
	RequestId string `json:"request_id,omitempty"`
	Code      int32  `json:"code,omitempty"`
}

type AutResponse200 struct {
	Token string `json:"token"`
}

type Flat struct {
	Id        uint32     `json:"id"`
	HouseId   uint32     `json:"house_id"`
	Price     *uint32    `json:"price"`
	Rooms     *uint32    `json:"rooms"`
	Status    Status     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
}

type House struct {
	Id        uint32    `json:"id"`
	Address   string    `json:"address"`
	Year      uint32    `json:"year"`
	Developer string    `json:"developer,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:",omitempty"`
}

type FlatCreateRequestBody struct {
	HouseId uint32 `json:"house_id"`
	Price   uint32 `json:"price"`
	Rooms   uint32 `json:"rooms,omitempty"`
}

type FlatCreateResponse200 struct {
	Flat *Flat
}

type FlatUpdateRequestBody struct {
	FlatId  int    `json:"id"`
	HouseId int    `json:"house_id"`
	Status  Status `json:"status"`
}
