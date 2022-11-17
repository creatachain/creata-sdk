package types

import (
	"fmt"

	sdk "github.com/creatachain/creata-sdk/types"
)

// The router is a map from module name to the ICPModule
// which contains all the module-defined callbacks required by ICS-26
type Router struct {
	routes map[string]ICPModule
	sealed bool
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]ICPModule),
	}
}

// Seal prevents the Router from any subsequent route handlers to be registered.
// Seal will panic if called more than once.
func (rtr *Router) Seal() {
	if rtr.sealed {
		panic("router already sealed")
	}
	rtr.sealed = true
}

// Sealed returns a boolean signifying if the Router is sealed or not.
func (rtr Router) Sealed() bool {
	return rtr.sealed
}

// AddRoute adds ICPModule for a given module name. It returns the Router
// so AddRoute calls can be linked. It will panic if the Router is sealed.
func (rtr *Router) AddRoute(module string, cbs ICPModule) *Router {
	if rtr.sealed {
		panic(fmt.Sprintf("router sealed; cannot register %s route callbacks", module))
	}
	if !sdk.IsAlphaNumeric(module) {
		panic("route expressions can only contain alphanumeric characters")
	}
	if rtr.HasRoute(module) {
		panic(fmt.Sprintf("route %s has already been registered", module))
	}

	rtr.routes[module] = cbs
	return rtr
}

// HasRoute returns true if the Router has a module registered or false otherwise.
func (rtr *Router) HasRoute(module string) bool {
	_, ok := rtr.routes[module]
	return ok
}

// GetRoute returns a ICPModule for a given module.
func (rtr *Router) GetRoute(module string) (ICPModule, bool) {
	if !rtr.HasRoute(module) {
		return nil, false
	}
	return rtr.routes[module], true
}
