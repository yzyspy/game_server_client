package zutils

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := NewShardLockMaps()
	m.Set("key1", "value1")
	m.Set("key2", "value2")

	items := m.Items()
	for k, v := range items {
		fmt.Printf("key: %s, value: %v\n", k, v)
	}
}
