package model

import (
	//"github.com/gin-gonic/gin"
	"time"
	"math/rand"

)

type URL struct {
	ID 			string	`db:"id" json:"id"`
	Original 	string 	`db:"original" json:"original"`
	Short 		string 	`db:"short" json:"short"`
	Visits 		int 	`db:"visits" json:"visits"`
}


func GenerateShortCode(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	code := make([]rune, n)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)
}