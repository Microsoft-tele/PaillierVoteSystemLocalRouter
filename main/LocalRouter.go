package main

import (
	"LocalRouter/CryptoUtils"
	"LocalRouter/VoteUtils"
	"LocalRouter/paillier"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"text/template"
)

var PublicKey paillier.PublicKey

var VoterName string

func main() {
	// 使用restful设计方式
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../css/"))))

	http.HandleFunc("/index", Index)

	http.HandleFunc("/", Index)

	http.HandleFunc("/uploadPaillier", UploadPaillier) // 请求上传文件的界面，并上传本地paillier公钥

	http.HandleFunc("/recvPaillierPubKey", RecvPaillierPubKey) // upload.html 提交的页面，本页面会返回302，并在后台对选票进行paillier加密，再上传至远程服务器

	http.HandleFunc("/recvTicket", RecvTicket) // 接收本地发送的选票并进行加密，再发送至远程服务器
	//http.HandleFunc("/")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println("监听错误:", err)
		return
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("../mod/index.html")
	if err != nil {
		fmt.Println("解析模版失败:", err)
		return
	}
	files.Execute(w, "")
}

func RecvTicket(w http.ResponseWriter, r *http.Request) {
	if PublicKey.N1 == nil {
		fmt.Println("用户还未上传公钥:")
		files, _ := template.ParseFiles("../mod/indexOfErrPaillier.html", "../mod/index.html")
		err := files.Execute(w, "您还未上传Paillier公钥:")
		if err != nil {
			fmt.Println("解析模版失败:", err)
			return
		}
		return
	}

	var CandidateNameList []string // 候选人姓名列表
	//var Option [][]byte            // 加密的选项
	CandidateNameList = make([]string, 0)
	//Option = make([][]byte, 0)
	var NameAndOption map[string][]byte // 候选人姓名和投票选项绑定在一个map中
	NameAndOption = make(map[string][]byte)

	err := r.ParseForm() // 解析post表单
	if err != nil {
		return
	}
	for k, v := range r.PostForm { // 遍历提交的选票表单
		if k == "VoetrName" {
			fmt.Println("监测到投票人姓名", v[0])
			VoterName = v[0]
			continue
		}
		total := 0            // 记录每一个人的被选状态
		for _, j := range v { // 计算每一个人的被选状态
			intJ, _ := strconv.Atoi(j)
			total += intJ
		}
		//CandidateNameList = append(CandidateNameList, k) // 添加候选人姓名
		mOption := new(big.Int).SetInt64(int64(total)) // 加密每一个人的结果
		cOption, err := paillier.Encrypt(&PublicKey, mOption.Bytes())
		if err != nil {
			fmt.Println("加密失败:", err)
		}
		//Option = append(Option, cOption)
		NameAndOption[k] = cOption
		fmt.Printf("[%v : %v]\n", k, total) // 曲线救国，打印调试信息
	}
	fmt.Println("Paillier加密成功:")

	// 将Rsa公钥写入选票 并将rsa私钥存返回给用户
	RsaPriKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Println("生成RSA密钥对失败:", err)
	}
	hashed := sha256.Sum256([]byte(VoterName))

	cipherOfNmae, err := rsa.SignPKCS1v15(rand.Reader, RsaPriKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Println("签名失败:", err)
	}
	fmt.Println("签名成功")

	RsaPubKey, err := json.Marshal(RsaPriKey.PublicKey)
	if err != nil {
		fmt.Println("Rsa公钥转换Json失败:", err)
	}

	//fmt.Println(CandidateNameList, Option)
	Ticket := VoteUtils.BallotTicket{} // 开始生成完整选票
	Ticket.InitBallotTicket(CandidateNameList, VoterName, NameAndOption, RsaPubKey, cipherOfNmae)
	TicketJson, err := json.Marshal(Ticket)
	//fmt.Println("TicketJson:")
	//fmt.Println(TicketJson)
	SendCipherToRemote(TicketJson, "https://404060p9q5.zicp.fun/recvTicket") // 向数据服务器发送选票数据
	SendCipherToRemote(TicketJson, "http://ivxq5w.natappfree.cc/data/new")   // 向区块链服务器发送选票数据

	files, _ := template.ParseFiles("../mod/index.html")
	files.Execute(w, "成功递交选票")
	//w.Header().Set("Location", "http://127.0.0.1:12345/index")
	//w.WriteHeader(302)
}

func UploadPaillier(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("../mod/upload.html")
	if err != nil {
		return
	}
	err = files.Execute(w, "recvPaillierPubKey")
	if err != nil {
		return
	}
}

func RecvPaillierPubKey(w http.ResponseWriter, r *http.Request) {
	fmt.Println("进入Paillier密钥存储")
	err := r.ParseMultipartForm(1024 * 10)
	if err != nil {
		return
	}
	file, m, err := r.FormFile("recvPaillierPubKey")
	if err != nil {
		return
	}
	savePath := "../paillierKey/" + m.Filename
	//fmt.Println(m.Filename)
	openFile, err := os.OpenFile(savePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666) // 这里最后要设置一下暂存密钥文件的地方
	if err != nil {
		fmt.Println("Create fil err:", err)
		return
	}
	n, err := io.Copy(openFile, file)
	if err != nil {
		fmt.Println("拷贝:", n)
		return
	}
	json, err1 := CryptoUtils.GetPubKeyFromJson()

	PublicKey = json
	if err1 != nil {
		return
	}
	//fmt.Println(PublicKey)
	fmt.Println("成功读取PaillierPubKey")
	files, err := template.ParseFiles("../mod/index.html")
	err = files.Execute(w, "成功上传Paillier公钥")
	if err != nil {
		fmt.Println("Execute false:", err)
		return
	}
	//w.Header().Set("Location", "http://127.0.0.1:12345/index")
	//w.WriteHeader(302)
}

func SendCipherToRemote(Ticket []byte, serverUrl string) {
	fmt.Println(string(Ticket))
	resp, err := http.PostForm(serverUrl,
		url.Values{"Ticket": {string(Ticket)}})

	if err != nil {
		fmt.Println("请求失败:", err)
	}
	buf := make([]byte, 1024*10)
	read, err := resp.Body.Read(buf)
	if err != nil {
		fmt.Println("读取数据失败:", err)
		return
	}
	buf = buf[:read]
	fmt.Println("读取到的数据:", buf)
}
