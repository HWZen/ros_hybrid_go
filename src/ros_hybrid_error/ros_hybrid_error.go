package ros_hybrid_go

type ros_hybrid_error struct{
	error_message string
}

func (e ros_hybrid_error) Error() string {
	return e.error_message
}

func NewError(message string) ros_hybrid_error {
	return ros_hybrid_error{message}
}