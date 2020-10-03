package singleton_lazy_test

import (
	lazySingleton "DesignPattern-go/singleton_lazy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, lazySingleton.GetLazyInstance(), lazySingleton.GetLazyInstance())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			assert.Equal(b, lazySingleton.GetLazyInstance(), lazySingleton.GetLazyInstance())
		}
	})
}
