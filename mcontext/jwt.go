package mcontext

import (
	"fmt"
	"golang.org/x/crypto/sha3"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"errors"
	"math/rand"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var signedKey string

func SetSignedKey(k string) {
	signedKey = k
}
func GetSignedKey() string {
	return signedKey
}

type UserClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// TODO: please use lib/email to generate token
func GenToken(email string, strClaim jwt.StandardClaims) (string, error) {
	claims := &UserClaim{
		email,
		strClaim,
	}
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(GetSignedKey()))
}

func ParseToken(encryptedToken string) (*UserClaim, error) {
	token, err := jwt.ParseWithClaims(encryptedToken, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		// TODO: Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(GetSignedKey()), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GenPasscode(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func HashPass(pass string, salt string) string {
	// TODO:
	k := []byte(salt)
	buf := []byte(pass)
	// A MAC with 32 bytes of output has 256-bit security strength -- if you use at least a 32-byte-long key.
	h := make([]byte, 32)
	d := sha3.NewShake256()
	// Write the key into the hash.
	d.Write(k)
	// Now write the data.
	d.Write(buf)
	// Read 32 bytes of output from the hash into h.
	d.Read(h)
	return base64.URLEncoding.EncodeToString(h)
}
func ValidatePass(pass string, hashedPass string, salt string) bool {
	encoded := HashPass(pass, salt)
	return encoded == hashedPass
}