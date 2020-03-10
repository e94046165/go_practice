package main

import(
	"fmt"
	"net/http"
	//用別人寫好的解析器來濾出 links 
	"github.com/jackdanger/collectlinks"
)

func main(){
	fmt.Println("Hello, crawler")
	url := "http://www.baidu.com/"
	download(url)
}

func download(url string){
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	response, err := client.Do(req)
	if err != nil{
		fmt.Println("http get error", err)
		return
	}
	defer response.Body.Close()
	links := collectlinks.All(response.Body)
	for _, link := range links{
		fmt.Println("parse url", link)
	}
}
