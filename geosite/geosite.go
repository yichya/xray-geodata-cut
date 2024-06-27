package geosite

import (
	"os"
	"strings"

	"github.com/xtls/xray-core/app/router"
	"google.golang.org/protobuf/proto"
)

func LoadGeoSite(fn string) (*router.GeoSiteList, error) {
	geoSiteBytes, err1 := os.ReadFile(fn)
	if err1 != nil {
		return nil, err1
	}
	var geoSiteList router.GeoSiteList
	if err2 := proto.Unmarshal(geoSiteBytes, &geoSiteList); err2 != nil {
		return nil, err2
	}
	return &geoSiteList, nil
}

func GetGeoSiteCodes(in *router.GeoSiteList) []string {
	result := make([]string, len(in.GetEntry()))
	for index, x := range in.GetEntry() {
		result[index] = x.CountryCode
	}
	return result
}

func CutGeoSiteCodes(in *router.GeoSiteList, codesToKeep []string) *router.GeoSiteList {
	out := &router.GeoSiteList{
		Entry: make([]*router.GeoSite, 0, len(codesToKeep)),
	}
	kept := make(map[string]bool, len(codesToKeep))
	for _, x := range in.GetEntry() {
		for _, y := range codesToKeep {
			u := strings.ToUpper(y)
			if x.CountryCode == u {
				if kept[u] {
					continue
				}
				out.Entry = append(out.Entry, x)
				kept[u] = true
			}
		}
	}

	return out
}

func SaveGeoSite(in *router.GeoSiteList, fn string) error {
	b, err := proto.Marshal(in)
	if err != nil {
		return err
	}
	return os.WriteFile(fn, b, 0644)
}
