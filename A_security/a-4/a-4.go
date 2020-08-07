package main

import (
	"fmt"
	"syscall"

	"github.com/tmc/keyring"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	secretValue, err := keyring.Get("proto-keyring-test", "password")
	if err == keyring.ErrNotFound {
		// 未登録だった
		fmt.Printf("Secret Value is not found. Please Type:")
		pw, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			panic(err)
		}
		// 登録
		err = keyring.Set("progo-keyring-test", "password", string(pw))
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		// 未知のエラー
		panic(err)
	} else {
		// 登録済みの値を表示
		fmt.Printf("Secret Value: %s\n", secretValue)
	}
}
