package layer

type fnError struct {
	message string
}

func (e fnError) Error() string {
	return e.message
}

func err(message string) fnError {
	return fnError{message: message}
}
