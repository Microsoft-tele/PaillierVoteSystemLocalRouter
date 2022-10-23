package CryptoUtils

import (
	"LocalRouter/FileUtils"
	"LocalRouter/ShellUtils"
	"LocalRouter/paillier"
	"encoding/json"
	"fmt"
)

func GetKeysFromJson() (key *paillier.PrivateKey) {
	home := ShellUtils.GetOutFromStdout("echo $HOME")[0]
	dirList := ShellUtils.GetOutFromStdout("ls " + "/" + home + "/paillier/keys")
	for i, v := range dirList {
		fmt.Printf("[%d : %v]\n", i, v)
	}
	choice := 0 //获取密钥文件
	fmt.Println("请输入您本次投票使用的密钥：")
	scanf, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println("您的输入有误:", err, scanf)
		return nil
	}
	fmt.Println("Choice:", choice)

	filepath := home + "/paillier/keys/" + dirList[choice]
	fmt.Println("filename:", filepath)

	var PrivateKey *paillier.PrivateKey

	PrivateKeysSlice := FileUtils.ReadFileContent(filepath)

	err = json.Unmarshal([]byte(PrivateKeysSlice[0]), &PrivateKey)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
	fmt.Println("成功从文件中恢复Paillier密钥对")
	return PrivateKey
}

// GetPubKeyFromJson 固定路径获取公钥
func GetPubKeyFromJson() (paillier.PublicKey, error) {
	filepath := "../paillierKey/key"
	var PubKey paillier.PublicKey
	PubKeysSlice := FileUtils.ReadFileContent(filepath)
	err := json.Unmarshal([]byte(PubKeysSlice[0]), &PubKey)
	if err != nil {
		fmt.Println("反射公钥失败:", err)
		if err != nil {
			fmt.Println("解析模版文件失败:", err)
			return PubKey, err
		}
		return PubKey, err
	}
	return PubKey, nil
}
