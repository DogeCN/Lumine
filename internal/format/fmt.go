package format

import (
	"fmt"
	"strconv"
	"unsafe"
)

func Byte(b byte) string {
	var hexBuf [2]byte
	const hexDigits = "0123456789abcdef"
	hexBuf[0] = hexDigits[b>>4]
	hexBuf[1] = hexDigits[b&0x0f]
	return string(hexBuf[:])
}

func Concat(values ...any) string {
	buf := make([]byte, 0, 32)
	for _, val := range values {
		switch v := val.(type) {
		case fmt.Stringer:
			buf = append(buf, v.String()...)
		case string:
			buf = append(buf, v...)
		case error:
			buf = append(buf, v.Error()...)
		case int:
			buf = strconv.AppendInt(buf, int64(v), 10)
		case int8:
			buf = strconv.AppendInt(buf, int64(v), 10)
		case int16:
			buf = strconv.AppendInt(buf, int64(v), 10)
		case int32:
			buf = strconv.AppendInt(buf, int64(v), 10)
		case int64:
			buf = strconv.AppendInt(buf, v, 10)
		case uint:
			buf = strconv.AppendUint(buf, uint64(v), 10)
		case uint8:
			buf = strconv.AppendUint(buf, uint64(v), 10)
		case uint16:
			buf = strconv.AppendUint(buf, uint64(v), 10)
		case uint32:
			buf = strconv.AppendUint(buf, uint64(v), 10)
		case uint64:
			buf = strconv.AppendUint(buf, v, 10)
		default:
			panic("format.Concat: unsupported type")
		}
	}
	return unsafe.String(unsafe.SliceData(buf), len(buf))
}

func Int[T int | int8 | int16 | int32 | int64](v T) string {
	return strconv.FormatInt(int64(v), 10)
}

func Uint[T uint | uint8 | uint16 | uint32 | uint64](v T) string {
	return strconv.FormatUint(uint64(v), 10)
}
