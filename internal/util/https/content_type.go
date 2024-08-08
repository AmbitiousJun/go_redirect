package https

import "strings"

var (
	ProxableContents = map[string]struct{}{
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
		"image/jpeg":                        {},
		"image/jpg":                         {},
		"image/png":                         {},
	}
)

// IsProxableContent 校验 contentType 是否是可代理格式, 大小写不区分
func IsProxableContent(contentType string) bool {
	contentType = strings.TrimSpace(contentType)
	contentType = strings.ToLower(contentType)
	for content := range ProxableContents {
		if strings.HasPrefix(contentType, content) {
			return true
		}
	}
	return false
}
