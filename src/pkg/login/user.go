package login

// User defines the properties of a user to be added
type User struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Timestamp string `json:"timestamp"`
}
