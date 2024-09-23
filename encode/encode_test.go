package encode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	assertions := assert.New(t)
	utf8 := []byte("顾小兔")
	assertions.Equal("[233 161 190 229 176 143 229 133 148]", fmt.Sprintf("%v", utf8))
	gbk, err := Utf8ToGbk(utf8)
	if err != nil {
		t.Error(err)
	}

	assertions.Equal("[185 203 208 161 205 195]", fmt.Sprintf("%v", gbk))
	utf8, err = GbkToUtf8(gbk)
	if err != nil {
		t.Error(err)
	}
	assertions.Equal("[233 161 190 229 176 143 229 133 148]", fmt.Sprintf("%v", utf8))
}
