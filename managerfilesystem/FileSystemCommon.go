package managerfilesystem

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"os"
)

func createDirectory(name string) (string) {

	dirName := RootDir + name

	println(dirName)

	_, err := os.Stat(dirName)

	if os.IsNotExist(err) {

		if err := os.MkdirAll(dirName, 0666); err != nil {
			panic("Unable to create directory! - " + err.Error())
		}

	}

	return dirName

}

func readNextBytes(file *os.File, number int) []byte {

	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func deleteFile(dir, fileName string)  {

	path := RootDir + dir + "/" + fileName

	err := os.Remove(path)

	if err != nil {
		log.Fatal(err)
	}
}

func GetUuid() string {
	return uuid.Must(uuid.NewV4()).String()
}