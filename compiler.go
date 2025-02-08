package latex

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/ag7if/go-files"
	"github.com/pkg/errors"
)

type Compiler struct {
	assetDir string
	buildDir string
}

func NewCompiler(assetDir, buildDir string) Compiler {
	return Compiler{
		assetDir: assetDir,
		buildDir: buildDir,
	}
}

func (lc Compiler) GenerateLaTeX(latex LaTeXer, outputFile files.File, assets []string) error {
	for _, asset := range assets {
		assetFile, err := files.NewFile(filepath.Join(lc.assetDir, asset))
		if err != nil {
			return errors.WithMessagef(err, "failed to acquire reference to asset: %s", asset)
		}

		_, err = assetFile.Copy(lc.buildDir)
		if err != nil {
			// TODO: React to whether this build asset has already been copied
			return errors.WithMessagef(err, "failed to copy build asset: %s", asset)
		}
	}

	texSourceFileName := fmt.Sprintf("%s.tex", outputFile.Base())
	texSource, err := files.NewFile(filepath.Join(lc.buildDir, texSourceFileName))
	if err != nil {
		return errors.WithStack(err)
	}
	err = texSource.WriteString(latex.LaTeX())
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (lc Compiler) CompileLaTeX(outputFile files.File) error {
	texSourceFile, err := files.NewFile(filepath.Join(lc.buildDir, fmt.Sprintf("%s.tex", outputFile.Base())))
	if err != nil {
		return errors.WithStack(err)
	}

	// First run
	cmd := exec.Command("xelatex", "-halt-on-error", texSourceFile.Name())
	cmd.Dir = texSourceFile.Dir()

	err = cmd.Run()
	if err != nil {
		return errors.WithStack(err)
	}

	// Second run (pdflatex usually needs two runs to get formatting right)
	cmd = exec.Command("pdflatex", "-halt-on-error", texSourceFile.Name())
	cmd.Dir = texSourceFile.Dir()

	cachedBuildfile, err := files.NewFile(filepath.Join(lc.buildDir, fmt.Sprintf("%s.pdf", outputFile.Base())))
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = cachedBuildfile.Move(outputFile.Dir())
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
