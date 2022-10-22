package VoteUtils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type BallotTicket struct {
	ID                string   //选票ID
	CandidateNum      int      //参选人数
	CandidateNameList []string //候选人列表
	Option            [][]byte // 选项
	RSAPublicKey      []byte   // RSA公钥，由投票者写入
	Signature         []byte   // 电子签名
}

func (b *BallotTicket) InitBallotTicket(CandidateNameList []string, VoterName string, Option [][]byte, RSAPublicKey []byte, Signature []byte) {
	bigNum, err := rand.Int(rand.Reader, new(big.Int).SetInt64(9999999999))
	if err != nil {
		return
	}
	b.ID = "Voter_" + fmt.Sprintf("%s", bigNum) + "_" + VoterName
	b.CandidateNum = len(CandidateNameList)
	b.CandidateNameList = CandidateNameList
	b.Option = Option
	b.RSAPublicKey = RSAPublicKey
	b.Signature = Signature
}
