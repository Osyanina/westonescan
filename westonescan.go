package main

import (
	"fmt"
	"github.com/malfunkt/iprange"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func iplist(ips string) ([]net.IP, error) {
	ad, err := iprange.ParseList(ips)
	if err != nil {
		return nil, err
	}
	list := ad.Expand()
	return list, err
}

func hostscan(ip string, count int) bool {
	var c = &exec.Cmd{}
	switch runtime.GOOS {
	case "windows":
		c = exec.Command("ping", "-n", strconv.Itoa(count), ip)
	default:
		c = exec.Command("ping", "-c", strconv.Itoa(count), ip)
	}
	out, err := c.StdoutPipe()
	defer out.Close()
	c.Start()
	if err != nil {
		log.Fatal(err)
	}else {
		re, err := ioutil.ReadAll(out)
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(string(re), "TTL") || strings.Contains(string(re), "ttl") {
			return true
		}else {
			return false
		}
	}
	return false
}

func portlist(sel string) ([]int, error) {
	ports := []int{}
	if sel == "" {
		return ports, nil
	}
	ranges := strings.Split(sel, ",")
	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if strings.Contains(r, "-") {
			parts := strings.Split(r, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("端口格式错误: '%s'", r)
			}
			a, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("端口错误: '%s'", parts[0])
			}
			b, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("端口错误: '%s'", parts[1])
			}
			if a > b {
				return nil, fmt.Errorf("端口范围错误: %d-%d", a, b)
			}
			for s := a; s <= b; s++ {
				ports = append(ports, s)
			}
		}else {
			if port, err := strconv.Atoi(r); err != nil {
				return nil, fmt.Errorf("端口错误: '%s'", r)
			}else {
				ports = append(ports, port)
			}
		}
	}
	return ports, nil
}

func main()  {
	fmt.Println(`
 __          __       _                   
 \ \        / /      | |                  
  \ \  /\  / /__  ___| |_ ___  _ __   ___ 
   \ \/  \/ / _ \/ __| __/ _ \| '_ \ / _ \
    \  /\  /  __/\__ \ || (_) | | | |  __/
     \/  \/ \___||___/\__\___/|_| |_|\___|

		磐石资产扫描工具
		版本       1.0
		`)
	if len(os.Args) == 3 {
		ips := os.Args[1]
		ps := os.Args[2]
		is, err := iplist(ips)
		_ = err
		p1, err := portlist(ps)
		for _, ip := range is {
			iptrue := hostscan(fmt.Sprintf("%v", ip), 3)
			if iptrue {
				fmt.Println("主机", ip, "存活,开始扫描对外开放端口")
				for _, port := range p1 {
					_, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ip, port), time.Second)
					if err == nil {
						fmt.Printf("主机 %v的%v端口开放\n", ip, port)
					}
				}
			}else {
				fmt.Println("主机", ip, "未上线")
			}
		}
		fmt.Println("扫描完成")
	}else {
		fmt.Println("格式输入错误，请输入要扫描的主机ip和端口，例如:")
		fmt.Println("westonescan.exe 192.168.0.0/16 1-65535")
		fmt.Println("westonescan.exe 192.168.0.1-254 80,3389")
		fmt.Println("westonescan.exe 192.168.0.1 3389")
	}
}