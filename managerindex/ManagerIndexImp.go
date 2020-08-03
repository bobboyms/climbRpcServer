package managerindex

import (
	"errors"
	"log"
	"message/managerfilesystem"
	"message/model"
	"sort"
)

type ManagerIndexImp struct {
	dir        string
	fileSystem managerfilesystem.FileSystem
}

func (m ManagerIndexImp) Init(dir string, fileSystem managerfilesystem.FileSystem) *ManagerIndexImp {
	m.dir = dir
	m.fileSystem = fileSystem
	return &m
}

func (m ManagerIndexImp) CreateIndex(indexFileName string) error {

	if !m.fileSystem.IsFileExist(m.dir, indexFileName) {

		err := m.fileSystem.CreateFile(indexFileName, m.dir, model.IndexFileStruct{
			Count: 0,
			Itens: make(map[uint16]string),
		})

		if err != nil {
			log.Fatal("Not created INDEX.m ", err.Error())
			panic(err)
		}

		return err
	}

	return nil

}

func (m ManagerIndexImp) GetIndexStruct(indexFileName string) (model.IndexFileStruct, error) {

	var i model.IndexFileStruct
	if m.fileSystem.OpenFile(m.dir, indexFileName, &i); &i == nil {
		return i, errors.New("Index not found")
	}

	return i, nil
}

func (m ManagerIndexImp) AddItem(fileName string, indexFileName string) error {

	i, err := m.GetIndexStruct(indexFileName)

	last := i.Count + 1
	i.Count = last
	i.Itens[last] = fileName

	m.fileSystem.CreateFile(indexFileName, m.dir, &i)

	return err

}

func (m ManagerIndexImp) RemoveItem(index uint16, indexFileName string) error {

	i, err := m.GetIndexStruct(indexFileName)

	delete(i.Itens, index)

	m.fileSystem.CreateFile(indexFileName, m.dir, &i)

	return err
}

func (m ManagerIndexImp) GetLastIndex(indexFileName string) (uint16, error) {

	i, err := m.GetIndexStruct(indexFileName)

	return i.Count, err
}

func (m ManagerIndexImp) GetFirstIndex(indexFileName string) (uint16, error) {

	i, _ := m.GetIndexStruct(indexFileName)

	return func(items map[uint16]string) (uint16, error) {
		keys := make([]int, 0, len(items))
		for k := range items {
			keys = append(keys, int(k))
		}
		sort.Ints(keys)

		if len(keys) > 0 {
			return uint16(keys[0]), nil
		}

		return 0, errors.New("Does not contain active index")
	}(i.Itens)

}

func (m ManagerIndexImp) GetItem(index uint16, indexFileName string) (string, error) {
	i, err := m.GetIndexStruct(indexFileName)
	return i.Itens[index], err
}
