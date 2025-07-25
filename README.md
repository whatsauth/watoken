# watoken

<img width="1920" height="1043" alt="image" src="https://github.com/user-attachments/assets/60f8236c-3ec3-46f6-bd27-281dfdb04b1f" />

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
   privateKey, publicKey := watoken.GenerateKey()
    
   //generate token for user awangga
   userid := "awangga"
   tokenstring, err := watoken.Encode(userid, privateKey)

    //decode token to get userid
   useridstring := watoken.DecodeGetId(publicKey, tokenstring)
   if useridstring == "" {
       fmt.Println("expire token")
   }

}

```

## Tagging

```sh
git tag v0.1.8
git push origin --tags
go list -m github.com/whatsauth/watoken@v0.1.8
```
