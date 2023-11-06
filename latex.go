package latex

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/ag7if/go-files"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type LaTeXer interface {
	LaTeX() string
}

type Compiler struct {
	assetDir string
	buildDir string
	logger   *zerolog.Logger
}

func NewCompiler(assetDir, buildDir string, logger *zerolog.Logger) Compiler {
	return Compiler{
		assetDir: assetDir,
		buildDir: buildDir,
		logger:   logger,
	}
}

func (lc Compiler) GenerateLaTeX(latex LaTeXer, outputFile files.File, assets []string) error {
	for _, asset := range assets {
		assetFile, err := files.NewFile(filepath.Join(lc.assetDir, asset), lc.logger)
		if err != nil {
			lc.logger.Warn().Err(err).Str("file", asset).Msg("failed to acquire reference to asset")
		}

		_, err = assetFile.Copy(filepath.Join(lc.buildDir, asset))
		if err != nil {
			// TODO: React to whether this build asset has already been copied
			lc.logger.Warn().Err(err).Str("file", asset).Msg("failed to copy build asset")
		}
	}

	texSourceFileName := fmt.Sprintf("%s.tex", outputFile.Base())
	texSource, err := files.NewFile(filepath.Join(lc.buildDir, texSourceFileName), lc.logger)
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
	texSourceFile, err := files.NewFile(filepath.Join(lc.buildDir, fmt.Sprintf("%s.tex", outputFile.Base())), lc.logger)

	// First run
	cmd := exec.Command("pdflatex", "-halt-on-error", texSourceFile.Name())
	cmd.Dir = texSourceFile.Dir()

	err = cmd.Run()
	if err != nil {
		return errors.WithStack(err)
	}

	// Second run (pdflatex usually needs two runs to get formatting right)
	cmd = exec.Command("pdflatex", "-halt-on-error", texSourceFile.Name())
	cmd.Dir = texSourceFile.Dir()

	cachedBuildfile, err := files.NewFile(filepath.Join(lc.buildDir, fmt.Sprintf("%s.pdf", outputFile.Base())), lc.logger)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = cachedBuildfile.Move(outputFile.Dir())
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
