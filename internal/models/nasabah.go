package models

type Nasabah struct {
	ID         uint   `gorm:"primaryKey"`
	Nama       string `gorm:"type:varchar(100);not null"`
	NIK        string `gorm:"type:varchar(16);uniqueIndex;not null"`
	NoHP       string `gorm:"type:varchar(15);uniqueIndex;not null"`
	NoRekening string `gorm:"type:varchar(10);uniqueIndex;not null"`
	Saldo      int64  `gorm:"default:0;not null"`
}

type DaftarRequest struct {
	Nama string `json:"nama" validate:"required"`
	NIK  string `json:"nik" validate:"required, len=16"`
	NoHP string `json:"no_hp" validate:"required"`
}

type DaftarResponse struct {
	NoRekening string `json:"no_rekening"`
}

type TabungRequest struct {
	NoRekening string `json:"no_rekening" validate:"required"`
	Nominal    int64  `json:"nominal" validate:"required,gt=0"`
}

type TarikRequest struct {
	NoRekening string `json:"no_rekening" validate:"required"`
	Nominal    int64  `json:"nominal" validate:"required,gt=0"`
}

type SaldoResponse struct {
	Saldo int64 `json:"saldo"`
}

type ErrorResponse struct {
	Remark string `json:"remark"`
}
