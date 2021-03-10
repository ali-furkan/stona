package defaults

import (
	"reflect"
)

// Get is ...
func Get(val interface{}, def interface{}) interface{} {
	d := val

	if reflect.ValueOf(val).IsZero() {
		d = def
	}

	return d
}

// GetString is ...
func GetString(val string, def string) string {
	d := def

	if val != "" {
		d = val
	}

	return d
}

// GetInt is ...
func GetInt(val int, def int) int {
	d := def

	if val != 0 {
		d = val
	}

	return d
}

func GetBool(val bool, def bool) bool {
	d := def

	if val != false {
		d = val
	}

	return d
}

func GetWithCond(cond bool, val interface{}, def interface{}) interface{} {
	d := def

	if cond && (val != nil || d != 0 || d != "") {
		d = val
	}

	return d
}
