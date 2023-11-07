package csy

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

func GetLocalIPList() (map[string]string, error) {
	interfaceList, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var byName *net.Interface
	var addrList []net.Addr
	var oneAddrs []string
	ipList := make(map[string]string, len(interfaceList))
	for i := range interfaceList {
		byName, err = net.InterfaceByName(interfaceList[i].Name)
		if err != nil {
			return nil, err
		}
		addrList, err = byName.Addrs()
		if err != nil {
			return nil, err
		}
		for ii := range addrList {
			oneAddrs = strings.SplitN(addrList[ii].String(), "/", 2)
			ipList[interfaceList[i].Name] = oneAddrs[0]
		}
	}
	return ipList, nil
}

func GetLocalIPv4List() ([]string, error) {
	var ipList []string

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		// 排除回环接口和无效接口
		if iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagUp != 0 {
			addrs, err := iface.Addrs()
			if err != nil {
				return nil, err
			}

			for _, addr := range addrs {
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					ipList = append(ipList, ipNet.IP.String())
				}
			}
		}
	}

	if len(ipList) == 0 {
		return nil, fmt.Errorf("无法获取本机IPv4地址")
	}

	return ipList, nil
}
func GetPublicIP() (string, error) {
	var externalIPServices = []string{
		"http://ipinfo.io/ip",
		"http://icanhazip.com",
		"http://ip.42.pl/raw",
		"http://myexternalip.com/raw",
		"http://ipecho.net/plain",
		"http://ident.me",
		// 添加其他获取IP的服务地址
	}
	rand.Seed(time.Now().UnixNano())
	selectedService := externalIPServices[rand.Intn(len(externalIPServices))]

	resp, err := http.Get(selectedService)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(body)), nil
}
