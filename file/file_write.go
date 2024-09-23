package file

import (
	"os"
	"sync"
	"time"
)

// ConcurrentFile 并发写文件
func ConcurrentFile(outFile, data string, times int) {
	const C = 10
	wg := sync.WaitGroup{}
	wg.Add(C)
	for i := 0; i < C; i++ {
		go func() {
			defer wg.Done()
			// 由于多次打开，所以使用你APPEND，而不是TRUNC
			fout, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				panic(err)
			}
			defer fout.Close()
			for j := 0; j < times; j++ {
				_, err := fout.WriteString(time.Now().Format("2006-01-02 15:04:05 ") + data + "\n")
				if err != nil {
					panic(err)
				}
			}
		}()
	}
	wg.Wait()
}
