package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/fernet/fernet-go"
)

func main() {
	cmd := "dd if=/dev/urandom bs=32 count=1 2>/dev/null | openssl base64"
	key, err := exec.Command("sh", "-c", cmd).Output()

	k := fernet.MustDecodeKeys(string(key))
	tok, err := fernet.EncryptAndSign([]byte("hello"), k[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(tok))

	// time.Sleep(5 * time.Second)

	msg := fernet.VerifyAndDecrypt(tok, 3*time.Second, k)
	fmt.Println(string(msg))
}
