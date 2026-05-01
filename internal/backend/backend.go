package backend

import "github.com/lucasmcclean/ghost/internal/target"

type Backend interface {
	Name() string
	Enter(target target.Target, command []string) error
}
