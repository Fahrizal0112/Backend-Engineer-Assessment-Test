package repository

import (
	"banking-service/internal/models"
	"banking-service/pkg/logger"
	"fmt"

	"gorm.io/gorm"
)

type NasabahRepository struct {
	db *gorm.DB
}

func NewNasabahRepository(db *gorm.DB) *NasabahRepository {
	return &NasabahRepository{
		db: db,
	}
}

func (r *NasabahRepository) CreateNasabah(nasabah models.Nasabah) error {
	result := r.db.Create(&nasabah)
	if result.Error != nil {
		logger.Error(fmt.Sprintf("Failed to create nasabah: %v", result.Error),
			"nama", nasabah.Nama, "nik", nasabah.NIK)
		return result.Error
	}
	logger.Info(fmt.Sprintf("Created nasabah with ID %d", nasabah.ID),
		"nama", nasabah.Nama, "no_rekening", nasabah.NoRekening)
	return nil
}

func (r *NasabahRepository) GetByNoRekening(noRekening string) (*models.Nasabah, error) {
	var nasabah models.Nasabah

	result := r.db.Where("no_rekening = ?", noRekening).First(&nasabah)
	if result.Error != nil {
		logger.Warning(fmt.Sprintf("Nasabah not found with no_rekening %s", noRekening))
		return nil, result.Error
	}
	logger.Info(fmt.Sprintf("Retrieved nasabah with no_rekening %s", noRekening),
		"id", nasabah.ID, "nama", nasabah.Nama)
	return &nasabah, nil
}

func (r *NasabahRepository) GetByNIK(nik string) (*models.Nasabah, error) {
	var nasabah models.Nasabah

	result := r.db.Where("nik = ?", nik).First(&nasabah)
	if result.Error != nil {
		return nil, result.Error
	}
	return &nasabah, nil
}

func (r *NasabahRepository) GetByNoHP(noHP string) (*models.Nasabah, error) {
	var nasabah models.Nasabah

	result := r.db.Where("no_hp = ?", noHP).First(&nasabah)
	if result.Error != nil {
		return nil, result.Error
	}
	return &nasabah, nil
}

func (r *NasabahRepository) UpdateSaldo(nasabah *models.Nasabah) error {
	result := r.db.Save(nasabah)
	if result.Error != nil {
		logger.Error(fmt.Sprintf("Failed to update saldo for nasabah with ID %d: %v", 
			nasabah.ID, result.Error))
		return result.Error
	}
	
	logger.Info(fmt.Sprintf("Updated saldo for nasabah with ID %d", nasabah.ID), 
		"no_rekening", nasabah.NoRekening, "new_saldo", nasabah.Saldo)
	return nil
}

func (r *NasabahRepository) CountNasabah() (int64, error) {
	var count int64
	result := r.db.Model(&models.Nasabah{}).Count(&count)
	if result.Error != nil {
		logger.Error(fmt.Sprintf("Failed to count nasabah: %v", result.Error))
		return 0, result.Error
	}
	return count, nil
}