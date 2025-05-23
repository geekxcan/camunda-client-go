package main

import (
	"fmt"
	"time"

	camundaclientgo "github.com/geekxcan/camunda-client-go/v3"
)

func main() {
	client := camundaclientgo.NewClient(camundaclientgo.ClientOptions{
		EndpointUrl: "http://localhost:8080/engine-rest",
		ApiUser:     "demo",
		ApiPassword: "demo",
		Timeout:     time.Second * 10,
	})

	result, err := client.ProcessInstance.GetList(nil)
	if err != nil {
		fmt.Printf("Error deploy process: %s\n", err)
		return
	}
	fmt.Printf("Result: %#+v\n", result)
}
