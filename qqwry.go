package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"
	"qqwry"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	datFile := flag.String("d", "./qqwry.dat", "specify the path to the qqwry.dat")
	port := flag.String("p", "8888", "server listening port")
	flag.Parse()

	qqwry.IPData.FilePath = *datFile
	startTime := time.Now()
	res := qqwry.IPData.InitIPData()

	if v, ok := res.(error); ok {
		log.Panic(v)
	}
	log.Printf("ip data load success, count: %d, elapsed time:%s\n", qqwry.IPData.IPNum, time.Since(startTime))

	// 下面开始加载 http 相关的服务
	http.HandleFunc("/", findIP)

	log.Printf("server listening at:%s", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", *port), nil); err != nil {
		log.Println(err)
	}
}

// findIP 查找 IP 地址的接口
func findIP(w http.ResponseWriter, r *http.Request) {
	res := qqwry.NewResponse(w, r)
	ip := r.Form.Get("ip")
	if ip == "" {
		res.ReturnError(http.StatusBadRequest, 400, "query example: http://127.0.0.1:8888?ip=8.8.8.8,114.114.114.114")
		return
	}

	ips := strings.Split(ip, ",")
	qqWry := qqwry.NewQQwry()
	rs := map[string]qqwry.ResultQQwry{}
	if len(ips) > 0 {
		for _, ip := range ips {
			rs[ip] = qqWry.Find(ip)
		}
	}
	res.ReturnSuccess(rs)
}
