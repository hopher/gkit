package log

import (
	"encoding"
	"encoding/json"
	"fmt"
	"io"
	"reflect"

	kitlog "github.com/go-kit/kit/log"
)

type aliLogger struct {
	io.Writer
}

// NewAliLogger 阿里云sls日志服务
func NewAliLogger(w io.Writer) kitlog.Logger {
	return &aliLogger{w}
}

func (l *aliLogger) Log(keyvals ...interface{}) error {
	n := (len(keyvals) + 1) / 2 // +1 to handle case when len is odd
	m := make(map[string]interface{}, n)
	for i := 0; i < len(keyvals); i += 2 {
		k := keyvals[i]
		var v interface{} = kitlog.ErrMissingValue
		if i+1 < len(keyvals) {
			v = keyvals[i+1]
		}
		merge(m, k, v)
	}
	enc := json.NewEncoder(l.Writer)
	enc.SetEscapeHTML(false)
	return enc.Encode(m)
}

func merge(dst map[string]interface{}, k, v interface{}) {
	var key string
	switch x := k.(type) {
	case string:
		key = x
	case fmt.Stringer:
		key = safeString(x)
	default:
		key = fmt.Sprint(x)
	}

	// We want json.Marshaler and encoding.TextMarshaller to take priority over
	// err.Error() and v.String(). But json.Marshall (called later) does that by
	// default so we force a no-op if it's one of those 2 case.
	switch x := v.(type) {
	case json.Marshaler:
	case encoding.TextMarshaler:
	case error:
		v = safeError(x)
	case fmt.Stringer:
		v = safeString(x)
	}

	dst[key] = v
}

func safeString(str fmt.Stringer) (s string) {
	defer func() {
		if panicVal := recover(); panicVal != nil {
			if v := reflect.ValueOf(str); v.Kind() == reflect.Ptr && v.IsNil() {
				s = "NULL"
			} else {
				s = fmt.Sprintf("PANIC in String method: %v", panicVal)
			}
		}
	}()
	s = str.String()
	return
}

func safeError(err error) (s interface{}) {
	defer func() {
		if panicVal := recover(); panicVal != nil {
			if v := reflect.ValueOf(err); v.Kind() == reflect.Ptr && v.IsNil() {
				s = nil
			} else {
				s = fmt.Sprintf("PANIC in Error method: %v", panicVal)
			}
		}
	}()
	s = err.Error()
	return
}
