// Code generated by "jade.go"; DO NOT EDIT.

package jade

import (
	"github.com/valyala/bytebufferpool"
)

func tpl_layout(title string, buffer *bytebufferpool.ByteBuffer) {

	buffer.WriteString(`<html><head><title>My Site - `)
	WriteEscString(title, buffer)
	buffer.WriteString(`</title><script src="/jquery.js"></script></head><body><div id="footer"><p>some footer content</p></div></body></html>`)

}