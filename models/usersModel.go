package databaseConnection

// Users Table struct
type UsersTable struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	SessionID   string `json:"sessionid"`
}
