package usrctx

import (
	"errors"
	"sync"
)

var (
	// ErrAttrExist represents errors that there is already attribute with specified name.
	ErrAttrExist = errors.New("Attribute already exists with this name")
	// ErrNoAttr represents errors that there is not attribute with specified name.
	ErrNoAttr = errors.New("There is no attribute with thsi name")
)

// Context represents the wrapper
type Context interface {
	SetAttribute(name string, attr interface{})
	Attribute(name string) (interface{}, error)
	Clone() Context
}

// ServerContext base Context will all required methods.
//
// It provides access to all required attrs e.g. chat structure and
// others.
type ServerContext struct {
	attrs map[string]interface{}
	mutex sync.RWMutex
}

// NewContext generates new context which can be provided into server.
func NewContext() Context {
	attrs := make(map[string]interface{})
	return &ServerContext{attrs: attrs}
}

// SetAttribute method allow to add new attribute into context.
//
// When attribute already exists in context operation will return error.
// All attrs when created then they've to be updated, never replaced with
// other.
func (ctx *ServerContext) SetAttribute(name string, attr interface{}) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	ctx.attrs[name] = attr
}

// Attribute returns attribute with specified name.
//
// When there is no attribute with specified name method will return error
// which describes the issue.
func (ctx *ServerContext) Attribute(name string) (interface{}, error) {
	ctx.mutex.RLock()
	defer ctx.mutex.RUnlock()

	attr, ok := ctx.attrs[name]
	if !ok {
		return nil, ErrNoAttr
	}

	return attr, nil
}

// Clone returns deep copy of the context.
//
// That copy is used to provide context to all the users with their data.
func (ctx *ServerContext) Clone() Context {
	attrs := make(map[string]interface{})

	for k, v := range ctx.attrs {
		attrs[k] = v
	}

	return &ServerContext{
		attrs: attrs,
	}
}
