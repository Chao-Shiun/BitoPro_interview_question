package model

type Gender string

const (
	Male   Gender = "male"
	Female        = "female"
)

func (s Gender) IsValid() bool {
	if s == Male || s == Female {
		return true
	}
	return false
}

type SinglePerson struct {
	Name           string `json:"name"`
	Height         int    `json:"height"`
	GenderType     Gender `json:"gender"`
	RemainingDates int    `json:"remainingDates"`
}
