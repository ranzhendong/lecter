package transformation

import (
	"reflect"
	"unsafe"
)

func Any2Str(any interface{}) string {
	switch any.(type) { //多选语句switch
	case string:
		return any.(interface{}).(string)
	case interface{}:
		return any.(interface{}).(string)
	}
	return "nil"
}

//self way
func B2S(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

//如不修改数据，仅转换类型
func StringToBytesNoSelf(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

//如不修改数据，仅转换类型
func BytesToStringNoSelf(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: bh.Data,
		Len:  bh.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}
