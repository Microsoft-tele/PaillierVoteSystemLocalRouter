package VoteUtils

type VoterUseBallotTicket interface { // 投票接口
	ShowBallotTicket()         //显示候选人信息
	ShowInfoByOrder(Order int) // 根据显示候选人信息
	SelectCandidate()
	SetBallotTicket() // 设置选票
}

//func (ballotTicket *BallotTicket) ShowBallotTicket() {
//	fmt.Println("打印选票:")
//	fmt.Printf("+------+----------+--------------------+----------------+\n")
//	fmt.Printf("| 序号 | 候选人ID |     候选人姓名     | 候选人自我介绍 |\n")
//	fmt.Printf("+------+----------+--------------------+----------------+\n")
//	for i, v := range ballotTicket.CandidateList {
//		fmt.Printf("| %4d | %9v| %18v | %v |\n", i, v.ID, v.Name, v.Introduction)
//		fmt.Printf("+------+----------+--------------------+----------------+\n")
//	}
//
//}
//func (ballotTicket *BallotTicket) ShowInfoByOrder(Order int) {
//	fmt.Printf("+------+----------+--------------------+----------------+\n")
//	fmt.Printf("| 序号 | 候选人ID |     候选人姓名     | 候选人自我介绍 |\n")
//	fmt.Printf("+------+----------+--------------------+----------------+\n")
//	fmt.Printf("| %4d | %9v| %18v | %v |\n", Order, ballotTicket.CandidateList[Order].ID, ballotTicket.CandidateList[Order].Name, ballotTicket.CandidateList[Order].Introduction)
//	fmt.Printf("+------+----------+--------------------+----------------+\n")
//}
//func (ballotTicket *BallotTicket) SelectCandidate() {
//	fmt.Println("请输入您的选择，为他投票填入1,否则填0")
//	for i := 0; i < ballotTicket.CandidateNum; i++ {
//		var option int
//		ballotTicket.ShowInfoByOrder(i)
//		for {
//			scanf, err := fmt.Scanf("%d", &option)
//			if err != nil {
//				fmt.Println(scanf)
//				return
//			} else if option == 1 || option == 0 {
//				break
//			} else {
//				fmt.Println("您的输入不合法，请重新输入")
//			}
//		}
//		ballotTicket.Option[i] = option
//	}
//}
