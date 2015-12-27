package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/wangtuanjie/ip17mon"
)

var (
	path string
	file string

	batch = 100
)

func init() {
	flag.StringVar(&path, "path", "17monipdb.dat", "17monipdb 数据文件路径")
	flag.StringVar(&file, "ips", "ips.txt", "要查询 ip 数据文件路径")
}

func main() {
	flag.Parse()

	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("打开日志[%s]出错: %s\n", file, err.Error())
		return
	}

	if err = ip17mon.Init(path); err != nil {
		fmt.Printf("初始化 ip 库失败: %s\n", err.Error())
		return
	}

	buf, output := bufio.NewReader(f), make([]string, 0, batch)
	for {
		ip, err := buf.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Printf("读取日志[%s]出错: %s\n", file, err.Error())
				return
			}
			break
		}

		ip = strings.TrimSpace(ip)
		loc, err := ip17mon.Find(ip)
		if err != nil {
			output = append(output, fmt.Sprintf("%s,%s,%s,%s,%s", ip, err.Error(), "", "", ""))
			continue
		}

		output = append(output, fmt.Sprintf("%s,%s,%s,%s,%s", ip, loc.Country, loc.Region, loc.City, loc.Isp))
		if len(output) < batch {
			continue
		}

		fmt.Println(strings.Join(output, "\n"))
		output = output[:0]
	}

	if len(output) > 0 {
		fmt.Println(strings.Join(output, "\n"))
	}
}
