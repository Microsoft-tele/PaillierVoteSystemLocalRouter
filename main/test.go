package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

func main() {
	message := "HelloWorld!"
	message1 := "Hello World!"
	hashed := sha256.Sum256([]byte(message))
	hashed1 := sha256.Sum256([]byte(message1))
	RsaPrivateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Println("生成密钥失败:", err)
		return
	}
	//fmt.Println("RSA Keys:", RsaPrivateKey)
	marshal, err := json.Marshal(RsaPrivateKey)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}
	fmt.Println("Json:", string(marshal))

	v15, err := rsa.SignPKCS1v15(rand.Reader, RsaPrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return
	}

	err = rsa.VerifyPKCS1v15(&RsaPrivateKey.PublicKey, crypto.SHA256, hashed1[:], v15)
	if err != nil {
		fmt.Println("验证签名失败:", err)
		return
	}
}
