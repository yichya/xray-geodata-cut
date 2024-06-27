package asn

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Subnets struct {
	Ipv4 []string `json:"ipv4"`
	Ipv6 []string `json:"ipv6"`
}

type AsnData struct {
	Asn         int32    `json:"asn"`
	Handle      string   `json:"handle"`
	Description string   `json:"description"`
	Subnets     *Subnets `json:"subnets"`
}

func GetAsnData(asn int32) (*AsnData, error) {
	if resp, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/ipverse/asn-ip/master/as/%d/aggregated.json", asn)); err != nil {
		return nil, err
	} else {
		var respBytes []byte
		if respBytes, err = io.ReadAll(resp.Body); err != nil {
			return nil, err
		}
		var asnData AsnData
		if err = json.Unmarshal(respBytes, &asnData); err != nil {
			return nil, err
		}
		return &asnData, resp.Body.Close()
	}
}
