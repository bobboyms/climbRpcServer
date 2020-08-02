package managerfilesystem

const (
	FileExtension = ".m"
	RootDir       = "./"
)

type FileSystem interface {
	CreateDirectory(name string) string
	CreateFile(fileName, dirName string, fileStruct interface{}) error
	OpenFile(dirName, fileName string, v interface{})
	DeleteFile(dir, fileName string)

	//CreateFileAndDir(fileName, folderName string, fileStruct interface{})
}
