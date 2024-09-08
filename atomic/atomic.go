package atomic

import (
	"errors"
	"sync"

	"github.com/grafana/sobek"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

type (
	// RootModule is the global module instance that will create module
	// instances for each VU.
	RootModule struct {
		counters sync.Map
	}

	// ModuleInstance represents an instance of the atomic module for every VU.
	ModuleInstance struct {
		vu       modules.VU
		counters *sync.Map
	}
)

var (
	_ modules.Module   = &RootModule{}
	_ modules.Instance = &ModuleInstance{}
)

// New returns a pointer to a new RootModule instance.
func New() *RootModule {
	return &RootModule{
		counters: sync.Map{},
	}
}

// NewModuleInstance implements the modules.Module interface to return
// a new instance for each VU.
func (rm *RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	mi := &ModuleInstance{
		vu:       vu,
		counters: &rm.counters,
	}

	return mi
}

// Exports returns the exports of the atomic module.
func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]interface{}{
			"Counter": mi.newCounter,
		},
	}
}

// newCounter is the constructor for the Counter class.
// it either returns a goja object wrapping the counter or throws an error.
// the counter either comes from the cache or is created.
func (mi *ModuleInstance) newCounter(c sobek.ConstructorCall) *sobek.Object {
	rt := mi.vu.Runtime()

	id := c.Argument(0).ToString().String()
	if id == "" {
		common.Throw(rt, errors.New("id is required, it must be a non-empty string"))
	}

	// try to get the counter from the cache/or create a new one
	raw, _ := mi.counters.LoadOrStore(id, &counter{})

	cnt, ok := raw.(*counter)
	if !ok {
		common.Throw(rt, errors.New("failed to cast counter"))
	}

	obj := rt.NewObject()
	must(rt, obj.DefineDataProperty(
		"add", rt.ToValue(cnt.Add), sobek.FLAG_FALSE, sobek.FLAG_FALSE, sobek.FLAG_TRUE))
	must(rt, obj.DefineDataProperty(
		"val", rt.ToValue(cnt.Val), sobek.FLAG_FALSE, sobek.FLAG_FALSE, sobek.FLAG_TRUE))
	must(rt, obj.DefineDataProperty(
		"inc", rt.ToValue(cnt.Inc), sobek.FLAG_FALSE, sobek.FLAG_FALSE, sobek.FLAG_TRUE))
	must(rt, obj.DefineDataProperty(
		"dec", rt.ToValue(cnt.Dec), sobek.FLAG_FALSE, sobek.FLAG_FALSE, sobek.FLAG_TRUE))

	return obj
}

// must is a small helper that will panic if err is not nil.
func must(rt *sobek.Runtime, err error) {
	if err != nil {
		common.Throw(rt, err)
	}
}
