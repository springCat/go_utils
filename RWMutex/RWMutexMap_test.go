package syn
import (
"testing"

)

func BenchmarkPut(b *testing.B) {
	var m = NewRWMutexMap()
	for i := 0; i < b.N; i++ {
		m.Put(i,i)
	}
	m.Recycle()
}

func BenchmarkGet(b *testing.B) {
	var m = NewRWMutexMap()
	for i := 0; i < b.N; i++ {
		m.Get(i)
	}
	m.Recycle()
}

func BenchmarkContains(b *testing.B) {
	var m = NewRWMutexMap()
	for i := 0; i < b.N; i++ {
		m.Contains(i)
	}
	m.Recycle()
}

func BenchmarkLen(b *testing.B) {
	var m = NewRWMutexMap()
	for i := 0; i < b.N; i++ {
		m.Len()
	}
	m.Recycle()
}

func BenchmarkKeys(b *testing.B) {
	var m = NewRWMutexMap()
	for i := 0; i < b.N; i++ {
		m.Keys()
	}
	m.Recycle()
}

func BenchmarkValues(b *testing.B) {
	var m = NewRWMutexMap()
	for i := 0; i < b.N; i++ {
		m.Values()
	}
	m.Recycle()
}

func BenchmarkClear(b *testing.B) {
	var m = NewRWMutexMap()
	for i := 0; i < b.N; i++ {
		m.Clear()
	}
	m.Recycle()
}

