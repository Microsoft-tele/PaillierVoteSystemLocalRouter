package VoteUtils

import (
	"LocalRouter/paillier"
	"encoding/json"
	"fmt"
	"net"
)

var CandidatesCnt int

type BallotOperateMachine struct { // 选票制作机器
	BallotTicketNum int
	BallotTickets   []BallotTicket // 等待收发选票
}
type BallotController interface {
	MakeBallotTickets(candidates []Candidate, PaillierPublicKey paillier.PublicKey) // 制作选票
	DistributeBallots(VotersConn *[]net.Conn)                                       // 分发选票
	//TakeBackBallots()   // 回收选票
	//StatisticResult()   // 统计结果
}

func (b *BallotOperateMachine) MakeBallotTickets(Candidates []Candidate, PaillierPublicKey paillier.PublicKey) {
	// 开始制作选票
	b.BallotTicketNum = len(Candidates)

	fmt.Println("按照当前已经加入的投票人数量制作选票:")
	for i := 0; i < b.BallotTicketNum; i++ {
		//tmpTicket := BallotTicket{}
		//tmpTicket.InitBallotTicket(len(Candidates), Candidates, PaillierPublicKey)
	}
}
func (b *BallotOperateMachine) DistributeBallots(VotersConn *[]net.Conn) {
	fmt.Println(*(VotersConn))
	for i := 0; i < b.BallotTicketNum; i++ {
		fmt.Printf("[%d]:向[ %v ]发送选票:\n", i, (*VotersConn)[i].RemoteAddr())
		write, err := (*VotersConn)[i].Write([]byte("请开始接收选票:"))
		if err != nil {
			fmt.Printf("向投票人[%v]分发选票时错误:[%v]\n", (*VotersConn)[i].RemoteAddr(), write)
			return
		}
		tmpJson, err := json.Marshal(b.BallotTickets[i])
		if err != nil {
			fmt.Println("选票转换为Json失败:", err)
		}
		write, err = (*VotersConn)[i].Write(tmpJson)
		if err != nil {
			fmt.Printf("向[ %v ]发送数据失败:[ %v ][ %v ]\n", (*VotersConn)[i].RemoteAddr(), err, write)
			return
		}
	}
}
