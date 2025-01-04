package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"

	"blockchain.com/bc-67/config"
	"blockchain.com/bc-67/entity"
	"gorm.io/gorm"
)

type Blockchain struct {
	Chain []entity.Block
}

// คำนวณ Hash ของบล็อก
func calculateHash(block entity.Block) string {
	record := string(block.Index) + block.Timestamp.String() + block.PreviousHash + block.Data + string(block.Nonce)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

// ตรวจสอบว่า Hash ตรงกับ Difficulty หรือไม่
func isValidHash(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

// ฟังก์ชันในการคำนวณ Nonce จนกว่าจะได้ Hash ที่ตรงตามเงื่อนไข
func (bc *Blockchain) mineBlock(block *entity.Block) {
	for !isValidHash(block.Hash, block.Difficulty) {
		block.Nonce++
		block.Hash = calculateHash(*block)
	}
}

// สร้าง Blockchain ใหม่และตั้งค่า Genesis Block
func NewBlockchain() *Blockchain {
	genesisBlock := entity.Block{
		Index:        0,
		Timestamp:    time.Now(),
		PreviousHash: "0", // Genesis block ไม่มี previous hash
		Data:         "Genesis Block",
		Hash:         "",
		Nonce:        0,
		Difficulty:   4, // ตั้งค่า Difficulty
	}

	genesisBlock.Hash = calculateHash(genesisBlock) // คำนวณ Hash ของ Genesis Block
	return &Blockchain{
		Chain: []entity.Block{genesisBlock}, // เริ่มต้นด้วย Genesis Block
	}
}

// SaveBlockToDB บันทึก Block ลงในฐานข้อมูล
func SaveBlockToDB(db *gorm.DB, block entity.Block) error {
	if err := db.Create(&block).Error; err != nil {
		return err
	}
	return nil
}

func (bc *Blockchain) CreateBlock(data string) (entity.Block, error) {
	previousBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := entity.Block{
		Index:        len(bc.Chain),
		Timestamp:    time.Now(),
		PreviousHash: previousBlock.Hash,
		Data:         data,
		Hash:         "",
		Nonce:        0,
		Difficulty:   4,
	}

	// คำนวณ Nonce และ Hash
	bc.mineBlock(&newBlock)

	// เพิ่มบล็อกใหม่ใน Blockchain
	bc.Chain = append(bc.Chain, newBlock)

	// บันทึกบล็อกลงในฐานข้อมูล
	if err := SaveBlockToDB(config.DB(), newBlock); err != nil {
		return entity.Block{}, err
	}

	return newBlock, nil
}

// ตรวจสอบว่า Blockchain ถูกต้องหรือไม่
func (bc *Blockchain) VerifyChain() bool {
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		previousBlock := bc.Chain[i-1]

		// ตรวจสอบว่า Hash ของบล็อกปัจจุบันถูกต้องหรือไม่
		if currentBlock.Hash != calculateHash(currentBlock) {
			return false
		}

		// ตรวจสอบว่า PreviousHash ตรงกับ hash ของบล็อกก่อนหน้า
		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}

		// ตรวจสอบว่า Hash ตรงตาม Difficulty หรือไม่
		if !isValidHash(currentBlock.Hash, currentBlock.Difficulty) {
			return false
		}
	}
	return true
}
