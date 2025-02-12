package latex

import (
	"github.com/ag7if/go-files"
)

type Package interface {
	Commander
	WriteStyToCache(dir files.Directory) error
}
