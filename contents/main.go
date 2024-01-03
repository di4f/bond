package contents

type Type string



const (
	// Using the UTF-8 by default.
	Unknown Type = "application/octet-stream"
	Binary = Unknown
	Plain Type = "text/plain; charset=utf-8"
	Css Type = "text/css"
	Html Type = "text/html"
	Json Type = "application/json"
)
