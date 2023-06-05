package files

import "myGoApp/storage"

type Storage struct {
	basePath string
}

func (s Storage) Save(p *storage.Messege) error {
	//TODO implement me
	panic("implement me")
}

func (s Storage) IsExist(p *storage.Messege) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}