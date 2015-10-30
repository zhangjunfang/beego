package webservice

import (
	base "cloudplatform_base/base"
	"fmt"
	"strings"
	"testing"
)

func TestMessageForwardYt(t *testing.T) {
	id1, _ := base.FromStr("1870747d-b26c-4507-9518-1ca62bc66e5d")
	id2 := base.MustFromStr("1870747db26c450795181ca62bc66e5d")
	fmt.Println(id1 == id2) // true
	fmt.Println(strings.Replace(base.Rand().Hex(), "-", "", -1))
	var m Message
	m.MessageForwardYt()
}
