package managerindex

type ManagerIndex interface {
	CreateIndex() error
	AddItem(fileName string) error
	RemoveItem(index uint16) error
	GetItem(index uint16) (string, error)
	GetLastIndex() (uint16, error)
	GetFirstIndex() (uint16, error)
}
