package models

import "time"

type Product struct {
	ID               uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	NamaBarang       string `json:"nama_barang"`
	Stok             int    `json:"stok"`
	JumlahTerjual    int    `json:"jumlah_terjual"`
	TanggalTransaksi string `json:"tanggal_transaksi"`
	// JenisBarang      string    `json:"jenis_barang"`
	CategoryID uint      `json:"category_id"`
	Category   Category  `json:"category" gorm:"foreignKey:CategoryID" `
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
