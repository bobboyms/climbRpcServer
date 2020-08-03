package managerfilesystem

import (
	"bytes"
	"github.com/kelindar/binary"
	"log"
	"os"
	"path/filepath"
)

type FileSystemBinaryEncoder struct{}

func (f *FileSystemBinaryEncoder) IsFileExist(dir, fileName string) bool {
	return isFileExist(dir, fileName)
}

func (f *FileSystemBinaryEncoder) DeleteFile(dir, fileName string) {
	deleteFile(dir, fileName)
}

func (f *FileSystemBinaryEncoder) CreateDirectory(name string) string {
	return createDirectory(name)
}

//
//func (f *FileSystemBinaryEncoder) CreateFileAndDir(fileName, folderName string, fileStruct interface{})  {
//
//	_, err := f.CreateDirectory(folderName)
//
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//
//	_, err = os.Stat(folderName)
//
//	if os.IsNotExist(err) {
//		log.Fatal("Folder dont created")
//	}
//
//	f.CreateFile(fileName, folderName, fileStruct)
//
//	if _, err := os.Stat(folderName + "/" + fileName); err != nil {
//		log.Fatal("File dont created")
//	}
//
//}

func (f *FileSystemBinaryEncoder) CreateFile(fileName, dirName string, fileStruct interface{}) error {

	encoded, err := binary.Marshal(fileStruct)

	_, err = os.Stat(filepath.Join(dirName))

	if os.IsNotExist(err) {
		println("Não existe karaio")
	}

	println(filepath.Join(RootDir+dirName, filepath.Base(fileName)))

	file, err := os.Create(filepath.Join(RootDir+dirName, filepath.Base(fileName)))

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err = file.Write(encoded)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (f *FileSystemBinaryEncoder) OpenFile(dirName, fileName string, v interface{}) {

	file, err := os.Open(RootDir + dirName + "/" + fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	stats, _ := file.Stat()
	var size int64 = stats.Size()
	data := readNextBytes(file, int(size))
	buffer := bytes.NewBuffer(data)

	err = binary.Unmarshal(buffer.Bytes(), v)

	if err != nil {
		log.Fatal("decode error: ", err)
	}

}
