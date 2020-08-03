package unittest

import (
	"message/managerfilesystem"
	"message/managerindex"
	"os"
	"testing"
)

func TestCreateIndex(t *testing.T) {

	var fileSystem managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)

	fileSystem.CreateDirectory(folderName)

	var index managerindex.ManagerIndex = new(managerindex.ManagerIndexImp).
		Init(folderName, fileSystem)

	index.CreateIndex()

	if _, err := os.Stat(folderName + "/INDEX.m"); err != nil {
		t.Error("File dont created")
	}

}

func TestCreateCreate(t *testing.T) {

	var fileSystem managerfilesystem.FileSystem = new(managerfilesystem.FileSystemGobEncoder)

	var index managerindex.ManagerIndex = new(managerindex.ManagerIndexImp).
		Init(folderName, fileSystem)

	index.AddItem("xxx-1111-tps14")
	if _, err := index.GetLastIndex(); err != nil {
		t.Error("Index error")
	}

	if _, err := index.GetFirstIndex(); err != nil {
		t.Error("Index error")
	}

	item, _ := index.GetFirstIndex()

	if _, err := index.GetItem(item); err != nil {
		t.Error("Index error")
	}

	index.RemoveItem(item)

}
