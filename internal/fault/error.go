package fault

type Kind int

const (
	KindValidation Kind = iota
	KindNotFound
	KindInternal
)

type Error struct {
	Kind    Kind
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func ValidationError(message string) Error {
	return Error{Kind: KindValidation, Message: message}
}

func NotFoundError(message string) Error {
	return Error{Kind: KindNotFound, Message: message}
}

func InternalError(message string) Error {
	return Error{Kind: KindInternal, Message: message}
}
