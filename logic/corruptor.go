package logic

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

const MaxPasses = 10

func CorruptFile(filePath string, passes int) error {
	if passes > MaxPasses {
		return fmt.Errorf("pass count (%d) exceeds maximum allowed (%d)", passes, MaxPasses)
	}

	info, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("failed to stat file: %v", err)
	}
	fileSize := info.Size()

	if fileSize == 0 {
		return fmt.Errorf("file is empty, nothing to process")
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY, 0)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	const chunkSize int64 = 4096
	buffer := make([]byte, chunkSize)

	for pass := 1; pass <= passes; pass++ {

		if _, err := file.Seek(0, io.SeekStart); err != nil {
			return fmt.Errorf("failed to seek file: %v", err)
		}
		remaining := fileSize
		for remaining > 0 {
			currentChunkSize := chunkSize
			if remaining < chunkSize {
				currentChunkSize = remaining
			}

			if _, err := rand.Read(buffer[:currentChunkSize]); err != nil {
				return fmt.Errorf("failed to generate random data: %v", err)
			}

			n, err := file.Write(buffer[:currentChunkSize])
			if err != nil {
				return fmt.Errorf("failed to write random data: %v", err)
			}
			if int64(n) != currentChunkSize {
				return fmt.Errorf("incomplete write: wrote %d bytes, expected %d", n, currentChunkSize)
			}
			remaining -= currentChunkSize
		}

		if err := file.Sync(); err != nil {
			return fmt.Errorf("failed to sync file data: %v", err)
		}
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("failed to seek file for header overwrite: %v", err)
	}

	headerSize := int64(1024)
	if fileSize < headerSize {
		headerSize = fileSize
	}
	headerBuffer := make([]byte, headerSize)
	n, err := file.Write(headerBuffer)
	if err != nil {
		return fmt.Errorf("failed to overwrite file header: %v", err)
	}
	if int64(n) != headerSize {
		return fmt.Errorf("incomplete header overwrite: wrote %d bytes, expected %d", n, headerSize)
	}
	if err := file.Sync(); err != nil {
		return fmt.Errorf("failed to sync header overwrite: %v", err)
	}

	return nil
}
