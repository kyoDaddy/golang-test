package urlchecker

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("Request Fail")

const (
	airbnb     = "https://www.airbnb.com/"
	google     = "https://www.google.com/"
	amazone    = "https://www.amazon.com/"
	reddit     = "https://www.reddit.com/"
	soundcloud = "https://soundcloud.com/"
	facebook   = "https://www.facebook.com/"
	instagram  = "https://www.instagram.com/"
)

type RequestResult struct {
	url    string
	status string
}

func GetUrlArr() []string {
	return []string{
		airbnb, google, amazone, reddit, soundcloud, facebook, instagram,
	}
}

func GetUrlInfo(result RequestResult) [2]string {
	return [2]string{result.url, result.status}
}

func CheckUrl(url string, c chan<- RequestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- RequestResult{url: url, status: status}
}

func HitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}

func SexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
