package main

import (
	"fmt"
	"net"
	"sort"
)

type Addrs []*net.IP

func (addrs Addrs) Len() int {
	return len(addrs)
}

func (addrs Addrs) Less(i, j int) bool {
	ip, jp := *addrs[i], *addrs[j]
	if ip[0] > jp[0] {
		return false
	} else if ip[0] < jp[0] {
		return true
	} else {
		if ip[1] > jp[1] {
			return false
		} else if ip[1] < jp[1] {
			return true
		} else {
			if ip[2] > jp[2] {
				return false
			} else if ip[2] < jp[2] {
				return true
			} else {
				return ip[3] < jp[3]
			}
		}
	}
}

func (addrs Addrs) Swap(i, j int) {
	addr := addrs[i]
	addrs[i] = addrs[j]
	addrs[j] = addr
}

func getIps(addrs *[]net.Addr) *[]*net.IP {
	ips := make([]*net.IP, 0)
	for _, address := range *addrs {
		if ipnet, flag := address.(*net.IPNet); flag && !ipnet.IP.IsLoopback() {
			ip := ipnet.IP.To4()
			if ip == nil {
				continue
			}
			ips = append(ips, &ip)
		}
	}
	return &ips
}

func showIps(_ips *[]*net.IP) {
	var ips []*net.IP = *_ips
	sort.Sort(Addrs(ips))
	for _, ip := range ips {
		fmt.Println(ip.String())
	}
}

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		println(err.Error())
		return
	}
	ips := getIps(&addrs)
	showIps(ips)
}
