package centricerrorwrapper

import (
	"encoding/json"
	"io/fs"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func (e *localeHolderWrapper) normalizeWrapper(path string) (*i18n.Bundle, error) {
	var (
		err    error
		bundle = i18n.NewBundle(language.English)
	)

	// Init locale as json.
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	err = filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			// Handle the error from filepath.WalkDir.
			return err
		}

		if d.IsDir() {
			// It's a directory, continue walking.
			return nil
		}

		_, err = bundle.LoadMessageFile(p)
		if err != nil {
			// Handle the error from processing the JSON file.
			return err
		}

		return nil
	})

	// Check for errors from filepath.WalkDir.
	if err != nil {
		return bundle, err
	}

	return bundle, nil
}
