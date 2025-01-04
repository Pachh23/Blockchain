package entity

import (
	"time"

	"gorm.io/gorm"
)

type Block struct {
	gorm.Model
	Index        int       `json:"index"`
	Timestamp    time.Time `json:"timestamp"`
	PreviousHash string    `json:"previous_hash"`
	Data         string    `json:"data"`
	Hash         string    `json:"hash"`
	Nonce        int       `json:"nonce"`      // Nonce สำหรับ Proof of Work
	Difficulty   int       `json:"difficulty"` // Difficulty สำหรับ Proof of Work
}
