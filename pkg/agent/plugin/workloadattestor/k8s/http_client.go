//go:generate mockgen -source=$GOFILE -destination=../../../test/mock/common/http/http_client.go -package=http_client_mock

package k8s

import "net/http"

type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}
