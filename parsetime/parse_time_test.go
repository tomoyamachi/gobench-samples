package gobench_samples

import (
	"regexp"
	"testing"
	"time"
)

var date = "2019/01/01 03:30:30"
var globalReg = regexp.MustCompile(`\d{4}(/\d{2}){2} \d{2}(:\d{2}){2}`)

func BenchmarkParseTime_InnerRegex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reg := regexp.MustCompile(`\d{4}(/\d{2}){2} \d{2}:\d{2}:\d{2}`)
		reg.MatchString(date)
	}
}

func BenchmarkParseTime_GlobalRegex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		globalReg.MatchString(date)
	}
}

func BenchmarkParseTime_Time(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		time.Parse("2006/01/02 15:04:05", date)
	}
}

// BenchmarkParseTime_InnerRegex-12                 1000000              5117 ns/op
// BenchmarkParseTime_GlobalRegex-12               20000000               232 ns/op
// BenchmarkParseTime_Time-12                      30000000               139 ns/op
