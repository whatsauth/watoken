package watoken

import (
	"encoding/json"
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
)

type Payload struct {
	Id  string    `json:"id"`
	Exp time.Time `json:"exp"`
	Iat time.Time `json:"iat"`
	Nbf time.Time `json:"nbf"`
}

func GenerateKey() (privateKey, publicKey string) {
	secretKey := paseto.NewV4AsymmetricSecretKey() // don't share this!!!
	publicKey = secretKey.Public().ExportHex()     // DO share this one
	privateKey = secretKey.ExportHex()
	return privateKey, publicKey
}

func Encode(id string, privateKey string) (string, error) {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))
	token.SetString("id", id)
	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKey)
	return token.V4Sign(secretKey, nil), err

}

func EncodeforHours(id string, privateKey string, hours int32) (string, error) {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(time.Duration(hours) * time.Hour))
	token.SetString("id", id)
	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKey)
	return token.V4Sign(secretKey, nil), err

}

func EncodeforMinutes(id string, privateKey string, minutes int32) (string, error) {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(time.Duration(minutes) * time.Minute))
	token.SetString("id", id)
	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKey)
	return token.V4Sign(secretKey, nil), err

}

func EncodeforSeconds(id string, privateKey string, seconds int32) (string, error) {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(time.Duration(seconds) * time.Second))
	token.SetString("id", id)
	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKey)
	return token.V4Sign(secretKey, nil), err

}

func Decode(publicKey string, tokenstring string) (payload Payload, err error) {
	var token *paseto.Token
	var pubKey paseto.V4AsymmetricPublicKey
	pubKey, err = paseto.NewV4AsymmetricPublicKeyFromHex(publicKey) // this wil fail if given key in an invalid format
	if err != nil {
		fmt.Println("Deocode NewV4AsymmetricPublicKeyFromHex : ", err)
	}
	parser := paseto.NewParser()                                // only used because this example token has expired, use NewParser() (which checks expiry by default)
	token, err = parser.ParseV4Public(pubKey, tokenstring, nil) // this will fail if parsing failes, cryptographic checks fail, or validation rules fail
	if err != nil {
		fmt.Println("Deocode ParseV4Public : ", err)
	} else {
		json.Unmarshal(token.ClaimsJSON(), &payload)
	}
	return payload, err
}
