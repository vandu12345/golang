package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://amazon.com",
		"https://google.com",
		"https://facebook.com/",
		"https://flipkart.com/",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	// for _, v := range links {
	// 	fmt.Println(v + " " + <-c)
	// }
	// for {
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		go checkLink(l, c)
	}
	// time.Sleep(10 * time.Second)
}
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "link is up")
	c <- link

}
