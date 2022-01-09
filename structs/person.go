package structs

type Person struct {
	Increment
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Timestamps
}
