package centricerrorwrapper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func New(errID ErrID, option ...ErrOptions) errWrapper {
	errVHolder, ok := errHolder.errHolderWrapper[errID]
	if !ok {
		return defaultErr
	}

	errVHolder.Message = errHolder.getMessage(errID, option...)
	return errVHolder
}

func Wrap(err error, errID ErrID, option ...ErrOptions) errWrapper {
	errVHolder, ok := errHolder.errHolderWrapper[errID]
	if !ok {
		return defaultErr
	}

	errVHolder.Message = errHolder.getMessage(errID, option...)
	errVHolder.error = err
	return errVHolder
}

func Unwrap(err error) error {
	u, ok := err.(interface {
		Unwrap() error
	})
	if !ok {
		return nil
	}

	return u.Unwrap()
}

func Cast(err error) *errWrapper {
	var errVHolder errWrapper
	if errors.As(err, &errVHolder) {
		return &errVHolder
	}

	return &errWrapper{error: err}
}

func SetLang(lang string) {
	errHolder.lang = lang
}

// Private function.
// Helper.
func (e *errHolderWrapper) getMessage(errID ErrID, option ...ErrOptions) string {
	var opt ErrOptions
	if len(option) > 0 {
		opt = option[0]
	}

	localizer := i18n.NewLocalizer(errHolder.bundler, errHolder.lang)
	locale, _, err := localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		MessageID:    string(errID),
		TemplateData: opt,
	})

	if err != nil {
		return defaultString
	}

	return locale
}

// Override from base-errs function.
func (e errWrapper) Error() string {
	return fmt.Sprintf(
		"Code: %s, StatusCode: %d, IsRetryable: %t, DevError: %s, Message: %s, Actual err: %v",
		e.Code, e.StatusCode, e.IsRetryable, e.DevMessage, e.Message, e.error,
	)
}

// Override from base-errs function.
// Unwrap is used to make it work with errors.Is, errors.As.
func (e errWrapper) Unwrap() error {
	// Return the inner error.
	return e.error
}

// Helper to construct the errors resource to a given struct.
func (e *errHolderWrapper) normalizeErrWrapper(path string) error {
	var (
		err error
	)

	err = filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			// Handle the error from filepath.WalkDir.
			return err
		}

		if d.IsDir() {
			// It's a directory, continue walking.
			return nil
		}

		err = e.parseErr(p)
		if err != nil {
			// Handle the error from processing the JSON file.
			return err
		}

		return nil
	})

	// Check for errors from filepath.WalkDir.
	if err != nil {
		return err
	}

	return nil
}

func (e *errHolderWrapper) parseErr(filePath string) error {
	raw, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, &e)
	return err
}
