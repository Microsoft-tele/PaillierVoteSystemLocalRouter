package CryptoUtils

import (
	"LocalRouter/paillier"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
)

func CreateKeys(bitNum int) (PrivateKey *paillier.PrivateKey) {
	seed := rand.Reader
	PrivateKey, err := paillier.GenerateKey(seed, bitNum)

	if err != nil {
		fmt.Println("Create keys err:", err)
	}
	pubmarshal, err1 := json.Marshal(PrivateKey.PublicKey)
	primarshal, err2 := json.Marshal(PrivateKey)
	if err1 != nil {
		fmt.Println("Keys to json err:", err1)
		return
	}

	fmt.Println("This is 公证人生成的 secret keys, don't explore them!")
	//NowTime := strings.Split(time.Now().String(), " ")[:2]
	//home := ShellUtils.GetOutFromStdout("echo $HOME")[0]
	pubfilename := "../paillierKeys/pub/" + "key" + ".json" // 需要改进
	prifilename := "../paillierKeys/pri/" + "key" + ".json" // 需要改进

	pubfile, err2 := os.OpenFile(pubfilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err2 != nil {
		fmt.Println("Open file err:", err2)
		return
	}
	defer func(file *os.File) {
		_, err := file.WriteString("\n")
		if err != nil {
			fmt.Println("Write \\n err:", err)
			return
		}
		err3 := file.Close()
		if err3 != nil {
			fmt.Println("Close file err:", err3)
		}
	}(pubfile)
	_, err4 := pubfile.Write(pubmarshal)
	if err4 != nil {
		fmt.Println("Write marshal err:", err4)
		return
	}
	prifile, err3 := os.OpenFile(prifilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err3 != nil {
		fmt.Println("Open file err:", err2)
		return
	}
	defer func(file *os.File) {
		_, err := file.WriteString("\n")
		if err != nil {
			fmt.Println("Write \\n err:", err)
			return
		}
		err3 := file.Close()
		if err3 != nil {
			fmt.Println("Close file err:", err3)
		}
	}(prifile)
	_, err5 := prifile.Write(primarshal)
	if err5 != nil {
		fmt.Println("Write marshal err:", err4)
		return
	}
	return
}
