package main

import (
	"flag"
	"fmt"
	"github.com/yichya/xray-geodata-cut/geoip"
	"github.com/yichya/xray-geodata-cut/geosite"
	"google.golang.org/protobuf/proto"
	"strings"
)

func main() {
	ft := flag.String("type", "", "GeoIP (geoip) or GeoSite (geosite)")
	in := flag.String("in", "", "Path to GeoIP file")
	show := flag.Bool("show", false, "Print codes in GeoIP or GeoSite file")
	keep := flag.String("keep", "cn,private,geolocation-!cn", "GeoIP or GeoSite codes to keep (private is always kept for GeoIP)")
	trimipv6 := flag.Bool("trimipv6", false, "Trim all IPv6 ranges in GeoIP file")
	out := flag.String("out", "", "Path to processed file")

	flag.Parse()
	if ft == nil {
		ft = proto.String("")
	}
	switch *ft {
	case "geoip":
		{
			gin, err := geoip.LoadGeoIP(*in)
			if err != nil {
				panic(err)
			}
			if *show {
				fmt.Println(geoip.GetGeoIPCodes(gin))
			} else {
				gout := geoip.CutGeoIPCodes(gin, strings.Split(*keep, ","), *trimipv6)
				if err = geoip.SaveGeoIP(gout, *out); err != nil {
					panic(err)
				}
			}
			return
		}
	case "geosite":
		{
			gin, err := geosite.LoadGeoSite(*in)
			if err != nil {
				panic(err)
			}
			if *show {
				fmt.Println(geosite.GetGeoSiteCodes(gin))
			} else {
				gout := geosite.CutGeoSiteCodes(gin, strings.Split(*keep, ","))
				if err = geosite.SaveGeoSite(gout, *out); err != nil {
					panic(err)
				}
			}
			return
		}
	default:
		{
			flag.Usage()
		}
	}
}
