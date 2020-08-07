package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"golang.org/x/crypto/ssh"
)

var hostKeyString string = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGn0x52p5aoN2MmmqPnujiDf0SBhKzYtEoGCgkCAT4wDjgh4ODph4SBKU9XuApf0Mlj6rp2d02pZG4Uhas4hCDk= takumi@X1C6"

func main() {
	// 秘密鍵の準備
	key, err := ioutil.ReadFile("id_sysprogo")
	if err != nil {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}

	// サーバーの鍵の準備
	hostKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(hostKeyString))
	if err != nil {
		panic(err)
	}

	// 接続設定
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	// 通信開始
	conn, err := ssh.Dial("tcp", "localhost:1222", config)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// コマンドを実行して出力結果を取得
	output, err := session.CombinedOutput("ssh -V")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
