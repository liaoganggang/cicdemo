package main

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"time"
)

type RespController struct {
	Time     string `json:"time"`
	Version  string `json:"version"`
	IPAdd    string `json:"ip address"`
	Hostname string `json:"hostname"`
}

func NewRespController(time, version, ipadd, hostname string) *RespController {
	return &RespController{
		Time:     time,
		Version:  version,
		IPAdd:    ipadd,
		Hostname: hostname,
	}
}

func GetIp() string {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			ip, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			if ip.IP.IsLoopback() {
				continue
			}
			if !ip.IP.IsGlobalUnicast() {
				continue
			}
			return ip.IP.String()
		}
	}
	return ""
}

func main() {
	addr := ":8080"
	version := "v1.0"
	ipadd := GetIp()
	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time := time.Now().Format("2006-01-02 15:04:05")
		respcontroller := NewRespController(time, version, ipadd, hostname)
		ctx, _ := json.Marshal(respcontroller)

		var buffer bytes.Buffer
		json.Indent(&buffer, ctx, "", " ")

		buffer.WriteTo(w)
	})

	http.ListenAndServe(addr, nil)
}
