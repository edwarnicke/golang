// +build go1.13

package object

// IsZero reports whether a value is a zero value of its kind.
// If value.Kind() is Struct, it traverses each field of the struct
// recursively calling IsZero, returning true only if each field's IsZero
// result is also true.
func IsZero(obj interface{}) bool {
	var v reflect.Value
	if vv, ok := obj.(reflect.Value); ok {
		v = vv
	} else {
		v = reflect.ValueOf(obj)
	}
	return v.IsZero()
}