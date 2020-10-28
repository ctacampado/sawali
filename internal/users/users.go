package users

// User contains user accound data
type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
