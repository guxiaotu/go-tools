package file

import "testing"

func TestReadFileText(t *testing.T) {
	ReadFileText("text/test.txt")
}

func TestSplitTextFile(t *testing.T) {
	SplitTextFile("text/test.txt", 4.0)
}
