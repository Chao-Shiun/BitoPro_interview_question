package model

type Gender int

const (
	Male Gender = iota
	Female
)

type SinglePerson struct {
	Name           string `json:"name"`
	Height         int    `json:"height"`
	GenderType     Gender `json:"gender"`
	RemainingDates int    `json:"remainingDates"`
}
