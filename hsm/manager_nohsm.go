//go:build !hsm
// +build !hsm

package hsm

import (
	"context"
	"sync"

	"github.com/ory/hydra/driver/config"
	"github.com/ory/x/logrusx"

	"github.com/pkg/errors"

	"github.com/ory/hydra/jwk"

	"gopkg.in/square/go-jose.v2"
)

type Context interface {
}

type KeyManager struct {
	jwk.Manager
	sync.RWMutex
	Context
	KeySetPrefix string
}

var ErrOpSysNotSupported = errors.New("Hardware Security Module is not supported on this platform.")

func NewContext(c *config.Provider, l *logrusx.Logger) Context {
	l.Fatalf("Hardware Security Module is not supported on this platform.")
	return nil
}

func NewKeyManager(hsm Context, config *config.Provider) *KeyManager {
	return nil
}

func (m *KeyManager) GenerateAndPersistKeySet(_ context.Context, set, kid, alg, use string) (*jose.JSONWebKeySet, error) {
	return nil, errors.WithStack(ErrOpSysNotSupported)
}

func (m *KeyManager) GetKey(_ context.Context, set, kid string) (*jose.JSONWebKeySet, error) {
	return nil, errors.WithStack(ErrOpSysNotSupported)
}

func (m *KeyManager) GetKeySet(_ context.Context, set string) (*jose.JSONWebKeySet, error) {
	return nil, errors.WithStack(ErrOpSysNotSupported)
}

func (m *KeyManager) DeleteKey(_ context.Context, set, kid string) error {
	return errors.WithStack(ErrOpSysNotSupported)
}

func (m *KeyManager) DeleteKeySet(_ context.Context, set string) error {
	return errors.WithStack(ErrOpSysNotSupported)
}

func (m *KeyManager) AddKey(_ context.Context, _ string, _ *jose.JSONWebKey) error {
	return errors.WithStack(ErrOpSysNotSupported)
}

func (m *KeyManager) AddKeySet(_ context.Context, _ string, _ *jose.JSONWebKeySet) error {
	return errors.WithStack(ErrOpSysNotSupported)
}

func (m *KeyManager) UpdateKey(_ context.Context, _ string, _ *jose.JSONWebKey) error {
	return errors.WithStack(ErrOpSysNotSupported)
}

func (m *KeyManager) UpdateKeySet(_ context.Context, _ string, _ *jose.JSONWebKeySet) error {
	return errors.WithStack(ErrOpSysNotSupported)
}
