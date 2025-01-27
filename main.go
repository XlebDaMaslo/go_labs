package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	resp, err := http.Get("https://www.anekdot.ru/last/anekdot/")
	if err != nil {
		log.Fatalf("Ошибка при запросе страницы: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка, код статуса: %v", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при разборе HTML: %v", err)
	}

	doc.Find(".text").Each(func(i int, s *goquery.Selection) {
		html, err := s.Html()
		if err != nil {
			log.Printf("Ошибка при получении HTML: %v", err)
			return
		}
		
		text := strings.ReplaceAll(html, "<br>", "\n")
		text = strings.ReplaceAll(text, "<br/>", "\n")
		text = strings.ReplaceAll(text, "<p>", "")
        text = strings.ReplaceAll(text, "</p>", "")

		re := regexp.MustCompile(`<[^>]*>`)
		text = re.ReplaceAllString(text, "")

		fmt.Println(text)
		fmt.Println("---ахахххаахаах---\n\n\n")
		time.Sleep(10 * time.Second)
	})
}