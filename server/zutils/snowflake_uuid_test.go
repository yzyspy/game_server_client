package zutils

import (
	"fmt"
	"testing"
)

func TestSnowflakeUUID(t *testing.T) {
	workerId := 1
	idWorker, err := NewIDWorker(int64(workerId))
	if err != nil {
		t.Fatal(err)
	}
	id, err := idWorker.NextID()
	fmt.Println(id)
}
