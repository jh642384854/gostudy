package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)
//设置的rateLimit是用来做速率限制
var rateLimit = time.Tick(100 * time.Millisecond)
func Fetch(url string) ([]byte, error) {
	<- rateLimit
	//生成client，参数默认
	client := &http.Client{}
	//提交请求
	request,err := http.NewRequest("GET",url,nil)
	//增加header选项,解决403禁止访问的问题，特别是下面的User-Agent值，这个可以查看浏览器发起请求的时候，得到这些值
	request.Header.Add("Cookie","sid=0ca4792d-26d9-49c9-b324-1f97234578ff; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1559027355; FSSBBIl1UgzbN7N80S=wVSyqoxsNmdZ42ossVrbvXAL5kZNAkm5SQqZbvtD0MPdYSaG90pY3HtmvOjMz9hc; __guid=228810700.3928245817306522000.1559027396929.942; monitor_count=6; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1559030753; FSSBBIl1UgzbN7N80T=3frvPGcPsziYs19H_ttZPUg4NDYRdJBdCyRZTnkD6Cg5hURmDHmkHX77_Fmr_iMTcbh3okiXzgR5ZGPkH7M_8UAegIky7zQ48ITcQ9dL6DE_HQQC5UBoRQtnUwKGtUBaWIfdkJVQJFIwc39HX62pH_kOpfhEIre1PVosUThaB8jsM6kNdyzdcn.77VTiAuiVqkfxGfTSeotbqt3EEUYyOyXeReX7.MoCPJpyLZ6jFkp5Lj030nhjDfyGXEMWpef4QgY6tf6XFnEZd99Vh3uu5JEbqqekbzW0C286.6YhpvbjACBtFiQmqBsbqqqpJw1Pwi7Sy2nDkSpE5728zL6RLWzkR2lJrNsp4aO3Bw1J.fsmjBq")
	request.Header.Add("Referer","http://www.zhenai.com/zhenghun/akesu")
	request.Header.Add("Host","album.zhenai.com")
	request.Header.Add("Upgrade-Insecure-Requests","1")
	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	if err != nil{
		panic(err)
	}
	response,err := client.Do(request)
	//response, err := http.Get(url)  //之所以没有用这个，是解决可能会遇到的403问题
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error:Status code:%d", response.StatusCode)
	}
	//抓取页面的编码判断
	pgencoding := determinEncoding(response.Body)
	utf8Reader := transform.NewReader(response.Body, pgencoding.NewEncoder())

	return ioutil.ReadAll(utf8Reader)
}

//检查页面的编码
func determinEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error :%v",err)
		//返回一个默认的编码
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
