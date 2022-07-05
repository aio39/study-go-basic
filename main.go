package main

import (
	"fmt"
	"time"
)

// func main() {

// 	// var results map[string]string  초기화 안됨
// 	// var results = map[string]string{}
// 	var results = make(map[string]string)

// 	urls := []string{
// 		"https://www.airbnb.com/",
// 		"https://www.google.com/",
// 		"https://www.amazon.com/",
// 		"https://www.reddit.com/",
// 		"https://www.google.com/",
// 		"https://soundcloud.com/",
// 		"https://www.facebook.com/",
// 		"https://www.instagram.com/",
// 		"https://academy.nomadcoders.co/",
// 	}

// 	for _, url := range urls {
// 		result := "OK"
// 		err := hitURL(url)
// 		if err != nil {
// 			result = "FAILED"
// 		}
// 		results[url] = result
// 	}
// 	// fmt.Println(results)
// 	for url, result := range results {
// 		fmt.Println(url, result)
// 	}
// }

// var errRequestFailed = errors.New("request failed")

// func hitURL(url string) error {
// 	fmt.Println("Checking:", url)
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		fmt.Println(err, resp.StatusCode)
// 		return errRequestFailed
// 	}
// 	return nil
// }

func main() {
	c := make(chan string)
	people := [4]string{"aio", "miku", "rin", "ichika"}
	for _, person := range people {
		go isCute(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

}

func isCute(person string, c chan string) {
	time.Sleep(time.Second * 1)
	fmt.Println(person)
	c <- person + " is cute" // go에서 채널에 메세지를 보내는 법
}
