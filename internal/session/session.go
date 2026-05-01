package session

import (
	"fmt"

	"github.com/lucasmcclean/ghost/internal/backend"
	"github.com/lucasmcclean/ghost/internal/target"
)

type State string

const (
	StateCreated State = "created"
	StateRunning State = "running"
	StateExited  State = "exited"
	StateFailed  State = "failed"
)

type Session struct {
	ID      string
	Target  target.Target
	Backend backend.Backend
	Tools   []string

	State State
	Err   error
}

func (s *Session) Start() error {
	if s.State != StateCreated {
		return fmt.Errorf("invalid state transition: %s → running", s.State)
	}

	s.State = StateRunning

	// Phase 1 (future): prepare tools
	if err := s.prepare(); err != nil {
		s.fail(err)
		return err
	}

	// Phase 2: enter runtime
	if err := s.enter(); err != nil {
		s.fail(err)
		return err
	}

	s.State = StateExited
	return nil
}

func (s *Session) prepare() error {
	// Placeholder for:
	// - Nix bundle mounting
	// - PATH modification
	// - environment setup
	return nil
}

func (s *Session) enter() error {
	cmd := []string{"/bin/sh"}

	return s.Backend.Enter(s.Target, cmd)
}

func (s *Session) fail(err error) {
	s.State = StateFailed
	s.Err = err
}
