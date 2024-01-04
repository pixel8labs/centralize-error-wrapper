package centricerrorwrapper

import (
	"errors"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type (
	ErrID               string
	errHolderWrapper    map[ErrID]errWrapper
	localeHolderWrapper map[ErrID]string

	errWrapper struct {
		Code        string `json:"code"`
		IsRetryable bool   `json:"is_retryable"`
		StatusCode  int    `json:"http_status_code"`
		DevMessage  string `json:"dev_message"`
		Message     string `json:"message"`

		error
	}
	errCentralizeWrapper struct {
		bundler *i18n.Bundle
		errHolderWrapper
	}

	ErrOptions map[string]interface{}
)

var (
	errHolder *errCentralizeWrapper
	// Default error value, throw if there's no err-mapping data found.
	defaultErr = errWrapper{
		error:       errors.New("Unexpected flow: Error occured"),
		Code:        "9999",
		IsRetryable: false,
		StatusCode:  500,
		DevMessage:  "Error key-definition is not found",
		Message:     "Error occured",
	}
)
