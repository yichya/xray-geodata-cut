package geoip

import (
	"github.com/xtls/xray-core/app/router"
	"github.com/xtls/xray-core/common/net"
)

func Search(gin *router.GeoIPList, addr string) []string {
	var result []string
	addrParsed := net.ParseAddress(addr)

	var container router.GeoIPMatcherContainer
	for _, x := range gin.Entry {
		m, err := container.Add(x)
		if err != nil {
			return result
		}
		if m.Match(addrParsed.IP()) {
			result = append(result, x.CountryCode)
		}
	}

	return result
}
