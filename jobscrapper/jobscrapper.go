package jobscrapper

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

type ExtractedJob struct {
	id       string
	title    string
	location string
	summary  string
}

func GetPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	// goquery : go get github.com/PuerkitoBio/goquery
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func GetPage(page int, mainC chan<- []ExtractedJob) {
	var jobs []ExtractedJob
	c := make(chan ExtractedJob)
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".job_seen_beacon")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
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

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func extractJob(card *goquery.Selection, c chan<- ExtractedJob) {
	id, _ := card.Find("a").Attr("data-jk")
	title := cleanString(card.Find("a>span").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	summary := cleanString(card.Find(".job-snippet").Text())

	c <- ExtractedJob{
		id:       id,
		title:    title,
		location: location,
		summary:  summary,
	}
}

var mu sync.Mutex

func WriteJobs(jobs []ExtractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title", "Location", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	// use goroutine
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	wg.Add(len(jobs))

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.summary}
		go WriteToFile(w, jobSlice, wg)
	}

	wg.Wait()

}

func WriteToFile(w *csv.Writer, jobSlice []string, wg *sync.WaitGroup) {
	// *bufio.Writer가 동시 액세스를 지원하지 않으므로, 뮤텍스로 보호필요... 삽질....
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()

	jwErr := w.Write(jobSlice)
	checkErr(jwErr)
}
