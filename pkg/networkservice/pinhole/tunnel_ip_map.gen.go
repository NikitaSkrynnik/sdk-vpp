// Code generated by "-output tunnel_ip_map.gen.go -type ipPortMap<IPPort,struct{}> -output tunnel_ip_map.gen.go -type ipPortMap<IPPort,struct{}>"; DO NOT EDIT.
// Install -output tunnel_ip_map.gen.go -type ipPortMap<IPPort,struct{}> by "go get -u github.com/searKing/golang/tools/-output tunnel_ip_map.gen.go -type ipPortMap<IPPort,struct{}>"

package pinhole

import (
	"sync" // Used by sync.Map.
)

// Generate code that will fail if the constants change value.
func _() {
	// An "cannot convert ipPortMap literal (type ipPortMap) to type sync.Map" compiler error signifies that the base type have changed.
	// Re-run the go-syncmap command to generate them again.
	_ = (sync.Map)(ipPortMap{})
}

var _nil_ipPortMap_struct___value = func() (val struct{}) { return }()

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *ipPortMap) Load(key IPPort) (struct{}, bool) {
	value, ok := (*sync.Map)(m).Load(key)
	if value == nil {
		return _nil_ipPortMap_struct___value, ok
	}
	return value.(struct{}), ok
}

// Store sets the value for a key.
func (m *ipPortMap) Store(key IPPort, value struct{}) {
	(*sync.Map)(m).Store(key, value)
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *ipPortMap) LoadOrStore(key IPPort, value struct{}) (struct{}, bool) {
	actual, loaded := (*sync.Map)(m).LoadOrStore(key, value)
	if actual == nil {
		return _nil_ipPortMap_struct___value, loaded
	}
	return actual.(struct{}), loaded
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *ipPortMap) LoadAndDelete(key IPPort) (value struct{}, loaded bool) {
	actual, loaded := (*sync.Map)(m).LoadAndDelete(key)
	if actual == nil {
		return _nil_ipPortMap_struct___value, loaded
	}
	return actual.(struct{}), loaded
}

// Delete deletes the value for a key.
func (m *ipPortMap) Delete(key IPPort) {
	(*sync.Map)(m).Delete(key)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently, Range may reflect any mapping for that key
// from any point during the Range call.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (m *ipPortMap) Range(f func(key IPPort, value struct{}) bool) {
	(*sync.Map)(m).Range(func(key, value interface{}) bool {
		return f(key.(IPPort), value.(struct{}))
	})
}
