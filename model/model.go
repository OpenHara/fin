package model

type Entry struct {
	Id           string `json:"id,omitempty" db:"id"`
	Description  string `json:"description" db:"description"`
	ValueInCents int64  `json:"valueInCents" db:"value_in_cents"`
	Type         string `json:"type" db:"type"`
}
