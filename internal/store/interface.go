package store

type GeneralStoreInterface interface {
	Put(code, url string) error
	Get(code string) (string, error)
}
