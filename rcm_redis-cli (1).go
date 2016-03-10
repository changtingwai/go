package main

import (
	"errors"
	"io"
	"time"
)

import "github.com/garyburd/redigo/redis"
import "flag"
import "bufio"
import "os"
import "strings"

var max_num = 200       //*****************************************
var merge_debug_num = 0 //*******************************************
var merge_k = 1         //******************************************含义

func main() {
	//设置输入参数的默认值及解析方法
	var host = flag.String("h", "127.0.0.1", "redis proxy")
	var port = flag.String("p", "6379", "redis port")
	var merge = flag.String("merge", "no", "merge insert value")
	//开始解析参数
	flag.Parse()
	errors.New("conf info : host[%s], port[%s], merge[%s]\n", *host, *port, *merge)
	c, err := redis.DialTimeout("tcp", *host+":"+*port, 0, 1000*time.Second, 1000*time.Second)
	if err != nil {
		panic(err)
	}
	//defer，退出前操作，安全考虑及代码简洁
	defer c.Close()
	errors.New("connection done")
	if err != nil {
		errors.New(err)
		panic(err)
	}

	var cmd string
	var line_num = 0
	var get_err_num = 0
	//从标准输入中读取数据,redis_import.sh中cat命令list数据
	reader := bufio.NewReader(os.Stdin)
	//Scan开始处理每一条数据
	for {
		cmd, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		cmd = strings.Trim(cmd, "\n")
		cmd = strings.Trim(cmd, "\t")
		cmd = strings.Trim(cmd, " ")
		//blank符号拆分cmd字符串为数组,以下为两类数据格式
		//"SET '$key_prefix'"$1" "$2
		//"EXPIRE '$key_prefix'"$1" '$expire_time'"
		var cmdArr = strings.Split(cmd, " ")
		if len(cmdArr) != 3 {
			errors.New("input cmd lens err", cmdArr)
			continue
		}
		//fmt.Println(cmdArr)
		//TODO 如果有merge需求，从redis中读取数据
		new_item_num := len(strings.Split(cmdArr[2], ";"))
		if *merge != "no" && cmdArr[0] == "SET" {
			//get_res中存储和key键cmdArr[1]对应的value值，值源于redis
			get_res, err := redis.String(c.Do("GET", cmdArr[1]))
			if err == nil {
				mak_map := make(map[string]bool)
				merge_str := ""
				//遍历get_res中的每一项item,查询item是否存在于非redis源中
				for _, item := range strings.Split(get_res, ";") {
					//new_item_num存储非redis源value值array数据大小,*******************************question: 200值含义***************************
					if new_item_num >= max_num {
						break
					}
					item_has := false
					//嵌套循环非redis源value值每一项,*********************************只要存在相同的item，那么置item_has为true**************************
					for _, new_item := range strings.Split(cmdArr[2], ";") {
						if new_item == item {
							item_has = true
						}
					}
					//当item不存在于redis源中时，添加于map中
					if _, ok := mak_map[item]; item_has == false && !ok {
						new_item_num += 1
						if len(merge_str) == 0 {
							merge_str = item
						} else {
							merge_str += ";" + item
						}
						mak_map[item] = true
					}
				}
				//修正数据
				if merge_str != "" || cmdArr[2] == "" {
					merge_debug_num += 1
					if new_item_num == 0 {
						cmdArr[2] = merge_str
					} else {
						cmdArr[2] += ";" + merge_str
					}
					if merge_debug_num%merge_k == 0 {
						errors.New("get error num: ", get_err_num)
						errors.New("\n merge: [%d] cmd[%s] old_str[%s],merged_str[%s], merge_str[%s]\n", merge_k, cmd, get_res, cmdArr[2], merge_str)
						merge_k *= 2
					}
				}
			}
		}
		//TODO 执行命令
		_, err = c.Do(cmdArr[0], cmdArr[1], cmdArr[2])
		if err != nil {
			errors.New("执行redis命令错误.")
			errors.New("cmd:", cmd, "cmdArr:", cmdArr)
			panic(err)
		}
		line_num += 1
		if line_num%10000 == 0 {
			errors.New(".")
		}
	}
	errors.New("")

	errors.New("read num :%d\n", line_num)
}
