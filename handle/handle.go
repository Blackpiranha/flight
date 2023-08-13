package handle

import (
	"encoding/json"
	"flight/gin"
	"io/ioutil"
	"net/http"
)

type Address struct {
	Start string `json:"start"`
	End   string `json:"end"`
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
		//log.Fatalf("unmarshal failed")
		c.Fail(http.StatusBadRequest, "invalid parameter;")
		return
	}
	if len(address.Address) == 0 {
		c.Fail(http.StatusBadRequest, "invalid parameter;")
		return
	} else if len(address.Address) == 1 {
		c.Success(http.StatusOK, address.Address[0])
	} else {
		starts := make([]string, len(address.Address))
		ends := make([]string, len(address.Address))
		for _, addr := range address.Address {
			starts = append(starts, addr.Start)
			ends = append(ends, addr.End)
		}
		position := Address{
			Start: findDifferences(starts, ends),
			End:   findDifferences(ends, starts),
		}
		c.Success(http.StatusOK, position)
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
