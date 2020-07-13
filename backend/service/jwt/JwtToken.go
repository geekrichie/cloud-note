package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

type Claims struct {
	UserId int `json:"userid"`
	jwt.StandardClaims
}

func CreateToken(userId int) (string, error){
	expirationTime := time.Now().Add(30 * time.Minute);

	claims :=&Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			//the expiry time is expressed
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
func ParseToken(tokenString string) (int, error){
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(interface{}, error) {
		return jwtKey,nil
	})
	if err != nil {
		return -1, err
	}
	if !token.Valid {
		return -1, err
	}
	return claims.UserId,nil
}