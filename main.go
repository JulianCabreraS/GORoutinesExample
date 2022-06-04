//Example of Couroutines
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
	}

	//Create a channel
	c := make(chan string)

	for _, link := range links {

		go checkLink(link, c) // Coroutines
	}
	//Receiving the value from the channel
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

	fmt.Println(link, "is up")
	c <- link
}
