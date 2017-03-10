package mayaclient

import (
        "io/ioutil"
        "net/http"
)

//Client : struct for maya client
type Client struct {
        URL string
}

//MayaClient : Client to connect to mayaserver and return response
func (c *Client) MayaClient() string {
        resp, err := http.Get(c.URL)
        if err != nil {
                panic(err)
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                panic(err)
        }
        return string(body[:])
}
