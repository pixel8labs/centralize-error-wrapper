package centricerrorwrapper

func (e *errWrapper) GetCode() string {
	return e.Code
}

func (e *errWrapper) GetHTTPStatusCode() int {
	return e.StatusCode
}

func (e *errWrapper) GetMessage() string {
	return e.Message
}

func (e *errWrapper) GetDevMessage() string {
	return e.DevMessage
}

func (e *errWrapper) GetIsRetryable() bool {
	return e.IsRetryable
}

func (e *errWrapper) GetErr() error {
	return e.error
}
