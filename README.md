### Simple HTTP Client for mayaserver
[![Go Report Card](https://goreportcard.com/badge/github.com/satyamz/mayaclient)](https://goreportcard.com/report/github.com/satyamz/mayaclient)

##### Example usage
1. Get the client `go get github.com/satyamz/mayaclient`
2. You can write `mayacli.go` as given below 


```golang
package main

import (
        "fmt"
        "github.com/satyamz/mayaclient" )

func main() {
        c := mayaclient.Client{ URL : "http://127.0.0.1:5656/latest/meta-data/instance-id",}
        fmt.Println(c.MayaClient())
}

```

3. Run `mayacli.go` using `go run mayacli.go` or `go build mayacli.go ` & `./mayacli` 

