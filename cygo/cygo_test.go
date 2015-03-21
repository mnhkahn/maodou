package cygo

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCygo(t *testing.T) {
	now := Now()
	res, err := json.Marshal(now)
	fmt.Println(res, err)
}
