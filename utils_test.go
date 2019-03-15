package gomygen

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestClean(t *testing.T) {
	fmt.Println(strings.Repeat("曾经沧海难为水，除却巫山不是云", 2000))
	time.Sleep(time.Millisecond * 700)
	Clean()
}
