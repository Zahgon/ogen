package http

import (
	"io"
	"mime/multipart"
	"net/textproto"
	"strings"
)

// MultipartFile is multipart form file.
type MultipartFile struct {
	Name   string
	File   io.Reader
	Size   int64
	Header textproto.MIMEHeader
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string { _ = "STUB: not implemented"; return "" }

// headers generates headers for multipart form file, similar to CreateFormFile, but this function does not
// overwrite Content-Type if it is already set.
func (m MultipartFile) headers(fieldName string) (h textproto.MIMEHeader) {
	_ = "STUB: not implemented"
	return *new(textproto.MIMEHeader)
}

// WriteMultipart writes data from reader to given multipart.Writer as a form file.
func (m MultipartFile) WriteMultipart(fieldName string, w *multipart.Writer) error {
	_ = "STUB: not implemented"
	return nil
}
