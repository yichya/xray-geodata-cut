# xray-geodata-cut

Cut unneeded data from geoip.dat or geosite.dat

```
Usage of xray-geodata-cut:
  -in string
        Path to GeoIP file
  -keep string
        GeoIP or GeoSite codes to keep (private is always kept for GeoIP) (default "cn,private,geolocation-!cn")
  -out string
        Path to processed file
  -show
        Print codes in GeoIP or GeoSite file
  -trimipv6
        Trim all IPv6 ranges in GeoIP file
  -type string
        GeoIP (geoip) or GeoSite (geosite)
```
