package repo

import "github.com/cdlinkin/marketplace/internal/models"

type FileUserRepo struct {
	filename string
}

func NewFileUserRepo(filename string) *FileUserRepo {
	return &FileUserRepo{filename: filename}
}

func (f *FileUserRepo) Save() {}
func (f *FileUserRepo) Get()  {}
func (f *FileUserRepo) GetByEmail(email string) (*models.User, error) {
	return &models.User{}, nil
}
