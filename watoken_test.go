package watoken

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWatoken(t *testing.T) {
	privateKey, publicKey := GenerateKey()
	fmt.Println("privateKey : ", privateKey)
	fmt.Println("publicKey : ", publicKey)
	userid := "awangga"

	signed, err := Encode(userid, privateKey)
	require.NoError(t, err)
	body, err := Decode(publicKey, signed)
	fmt.Println("signed : ", signed)
	fmt.Println("isi : ", body)
	require.NoError(t, err)
}

func TestWacipher(t *testing.T) {
	n := 100
	rnd := RandomString(n)
	require.Len(t, rnd, n)
	fmt.Println("rnd : ", rnd)
}
