package centricerrorwrapper

const (
	errPath    = "./error"
	localePath = "./locale"
)

// Default value of data.
const (
	defaultLang   = "en"
	defaultString = ""

	httpHeaderAcceptLanguage = "Accept-Lang"
)

// List of errors key.
// Identifier error based on key.
const (
	// Generic errors.
	// Code: 1...x.
	ErrIDMarshall   ErrID = "ErrIDMarshal"
	ErrIDUnmarshall ErrID = "ErrIDUnmarshal"

	// Specific errors.
)
