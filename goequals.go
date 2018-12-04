package goequals

import (
	"reflect"
)

func Equals(v1, v2 interface{}) bool {
	switch v := v1.(type) {
	case int:
		return equalsInt(int64(v), v2)
	case int8:
		return equalsInt(int64(v), v2)
	case int16:
		return equalsInt(int64(v), v2)
	case int32:
		return equalsInt(int64(v), v2)
	case int64:
		return equalsInt(int64(v), v2)
	case uint:
		return equalsInt(int64(v), v2)
	case uint8:
		return equalsInt(int64(v), v2)
	case uint16:
		return equalsInt(int64(v), v2)
	case uint32:
		return equalsInt(int64(v), v2)
	case uint64:
		return compareUint(uint64(v), v2)
	case float32:
		return compareFloat(float64(v), v2)
	case float64:
		return compareFloat(float64(v), v2)
	case string:
		return compareString(v, v2)
	case bool:
		return compareBool(v, v2)
	case complex64:
		return compareComplex(complex128(v), v2)
	case complex128:
		return compareComplex(v, v2)
	}

	val1 := reflect.ValueOf(v1)
	val2 := reflect.ValueOf(v2)
	switch val1.Kind() {
	case reflect.Slice:
		return compareSlice(val1, val2)
	case reflect.Map:
		return compareMap(val1, val2)
	case reflect.Array:
		return compareArray(val1, val2)
	default:
		return reflect.DeepEqual(v1, v2)
	}
}

func compareComplex(v1 complex128, _v2 interface{}) bool {
	switch v2 := _v2.(type) {
	case complex64:
		return v1 == complex128(v2)
	case complex128:
		return v1 == v2
	}
	return false
}

func compareArray(v1 reflect.Value, v2 reflect.Value) bool {
	if v2.Kind() != reflect.Array {
		return false
	}

	dv1 := downgradeArray(v1)
	dv2 := downgradeArray(v2)
	if len(dv1) != len(dv2) {
		return false
	}

	for i, elem1 := range dv1 {
		//log.Printf("dv1 %+v dv2 %+v elem %d = %+v, %+v", dv1, dv2, i, elem1, dv2[i])
		if !Equals(elem1, dv2[i]) {
			return false
		}
	}
	return true
}

func compareMap(v1, v2 reflect.Value) bool {
	if v2.Kind() != reflect.Map {
		return false
	}

	dv1 := downgradeMap(v1)
	dv2 := downgradeMap(v2)
	if len(dv1) != len(dv2) {
		return false
	}

	for k1, elem1 := range dv1 {
		for k2, elem2 := range dv2 {
			if Equals(k1, k2) {
				if !Equals(elem1, elem2) {
					return false
				}
				delete(dv2, k2)
				break
			}
		}
	}
	return true
}

func compareSlice(v1, v2 reflect.Value) bool {
	if v2.Kind() != reflect.Slice {
		return false
	}

	dv1 := downgradeSlice(v1)
	dv2 := downgradeSlice(v2)
	if len(dv1) != len(dv2) {
		return false
	}

	for i, elem1 := range dv1 {
		//log.Printf("dv1 %+v dv2 %+v elem %d = %+v, %+v", dv1, dv2, i, elem1, dv2[i])
		if !Equals(elem1, dv2[i]) {
			return false
		}
	}
	return true
}

func downgradeSlice(v reflect.Value) (dv []interface{}) {
	dv = make([]interface{}, v.Len())
	for i := 0; i < len(dv); i++ {
		dv[i] = v.Index(i).Interface()
	}
	return
}

func downgradeMap(v reflect.Value) (dv map[interface{}]interface{}) {
	dv = make(map[interface{}]interface{}, v.Len())
	for _, key := range v.MapKeys() {
		elem := v.MapIndex(key)
		dv[key.Interface()] = elem.Interface()
	}
	return
}

func downgradeArray(v reflect.Value) (dv []interface{}) {
	dv = make([]interface{}, v.Len())
	for i := 0; i < len(dv); i++ {
		dv[i] = v.Index(i).Interface()
	}
	return
}

func compareBool(v bool, v2 interface{}) bool {
	switch v2 := v2.(type) {
	case bool:
		return v == v2
	default:
		return false
	}
}

func compareString(v string, v2 interface{}) bool {
	switch v2 := v2.(type) {
	case string:
		return v == v2
	default:
		return false
	}
}

func compareFloat(v float64, v2 interface{}) bool {
	switch v2 := v2.(type) {
	case float32:
		return v-0.000001 < float64(v2) && float64(v2) < v+0.000001
	case float64:
		return v-0.000001 < float64(v2) && float64(v2) < v+0.000001
	default:
		return false
	}
}

func equalsInt(v int64, r interface{}) bool {
	switch v2 := r.(type) {
	case int:
		return int64(v2) == v
	case int8:
		return int64(v2) == v
	case int16:
		return int64(v2) == v
	case int32:
		return int64(v2) == v
	case int64:
		return int64(v2) == v
	case uint:
		return int64(v2) == v
	case uint8:
		return int64(v2) == v
	case uint16:
		return int64(v2) == v
	case uint32:
		return int64(v2) == v
	case uint64:
		if v < 0 {
			return false
		}
		return int64(v2) == v
	default:
		return false
	}
}

func compareUint(v uint64, r interface{}) bool {
	switch v2 := r.(type) {
	case int:
		return v2 >= 0 && uint64(v2) == v
	case int8:
		return v2 >= 0 && uint64(v2) == v
	case int16:
		return v2 >= 0 && uint64(v2) == v
	case int32:
		return v2 >= 0 && uint64(v2) == v
	case int64:
		return v2 >= 0 && uint64(v2) == v
	case uint:
		return v2 >= 0 && uint64(v2) == v
	case uint8:
		return v2 >= 0 && uint64(v2) == v
	case uint16:
		return v2 >= 0 && uint64(v2) == v
	case uint32:
		return v2 >= 0 && uint64(v2) == v
	case uint64:
		return v2 >= 0 && uint64(v2) == v
	default:
		return false
	}
}
