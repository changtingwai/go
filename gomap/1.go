package main

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

import "github.com/garyburd/redigo/redis"
import "flag"
import "os"

import "strings"

var max_num = 200       //*****************************************
var merge_debug_num = 0 //*******************************************
var merge_k = 1         //******************************************含义

func main() {
	//设置输入参数的默认值及解析方
	var host = flag.String("h", "127.0.0.1", "redis proxy")
	var port = flag.String("p", "6379", "redis port")
	var merge = flag.String("merge", "no", "merge insert value")
	//开始解析参数
	flag.Parse()
	fmt.Fprintf(os.Stderr, "conf info : host[%s], port[%s], merge[%s]\n", *host, *port, *merge)
	c, err := redis.DialTimeout("tcp", *host+":"+*port, 0, 1000*time.Second, 1000*time.Second)
	if err != nil {
		panic(err)
	}
	//defer，退出前操作，安全考虑及代码简洁
	defer c.Close()
	fmt.Fprintln(os.Stderr, "connection done")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		panic(err)
	}
	var cmd string
	//总key数
	var count_all int = 0
	//存在的key数
	var count_exi int = 0
	//存在的key里面总有效时间
	var Time_all int = 0
	//从标准输入中读取数据,redis_import.sh中cat命令list数据
	reader := bufio.NewReader(os.Stdin)
	//Scan开始处理每一条数据
	for {
		count_all = 0
		count_exi = 0
		Time_all = 0
		cmd, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		count_all = 1
		cmd = strings.Trim(cmd, "\n")
		cmd = strings.Trim(cmd, "\t")
		cmd = strings.Trim(cmd, " ")
		if cmd == "" {
			continue
		}

		exi_res, err := redis.Int(c.Do("exists", cmd))
		if err == nil {
			//累加存在的key
			count_exi = exi_res
		}
		if exi_res == 1 {
			//累加存在的key的剩余时间
			count_time, err := redis.Int(c.Do("ttl", cmd))
			if err == nil {
				Time_all = count_time
			}
		}
		rescmd := strings.Split(cmd, "_")
		fmt.Printf("%v\t%v\t%v\t%v\n", rescmd[0], count_all, count_exi, Time_all)

	}

}
