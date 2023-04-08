package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	// 定义命令行选项
	var (
		targetURLStr string
		port         string
	)

	// 设置默认值
	defaultTargetURL := "https://chat.gateway.do/"
	defaultPort := "8080"

	// 解析命令行选项
	flag.StringVar(&targetURLStr, "target-url", defaultTargetURL, "the target URL to forward requests to")
	flag.StringVar(&targetURLStr, "t", defaultTargetURL, "shorthand for --target-url")
	flag.StringVar(&port, "port", defaultPort, "the port number to listen on")
	flag.StringVar(&port, "p", defaultPort, "shorthand for --port")
	flag.Parse()

	// 获取环境变量
	envTargetURL := os.Getenv("TARGET_URL")
	envPort := os.Getenv("PORT")

	// 检查环境变量是否存在
	if envTargetURL != "" {
		targetURLStr = envTargetURL
	}
	if envPort != "" {
		port = envPort
	}

	// 解析目标 URL
	targetURL, err := url.Parse(targetURLStr)
	if err != nil {
		log.Fatal(err)
	}

	// 创建反向代理处理程序
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// 启动 HTTP 服务器并将请求转发到目标 URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 设置请求头中的 Host 字段，以便目标服务器能正确识别请求的主机名
		r.Host = targetURL.Host

		// 将请求转发到目标 URL
		proxy.ServeHTTP(w, r)
	})

	fmt.Printf("Listening on :%s, forwarding requests to %s\n", port, targetURLStr)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
