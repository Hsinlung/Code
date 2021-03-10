package main

import (
	"fmt"
	"io/ioutil" // 实现了一些I/O实用程序功能
	"net/http" 	// HTTP 客户端和服务器实现
	"regexp" 	// 实现正则表达式搜索
	"strconv" 	// 实现了对基本数据类型的字符串表示的转换
	"strings" 	// 打包字符串实现简单的函数来操纵 UTF-8 编码的字符串
	"sync" 		// 提供基本的同步原语，如互斥锁
	"time" 		// 提供了测量和显示时间的功能
)

var (

	reQQWmail = `(\d+)@qq.com`

	url_1 = "https://tieba.baidu.com/p/2748780775?red_tag=3039888811"

	// 存放图片链接的数据管道
	chanImageUrls chan string
	waitGroup 	  sync.WaitGroup

	// 用于监控协程
	chanTask 	  chan string

	// 正则
	reImg		  = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`

)

// 爬取邮箱

/**
 *	关键字defer特性:
 *	用于注册延迟调用
 *	调用直到 return 前才被执,可以用来做资源清理
 *	多个defer语句，按先进后出的方式执行
 *	defer语句中的变量，在defer声明时就决定了

 *	defer用途：
 *	关闭文件句柄
 *	锁资源释放
 *	数据库连接释放
 */

func GetEmail()  {

	// 1.请求数据
	resp, err := http.Get(url_1)

	// 异常处理
	HandleError(err, "http.Get url")

	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)

	// 异常处理
	HandleError(err, "ioutil.ReadAll")

	// 字节转字符串
	pageStr := string(pageBytes)

	//fmt.Println(pageStr)

	// 3.过滤数据，过滤qq邮箱
	re := regexp.MustCompile(reQQWmail)

	// -1代表获取全部
	results := re.FindAllStringSubmatch(pageStr, -1)

	//fmt.Println(results)

	// 4.遍历结果
	for _, result := range results {
		fmt.Println("email:", result[0])
		//fmt.Println("qq:", result[1])
	}

}


// 并发爬取美图

/**
 * 并发爬思路
 *
 * 1.初始化数据管道
 * 2.爬虫写出：26个协程向管道中添加图片链接
 * 3.任务统计协程：检查26个任务是否都完成，完成则关闭数据管道
 * 4.下载协程：从管道里读取链接并下载
 */





// 爬虫图片链接到管道
func getImgUrls(url string)  {

	urls := getImgs(url)

	// 遍历切片里面所有链接 存入到数据管道
	for _, url := range urls {

		chanImageUrls <- url

	}

	// 标识当前协程完成
	// 每完成一个任务 写一条数据
	// 用于监控携程知道已经完成了几个任务

	chanTask <- url
	waitGroup.Done()

}

// 爬取当前图片链接
func getImgs(url string) (urls []string) {

	pageStr := GetPageStr(url)

	re := regexp.MustCompile(reImg)

	results := re.FindAllStringSubmatch(pageStr, -1)

	fmt.Println("共找到%d条结果\n", len(results))

	for _, result := range results {

		url  := result[0]
		urls = append(urls, url)

	}

	return

}


// 根据url获取内容

func GetPageStr(url string) (pageStr string)  {

	resp, err := http.Get(url)

	HandleError(err, "http.Get url")

	defer resp.Body.Close()

	// 读取内容
	pageBytes, err := ioutil.ReadAll(resp.Body)

	HandleError(err, "ioutil.ReadAll")

	// 字节转字符串
	pageStr = string(pageBytes)

	return pageStr

}

// 任务统计协程
func CheckOK()  {

	var count int

	for {

		url := <-chanTask

		fmt.Printf("%s 完成了爬取任务\n", url)

		count ++

		if count == 3 {

			close(chanImageUrls)
			break
		}
	}

	waitGroup.Done()

}

// 截取url名字

func GetFilenameFromUrl(url string) (filename string) {

	// 返回最后一个/位置
	lastIndex := strings.LastIndex(url, "/")

	// 切出来
	filename  = url[lastIndex + 1:]

	// 时间戳解决重名
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))

	filename    = timePrefix + "_" + filename

	return

}

// 下载图片 传入的是图片见什么

func DownloadFile(url string, filename string) (ok bool) {

	resp, err := http.Get(url)

	HandleError(err, "http.get.url")

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	HandleError(err, "resp.body")

	filename = "./img/" + filename

	// 写出数据
	err = ioutil.WriteFile(filename, bytes, 0666)

	if err != nil {

		fmt.Println(err)

		return false

	} else {

		return true
	}


}

// 下载图片

func DownloadImg()  {

	for url := range chanImageUrls{

		filename := GetFilenameFromUrl(url)

		ok := DownloadFile(url, filename)

		if ok {

			fmt.Printf("%s 下载完成\n", filename)

		} else {

			fmt.Printf("%s 下载失败\n", filename)

		}

	}

	waitGroup.Done()

}



// 处理异常
func HandleError(err error, why string)  {

	if err != nil {
		fmt.Println(why, err)
	}

}

func main()  {

	//GetEmail()

	// 初始化数据管道
	chanImageUrls = make(chan string, 1000000)
	chanTask	  = make(chan string, 2)

	// 爬虫协程
	for i := 1; i < 3; i++ {

		waitGroup.Add(1)
		go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}

	// 任务统计协程 统计26个任务是否都完成 完成则关闭通道
	waitGroup.Add(1)

	go CheckOK()

	// 下载协程：从管道中读取链接并下载
	for i := 0; i < 5; i++ {

		waitGroup.Add(1)
		go DownloadImg()

	}

	waitGroup.Wait()

}