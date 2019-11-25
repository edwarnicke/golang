// Code generated by "go-syncmap -type Globals<uintptr, onSignalHandler>"; DO NOT EDIT.

package cgo

import "sync"

func _() {
	// An "cannot convert Globals literal (type Globals) to type sync.Map" compiler error signifies that the base type have changed.
	// Re-run the go-syncmap command to generate them again.
	_ = (sync.Map)(Globals{})
}

var _nil_Globals_onSignalHandler_value = func() (val onSignalHandler) { return }()

func (m *Globals) Store(key uintptr, value onSignalHandler) {
	(*sync.Map)(m).Store(key, value)
}

func (m *Globals) LoadOrStore(key uintptr, value onSignalHandler) (onSignalHandler, bool) {
	actual, loaded := (*sync.Map)(m).LoadOrStore(key, value)
	if actual == nil {
		return _nil_Globals_onSignalHandler_value, loaded
	}
	return actual.(onSignalHandler), loaded
}

func (m *Globals) Load(key uintptr) (onSignalHandler, bool) {
	value, ok := (*sync.Map)(m).Load(key)
	if value == nil {
		return _nil_Globals_onSignalHandler_value, ok
	}
	return value.(onSignalHandler), ok
}

func (m *Globals) Delete(key uintptr) {
	(*sync.Map)(m).Delete(key)
}

func (m *Globals) Range(f func(key uintptr, value onSignalHandler) bool) {
	(*sync.Map)(m).Range(func(key, value interface{}) bool {
		return f(key.(uintptr), value.(onSignalHandler))
	})
}