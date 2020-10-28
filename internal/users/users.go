package users

// User contains user accound data
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	// salted hash of password
	Hash string `json:"hash,omitempty"`
}
