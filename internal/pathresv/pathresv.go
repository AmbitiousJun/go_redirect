package pathresv

import (
	"fmt"
	"strings"
)

// Handle 路径映射, 返回最终需要重定向的路径
//
//	idx: 使用第几个模板
//	url: 请求参数 url, 例如 /a/b/c, 最终 a, b, c 会替换到模板中
func Handle(idx int, url string) (string, error) {
	var tpl *template
	var ok bool
	if tpl, ok = cloneTemplate(idx); !ok {
		return "", fmt.Errorf("匹配不到模板, idx: %d", idx)
	}
	vars := strings.Split(strings.Trim(url, "/"), "/")
	for i := 0; i < len(vars); i++ {
		if vars[i] == "" {
			vars = append(vars[:i], vars[i+1:]...)
			i--
		}
	}
	if len(vars) > tpl.maxVar || len(vars) < tpl.minVar {
		return "", fmt.Errorf("变量个数不合法, min: %d, max: %d", tpl.minVar, tpl.maxVar)
	}
	j := len(vars) - 1
	dftValCoverNum := len(vars) - tpl.minVar
	for i := len(tpl.tplSegments) - 1; i >= 0; i-- {
		if j < 0 {
			break
		}
		curSeg := tpl.tplSegments[i]
		if !curSeg.hasDefault {
			curSeg.value = vars[j]
			j--
			continue
		}
		if dftValCoverNum > 0 {
			curSeg.value = vars[j]
			j--
			dftValCoverNum--
		}
	}
	res := strings.Builder{}
	for _, t := range tpl.segments {
		res.WriteString(t.value)
	}
	return res.String(), nil
}
