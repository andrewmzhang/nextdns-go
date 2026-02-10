package services

import (
	"context"
	"encoding/json"
	"fmt"
	"iter"

	"github.com/andrewmzhang/nextdns-go/models"
	"resty.dev/v3"
)

type AnalyticsBaseQueryConfig struct {
	// Query parameters
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
	Device string `json:"device,omitempty"`
	// Time series parameter
	Interval  string `json:"interval,omitempty"`
	Alignment string `json:"alignment,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
	Partials  string `json:"partials,omitempty"`
}

type TimeSeriesService[T any, C any] struct {
	path          string
	nextDNSClient *NextDNSClient
}

func fetchPage[T any, Config any](
	ctx context.Context,
	resty *resty.Client,
	path string,
	isTimeSeries bool,
	config Config,
	cursor string,
) (*DataListMetaWrapper[T], error) {
	// Convert config to map[string]string
	b, _ := json.Marshal(config)
	var tmp map[string]interface{}
	_ = json.Unmarshal(b, &tmp)
	configMap := make(map[string]string)
	for k, v := range tmp {
		configMap[k] = fmt.Sprint(v)
	}

	var results DataListMetaWrapper[T]

	req := resty.R().
		SetContext(ctx).
		SetResult(&results).
		SetQueryParams(configMap)

	if cursor != "" {
		req.SetQueryParam("cursor", cursor)
	}

	finalPath := path
	if isTimeSeries {
		finalPath = finalPath + ";series"
	}
	resp, err := req.Get(path)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	err = handleResponse(resp, &results.Data)
	return &results, err
}

type AnalyticIterator[Service any, Config any, T any] struct {
	nextDNSClient *NextDNSClient
	path          string
	isTimeSeries  bool
	ctx           context.Context
	service       *Service
	config        Config

	buffer []T
	cursor string
	err    error
	done   bool
	index  int
}

func (s *TimeSeriesService[T, C]) Query(ctx context.Context, config C, isTimeSeries bool) *AnalyticIterator[TimeSeriesService[T, C], C, T] {
	return &AnalyticIterator[TimeSeriesService[T, C], C, T]{
		nextDNSClient: s.nextDNSClient,
		path:          s.path,
		isTimeSeries:  isTimeSeries,
		ctx:           ctx,
		service:       s,
		config:        config,
	}
}

func (it *AnalyticIterator[Service, Config, T]) Next() bool {
	// Exit immediately if there's an error
	if it.err != nil {
		return false
	}

	// Still items left in buffer
	if it.index < len(it.buffer) {
		it.index++
		return true
	}

	// Check if cursor is not nil. Make an exception if index == 0 (first run of Next())
	if it.index != 0 && it.cursor == "" {
		return false
	}

	// Cursor is not nil, we need to fetch next page
	resp, err := fetchPage[T, Config](it.ctx, it.nextDNSClient.restyClient, it.path, it.isTimeSeries, it.config, it.cursor)
	if err != nil {
		it.err = err
		fmt.Println("Error fetching next DNS page")
		return false
	}

	// No more data
	if len(resp.Data) == 0 {
		it.done = true
		fmt.Println("No more data to read")
		return false
	}

	fmt.Println(len(resp.Data))
	it.buffer = resp.Data
	it.cursor = resp.Meta.Pagination.Cursor
	it.index = 1

	if it.cursor == "" {
		it.done = true
	}

	return true
}

func (it *AnalyticIterator[Service, Config, T]) Value() T {
	return it.buffer[it.index-1]
}

func (it *AnalyticIterator[Service, Config, T]) Err() error {
	return it.err
}

func (c *TimeSeriesService[T, C]) ListAll(ctx context.Context, config C, isTimeSeries bool) ([]T, error) {
	it := c.Query(ctx, config, isTimeSeries)

	var out []T
	for it.Next() {
		out = append(out, it.Value())
	}

	if err := it.Err(); err != nil {
		return nil, err
	}

	return out, nil
}

func (c *TimeSeriesService[T, C]) AnalyticsRange(ctx context.Context, config C, isTimeSeries bool) iter.Seq[T] {
	return func(yield func(status T) bool) {
		it := c.Query(ctx, config, isTimeSeries)

		for it.Next() {
			if !yield(it.Value()) {
				return
			}
		}
	}
}

/** Concrete class below **/

type AnalyticsStatusQueryConfig struct {
	AnalyticsBaseQueryConfig
}

func (c *NextDNSClient) AnalyticStatus(profileId string) *TimeSeriesService[models.AnalyticStatus, AnalyticsStatusQueryConfig] {
	return &TimeSeriesService[models.AnalyticStatus, AnalyticsStatusQueryConfig]{
		path:          "/profiles/" + profileId + "/analytics/status",
		nextDNSClient: c,
	}
}

type AnalyticsDomainsQueryConfig struct {
	AnalyticsBaseQueryConfig
}

func (c *NextDNSClient) AnalyticDomains(profileId string) *TimeSeriesService[models.AnalyticDomain, AnalyticsDomainsQueryConfig] {
	return &TimeSeriesService[models.AnalyticDomain, AnalyticsDomainsQueryConfig]{
		path:          "/profiles/" + profileId + "/analytics/status",
		nextDNSClient: c,
	}
}

type AnalyticsReasonsQueryConfig struct {
	AnalyticsBaseQueryConfig
}

func (c *NextDNSClient) AnalyticReasons(profileId string) *TimeSeriesService[models.AnalyticReason, AnalyticsReasonsQueryConfig] {
	return &TimeSeriesService[models.AnalyticReason, AnalyticsReasonsQueryConfig]{
		path:          "/profiles/" + profileId + "/analytics/reasons",
		nextDNSClient: c,
	}
}

type AnalyticsIpsQueryConfig struct {
	AnalyticsBaseQueryConfig
}

func (c *NextDNSClient) AnalyticIps(profileId string) *TimeSeriesService[models.AnalyticIp, AnalyticsIpsQueryConfig] {
	return &TimeSeriesService[models.AnalyticIp, AnalyticsIpsQueryConfig]{
		path:          "/profiles/" + profileId + "/analytics/ips",
		nextDNSClient: c,
	}
}

type AnalyticsDevicesQueryConfig struct {
	AnalyticsBaseQueryConfig
}

func (c *NextDNSClient) AnalyticDevices(profileId string) *TimeSeriesService[models.AnalyticDevice, AnalyticsDevicesQueryConfig] {
	return &TimeSeriesService[models.AnalyticDevice, AnalyticsDevicesQueryConfig]{
		path:          "/profiles/" + profileId + "/analytics/devices",
		nextDNSClient: c,
	}
}

type AnalyticsProtocolsQueryConfig struct {
	AnalyticsBaseQueryConfig
}

func (c *NextDNSClient) AnalyticProtocols(profileId string) *TimeSeriesService[models.AnalyticProtocol, AnalyticsProtocolsQueryConfig] {
	return &TimeSeriesService[models.AnalyticProtocol, AnalyticsProtocolsQueryConfig]{
		path:          "/profiles/" + profileId + "/analytics/protocols",
		nextDNSClient: c,
	}
}

type AnalyticsQueryTypesQueryConfig struct {
	AnalyticsBaseQueryConfig
}

func (c *NextDNSClient) AnalyticQueryTypes(profileId string) *TimeSeriesService[models.AnalyticQueryType, AnalyticsQueryTypesQueryConfig] {
	return &TimeSeriesService[models.AnalyticQueryType, AnalyticsQueryTypesQueryConfig]{
		path:          "/profiles/" + profileId + "/analytics/queryTypes",
		nextDNSClient: c,
	}
}

type AnalyticsIpVersionsQueryConfig struct {
	AnalyticsBaseQueryConfig
}

func (c *NextDNSClient) AnalyticIpVersions(profileId string) *TimeSeriesService[models.AnalyticIpVersion, AnalyticsIpVersionsQueryConfig] {
	return &TimeSeriesService[models.AnalyticIpVersion, AnalyticsIpVersionsQueryConfig]{
		path:          "/profiles/" + profileId + "/analytics/ipVersions",
		nextDNSClient: c,
	}
}

type AnalyticsDnssecQueryConfig struct {
	AnalyticsBaseQueryConfig
}

func (c *NextDNSClient) AnalyticDnssec(profileId string) *TimeSeriesService[models.AnalyticDnssec, AnalyticsDnssecQueryConfig] {
	return &TimeSeriesService[models.AnalyticDnssec, AnalyticsDnssecQueryConfig]{
		path:          "/profiles/" + profileId + "/analytics/dnssec",
		nextDNSClient: c,
	}
}

type AnalyticsEncryptionsQueryConfig struct {
	AnalyticsBaseQueryConfig
}

func (c *NextDNSClient) AnalyticEncryptions(profileId string) *TimeSeriesService[models.AnalyticEncryption, AnalyticsEncryptionsQueryConfig] {
	return &TimeSeriesService[models.AnalyticEncryption, AnalyticsEncryptionsQueryConfig]{
		path:          "/profiles/" + profileId + "/analytics/encryption",
		nextDNSClient: c,
	}
}

type AnalyticsDestinationsCountriesQueryConfig struct {
	AnalyticsBaseQueryConfig
}

func (c *NextDNSClient) AnalyticDesintationsCountrys(profileId string) *TimeSeriesService[models.AnalyticsDestinationCountries, AnalyticsDestinationsCountriesQueryConfig] {
	return &TimeSeriesService[models.AnalyticsDestinationCountries, AnalyticsDestinationsCountriesQueryConfig]{
		path:          "/profiles/" + profileId + "/analytics/destinations?type=countries",
		nextDNSClient: c,
	}
}

type AnalyticsDestinationsGafamQueryConfig struct {
	AnalyticsBaseQueryConfig
}

func (c *NextDNSClient) AnalyticDestination(profileId string) *TimeSeriesService[models.AnalyticDestinationGafam, AnalyticsDestinationsGafamQueryConfig] {
	return &TimeSeriesService[models.AnalyticDestinationGafam, AnalyticsDestinationsGafamQueryConfig]{
		path:          "/profiles/" + profileId + "/analytics/destinations?type=gafam",
		nextDNSClient: c,
	}
}
