package mysession

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
)

func crawl() io.ReadCloser {
	client := &http.Client{}
	url := "http://www.ganji.com/index.htm"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", "Hm_lvt_19f2885c71d43418b0d09756de28e5f5=1540965917; ganji_uuid=7552725401734458326893; _gl_tracker=%7B%22ca_source%22%3A%22www.hao123.com%22%2C%22ca_name%22%3A%22-%22%2C%22ca_kw%22%3A%22-%22%2C%22ca_id%22%3A%22-%22%2C%22ca_s%22%3A%22other_www.hao123.com%22%2C%22ca_n%22%3A%22-%22%2C%22ca_i%22%3A%22-%22%2C%22sid%22%3A50717290460%7D; xxzl_deviceid=R7RPJBvR6EEjn%2B7WgPl%2Bf4K4lb6%2FSMQ6297meCml9lppE14oEGOn68cBxkgbgCkj; ganji_xuuid=dc4df1e9-e657-4725-db27-c53d366d8b1c.1540965927632; __utma=32156897.1861903173.1540965928.1540965928.1540965928.1; __utmc=32156897; __utmz=32156897.1540965928.1.1.utmcsr=ganji.com|utmccn=(referral)|utmcmd=referral|utmcct=/zhaopin/; GANJISESSID=pplf3e5m318122km295j84luuk; lg=1; sscode=7MQgl4jcM5iqSWoI7MsQ8%2BOl; GanjiUserName=%23t_802472249; GanjiUserInfo=%7B%22user_id%22%3A802472249%2C%22email%22%3A%22%22%2C%22username%22%3A%22%23t_802472249%22%2C%22user_name%22%3A%22%23t_802472249%22%2C%22nickname%22%3A%22%22%7D; bizs=%5B%5D; supercookie=BQNlAQplZwD5WQt1ZmV5LwH0BGx0MTV5LmWvLJR1L2HmAJWyZmt2ZTZlZGN2MzAwZQx%3D; GanjiLoginType=1; cityDomain=hz; xxzl_smartid=793605680a068cc11c4bbf51977276bf; gj_footprint=%5B%5B%22%5Cu9500%5Cu552e%22%2C%22%5C%2Fzpshichangyingxiao%5C%2F%22%5D%5D; _wap__utmganji_wap_newCaInfo_V2=%7B%22ca_n%22%3A%22-%22%2C%22ca_s%22%3A%22self%22%2C%22ca_i%22%3A%22-%22%7D; WantedListPageScreenType=1440; Hm_lpvt_19f2885c71d43418b0d09756de28e5f5=1540971998; citydomain=xa; ganji_login_act=1540975159886; __utmb=32156897.44.10.1540965928")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
	req.Header.Add("Referer", "http://xa.ganji.com/")
	if err != nil {
		panic(err)
	}
	response, _ := client.Do(req)
	return response.Body
}

func ParseHtml() {
	var body io.ReadCloser
	body = crawl()
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		panic(err)
	}
	defer body.Close()
	doc.Find("body > div.wrap > div.all-city dl").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, selection *goquery.Selection) {
			fmt.Println(selection.Text())
			url, _ := selection.Attr("href")
			fmt.Println(url)
		})
	})
	doc.Find("body > div.wrap > div.all-city dl a").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
		url, _ := selection.Attr("href")
		fmt.Println(url)
	})
}
