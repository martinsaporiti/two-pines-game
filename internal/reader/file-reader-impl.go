package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type FileReaderImpl struct {
	filePath string
}

func NewFileReader(filePath string) *FileReaderImpl {
	return &FileReaderImpl{
		filePath,
	}
}

// Load the content of the file and return a list, one element per row
func (fileReader *FileReaderImpl) GetContent() []string {
	file, err := os.Open(fileReader.filePath)
	if err != nil {
		log.Panicln("Error reading the file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content []string
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		content = append(content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Panicln("Error reading the file: ", err)
	}
	return content
}
