package runtime

import "github.com/fluxninja/aperture/pkg/policies/controlplane/iface"

// TickEndCallback is a function that is called when a tick ends.
type TickEndCallback func(tickInfo TickInfo) error

// TickStartCallback is a function that is called when a tick starts.
type TickStartCallback func(tickInfo TickInfo) error

// CircuitAPI is for read only access to policy and also provides methods for acquiring & releasing circuit execution lock.
type CircuitAPI interface {
	iface.Policy
	RegisterTickEndCallback(ec TickEndCallback)
	RegisterTickStartCallback(sc TickStartCallback)
	LockExecution()
	UnlockExecution()
}
