package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){
	fmt.Println("Hello, world")
	resp, err := http.Get("http://www.baidu.com/")
	defer resp.Body.Close()
	if err != nil{
		fmt.Println("http get error:", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("read error:", err)
		return
	}
	fmt.Println(string(body))
	return
}
