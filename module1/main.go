package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/healthz", handleHealthy)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 将请求中的 header 写入 response header
	for key, values := range r.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// 读取环境变量中的 VERSION 配置并写入 response header
	version := os.Getenv("VERSION")
	if version != "" {
		w.Header().Add("VERSION", version)
	}

	// 记录访问日志
	clientIP := r.RemoteAddr
	httpCode := http.StatusOK // 默认为 200
	fmt.Printf("Client IP: %s | HTTP Code: %d\n", clientIP, httpCode)

	// 返回响应
	w.WriteHeader(httpCode)
	_, err := fmt.Fprintln(w, "Hello, this is the response.")
	if err != nil {
		return
	}
}

func handleHealthy(w http.ResponseWriter, r *http.Request) {
	// 将请求中的 header 写入 response header
	for key, values := range r.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	// 返回 200 响应
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintln(w, "OK")
	if err != nil {
		return
	}
}

//测试推送git
