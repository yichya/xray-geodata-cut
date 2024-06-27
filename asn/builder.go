package asn

import (
	"fmt"
	"net/netip"

	"github.com/xtls/xray-core/app/router"
)

func BuildGeoIp(asn []int32, trimIpv6 bool) (*router.GeoIPList, error) {
	result := &router.GeoIPList{}
	for _, x := range asn {
		data, err := GetAsnData(x)
		if err != nil {
			return nil, err
		}
		entry := &router.GeoIP{
			CountryCode: fmt.Sprintf("AS%d", x),
			Cidr:        make([]*router.CIDR, 0),
		}
		for _, y := range data.Subnets.Ipv4 {
			ip, e1 := netip.ParsePrefix(y)
			if e1 != nil {
				return nil, e1
			}
			b, e2 := ip.Addr().MarshalBinary()
			if e2 != nil {
				return nil, e2
			}
			entry.Cidr = append(entry.Cidr, &router.CIDR{
				Ip:     b,
				Prefix: uint32(ip.Bits()),
			})
		}
		if !trimIpv6 {
			for _, y := range data.Subnets.Ipv6 {
				ip, e1 := netip.ParsePrefix(y)
				if e1 != nil {
					return nil, e1
				}
				b, e2 := ip.Addr().MarshalBinary()
				if e2 != nil {
					return nil, e2
				}
				entry.Cidr = append(entry.Cidr, &router.CIDR{
					Ip:     b,
					Prefix: uint32(ip.Bits()),
				})
			}
		}
		result.Entry = append(result.Entry, entry)
	}
	return result, nil
}
