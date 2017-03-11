package mayaclient

import "testing"

func TestMayaClient(t *testing.T) {
	instanceID := "\"any-compute\""
	c := Client{URL: "http://127.0.0.1:5656/latest/meta-data/instance-id"}

	response := c.MayaClient()
	if response != instanceID {
		t.Error("Expected response ", instanceID, " got ", response)
	}
}
