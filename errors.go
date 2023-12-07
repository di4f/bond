package bond

import (
	"errors"
)

var (
	DupDefErr = errors.New("duplicate route define")
	UnknownContentTypeErr = errors.New("unknown content type")
)
