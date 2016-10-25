package compiler

import (
	"io"
	"sync"

	"github.com/mibk/hawk/compiler/internal/hawkc"
)

var mu sync.Mutex

// Compile compiles a Hawk program from src.
func Compile(src io.Reader) (*hawkc.Program, error) {
	mu.Lock()
	defer mu.Unlock()
	return hawkc.Compile(src)
}
