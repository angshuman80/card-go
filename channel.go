package main

import (
	"fmt"
	"net/http"
	"time"
)

func callChannel() {
	urlList := []string{"http://google.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://golang.org",
	"http://amazon.com",
}
    linkChannel := make(chan string)

	for _, urlname := range urlList {
		go checkLink(linkChannel,urlname)
	}

	for l := range linkChannel{
		go func (link string)  {
			time.Sleep(time.Second * 3)
			checkLink(linkChannel,link)
		}(l)
	}

/*	value1 := <- linkChannel
	value2 := <- linkChannel
	value3 := <- linkChannel
	value4 := <- linkChannel
	value5 := <- linkChannel

	close(linkChannel)

	fmt.Println("result = ",value1," ",value2," ",value3," ",value4," ",value5)*/

}

func checkLink(c chan string,urlname string) {
	time.Sleep(time.Second)
	resp,err := http.Get(urlname)
	if err!=nil{
		fmt.Println("Unable to make connection to url ", urlname)
		c <- urlname
		return
	}
	if(err==nil){
		if resp.StatusCode==200{
			fmt.Println("Success for URl ", urlname)
		}
	}
	c <- urlname
}
