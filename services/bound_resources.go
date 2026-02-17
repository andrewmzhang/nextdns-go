package services

type HasInit[BoundResourcePtr any] interface {
	InitBoundResource(nextDNSClient *NextDNSClient, pathFmt string, pathArgs map[string]interface{}, err error) BoundResourcePtr
}

type BoundResource interface {
	GetError() error
	GetNextDNSClient() *NextDNSClient
	GetBoundPath() (string, error)
}
