package toscalib

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
)

// Original source:
// https://gist.github.com/hvoecking/10772475
// There is an implied assumption that all attributes on a Struct are exported.
// For toscalib that will be the case so that assumption should work for us.

func _deepClone(to, from reflect.Value) {
	switch from.Kind() {
	// The first cases handle nested structures and translate them recursively

	// If it is a pointer we need to unwrap and call once again
	case reflect.Ptr:
		// To get the actual value of the from we have to call Elem()
		// At the same time this unwraps the pointer so we don't end up in
		// an infinite recursion
		fromValue := from.Elem()
		// Check if the pointer is nil
		if !fromValue.IsValid() {
			return
		}
		// Allocate a new object and set the pointer to it
		to.Set(reflect.New(fromValue.Type()))
		// Unwrap the newly created pointer
		_deepClone(to.Elem(), fromValue)

	// If it is an interface (which is very similar to a pointer), do basically the
	// same as for the pointer. Though a pointer is not the same as an interface so
	// note that we have to call Elem() after creating a new object because otherwise
	// we would end up with an actual pointer
	case reflect.Interface:
		// Get rid of the wrapping interface
		fromValue := from.Elem()
		if !fromValue.IsValid() {
			return
		}

		// Create a new object. Now new gives us a pointer, but we want the value it
		// points to, so we have to call Elem() to unwrap it
		toValue := reflect.New(fromValue.Type()).Elem()
		_deepClone(toValue, fromValue)
		to.Set(toValue)

	// If it is a struct we translate each field
	case reflect.Struct:
		for i := 0; i < from.NumField(); i++ {
			_deepClone(to.Field(i), from.Field(i))
		}

	// If it is a slice we create a new slice and translate each element
	case reflect.Slice:
		if from.IsNil() {
			return
		}
		to.Set(reflect.MakeSlice(from.Type(), from.Len(), from.Cap()))
		for i := 0; i < from.Len(); i++ {
			_deepClone(to.Index(i), from.Index(i))
		}

	// If it is a map we create a new map and translate each value
	case reflect.Map:
		if from.IsNil() {
			return
		}
		to.Set(reflect.MakeMap(from.Type()))
		for _, key := range from.MapKeys() {
			fromValue := from.MapIndex(key)
			// New gives us a pointer, but again we want the value
			toValue := reflect.New(fromValue.Type()).Elem()
			_deepClone(toValue, fromValue)
			to.SetMapIndex(key, toValue)
		}

	// And everything else will simply be taken from the from
	default:
		to.Set(from)
	}
}

func clone(obj interface{}) interface{} {
	// Wrap the from in a reflect.Value
	from := reflect.ValueOf(obj)

	to := reflect.New(from.Type()).Elem()
	_deepClone(to, from)

	// Remove the reflection wrapper
	return to.Interface()
}

func get(k int, list []interface{}) string {
	if len(list) <= k {
		return ""
	}
	if v, ok := list[k].(string); ok {
		return v
	}
	return ""
}

func remainder(k int, list []interface{}) []interface{} {
	if len(list) <= k {
		return make([]interface{}, 0)
	}
	return list[k+1:]
}

func copyFile(src, destDir string) (string, error) {
	absSrc, err := filepath.Abs(src)
	if err != nil {
		return "", err
	}
	srcFilename := filepath.Base(absSrc)

	dest, err := filepath.Abs(filepath.Join(destDir, srcFilename))
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(absSrc)
	if err != nil {
		return "", err
	}
	// Write data to dst
	err = ioutil.WriteFile(dest, data, 0644)
	if err != nil {
		return "", err
	}

	return dest, nil
}

func isAbsLocalPath(location string) bool {
	if filepath.IsAbs(location) {
		_, err := os.Stat(location)
		return !os.IsNotExist(err)
	}
	return false
}
