package utils

type ReasonableError struct {
	Reason  error
	Message string
}

func (err ReasonableError) Error() string {
	if err.Reason != nil {
		return err.Message + ":\n" + err.Reason.Error()
	}
	return err.Message
}
