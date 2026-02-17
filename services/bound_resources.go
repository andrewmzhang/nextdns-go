package services

type HasBind[T any] interface {
	SetBind(nextDNSClient *NextDNSClient, pathFmt string, pathArgs map[string]interface{}, err error) T
}

type BoundResource interface {
	GetError() error
	GetNextDNSClient() *NextDNSClient
	GetBoundPath() (string, error)
}
