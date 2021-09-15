package loggers

import "testing"

func BenchmarkZerolog(b *testing.B) {benchmarkZerolog(100, b)}
func BenchmarkLogrus(b *testing.B) {benchmarkLogrus(100, b)}
