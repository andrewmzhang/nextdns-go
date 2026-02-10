package services

type BoundResource[T any] struct {
	nextDNSClient *NextDNSClient
	path          string
}

type HasBind[T any] interface {
	SetBind(nextDNSClient *NextDNSClient, bindpath string) T
}

// func (r *BoundResource[T]) Get(ctx context.Context) (*T, error) {
// 	var result DataWrapper[T]
// 	resp, err := r.nextDNSClient.nextDNSClient.R().SetContext(ctx).
// 		SetResult(&result).
// 		Get(r.path)
//
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = handleResponse(resp, &result)
// 	return &result.Data, err
// }
//
// func (r *BoundResource[T]) Update(ctx context.Context, payload any) error {
// 	var result DataWrapper[T]
// 	resp, err := r.nextDNSClient.nextDNSClient.R().SetContext(ctx).
// 		SetBody(payload).
// 		Patch(r.path)
//
// 	err = handleResponse(resp, &result)
// 	return err
// }
//
// func (r *BoundResource[T]) Delete(ctx context.Context) error {
// 	var result DataWrapper[T]
// 	resp, err := r.nextDNSClient.nextDNSClient.R().SetContext(ctx).
// 		Delete(r.path)
// 	if err != nil {
// 		return err
// 	}
// 	err = handleResponse(resp, &result)
// 	return err
// }
