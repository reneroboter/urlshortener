package infrastructure

type RepositoryInterface interface {
	Put(code, url string) error
	Get(code string) (string, error)
}
