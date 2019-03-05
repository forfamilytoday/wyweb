package utils

import (
	_ "wyweb/models/user"
	"github.com/astaxie/beego/orm"
	"fmt"
	"os"
	"github.com/astaxie/beego"
	"strconv"
	"sync"
	"net/http"
	"log"
	"github.com/axgle/mahonia"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"math"
	"time"
	"github.com/go-redis/redis"
)

func CreateDB() {
	file := "logs/lock.log"

	if FileExist(file) {
		return
	}

	fb, err := os.Create(file)
	defer fb.Close()

	if err != nil {
		beego.Error(err)
		return
	}

	orm.RunCommand()

	// 数据库别名
	name := "default"

	// drop table 后再建表
	force := true

	// 打印执行过程
	verbose := true

	// 遇到错误立即返回
	err1 := orm.RunSyncdb(name, force, verbose)
	if err1 != nil {
		fmt.Println(err1)
	}
}

func FileExist(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

func Goquery() {
	var WG sync.WaitGroup
	var flag bool
	die := make(chan bool)

	//控制跳出for循环
	go func() {
		if ok := <-die; ok {
			flag = true
		}
	}()

	res, err := http.Get("https://search.51job.com/list/020000,000000,0000,00,9,99,php,2,1.html")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	//解析gbk，这里相当于注册下
	enc := mahonia.NewDecoder("gbk")

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	//获取总条数
	sum, _ := strconv.ParseFloat(strings.Replace(strings.Replace(strings.TrimSpace(enc.ConvertString(doc.Find("#resultList .dw_tlc .rt").First().Text())), "共", "", -1), "条职位", "", -1), 10)

	count := int(math.Ceil(sum / 50))
	//连接redis
	conn := RedisConn("127.0.0.1:6379", 2)

	defer conn.Close()

	//加入redis队列
	RedisPush(count, conn)

	for {
		func() {
			length, _ := conn.LLen("php").Result()
			if length < 20 {
				s, _ := conn.LRange("php", 0, length).Result()
				for _, val := range s {
					conn.LPop("php")
					//拼接url
					url := fmt.Sprintf("https://search.51job.com/list/020000,000000,0000,00,9,99,php,2,%v.html", val)
					//开协程去处理爬虫
					WG.Add(1)
					go Work(url, val, conn, &WG)
				}

				die <- true
			} else {
				s, _ := conn.LRange("php", 0, 19).Result()
				for _, val := range s {
					conn.LPop("php")
					//拼接url
					url := fmt.Sprintf("https://search.51job.com/list/020000,000000,0000,00,9,99,php,2,%v.html", val)
					WG.Add(1)
					go Work(url, val, conn, &WG)
					//开协程去处理爬虫
				}
			}
		}()

		//结束堵塞
		time.Sleep(time.Second * 1)

		if flag {
			break;
		}
	}
	WG.Wait()

}

func unitBranch(s string) (small int, big int) {
	if strings.Contains(s, "万/月") {
		tempSlice := strings.Split(strings.Replace(s, "万/月", "", -1), "-")
		temSmall, _ := strconv.ParseFloat(tempSlice[0], 10)
		temBig, _ := strconv.ParseFloat(tempSlice[1], 10)
		fmt.Println(temSmall, temBig, tempSlice[1])
		small = int(temSmall * 10)
		big = int(temBig * 10)

	} else if strings.Contains(s, "千/月") {
		tempSlice := strings.Split(strings.Replace(s, "千/月", "", -1), "-")
		small, _ = strconv.Atoi(tempSlice[0])
		big, _ = strconv.Atoi(tempSlice[1])
	}

	return
}

func RedisConn(addr string, db int) (client *redis.Client) {
	return redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   db,
	})
}

func RedisPush(sum int, conn *redis.Client) {
	for i := 1; i <= sum; i++ {
		conn.RPush("php", i)
	}
}

func Work(url string, page string, conn *redis.Client, group *sync.WaitGroup) {

	res, err := http.Get(url)
	enc := mahonia.NewDecoder("gbk")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#resultList .el").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		if i >= 1 {
			//职位链接
			href, _ := selection.Find(".el .t1").Find("a").Attr("href")
			//职位名称
			name := enc.ConvertString(selection.Find(".el .t2").Find("a").Text())
			address := enc.ConvertString(selection.Find(".el .t3").Text())
			salary := enc.ConvertString(selection.Find(".el .t4").Text())

			small, big := unitBranch(salary)

			key := page + "-" + strconv.Itoa(i)
			temp := make(map[string]interface{})
			temp["url"] = href
			temp["name"] = name
			temp["address"] = address
			temp["salarySma"] = small
			temp["salaryBig"] = big

			conn.HMSet(key, temp)
			return true
		}
		return true

	})
	group.Done()
}
