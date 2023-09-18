package user_auth_util

import (
	"crypto/rand"
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func EncryptPass(password string) ([]byte,[]byte,error){
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
			return nil,nil,err
	}
	passwordBytes := []byte(password)
	iterations := 10000 
	hashBytes := pbkdf2.Key(passwordBytes, salt, iterations, 32, sha256.New)
	
	return hashBytes,salt,nil
}