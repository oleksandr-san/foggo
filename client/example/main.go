package main

import (
	"time"
	"fmt"
	"golang.org/x/net/context"
	sw ".."
)

func main() {
	config := sw.NewConfiguration()
	client := sw.NewAPIClient(config)

	client.ChangeBasePath("http://192.168.88.172:3001")

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	
	_, _, err := client.DefaultApi.HelloPost(ctx, "foggo-client", 1488)
	if err != nil {
		fmt.Println(err)
	}

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("done") // prints "context deadline exceeded"
	}
}