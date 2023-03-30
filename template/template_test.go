package template

import (
	"fmt"
	"testing"
	"time"
)

func TestCook(t *testing.T) {
	var tomato Tomato
	Cooking(tomato)

	var potato Potato
	Cooking(potato)

	nt := time.Now()
	fmt.Println(nt.Format("2006-01-02 15:04:05"))
}
