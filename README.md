# watoken

Whatsauth Token using PASETO, How to Use :

```go
package controller

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/whatsauth/watoken"
)

func TestWatoken(t *testing.T) {
    // privateKey, publicKey := watoken.GenerateKey()
	publicKey := "25ae7f7b1a959a94f977e6a334e62bc480da66dbf4de01718b252029e93917b8"
	privateKey := "4c90438603a368b45aa70d918416aabb40a14f9477e478fb990402ec1def6b4f25ae7f7b1a959a94f977e6a334e62bc480da66dbf4de01718b252029e93917b8"
    
    //generate token for user awangga
	userid := "awangga"
	signed, err := watoken.Encode(userid, privateKey)

    //decode token to get userid
	require.NoError(t, err)
	body, err := watoken.Decode(publicKey, signed)


	fmt.Println("signed : ", signed)
	fmt.Println("isi : ", body)
	require.NoError(t, err)
}

```

## Tagging

```sh
git tag v0.0.1
git push origin --tags
go list -m github.com/whatsauth/watoken@v0.0.1
```
