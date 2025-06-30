package types

type Disposable interface {
	Dispose() error // Dispose releases any resources held by the object
}
