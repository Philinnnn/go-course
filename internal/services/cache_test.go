package services

import (
	"os"
	"reflect"
	"runtime/pprof"
	"testing"
	"unsafe"
)

func TestCacheHits(t *testing.T) {
	c := newCache()

	t.Run("manualChange", func(t *testing.T) {
		v := reflect.ValueOf(c).Elem()
		hitsField := v.FieldByName("hits")
		ptr := (*int)(unsafe.Pointer(hitsField.UnsafeAddr()))
		*ptr = 42
		if c.hits != 42 {
			t.Errorf("ожидалось 42, получили %d", c.hits)
		}
	})

	t.Run("methodChange", func(t *testing.T) {
		c.data["foo"] = "bar"
		val := c.get("foo")
		if val != "bar" {
			t.Errorf("ожидалось bar, получили %s", val)
		}
		if c.hits != 43 {
			t.Errorf("ожидалось 43, получили %d", c.hits)
		}
	})
}

func TestCacheParallel(t *testing.T) {
	c := newCache()
	c.data["x"] = "y"
	t.Run("parallel", func(t *testing.T) {
		t.Parallel()
		for i := 0; i < 1000; i++ {
			_ = c.get("x")
		}
	})
}

func TestCachePprof(t *testing.T) {
	c := newCache()
	c.data["key"] = "val"

	f, err := os.Create("cpu.prof")
	if err != nil {
		t.Fatalf("не удалось создать файл профиля: %v", err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 100000; i++ {
		_ = c.get("key")
	}
	t.Logf("hits: %d", c.hits)
}
