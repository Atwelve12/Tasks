package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	dir := "/home/aaa12/demo/test_demo_files"
	frontnote := "aaa12_"

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("读取目录失败:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".txt" {
			oldPath := filepath.Join(dir, file.Name())
			newPath := filepath.Join(dir, frontnote + file.Name())

			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Printf("重命名 %s 失败：%v\n", file.Name(), err)
				continue
			}
			fmt.Printf("成功重命名：%s -> %s\n", file.Name(), frontnote+file.Name())
		}
	}
	fmt.Println("批量重命名完成")
}


