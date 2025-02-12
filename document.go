package latex

import (
	"github.com/ag7if/go-files"
	"github.com/pkg/errors"
)

type Document interface {
	LaTeXer
	AssetDir() *files.Directory
	Assets() []files.File
	BuildEngine() BuildEngine
	IncludePackage(pkg Package, args ...string)
	WriteClass() error
	WritePackages() error
}

func BuildDocument(document Document, outputFile files.File, buildDir *files.Directory) error {
	c := NewCompiler(document.BuildEngine(), document.AssetDir(), buildDir)

	inputFile, err := c.GenerateLaTeX(document, outputFile, document.Assets())
	if err != nil {
		return errors.WithStack(err)
	}

	err = c.CompileLaTeX(inputFile, outputFile)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
