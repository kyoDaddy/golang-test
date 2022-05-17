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

var baseURL string = "https://kr.indeed.com/jobs?q=%s&limit=50"
var BaseFileName string = "jobs.csv"

type ExtractedJob struct {
	id       string
	title    string
	location string
	summary  string
}

func Scrape(term string) []ExtractedJob {
	if len(term) == 0 {
		term = "python"
	}

	var jobs []ExtractedJob
	c := make(chan []ExtractedJob)
	// step1 전체페이지
	totalPages := GetPages(term)
	// step2 각페이지 정보 고루틴으로 병행 작업후 채널에 전달
	for i := 0; i < totalPages; i++ {
		go GetPage(i, term, c)
	}
	// step3 각페이지 정보를 합치기
	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}
	// step4 excel 생성
	if len(jobs) > 0 {
		WriteJobs(jobs, term)
	}
	fmt.Println("Done, extracted", len(jobs))
	return jobs

}

func GetPages(term string) int {
	pages := 0
	baseURL = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"
	//res, err := http.Get(baseURL)
	fmt.Println("GetPages", baseURL)
	res, err := callHttp(baseURL)
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

func GetPage(page int, term string, mainC chan<- []ExtractedJob) {
	var jobs []ExtractedJob
	c := make(chan ExtractedJob)
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)

	//res, err := http.Get(pageURL)
	res, err := callHttp(pageURL)
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

func Info(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func callHttp(url string) (*http.Response, error) {
	// Request 객체 생성
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	// 필요시 헤더 추가 가능 (user-agent -> http://www.useragentstring.com/)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.57 Whale/3.14.133.23 Safari/537.36")
	// Client객체에서 Request 실행
	client := &http.Client{}
	res, err := client.Do(req)

	return res, err
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
		//fmt.Println(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
		//fmt.Println("Request failed with Status:", res.StatusCode)
	}
}

func extractJob(card *goquery.Selection, c chan<- ExtractedJob) {
	id, _ := card.Find("a").Attr("data-jk")
	title := CleanString(card.Find("a>span").Text())
	location := CleanString(card.Find(".companyLocation").Text())
	summary := CleanString(card.Find(".job-snippet").Text())

	c <- ExtractedJob{
		id:       id,
		title:    title,
		location: location,
		summary:  summary,
	}
}

var mu sync.Mutex

func WriteJobs(jobs []ExtractedJob, term string) {
	file, err := os.Create(BaseFileName)
	checkErr(err)
	//utf8bom := []byte{0xEF, 0xBB, 0xBF} // utf-8 한글깨짐
	//file.Write(utf8bom)

	w := csv.NewWriter(file)
	defer w.Flush()
	defer file.Close()

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

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
