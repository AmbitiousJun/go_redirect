package https

import "strings"

var (
	TextContents = map[string]struct{}{
		"text/plain":                        {},
		"text/html":                         {},
		"text/css":                          {},
		"text/javascript":                   {},
		"text/xml":                          {},
		"application/json":                  {},
		"application/xml":                   {},
		"application/javascript":            {},
		"application/xhtml+xml":             {},
		"application/x-www-form-urlencoded": {},
	}
)

// IsTextContent 校验 contentType 是否是文本格式, 大小写不区分
func IsTextContent(contentType string) bool {
	contentType = strings.TrimSpace(contentType)
	contentType = strings.ToLower(contentType)
	for content := range TextContents {
		if strings.HasPrefix(contentType, content) {
			return true
		}
	}
	return false
}
