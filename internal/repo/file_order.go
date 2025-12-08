// save | get | listbyuser
package repo

type FileOrderRepo struct {
	filename string
}

func NewFileOrderRepo(filename string) *FileOrderRepo {
	return &FileOrderRepo{filename: filename}
}

func (f *FileOrderRepo) Save()       {}
func (f *FileOrderRepo) Get()        {}
func (f *FileOrderRepo) ListByUser() {}
