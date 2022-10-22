package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	message := "liweijun"
	//random, _ := rand.Int(rand.Reader, big.NewInt(1000000000000000000))
	hashed := sha256.New()
	hashed.Write([]byte(message))
	hashcode := hashed.Sum(nil)
	fmt.Println(hashcode)
	for _, v := range hashcode {
		fmt.Printf("%08b", v)
	}
}
