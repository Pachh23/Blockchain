package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int       `json:"index"`
	Timestamp    time.Time `json:"timestamp"`
	PreviousHash string    `json:"previous_hash"`
	Data         string    `json:"data"`
	Hash         string    `json:"hash"`
}

type Blockchain struct {
	Chain []Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Now(),
		PreviousHash: "0",
		Data:         "Genesis Block",
		Hash:         "",
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	return &Blockchain{
		Chain: []Block{genesisBlock},
	}
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp.String() + block.PreviousHash + block.Data
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func (bc *Blockchain) CreateBlock(data string) Block {
	previousBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := Block{
		Index:        len(bc.Chain),
		Timestamp:    time.Now(),
		PreviousHash: previousBlock.Hash,
		Data:         data,
		Hash:         "",
	}
	newBlock.Hash = calculateHash(newBlock)
	bc.Chain = append(bc.Chain, newBlock)
	return newBlock
}

// VerifyChain ตรวจสอบว่า blockchain นี้ถูกต้องหรือไม่
func (bc *Blockchain) VerifyChain() bool {
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		previousBlock := bc.Chain[i-1]
		// ตรวจสอบว่า hash ของบล็อกปัจจุบันถูกต้องหรือไม่
		if currentBlock.Hash != calculateHash(currentBlock) {
			return false
		}
		// ตรวจสอบว่า PreviousHash ตรงกับ hash ของบล็อกก่อนหน้า
		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}
