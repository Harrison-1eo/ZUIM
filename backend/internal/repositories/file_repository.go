package repositories

import "backend/internal/models"

type FileRepository struct {
}

func NewFileRepository() *FileRepository {
	return &FileRepository{}
}

func (repo *FileRepository) CreateFile(file models.File) (*models.File, error) {
	if err := db.Model(models.File{}).Create(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (repo *FileRepository) GetFileByPath(path string) (*models.File, error) {
	var file models.File
	if err := db.Model(models.File{}).Where("file_path = ?", path).First(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}
