package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type Robot struct {
	Result  int    `json:"result"`
	Content string `json:"content"`
}

// 获取回复
func getResponce(msg string) (content string, err error) {
	// 请求机器人链接
	robotUrl := "http://api.qingyunke.com/api.php?key=free&appid=0&msg=%s"
	robotUrl = fmt.Sprintf(robotUrl, msg)
	rsp, err := http.Get(robotUrl)
	if err != nil {
		return
	}
	// 读取信息
	body, _ := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	// 转义信息
	var robotRsp Robot
	json.Unmarshal([]byte(body), &robotRsp)
	// 判断是否成功
	if robotRsp.Result != 0 {
		err = errors.New("不知道你在说什么")
		return
	}
	// 返回内容
	content = robotRsp.Content
	return
}

// 查询数组内字段
func in(target string, str_array []string) bool { 
    sort.Strings(str_array) 
    index := sort.SearchStrings(str_array, target) 
    if index < len(str_array) && str_array[index] == target { 
        return true 
    }
    return false 
} 

func main() {
	var input string
	keywords := []string{"拜拜", "再见", "bye"}
	for {
		fmt.Print("请输入你要对我说的话：")
		// 监听输入内容
		fmt.Scanln(&input)
		fmt.Println("机器人酝酿中...")
		cont, err := getResponce(input)
		if err != nil {
			fmt.Println(err)
			break
		}
		// 输出回复
		fmt.Println(cont)
		// 结束关键词
		if in(input, keywords) == true {
			break
		}
	}
}
