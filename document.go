package latex

import (
	"github.com/ag7if/go-files"
	"github.com/pkg/errors"
)

type Style interface {
	AddCommand(command Command, args ...string)
	AddDef(command CustomCommand)
	StyleDef() string
}

type Document interface {
	LaTeXer
	AddAsset(asset string)
	AddInclude(packageName, options string)
	AddStyle(style Style)
	AssetDir() string
	Assets() []string
	BuildDir() string
	SetClass(className, options string)
}

func BuildDocument(document Document, outputFile files.File) error {
	c := NewCompiler(document.AssetDir(), document.BuildDir())

	err := c.GenerateLaTeX(document, outputFile, document.Assets())
	if err != nil {
		return errors.WithStack(err)
	}

	err = c.CompileLaTeX(outputFile)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
