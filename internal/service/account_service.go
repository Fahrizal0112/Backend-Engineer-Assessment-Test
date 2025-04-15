package service

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"banking-service/internal/models"
	"banking-service/internal/repository"
	"banking-service/pkg/logger"
)

type AccountService struct {
	repo *repository.NasabahRepository
}

func NewAccountService(repo *repository.NasabahRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) RegisterNasabah(request models.DaftarRequest) (*models.DaftarResponse, error) {
	logger.Info("Processing registration request", "nama", request.Nama, "nik", request.NIK)

	if request.Nama == "" || request.NIK == "" || request.NoHP == "" {
		return nil, errors.New("semua field harus diisi")
	}

	existingNasabah, err := s.repo.GetByNIK(request.NIK)
	if err == nil && existingNasabah != nil {
		logger.Warning("NIK already registered", "nik", request.NIK)
		return nil, errors.New("NIK sudah terdaftar")
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(fmt.Sprintf("Error checking NIK: %v", err))
		return nil, err
	}

	existingNasabah, err = s.repo.GetByNoHP(request.NoHP)
	if err == nil && existingNasabah != nil {
		logger.Warning("Phone number already registered", "no_hp", request.NoHP)
		return nil, errors.New("nomor HP sudah terdaftar")
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(fmt.Sprintf("Error checking NoHP: %v", err))
		return nil, err
	}

	noRekening, err := s.generateNoRekening()
	if err != nil {
		logger.Error(fmt.Sprintf("Error generating account number: %v", err))
		return nil, err
	}

	nasabah := models.Nasabah{
		Nama:       request.Nama,
		NIK:        request.NIK,
		NoHP:       request.NoHP,
		NoRekening: noRekening,
		Saldo:      0,
	}

	if err := s.repo.CreateNasabah(nasabah); err != nil {
		logger.Error(fmt.Sprintf("Error creating nasabah: %v", err))
		return nil, err
	}

	logger.Info("Registration successful", "nama", request.Nama, "no_rekening", noRekening)
	return &models.DaftarResponse{NoRekening: noRekening}, nil
}

func (s *AccountService) Deposit(request models.TabungRequest) (*models.SaldoResponse, error) {
	logger.Info("Processing deposit request", "no_rekening", request.NoRekening, "nominal", request.Nominal)

	if request.Nominal <= 0 {
		logger.Warning("Invalid deposit amount", "nominal", request.Nominal)
		return nil, errors.New("nominal harus lebih dari 0")
	}

	nasabah, err := s.repo.GetByNoRekening(request.NoRekening)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warning("Account not found", "no_rekening", request.NoRekening)
			return nil, errors.New("nomor rekening tidak ditemukan")
		}
		logger.Error(fmt.Sprintf("Error finding account: %v", err))
		return nil, err
	}

	nasabah.Saldo += request.Nominal

	if err := s.repo.UpdateSaldo(nasabah); err != nil {
		logger.Error(fmt.Sprintf("Error updating balance: %v", err))
		return nil, err
	}

	logger.Info("Deposit successful", "no_rekening", request.NoRekening,
		"amount", request.Nominal, "new_balance", nasabah.Saldo)
	return &models.SaldoResponse{Saldo: nasabah.Saldo}, nil
}

func (s *AccountService) Withdraw(request models.TarikRequest) (*models.SaldoResponse, error) {
	logger.Info("Processing withdrawal request", "no_rekening", request.NoRekening, "nominal", request.Nominal)

	if request.Nominal <= 0 {
		logger.Warning("Invalid withdrawal amount", "nominal", request.Nominal)
		return nil, errors.New("nominal harus lebih dari 0")
	}

	nasabah, err := s.repo.GetByNoRekening(request.NoRekening)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warning("Account not found", "no_rekening", request.NoRekening)
			return nil, errors.New("nomor rekening tidak ditemukan")
		}
		logger.Error(fmt.Sprintf("Error finding account: %v", err))
		return nil, err
	}

	if nasabah.Saldo < request.Nominal {
		logger.Warning("Insufficient balance", "no_rekening", request.NoRekening,
			"requested", request.Nominal, "available", nasabah.Saldo)
		return nil, errors.New("saldo tidak cukup")
	}

	nasabah.Saldo -= request.Nominal

	if err := s.repo.UpdateSaldo(nasabah); err != nil {
		logger.Error(fmt.Sprintf("Error updating balance: %v", err))
		return nil, err
	}

	logger.Info("Withdrawal successful", "no_rekening", request.NoRekening,
		"amount", request.Nominal, "new_balance", nasabah.Saldo)
	return &models.SaldoResponse{Saldo: nasabah.Saldo}, nil
}

func (s *AccountService) GetBalance(noRekening string) (*models.SaldoResponse, error) {
	logger.Info("Processing balance inquiry", "no_rekening", noRekening)

	nasabah, err := s.repo.GetByNoRekening(noRekening)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warning("Account not found", "no_rekening", noRekening)
			return nil, errors.New("nomor rekening tidak ditemukan")
		}
		logger.Error(fmt.Sprintf("Error finding account: %v", err))
		return nil, err
	}

	logger.Info("Balance inquiry successful", "no_rekening", noRekening, "balance", nasabah.Saldo)
	return &models.SaldoResponse{Saldo: nasabah.Saldo}, nil
}

func (s *AccountService) generateNoRekening() (string, error) {
	count, err := s.repo.CountNasabah()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%010d", count+1), nil
}
