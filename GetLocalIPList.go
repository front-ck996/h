package csy

import (
	"fmt"
	"net"
	"strings"
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
