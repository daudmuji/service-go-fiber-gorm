package repository

import (
	"golang-template-service/model"
	"gorm.io/gorm"
)

type SampleUploadPostgresRepository struct {
	Conn *gorm.DB
}

func NewSampleUploadPostgresRepository(conn *gorm.DB) *SampleUploadPostgresRepository {
	return &SampleUploadPostgresRepository{Conn: conn}
}

func (r *SampleUploadPostgresRepository) CreateDataUpload(dataUpload *model.SampleUpload) (err error, result int64) {
	tx := r.Conn.Table("tbl_temporary_upload").Create(&dataUpload)
	err = tx.Error
	if err != nil {
		return err, result
	}

	rowsAffected := tx.RowsAffected
	if rowsAffected == 0 {
		return err, rowsAffected
	}

	return nil, rowsAffected
}
