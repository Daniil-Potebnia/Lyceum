package models

type Position struct {
	ID       int32 `json:"id" db:"id"`
	Item     string `json:"item" db:"item"`
	Quantity int32  `json:"quantity" db:"quantity"`
}
