package enum

type HTTPResponseStatus string

const (
	HTTPResponseStatusSuccess HTTPResponseStatus = "success"
	HTTPResponseStatusFailed  HTTPResponseStatus = "failed"
)
