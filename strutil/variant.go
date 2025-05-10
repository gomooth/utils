package strutil

import (
	"strings"
	"unicode"
)

// Snake 转成蛇形字符串。 XxYy to xx_yy , XxYY to xx_yy
func Snake(s string) string {
	var builder strings.Builder
	builder.Grow(len(s) * 2) // Pre-allocate space for efficiency

	prevLower := false
	prevUpper := false

	for i, r := range s {
		currentUpper := unicode.IsUpper(r)

		// 添加下划线条件：
		// 1. 当前为大写，之前为小写（Aa -> a_a）
		// 2. 当前为大写，下一个为小写（HTMLElement -> html_element）
		if i > 0 {
			if currentUpper && (prevLower || (prevUpper && i < len(s)-1 && unicode.IsLower(rune(s[i+1])))) {
				builder.WriteByte('_')
			}
		}

		builder.WriteRune(unicode.ToLower(r))
		prevLower = !currentUpper
		prevUpper = currentUpper
	}

	return builder.String()
}

// Camel 转成驼峰字符串。 xx_yy to XxYy
func Camel(s string) string {
	var builder strings.Builder
	builder.Grow(len(s))

	nextUpper := false
	firstChar := true

	for _, r := range s {
		if r == '_' {
			nextUpper = true
			continue
		}

		if nextUpper || firstChar {
			builder.WriteRune(unicode.ToUpper(r))
			nextUpper = false
			firstChar = false
		} else {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
