package handle

import (
	"encoding/json"
	"flight/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type Address struct {
	Start string
	End   string
}
type AddressList struct {
	Address []Address
}

func Track(c *gin.Context) {
	address := &AddressList{}
	defer c.Req.Body.Close()
	data, _ := ioutil.ReadAll(c.Req.Body)
	err := json.Unmarshal(data, address)
	if err != nil {
		log.Fatalf("unmarshal failed")
	}
	if len(address.Address) == 1 {
		c.JSON(http.StatusOK, address.Address[0])
	} else {
		starts := make([]string, len(address.Address))
		ends := make([]string, len(address.Address))
		for _, addr := range address.Address {
			starts = append(starts, addr.Start)
			ends = append(ends, addr.End)
		}
		c.JSON(http.StatusOK, gin.H{
			"Start": findDifferences(starts, ends),
			"End":   findDifferences(ends, starts),
		})
	}

}

func findDifferences(starts, ends []string) string {
	hashMap := make(map[string]bool)
	for _, v := range ends {
		hashMap[v] = true
	}
	differ := ""
	for _, v := range starts {
		if _, exists := hashMap[v]; !exists {
			differ = v
			break
		}
	}
	return differ
}
