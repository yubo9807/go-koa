package number

import (
	"math/rand"
	"time"
)

func Random(num int) int {
	var timestamp = time.Now().UnixNano()
	rand.Seed(timestamp)
	return rand.Intn(num)
}
