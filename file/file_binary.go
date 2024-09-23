package file

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

// ReadFileBinary 读取文件内容
func ReadFileBinary(file string) {
	f, err := os.Open(file)
	// 打开文件失败
	if err != nil {
		log.Fatal(err)
	}
	// 关闭文件
	defer f.Close()

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

// SplitFileBinary 分割二进制文件
func SplitFileBinary(file, outDir string, count int) []string {
	baseName := strings.Split(file, "/")[1]

	// 创建输出目录
	CreateDir(outDir)

	fin, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	fileInfo, err := fin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	size := fileInfo.Size()                               //获取文件大小
	seg := int(math.Ceil(float64(size) / float64(count))) //分割成count段，计算每段大小
	fmt.Printf("文件总大小: %d，每段大小: %d\n", size, seg)

	//对于二进制文件，文件开头存储了一些元信息，比如文件格式、文件总大小等
	files := make([]string, 0, count)

	for i := 0; i < count; i++ {
		outFile := fmt.Sprintf("%s.part%d", outDir+"/"+baseName, i)
		files = append(files, outFile)
		fout, err := os.Create(outFile)
		if err != nil {
			log.Fatal(err)
		}
		defer fout.Close()

		buffer := make([]byte, seg)
		n, err := fin.Read(buffer)
		if err != nil {
			switch {
			case err != io.EOF:
				log.Fatal(err)
			case err == io.EOF && len(buffer) > 0:
				// 读取最后一段
				_, _ = fout.Write(buffer[:n])
			}
			break
		}
		// 读取每一段
		_, _ = fout.Write(buffer[:n])
	}

	return files
}

// MergeBinary 合并二进制文件
func MergeBinary(output, mergedFile string) {
	// 注意是APPEND模式
	fout, err := os.OpenFile(mergedFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	defer fout.Close()

	files := ListDir(output)
	// 遍历每一个小的二进制文件
	for _, file := range files {
		fin, err := os.Open(file)
		if err != nil {
			log.Panic(err)
		}
		defer fin.Close()

		buffer := make([]byte, 1024)
		// 循环读取一个二进制文件
		for {
			n, err := fin.Read(buffer)
			if err != nil {
				switch {
				case err != io.EOF:
					log.Println(err)
				case err == io.EOF && n > 0:
					// 追加到合并后的大文件中去
					_, _ = fout.Write(buffer[:n])
				}
				break
			}
			_, _ = fout.Write(buffer[:n])
		}
	}
}

// CreateDir 创建目录
func CreateDir(outDir string) {
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		if err := os.Mkdir(outDir, 0755); err != nil {
			log.Fatal(err)
		}
	}
}

func ListDir(root string) []string {
	files := make([]string, 0)
	_ = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files
}
