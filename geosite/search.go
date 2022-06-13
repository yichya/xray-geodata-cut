package geosite

import (
	"github.com/xtls/xray-core/app/router"
	"github.com/xtls/xray-core/common/strmatcher"
)

func Search(gin *router.GeoSiteList, domain string) []string {
	var result []string
	for _, x := range gin.Entry {
		rootMatcher := &strmatcher.MatcherGroup{}
		for _, y := range x.Domain {
			domainType := strmatcher.Type(y.Type)
			matcher, err := domainType.New(y.Value)
			if err != nil {
				return result
			}
			rootMatcher.Add(matcher)
		}
		matchResult := rootMatcher.Match(domain)
		if len(matchResult) > 0 {
			result = append(result, x.CountryCode)
		}
	}

	return result
}
