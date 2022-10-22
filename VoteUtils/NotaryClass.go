package VoteUtils

import (
	"LocalRouter/CryptoUtils"
	"LocalRouter/paillier"
	"crypto/rand"
	"fmt"
	"math/big"
	"net"
	"os"
)

type Notary struct { // 公证人
	ID                 string
	Name               string
	PaillierPublicKey  paillier.PublicKey
	PaillierPrivatekey *paillier.PrivateKey
	RsaPublicKey       os.File
}

func (n *Notary) InitNotary() {
	b, err := rand.Int(rand.Reader, new(big.Int).SetInt64(9999999999))
	if err != nil {
		return
	}
	n.ID = "Notary_" + fmt.Sprintf("%s", b)
	for { // 加载公证人姓名
		var Name string
		fmt.Println("请输入公证人的名字：")
		scanf, err := fmt.Scanf("%s", &Name)
		if err != nil {
			fmt.Println("您的输入不合法，请重新输入", scanf, err)
		} else {
			n.Name = Name
			break
		}
	}

	for { // 加载Pailler密钥对
		var op string
		fmt.Println("是否生成新的 Paillier 密钥 (Y/n)")
		scanf, err := fmt.Scanf("%s", &op)
		if err != nil {
			fmt.Println("您的输入不合法，请重新输入", scanf, err)
			return
		} else if op == "Y" || op == "y" {
			tmpPrivateKey := CryptoUtils.CreateKeys(1024)
			n.PaillierPrivatekey = tmpPrivateKey
			n.PaillierPublicKey = n.PaillierPrivatekey.PublicKey
			break
		} else if op == "n" || op == "N" {
			n.PaillierPrivatekey = CryptoUtils.GetKeysFromJson()
			n.PaillierPublicKey = n.PaillierPrivatekey.PublicKey
			break
		} else {
			fmt.Println("您的输入不合法，请重新输入", scanf)
		}
	}
}

func (n *Notary) Work(I BallotController, Candidates []Candidate, VotersConn *[]net.Conn) {
	I.MakeBallotTickets(Candidates, n.PaillierPublicKey)
	I.DistributeBallots(VotersConn)
}
