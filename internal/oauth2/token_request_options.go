package oauth2

type TokenOption[T any] func(T) error

type GenericTokenOption = TokenOption[any]
