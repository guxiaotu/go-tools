package file

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

// ReadFileText 读取文件内容
func ReadFileText(file string) {
	f, err := os.Open(file)

	// 打开文件失败
	if err != nil {
		log.Fatal(err)
	}

	// 关闭文件
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()

		// 当err不为空时两种情况（打印最后一行和读文件错误）
		if err != nil {
			switch {
			// 打印读取错误
			case err != io.EOF:
				log.Fatal(err)

			// 打印最后一行，表示已经读取到文件尾部
			case err == io.EOF && len(line) > 0:
				fmt.Println(string(line))
			}
			// 读完最后一行后结束循环读取
			break
		}

		// 打印每一行
		fmt.Println(string(line))
	}
}

// SplitTextFile 拆分文件
func SplitTextFile(file string, count float64) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	// 文件大小
	fileSize := info.Size()
	// 根据count值进行文件分割，计算每段大小（向上取整）
	seg := int(math.Ceil(float64(fileSize) / count))
	fmt.Printf("文件总大小：%d，每段大小：%d\n", fileSize, seg)

	// 对于文本文件，没有存储额外的元信息，存储的全都是内容数据本身
	// 同一个字符，不同的编码方式将其转为不同的byte数字

	for {
		buffer := make([]byte, seg)
		// n是成功读取的字节数
		n, err := f.Read(buffer)
		if err != nil {
			switch {
			case err != io.EOF:
				log.Fatal(err)
			case err == io.EOF && len(buffer) > 0:
				fmt.Println(buffer[:n])
				fmt.Print(string(buffer[:n]))
			}
			// 读完最后一行
			break
		}
		// 读取每一行
		fmt.Println(buffer[:n])
		fmt.Print(string(buffer[:n]))
	}
}
