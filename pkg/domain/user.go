package domain

type User struct {
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Midname   string    `json:"midname"`
	LastName  string    `json:"last_name"`
	Age       int8      `json:"age"`
	Gender    string    `json:"gender"`
	City      string    `json:"city"`
	Interests []string  `json:"interests"`
}
