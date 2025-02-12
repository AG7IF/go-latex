package latex

import (
	"github.com/ag7if/go-files"
)

type Class interface {
	Commander
	WriteClassToCache(dir files.Directory) error
}
