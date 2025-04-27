package application

type UseCase[P any, R any] interface {
	Execute(params P) (result R, err error)
}
