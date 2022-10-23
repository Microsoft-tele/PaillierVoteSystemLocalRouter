package Blockchain

import (
	"LocalRouter/VoteUtils"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Header struct {
	TimeStamp string
	RandNum   string
	PreSha    string
}

type Block struct {
	Header Header
	Ticket VoteUtils.BallotTicket
	Tail   Tail
}

type Tail struct {
	PostSha string
}

func (b *Block) InitBlock(isFirstBlock bool, PreSha string, Ticket VoteUtils.BallotTicket) {
	flag := 0
	if isFirstBlock == true {
		flag = 1
	}
	if flag == 1 {
		b.Header.PreSha = "0"
	} else {
		b.Header.PreSha = PreSha
	}
	b.Header.TimeStamp = "2022-10-23 22:30:11.518173 +0800 CST m=+0.001693126"
	b.Header.RandNum = "0" // strconv.Itoa(rand.Intn(999999999999999))
	b.Ticket = Ticket
	b.Tail.PostSha = ""
}

func (b *Block) CalPostSha(NowDigit int) {
	//var InputChan chan string
	InputChan := make(chan string, 1000)
	ExitChan := make(chan bool, 1)
	fmt.Println("开始计算")
	for i := 0; i < 3; i++ {
		go InputRandToChan(InputChan)
	}
	for i := 0; i < 20; i++ {
		go CalSha(*b, InputChan, ExitChan, NowDigit)
	}
	for {
		state, ok := <-ExitChan
		if !ok {
			fmt.Println("!ok")
		}
		if state == true {
			close(InputChan)
			close(ExitChan)
		}
	}
}
func InputRandToChan(InputChan chan string) {
	for {
		//fmt.Println("进入输入")
		rand.Seed(time.Now().UnixNano())
		RandNum := strconv.Itoa(rand.Intn(999999999999999))
		InputChan <- RandNum
		//fmt.Println("进入管道")
	}
}
func CalSha(Block Block, OutputChan chan string, ExitChan chan bool, NowDigit int) {
	//fmt.Println("开始计算:")
	for {
		//Block.Header.RandNum, ok := <-OutputChan
		num, ok := <-OutputChan
		if !ok {
			fmt.Println("通道关闭")
			break
		}
		Block.Header.RandNum = num
		marshal, err := json.Marshal(Block)
		if err != nil {
			fmt.Println("转换为JSON错误:", err)
			return
		}
		hashed := sha256.Sum256(marshal)
		//fmt.Println(hashed)
		hashbin := ""
		for _, v := range hashed {
			sprintf := fmt.Sprintf("%08b", v)
			hashbin += sprintf
		}
		//fmt.Println(hasedbin)
		test := hashbin[:NowDigit]
		fmt.Println(test)
		testnum, _ := strconv.Atoi(test)
		if testnum == 0 {
			fmt.Println("随机数是:", Block.Header.RandNum)
			break
		}
	}
	fmt.Println("计算完毕")
	ExitChan <- true
}

func (b *Block) VerifyRand(rand string) {
	b.Header.RandNum = rand
	marshal, err := json.Marshal(*b)
	if err != nil {
		fmt.Println("转换为JSON错误:", err)
		return
	}
	hashed := sha256.Sum256(marshal)
	hashbin := ""
	for _, v := range hashed {
		sprintf := fmt.Sprintf("%08b", v)
		hashbin += sprintf
	}
	fmt.Println(hashbin)
}
