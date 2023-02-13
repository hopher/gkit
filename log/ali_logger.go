package log

import (
	"bytes"
	"encoding"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/aliyun/aliyun-log-go-sdk/producer"
	kitlog "github.com/go-kit/kit/log"
)

// aliLogger 阿里云日志服务
//
// 文档:
//   - 安装Go SDK https://help.aliyun.com/document_detail/286949.html
//   - 保留字段 https://help.aliyun.com/document_detail/92264.html
type aliLogger struct {
	producer *producer.Producer
	callBack producer.CallBack
	project  string
	logStore string
	topic    string
	source   string
}

// Option 可选参数定义
type Option func(o *aliLogger)

func WithAliLoggerCallBack(callback producer.CallBack) Option {
	return func(o *aliLogger) {
		o.callBack = callback
	}
}

// NewAliLogger 阿里云sls日志服务
func NewAliLogger(project, logStore, topic, source string, producer *producer.Producer, options ...Option) kitlog.Logger {

	logger := &aliLogger{
		project:  project,
		logStore: logStore,
		topic:    topic,
		source:   source,
		producer: producer,
	}

	for _, option := range options {
		option(logger)
	}

	return logger
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

	buffer := &bytes.Buffer{}
	enc := json.NewEncoder(buffer)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(m); err != nil {
		return err
	}

	log := producer.GenerateLog(uint32(time.Now().Unix()), map[string]string{
		"content": buffer.String(),
	})

	if l.callBack != nil {
		return l.producer.SendLogWithCallBack(l.project, l.logStore, l.topic, l.source, log, l.callBack)
	}

	return l.producer.SendLog(l.project, l.logStore, l.topic, l.source, log)
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
