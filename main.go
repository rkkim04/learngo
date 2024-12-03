package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// https://kr.indeed.com/jobs?q=python&l=&from=searchOnHP&vjk=875d4550ab104a22

// var baseURL string = "https://tech.kakaobank.com/"
var baseURL string = "https://www.jobkorea.co.kr/Search/?stext=python&tabType=recruit"

func main() {
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}

}

func getPage(page int) {
	// pageURL := baseURL + "page/" + strconv.Itoa(page)
	pageURL := baseURL + "&Page_No=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	// to prevent memory leak, close the connection
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find("post")
	searchCards.Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("data-listno")
		fmt.Println(id)
	})

}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	// to prevent memory leak, close the connection
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
