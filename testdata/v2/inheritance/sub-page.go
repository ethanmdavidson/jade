// Code generated by "jade.go"; DO NOT EDIT.

package jade

import (
	"github.com/valyala/bytebufferpool"
)

func SubPage(title string, buffer *bytebufferpool.ByteBuffer) {

	buffer.WriteString(`<html><head><title>My Site - `)
	WriteEscString(title, buffer)
	buffer.WriteString(`</title><script src="/jquery.js"></script></head><body><p>no</p><div class="sidebar"><p>thing</p></div><div id="footer"><p>some footer content</p></div></body></html>`)

}