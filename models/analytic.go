package models

type AnalyticStatus struct {
	Status  string `json:"status"`
	Queries int    `json:"queries"`
}

type AnalyticDomain struct {
	Domain  string `json:"domain"`
	Root    string `json:"root"`
	Queries int    `json:"queries"`
}

type AnalyticReason struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Queries int    `json:"queries"`
}

type AnalyticIpNetwork struct {
	Cellular bool   `json:"cellular"`
	Vpn      bool   `json:"vpn"`
	Isp      string `json:"isp"`
	Asn      string `json:"asn"`
}

type AnalyticIpGeo struct {
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	CountryCode string  `json:"countryCode"`
	Country     string  `json:"country"`
	City        string  `json:"city"`
}

type AnalyticIp struct {
	IP      string            `json:"ip"`
	Network AnalyticIpNetwork `json:"network"`
	Geo     AnalyticIpGeo     `json:"geo"`
	Queries int               `json:"queries"`
}

type AnalyticDevice struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Model   string `json:"model"`
	LocalIp string `json:"localIp"`
	Queries int    `json:"queries"`
}

type AnalyticProtocol struct {
	Protocol string `json:"protocol"`
	Queries  int    `json:"queries"`
}

type AnalyticQueryType struct {
	Type    int    `json:"type"`
	Name    string `json:"name"`
	Queries int    `json:"queries"`
}

type AnalyticIpVersion struct {
	Version string `json:"version"`
	Queries int    `json:"queries"`
}

type AnalyticDnssec struct {
	Validated bool `json:"validated"`
	Queries   int  `json:"queries"`
}

type AnalyticEncryption struct {
	Encrypted bool `json:"encrypted"`
	Queries   int  `json:"queries"`
}

type AnalyticsDestinationCountries struct {
	Code    string   `json:"code"`
	Domains []string `json:"domains"`
	Queries int      `json:"queries"`
}

type AnalyticDestinationGafam struct {
	Company string `json:"company"`
	Queries int    `json:"queries"`
}
