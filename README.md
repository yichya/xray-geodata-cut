# xray-geodata-cut

Cut unneeded data from geoip.dat or geosite.dat, or build geoip.dat from ASNs

```
Usage of xray-geodata-cut:
  -in string
        Path to GeoData file / ASNs split by comma
  -keep string
        GeoIP or GeoSite codes to keep (private is always kept for GeoIP) (default "cn,private,geolocation-!cn")
  -out string
        Path to processed file
  -search string
        Search GeoIP or GeoSite Item
  -show
        Print codes in GeoIP or GeoSite file
  -trimipv6
        Trim all IPv6 ranges in GeoIP file
  -type string
        ASN (asn), GeoIP (geoip) or GeoSite (geosite)

```

ASN information comes from [https://github.com/ipverse/asn-ip/](https://github.com/ipverse/asn-ip/)

Examples for search: 

```
sh-5.1$ go run . -type asn -in 24429,4134 -search 106.124.1.2
AS4134
sh-5.1$ go run . -in /usr/local/share/xray/geoip.dat -type geoip -search 114.114.114.114
CN
sh-5.1$ go run . -in /usr/local/share/xray/geoip.dat -type geoip -search 192.0.2.1
PRIVATE
sh-5.1$ go run . -in /usr/local/share/xray/geoip.dat -type geoip -search 127.0.0.1
PRIVATE
TEST
sh-5.1$ go run . -in /usr/local/share/xray/geosite.dat -type geosite -search bilibili.com
BILIBILI
CN
GEOLOCATION-CN
sh-5.1$ go run . -in /usr/local/share/xray/geosite.dat -type geosite -search baidu.com
BAIDU
CN
GEOLOCATION-CN
sh-5.1$ go run . -in /usr/local/share/xray/geosite.dat -type geosite -search youtube.com
CATEGORY-COMPANIES
GEOLOCATION-!CN
GOOGLE
YOUTUBE
sh-5.1$ go run . -in /usr/local/share/xray/geosite.dat -type geosite -search www.netflix.com
CATEGORY-ENTERTAINMENT
GEOLOCATION-!CN
NETFLIX
```
