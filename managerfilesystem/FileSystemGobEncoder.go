package managerfilesystem

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"path/filepath"
)

type FileSystemGobEncoder struct{}

func (f *FileSystemGobEncoder) DeleteFile(dir, fileName string) {
	deleteFile(dir, fileName)
}

func (f *FileSystemGobEncoder) CreateDirectory(name string) string {
	return createDirectory(name)
}

func (f *FileSystemGobEncoder) CreateFile(fileName, dirName string, fileStruct interface{}) error {

	var byteBuffer bytes.Buffer
	enc := gob.NewEncoder(&byteBuffer)

	err := enc.Encode(fileStruct)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	file, _ := os.Create(filepath.Join(RootDir + dirName, filepath.Base(fileName)))

	_, err = file.Write(byteBuffer.Bytes())

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (f *FileSystemGobEncoder) OpenFile(dirName, fileName string, v interface{}) {

	file, err := os.Open(RootDir + dirName + "/" + fileName)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	stats, _ := file.Stat()
	var size int64 = stats.Size()
	data := readNextBytes(file, int(size))
	buffer := bytes.NewBuffer(data)

	dec := gob.NewDecoder(buffer)

	err = dec.Decode(v)
	if err != nil {
		log.Fatal("decode error: ", err)
	}
}

