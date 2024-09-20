package lib

import "github.com/microcosm-cc/bluemonday"

var (
	// no you ain't getting any lfi
	LFI_filters = []string{"..", "\\"}
	// and no you ain't gettin any xss
	XSS_sanitizer = bluemonday.UGCPolicy()
)
