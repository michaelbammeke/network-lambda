package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/yaml.v2"
)

const (
	timeoutSecs = 5
)

type Conf struct {
	Endpoints []Endpoint
}

type Endpoint struct {
	Host string
	Port string
}

func HandleRequest(ctx context.Context) (err error) {
	// Read yaml file
	yamlFile, _ := ioutil.ReadFile("test.yaml")

	urlList := &Conf{}
	_ = yaml.Unmarshal(yamlFile, urlList)

	for _, endpoint := range urlList.Endpoints {
		timeout := timeoutSecs * time.Second
		url := fmt.Sprintf("%s:%s", endpoint.Host, endpoint.Port)
		_, err = net.DialTimeout("tcp", url, timeout)

		if err == nil {
			fmt.Printf("connection to %s SUCCESSFUL!!!\n", url)
		} else {
			fmt.Printf("connection to %s FAILED!!! -- REASON: %v\n", url, err)
		}
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
