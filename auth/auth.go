package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"time"
)

//var sampleSecretKey = []byte("SecretYouShouldHide")

//var key := *rsa.PrivateKey{
//
//}

//func generateJWT() (string, error) {
//	//token := jwt.New(jwt.SigningMethodHS256)
//	token := jwt.New(jwt.SigningMethodRS256)
//	log.Print(token)
//
//	claims := token.Claims.(jwt.MapClaims)
//	claims["exp"] = time.Now().Add(10 * time.Minute)
//	claims["authorized"] = true
//	claims["user"] = "username"
//
//	//tokenString, err := token.SignedString(sampleSecretKey)
//	tokenString, err := token.SignedString(sampleSecretKey)
//	if err != nil {
//		return "", err
//	}
//
//	return tokenString, nil
//}
//
//func verifyJWT2(token string) {
//	type MyCustomClaims struct {
//		Foo string `json:"foo"`
//		jwt.RegisteredClaims
//	}
//	// Parse the token
//	t, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//		// since we only use the one private key to sign the tokens,
//		// we also only use its public counter part to verify
//		return []byte("AllYourBase"), nil
//	})
//	if err != nil {
//		log.Println("cannot parse")
//	}
//	log.Print(t.Valid)
//	log.Printf("%+v\n", t)
//}

//func verifyJWT(token string) {
//	type MyCustomClaims struct {
//		Foo string `json:"foo"`
//		jwt.RegisteredClaims
//	}
//	t, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//		// since we only use the one private key to sign the tokens,
//		// we also only use its public counter part to verify
//		return []byte("AllYourBase"), nil
//	})
//	//if t.Valid {
//	//
//	//}
//	if err != nil {
//		log.Println("cannot parse")
//	}
//	log.Print(t.Valid)
//	log.Print(t)
//}

func ExtractAuthToken(r *http.Request) error {
	bearer := r.Header.Get("authorization")
	//fmt.Println("e1", bearer)
	if len(bearer) < 10 {
		return errors.New("authorization token has a invalid length")
	}
	tokenString := bearer[7:]

	token, err := ParseJWT(tokenString)
	//fmt.Println("e", token, err)
	if err != nil {
		return err
	}

	//fmt.Println("e2", err)
	err = ValidateToken(token)
	if err != nil {
		return err
	}
	return nil
}

func ParseJWT(token string) (*jwt.Token, error) {
	// Parse the token
	t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counterpart to verify
		return []byte{}, nil
	})
	if t != nil {
		return t, nil
	}
	if err != nil {
		log.Println("cannot parse: ", err)
		return nil, err
	}
	return t, nil
}

func ValidateToken(token *jwt.Token) error {
	claims := token.Claims
	auds, err := claims.GetAudience()
	if err != nil {
		return errors.New("unable to parse token audience")
	}

	// check the platform audience
	platformFound := false
	for _, aud := range auds {
		if aud == "" {
			platformFound = true
		}
	}
	if !platformFound {
		return errors.New("platform audience not found")
	}

	// check the platform
	sub, err := claims.GetSubject()
	if err != nil {
		return errors.New("unable to parse token subject")
	}
	if sub != "5d1710b1-6bb4-44b3-bd4b-f9edd50b1c10" {
		return errors.New("subject does not match")
	}

	// check that the token has not expired
	now := time.Now()
	numericDate, err := claims.GetExpirationTime()
	if now.After(numericDate.Local()) {
		return errors.New("token has expired")
	}
	return nil
}
