package VoteUtils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Candidate struct { // 候选人
	ID           string
	Name         string
	Introduction string //候选人自我介绍
}

func (c *Candidate) SetCandidateInfo() {
	fmt.Println("开始录入候选人信息：")
	fmt.Printf("+--------------------+----------------+\n")
	fmt.Printf("|     候选人姓名     | 候选人自我介绍 |\n")
	fmt.Printf("+--------------------+----------------+\n")
	fmt.Println("请输入对应信息：")

	fmt.Printf("候选人姓名：")
	var Name string
	for {
		scanf, err := fmt.Scanf("%s", &Name)
		if err != nil {
			fmt.Println("您的输入不合法，请重新输入:", err, scanf)
		} else {
			c.Name = Name
			break
		}
	}
	fmt.Printf("候选人简介：")
	var Introduction string
	for {
		scanf, err := fmt.Scanf("%s", &Introduction)
		if err != nil {
			fmt.Println("您的输入不合法，请重新输入:", err, scanf)
		} else {
			c.Introduction = Introduction
			break
		}
	}

	b, err := rand.Int(rand.Reader, new(big.Int).SetInt64(9999999999))
	if err != nil {
		return
	}
	c.ID = "Candidate_" + fmt.Sprintf("%s", b)
}
