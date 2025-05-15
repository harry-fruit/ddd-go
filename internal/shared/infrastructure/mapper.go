package sharedinfra

type Mapper[D any, I any] interface {
	ToDomain(I) D
	ToDB(D) I
}
