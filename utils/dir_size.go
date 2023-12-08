package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func DirSize(dirPath string) (dirSize int64) {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if !info.IsDir() {
			dirSize += info.Size()
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return dirSize
}

// Funkce pro formátování velikosti v přehledné formě
func formatSize(size int64) string {
	const (
		B  = 1
		KB = B * 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%d B", size)
	}
}
