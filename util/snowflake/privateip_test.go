package snowflake_test

import (
	"testing"

	"github.com/micro/go-micro/v2/util/snowflake"
)

func TestPrivateIPToMachineID(t *testing.T) {
	mid := snowflake.PrivateIPToMachineID()
	if mid <= 0 {
		t.Error("MachineID should be > 0")
	}
}
