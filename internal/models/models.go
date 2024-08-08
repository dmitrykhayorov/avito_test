package models

type UserRole string
type Status string

const (
	Moderator    UserRole = "moderator"
	Client       UserRole = "client"
	Created      Status   = "created"
	Approved     Status   = "approved"
	Declined     Status   = "declined"
	OnModeration Status   = "on moderation"
)

type Response500 struct {
	// Описание ошибки
	Message string `json:"message"`
	// Идентификатор запроса. Предназначен для более быстрого поиска проблем.
	RequestId string `json:"request_id,omitempty"`
	// Код ошибки. Предназначен для классификации проблем и более быстрого решения проблем.
	Code int32 `json:"code,omitempty"`
}

type AutResponse200 struct {
	Token string `json:"token"`
}

type Flat struct {
	// Number of a flat
	Number uint32 `json:"id"`
	// HouseId id of a house
	HouseId uint32 `json:"house_id"`
	// Price is a price of a flat
	Price uint32 `json:"price"`
	// Rooms is number of rooms in a flat
	Rooms uint32 `json:"rooms"`
	// Status of a flat
	Status Status `json:"status"`
}

type FlatCreateResponse200 struct {
	Flat *Flat
}
