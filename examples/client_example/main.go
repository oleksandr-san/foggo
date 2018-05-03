package main

import (
	"time"
	"fmt"
	"net/http"
	"golang.org/x/net/context"
	cl "../../client"
	"../../common"
	"io/ioutil"
	"encoding/json"
)

func getIPAdresses() ([]string, error) {
	response, err := http.Get("http://lab.ytdev.com/")
    if err != nil {
        return nil, err
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            return nil, err
		}

		var addresses []string
		if err := json.Unmarshal(contents, &addresses); err != nil {
			return nil, err
		}

		return addresses, nil
    }
}

func getData(ip string) ([]common.Data, error) {
	config := cl.NewConfiguration()
	client := cl.NewAPIClient(config) 

	client.ChangeBasePath("http://" + ip + ":3001")

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	
	data, _, err := client.DefaultApi.ListGet(ctx)
	if err != nil {
		return nil, err
	}

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("done") // prints "context deadline exceeded"
	}

	return data, nil
}

func main() {
	ips, err := getIPAdresses()
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(ips)

	for _, ip := range ips {
		fmt.Printf("Getting data from %v ...\n", ip)

		data, err := getData(ip)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(data)
	}
}