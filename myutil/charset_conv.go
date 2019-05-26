// charset_conv.go 用于UTF-8与GBK编码相互转换
package myutil

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func GBKToUTF8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func UTF8ToGBK(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//func main() {
//	s := "GBK 与 UTF-8 编码转换测试"
//	gbk, err := UTF8ToGBK([]byte(s))
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(string(gbk))
//	}

//	ioutil.WriteFile("gbk.txt", gbk, 0644)

//	utf8, err := GBKToUTF8(gbk)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(string(utf8))
//	}

//	ioutil.WriteFile("utf8.txt", utf8, 0644)
//}
