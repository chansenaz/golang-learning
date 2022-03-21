package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// this is useful for running every check once.
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// but what if we want to check on a continuous loop?
	// for {
	// 	go checkLink(<-c, c)
	// }

	// go actually provides an alternative syntax that does the same as the code above
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	// now we can use a function literal (known as a lambda in C#/Python, or an anonymouse function in javascript)
	// to sleep for 5 seconds between calls for each routine
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
