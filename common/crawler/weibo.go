/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: weibo
* @Date: 2021-9-15 9:43
 */
package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

const HOSTNAME = "https://s.weibo.com"

func Weibo() string {

	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://s.weibo.com/top/summary?cate=realtimehot", nil)
	if err != nil {
		fmt.Println("err")
		return ""
	}
	// 添加请求头
	req.Header.Add("Host", "s.weibo.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3880.400 QQBrowser/10.8.4554.400")
	req.Header.Add("Cookie", "SINAGLOBAL=1128319787547.8801.1569292994252; _ga=GA1.2.475239693.1572926506; SCF=AvdKMnu-HG4cwIn8FaKCvy6CmHahnbZI2azIav75-xNphQgp8rdtFpQObmP8FMK_IAQtbwiC1SnGDMGjOQ5Y5uU.; ALF=1668822713; SUB=_2AkMWnMXVf8NxqwJRmfsWxW3mbYpzzgDEieKgwDQOJRMxHRl-yT9jqkdatRB6PRzrOq_FIaEjGwk0QC4lkhhza4nwX4Mk; SUBP=0033WrSXqPxfM72-Ws9jqgMF55529P9D9WhHfB2q-WCde29RXkOjuopJ; _s_tentry=weibo.com; Apache=7360736401460.694.1639992063755; ULV=1639992063973:67:1:1:7360736401460.694.1639992063755:1635493974091; UOR=www.sendong.com,widget.weibo.com,www.baidu.com; WBtopGlobal_register_version=2021122209")
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err")
		return ""
	}
	defer resp.Body.Close()
	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("err")
	//}
	//fmt.Println(string(b))
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)
	var str = ""
	for _, v := range []string{"热", "爆", "新", "沸"} {

		doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
			redu := s.Find(".td-03 i").Text()

			if redu == v {

				//跳过置顶推荐
				if i == 0 && redu == "热" {
					return
				}
				href, _ := s.Find(".td-02 a").Attr("href")
				herfText := s.Find(".td-02 a").Text()
				//redu := s.Find(".td-03 i").Text()
				if redu != "商" {
					//fmt.Println(href, herfText)
					//href, _ = url.QueryUnescape(href)
					str += fmt.Sprintf("%s%s %s %s\n", HOSTNAME, href, herfText, redu)

				}
			}
		})
	}

	if len(str) > 0 {
		return str
	} else {
		return ""
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
