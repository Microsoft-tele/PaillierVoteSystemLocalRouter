package VoteUtils

import (
	"LocalRouter/paillier"
	"crypto/rand"
	"fmt"
	"math/big"
)

type Voter struct { // 投票人
	ID                    string
	Name                  string
	SelfBallotTicket      BallotTicket // 每个选民自己的票
	RSAPublicKeyFilePath  string
	RSAPrivateKeyFilePath string             // RSA的私钥,用于对选票进行签名，RSA公钥在候选人注册时生成唯一的密钥对，公钥由第三方进行保存，以备后续验证使用
	PaillierPublicKey     paillier.PublicKey // paillier的公钥，由公证第三方发布
}

func (v *Voter) InitVoter(name string, RSAPrivateKeyFilePath string, RSAPublicKeyFilePath string) {
	b, err := rand.Int(rand.Reader, new(big.Int).SetInt64(9999999999))
	if err != nil {
		return
	}
	v.ID = "Voter_" + fmt.Sprintf("%s", b)
	v.Name = name
	v.RSAPrivateKeyFilePath = RSAPrivateKeyFilePath
}
