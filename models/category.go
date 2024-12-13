package models

import "time"

type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	JenisBarang string    `json:"jenis_barang"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
