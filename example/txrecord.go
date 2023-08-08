package example

import (
	"context"

	"gorm.io/gorm"
)

const (
	HangingStatus    = 1
	SuccessfulStatus = 2
	FailedStatus     = 3
)

type TXRecordPO struct {
	gorm.Model
	Status               int    `gorm:"status"`
	ComponentTryStatuses string `gorm:"componentTryStatuses"`
}

func (t TXRecordPO) TableName() string {
	return "tx_record"
}

type ComponentTryStatus struct {
	ComponentID string `json:"componentID"`
	TryStatus   string `json:"tryStatus"`
}

type TXRecordDAO struct {
	db *gorm.DB
}

func NewTXRecordDAO(db *gorm.DB) *TXRecordDAO {
	return &TXRecordDAO{
		db: db,
	}
}

func (t *TXRecordDAO) GetTXRecords(ctx context.Context, opts ...QueryOption) ([]*TXRecordPO, error) {
	db := t.db.WithContext(ctx).Model(&TXRecordPO{})
	for _, opt := range opts {
		db = opt(db)
	}

	var records []*TXRecordPO
	return records, db.Scan(&records).Error
}

func (t *TXRecordDAO) CreateTXRecord(ctx context.Context, record *TXRecordPO) (uint, error) {
	return record.ID, t.db.WithContext(ctx).Model(&TXRecordPO{}).Create(record).Error
}

func (t *TXRecordDAO) UpdateTXRecord(ctx context.Context, record *TXRecordPO) error {
	return t.db.WithContext(ctx).Model(&TXRecordPO{}).Updates(record).Error
}
