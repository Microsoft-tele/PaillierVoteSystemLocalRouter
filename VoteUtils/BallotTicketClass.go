package VoteUtils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type BallotTicket struct {
	ID            string //选票ID
	CandidateNum  int    //参选人数
	NameAndOption map[string][]byte
	RSAPublicKey  []byte // RSA公钥，由投票者写入
	Signature     []byte // 电子签名
}

func (b *BallotTicket) InitBallotTicket(CandidateNameList []string, VoterName string, NameAndOption map[string][]byte, RSAPublicKey []byte, Signature []byte) {
	bigNum, err := rand.Int(rand.Reader, new(big.Int).SetInt64(9999999999))
	if err != nil {
		return
	}
	b.ID = "Voter_" + fmt.Sprintf("%s", bigNum) + "_" + VoterName
	b.CandidateNum = len(CandidateNameList)
	b.NameAndOption = NameAndOption
	b.RSAPublicKey = RSAPublicKey
	b.Signature = Signature
}
