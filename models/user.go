package models

// type User struct {
// 	Id       int    `json:"-" db:"id"`
// 	Username string `json:"username" db:"username"`
// 	Password string `json:"password" db:"password"`
// }

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
	Coins    int    `json:"coins" db:"coins"`
}
