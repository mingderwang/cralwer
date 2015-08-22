package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	//"github.com/davecgh/go-spew/spew"
)

func main() {
	rand.Seed(int64(time.Now().Unix()))
	maxNumber := 100000000

	for i := 0; i < 100000; i++ {

		id := fmt.Sprintf("%08d\n", rand.Intn(maxNumber))
		scan(id)
	}

}

func scan(id string) {
	response, err := http.Get("http://gcis.nat.g0v.tw/api/show/" + id)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		str := "{\"error\":true,\"message\":\"Company not found\"}"
		str2 := string(contents)
		if str2 != str {
			fmt.Printf(id)
			//fmt.Printf("%s\n", str2)
			//spew.Dump(contents)
		}
	}
}
