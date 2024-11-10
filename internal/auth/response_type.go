package auth

type ResponseType int64

const (
	ResponseTypeToken ResponseType = iota
)

func (rT ResponseType) String() string {
	return map[ResponseType]string{
		ResponseTypeToken: "token",
	}[rT]
}
