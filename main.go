package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

const ChunkSize = 64 * 1024 * 1024

var isVerbose = false

func getFileSize(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

func readBlock(file *os.File, offset int64, size int64) (int, []byte, error) {

	_, err := file.Seek(offset, 0)
	if err != nil {
		return 0, nil, err
	}

	block := make([]byte, size)
	bytesRead, err := file.Read(block)
	if err != nil {
		return 0, nil, err
	}
	return bytesRead, block, nil
}

func appendBlock(file *os.File, sizeOfBlock int, block []byte) error {
	_, err := file.Write(block[:sizeOfBlock])
	if err != nil {
		return err
	}
	// flush write
	if isVerbose {
		log.Printf("Flushing write to disk")
	}
	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}

func openFileRead(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func openFileWrite(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getPercentageRead(fileSize int64, fileByteCount int64) int {
	return int((fileByteCount * 100) / fileSize)
}

func checkDestinationFileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(1)
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Printf("Usage: %s <source> <destination>\n", os.Args[0])
		os.Exit(1)
	}

	sourceFile := args[0]
	destinationFilePath := args[1]

	// check if file exists
	if _, err := os.Stat(sourceFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// check if destination path exists
	if _, err := os.Stat(destinationFilePath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var sourceFilenName = path.Base(sourceFile)
	var destinationFile = path.Join(destinationFilePath, sourceFilenName)

	size, err := getFileSize(sourceFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f, e := openFileRead(sourceFile)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	defer f.Close()

	fdest, e := openFileWrite(destinationFile)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	var fileByteCount = int64(0)

	for {
		count, b, e := readBlock(f, int64(fileByteCount), ChunkSize)
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
		appendBlock(fdest, count, b)
		if count == 0 || count < ChunkSize {
			break
		}
		fileByteCount += int64(count)
		fmt.Printf("Read %d bytes of %d bytes (%d%%)\n", fileByteCount, size, getPercentageRead(size, fileByteCount))
	}
	fmt.Println("Done")
}
