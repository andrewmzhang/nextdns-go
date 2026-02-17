package services

type BoundResource[T any] struct {
	nextDNSClient *NextDNSClient
	path          string
}

type HasBind[T any] interface {
	SetBind(nextDNSClient *NextDNSClient, bindpath string, err error) T
}
