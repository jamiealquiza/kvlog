package kvlog_test

import (
	"testing"

	"github.com/jamiealquiza/kvlog"
)

func BenchmarkLog(b *testing.B) {
    for i := 0; i < b.N; i++ {
        kvlog.Log("test", map[string]string{
        	"field1": "value1",
        })
    }
}

func BenchmarkLogInfo(b *testing.B) {
    for i := 0; i < b.N; i++ {
        kvlog.LogInfo("test", "an info log")
    }
}

func BenchmarkLogErr(b *testing.B) {
    for i := 0; i < b.N; i++ {
        kvlog.LogErr("test", "an err log")
    }
}