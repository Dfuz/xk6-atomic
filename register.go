// Package atomic exist just to register the atomic extension
package atomic

import (
	"github.com/Dfuz/xk6-atomic/atomic"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/atomic", new(atomic.RootModule))
}
