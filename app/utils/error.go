package utils

type StatusError struct {
	Err     error
	Message string
}

func (se StatusError) Error() string {
	return se.Message
}

func (se StatusError) Unwrap() error {
	return se.Err
}
