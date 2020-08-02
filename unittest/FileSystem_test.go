package unittest

import (
	"climbmessage/filesystem"
	"message/managerfilesystem"
	"os"
	"testing"
)

const (
	fileName = "myfile.m"
	folderName = "testediretorio"
)

func GlobalTest(t *testing.T) {

	var fs managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)

	fs.CreateDirectory(folderName)

	_, err := os.Stat(folderName)

	if os.IsNotExist(err) {
		t.Error("Folder dont created")
	}
}

func TestCreateDirectory(t *testing.T) {

	var fs managerfilesystem.FileSystem = new(managerfilesystem.FileSystemBinaryEncoder)

	fs.CreateDirectory(folderName)

	_, err := os.Stat(folderName)

	if os.IsNotExist(err) {
		t.Error("Folder dont created")
	}

	fs.CreateFile(fileName, folderName, filesystem.MessageFileStruct{
		"tes",
		"sdsdsd",
	})

	if _, err := os.Stat(folderName + "/" + fileName); err != nil {
		t.Error("File dont created")
	}

}

func TestCreateFile(t *testing.T) {

	var fs managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)

	fs.CreateFile(fileName, folderName, filesystem.MessageFileStruct{
		"tes",
		"sdsdsd",
	})

	if _, err := os.Stat(folderName + "/" + fileName); err != nil {
		t.Error("File dont created")
	}
}

func TestOpenFile(t *testing.T)  {
	var fs managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)

	var m filesystem.MessageFileStruct
	fs.OpenFile(folderName, fileName, &m)

	if &m == nil {
		t.Error("File dont open")
	}

}

func TestDeleteFile(t *testing.T)  {

	var fs managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)
	fs.DeleteFile(folderName, fileName)

}