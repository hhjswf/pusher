package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Person struct {
	Name string
}

func (p Person) SayHi() {
	fmt.Println("Hi, my name is", p.Name)
}

func spider() {
	// 设置搜索关键字
	query := "嚞猫闪充客服电话"

	// 构造搜索 URL
	searchUrl := "http://www.baidu.com/s?wd=" + url.QueryEscape(query)

	// 创建 HTTP 客户端
	client := &http.Client{}

	// 构造 HTTP 请求
	req, err := http.NewRequest("GET", searchUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 发送 HTTP 请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 将响应内容转换为字符串
	bodyStr := string(body)

	// 输出搜索结果
	fmt.Println(bodyStr)
}
func main() {

	for i := 0; i < 200; i++ {
		go spider()
	}
	select {}
}

func deferMe() (x int) {
	defer func() { x *= 2 }()

	x = 1

	return
}

func intSeq() func() func() {
	return func() func() {

		return func() {

		}
	}
}
