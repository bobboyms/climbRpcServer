package managerindex

type ManagerIndex interface {
	CreateIndex(indexFileName string) error
	AddItem(fileName string, indexFileName string) error
	RemoveItem(index uint16, indexFileName string) error
	GetItem(index uint16, indexFileName string) (string, error)
	GetIndex(id string, indexFileName string) (uint16, error)
	GetLastIndex(indexFileName string) (uint16, error)
	GetFirstIndex(indexFileName string) (uint16, error)
}
