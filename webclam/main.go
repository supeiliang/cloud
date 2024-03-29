// 导入需要的包
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Spider 定义一个爬虫工具的结构体
type Spider struct {
	// 要爬取的网站的URL
	URL string
	// 保存爬取到的数据
	Data []byte
}

// Crawl 定义一个爬取网页数据的方法
func (s *Spider) Crawl() {
	// 使用http包发送GET请求
	resp, err := http.Get(s.URL)
	if err != nil {
		log.Fatal(err)
	}
	// 关闭响应体
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err)
		}
	}(resp.Body)
	// 读取响应体数据
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 将数据赋值给Spider结构体的Data字段
	s.Data = data
}

// Print 定义一个打印数据的方法
func (s *Spider) Print() {
	// 将数据转换为字符串并打印
	fmt.Println(string(s.Data))
}

// 主函数
func main() {
	// 创建一个Spider结构体实例，指定要爬取的网站URL
	s := &Spider{
		URL: "https://www.bilibili.com/",
	}
	// 调用Crawl方法爬取数据
	s.Crawl()
	// 调用Print方法打印数据
	s.Print()
}
