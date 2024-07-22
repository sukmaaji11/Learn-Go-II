package models

// User Type Model
type User struct {
	ID int `json: "id"`
	Nama string `json: "nama"`
}

// Users Type Model (Struktur >1 User)
type Users []User