package exception

type ConflictError struct {
	Message string
}

func (n *ConflictError) Error() string {
	return n.Message
}
