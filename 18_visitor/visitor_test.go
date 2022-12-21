package visitor

import "testing"

func TestCompressor_Visit(t *testing.T) {
	c := &Compressor{}
	pdf := &PdfFile{}
	ppt := &PPTFile{}
	pdf.Accept(c)
	ppt.Accept(c)
}
