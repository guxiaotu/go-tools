package encode

import (
	"bytes"
	"io"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func convertEncode(reader *transform.Reader) ([]byte, error) {
	d, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// Utf8ToGbk 将utf8（三个字节）编码的字符串转换为gbk编码
func Utf8ToGbk(utf8 []byte) ([]byte, error) {
	gbkEncodeReader := transform.NewReader(bytes.NewReader(utf8), simplifiedchinese.GBK.NewEncoder())
	return convertEncode(gbkEncodeReader)
}

// GbkToUtf8 将gbk（两个字节）编码的字符串转换为utf8编码
func GbkToUtf8(gbk []byte) ([]byte, error) {
	gbkDecodeReader := transform.NewReader(bytes.NewReader(gbk), simplifiedchinese.GBK.NewDecoder())
	return convertEncode(gbkDecodeReader)
}
