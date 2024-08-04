package pathresv

import (
	"go_redirect/internal/config"
)

const (
	Prefix = "${"
	Suffix = "}"
)

// tpls 程序初始化时, 先从配置文件中转换模板
var tpls map[int]*template

// template url 模板
type template struct {
	raw         string     // 原始 url
	segments    []*segment // 根据模板划分成的分片数组
	tplSegments []*segment // 需要进行过模板替换的分片数组
	maxVar      int        // 最多可传递的变量
	minVar      int        // 最少需传递的变量
}

// segment 每一个路径段
type segment struct {
	isTemplate bool
	hasDefault bool
	value      string
}

// clone 拷贝自身, 返回一个新对象
func (s *segment) clone() *segment {
	return &segment{
		isTemplate: s.isTemplate,
		hasDefault: s.hasDefault,
		value:      s.value,
	}
}

// cloneTemplate 拷贝模板
func cloneTemplate(idx int) (*template, bool) {
	var origin *template
	var ok bool
	if origin, ok = tpls[idx]; !ok {
		return nil, false
	}
	dest := &template{
		raw:         origin.raw,
		minVar:      origin.minVar,
		maxVar:      origin.maxVar,
		segments:    make([]*segment, 0),
		tplSegments: make([]*segment, 0),
	}
	for _, seg := range origin.segments {
		newSeg := seg.clone()
		dest.segments = append(dest.segments, newSeg)
		if newSeg.isTemplate {
			dest.tplSegments = append(dest.tplSegments, newSeg)
		}
	}
	return dest, true
}

// InitTemplates 将配置文件中的链接解析成模板对象
func InitTemplates() {
	groups := config.C.Groups
	tpls = make(map[int]*template)
	for idx, group := range groups {
		segments := make([]*segment, 0)
		tplSegments := make([]*segment, 0)
		minVar, maxVar := 0, 0
		preIdx, dftValStartPos := 0, 0
		for curIdx := 0; curIdx < len(group); curIdx++ {
			// 统计模板
			if group[curIdx:curIdx+1] == Suffix && dftValStartPos > preIdx {
				seg := &segment{isTemplate: true, value: group[dftValStartPos:curIdx]}
				seg.hasDefault = seg.value != ""
				segments = append(segments, seg)
				tplSegments = append(tplSegments, seg)
				maxVar++
				if !seg.hasDefault {
					minVar++
				}
				preIdx = curIdx + 1
				continue
			}
			// 统计非模板
			if curIdx <= len(group)-len(Prefix) && group[curIdx:curIdx+len(Prefix)] == Prefix {
				seg := &segment{isTemplate: false, value: group[preIdx:curIdx]}
				segments = append(segments, seg)
				preIdx = curIdx
				dftValStartPos = preIdx + len(Prefix)
				curIdx += len(Prefix) - 1
			}
		}
		// 尾串处理
		if preIdx < len(group) {
			seg := &segment{isTemplate: false, value: group[preIdx:]}
			segments = append(segments, seg)
		}
		tpls[idx+1] = &template{
			raw:         group,
			segments:    segments,
			tplSegments: tplSegments,
			maxVar:      maxVar,
			minVar:      minVar,
		}
	}
}
