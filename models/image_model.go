package models

import (
	"fmt"
)

type ImageModel struct {
	Model
	FileName string `gorm:"type:varchar(64);not null" json:"fileName"`  // 文件名
	FilePath string `gorm:"type:varchar(255);not null" json:"filePath"` // 文件路径
	Size     int64  `gorm:"type:bigint;not null" json:"size"`           // 文件大小
	Hash     string `gorm:"size:32" json:"hash"`                        // 文件hash
}

func (i ImageModel) WebPath() string {
	return fmt.Sprintf("/static/%s", i.FilePath)
}
