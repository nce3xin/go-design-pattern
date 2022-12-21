package visitor

import "fmt"

type IVisitor interface {
	Visit(file IResourceFile) error
}

type IResourceFile interface {
	Accept(IVisitor) error
}
type PdfFile struct {
	path string
}

func (p *PdfFile) Accept(visitor IVisitor) error {
	return visitor.Visit(p)
}

type PPTFile struct {
	path string
}

func (P *PPTFile) Accept(visitor IVisitor) error {
	return visitor.Visit(P)
}

// Compressor 实现压缩功能
type Compressor struct{}

func (c *Compressor) Visit(file IResourceFile) error {
	switch t := file.(type) {
	case *PdfFile:
		return c.VisitPdfFile(t)
	case *PPTFile:
		return c.VisitPPTFile(t)
	default:
		return fmt.Errorf("no such file type %v\n", file)
	}
}

func (c *Compressor) VisitPdfFile(file *PdfFile) error {
	fmt.Printf("this is pdf file.\n")
	return nil
}

func (c *Compressor) VisitPPTFile(file *PPTFile) error {
	fmt.Printf("this is ppt file.\n")
	return nil
}
