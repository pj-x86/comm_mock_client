// client.go 用于验证comm_mock_svr 程序的正确性
package main

import (
	"bufio"
	"comm_mock_client/myutil"
	"flag"
	"fmt"
	"io"

	//"io"
	"log"
	"net"
	"os"
)

var charset = flag.String("charset", "UTF8", "message's charset, GBK|UTF8")

//!+
func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", "localhost:6610")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//go mustCopy(os.Stdout, conn)
	//独立go线程处理返回报文
	go handleResp(conn)

	//mustCopy(conn, os.Stdin)
	//处理请求报文
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		var reqmsg []byte

		reqmsg = input.Bytes() //读入的是UTF8
		if *charset == "GBK" {
			gbkmsg, err := myutil.UTF8ToGBK(reqmsg)
			if err != nil {
				fmt.Printf("转成GBK编码失败\n")
				continue
			}
			fmt.Printf("转成GBK编码再发送\n")
			reqmsg = gbkmsg
		}

		//计算报文长度
		lenstr := fmt.Sprintf("%04d", len(reqmsg)) //基于最后要发送的报文内容+MD5计算
		fmt.Printf("发送报文字节长度: [%s]", lenstr)

		reqmsg1 := []byte(lenstr)
		reqmsg1 = append(reqmsg1, reqmsg...)
		conn.Write(reqmsg1)
		fmt.Printf("发送成功\n")
	}

}

func handleResp(conn net.Conn) error {

	for {
		var recvdata []byte
		recvdata = make([]byte, 8192)
		count, err := conn.Read(recvdata)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Read data from socket error. err=%s", err.Error())
				os.Exit(-1)
			} else {
				fmt.Printf("数据接收完毕[%s]\n", err.Error())
				break
			}
		}
		var reqmsg string
		if *charset == "GBK" {
			gbkmsg := recvdata[:count]
			utf8msg, err := myutil.GBKToUTF8(gbkmsg)
			if err != nil {
				fmt.Printf("转为UTF8编码失败\n")
				continue
			}
			fmt.Printf("原始报文编码为GBK，转为UTF8\n")
			reqmsg = string(utf8msg)
		} else { //UTF-8
			reqmsg = string(recvdata[:count])
		}

		fmt.Printf("返回报文: [%s]\n", reqmsg)
	}

	return nil
}

//!-

//func mustCopy(dst io.Writer, src io.Reader) {
//	if _, err := io.Copy(dst, src); err != nil {
//		log.Fatal(err)
//	}
//}
