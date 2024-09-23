package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileBinary(t *testing.T) {
	ReadFileBinary("img/penguin.png")
}

func TestSplitFileBinary(t *testing.T) {
	//_ = chatgpt("img/penguin.png", 10000.0)
	files := SplitFileBinary("img/penguin.png", "d:/output", 10.0)
	assert.Equal(t, []string{"d:/output/penguin.png.part0", "d:/output/penguin.png.part1", "d:/output/penguin.png.part2", "d:/output/penguin.png.part3",
		"d:/output/penguin.png.part4", "d:/output/penguin.png.part5", "d:/output/penguin.png.part6", "d:/output/penguin.png.part7",
		"d:/output/penguin.png.part8", "d:/output/penguin.png.part9"}, files)
}

func TestMergeBinaryFileBinary(t *testing.T) {
	MergeBinary("d:/output", "img/企鹅.png")
}
