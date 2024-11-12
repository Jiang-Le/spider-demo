package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println(fetchAddressLabels("0x50D1c9771902476076eCFc8B2A83Ad6b9355a4c9"))
	fmt.Println(fetchAddressLabels("0xf89d7b9c864f589bbF53a82105107622B35EaA40"))
	fmt.Println(fetchAddressLabels("0x95aD61b0a150d79219dCF64E1E6Cc01f0B64C4cE"))
	fmt.Println(fetchAddressLabels("0x60e4d786628fea6478f785a6d7e704777c86a7c6"))
}

func fetchAddressLabels(address string) []string {
	// 目标URL
	url := fmt.Sprintf("https://etherscan.io/address/%s", address)

	// 发送HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 检查HTTP响应状态码
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// 使用goquery解析HTML文档
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var labels []string
	// 查找id为ContentPlaceHolder1_divSummary的元素下的第一个div中的第一个div中的所有span
	doc.Find("#ContentPlaceHolder1_divSummary > div:first-child > div:first-child .hash-tag ").Each(func(i int, s *goquery.Selection) {
		// 提取span的文本内容
		text := s.Text()
		labels = append(labels, text)
	})
	return labels
}