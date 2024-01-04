package centricerrorwrapper

func NewCentralizeErrors() error {
	var (
		holderErr    errHolderWrapper
		holderLocale localeHolderWrapper
		err          error
	)

	err = holderErr.normalizeErrWrapper(errPath)
	if err != nil {
		return err
	}

	bundle, err := holderLocale.normalizeLocaleWrapper(localePath)
	if err != nil {
		return err
	}

	// Passing to value.
	errHolder = &errCentralizeWrapper{
		errHolderWrapper: holderErr,
		bundler:          bundle,
		lang:             defaultLang,
	}

	return nil
}
