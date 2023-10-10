package exporter

import (
	"bytes"
	"fmt"
	"strings"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

type TextLogsMarshaler struct{}

type dataBuffer struct {
	buf bytes.Buffer
}

func (b *dataBuffer) logEntry(format string, a ...interface{}) {
	b.buf.WriteString(fmt.Sprintf(format, a...))
	b.buf.WriteString("\n")
}

func (b *dataBuffer) logAttributes(header string, m pcommon.Map) {
	if m.Len() == 0 {
		return
	}

	b.logEntry("%s:", header)
	attrPrefix := "     ->"

	// Add offset to attributes if needed.
	headerParts := strings.Split(header, "->")
	if len(headerParts) > 1 {
		attrPrefix = headerParts[0] + attrPrefix
	}

	m.Range(func(k string, v pcommon.Value) bool {
		b.logEntry("%s %s: %s", attrPrefix, k, valueToString(v))
		return true
	})
}

func (b *dataBuffer) logInstrumentationScope(il pcommon.InstrumentationScope) {
	b.logEntry(
		"InstrumentationScope %s %s",
		il.Name(),
		il.Version())
	b.logAttributes("InstrumentationScope attributes", il.Attributes())
}

func valueToString(v pcommon.Value) string {
	return fmt.Sprintf("%s(%s)", v.Type().String(), v.AsString())
}

func NewTextLogsMarshaler() plog.Marshaler {
	return TextLogsMarshaler{}
}

func (m TextLogsMarshaler) MarshalLogs(ld plog.Logs) ([]byte, error) {
	buf := dataBuffer{}
	rls := ld.ResourceLogs()
	for i := 0; i < rls.Len(); i++ {
		buf.logEntry("ResourceLog #%d", i)
		rl := rls.At(i)
		buf.logEntry("Resource SchemaURL: %s", rl.SchemaUrl())
		buf.logAttributes("Resource attributes", rl.Resource().Attributes())
		ills := rl.ScopeLogs()
		for j := 0; j < ills.Len(); j++ {
			buf.logEntry("ScopeLogs #%d", j)
			ils := ills.At(j)
			buf.logEntry("ScopeLogs SchemaURL: %s", ils.SchemaUrl())
			buf.logInstrumentationScope(ils.Scope())

			logs := ils.LogRecords()
			for k := 0; k < logs.Len(); k++ {
				buf.logEntry("LogRecord #%d", k)
				lr := logs.At(k)
				buf.logEntry("ObservedTimestamp: %s", lr.ObservedTimestamp())
				buf.logEntry("Timestamp: %s", lr.Timestamp())
				buf.logEntry("SeverityText: %s", lr.SeverityText())
				buf.logEntry("SeverityNumber: %s(%d)", lr.SeverityNumber(), lr.SeverityNumber())
				buf.logEntry("Body: %s", valueToString(lr.Body()))
				buf.logAttributes("Attributes", lr.Attributes())
				buf.logEntry("Trace ID: %s", lr.TraceID())
				buf.logEntry("Span ID: %s", lr.SpanID())
				buf.logEntry("Flags: %d", lr.Flags())
			}
		}
	}

	return buf.buf.Bytes(), nil
}
