// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameIndexCollect = "chii_index_collects"

// IndexCollect 目录收藏
type IndexCollect struct {
	CltID       uint32 `gorm:"column:idx_clt_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true"`
	IndexID     uint32 `gorm:"column:idx_clt_mid;type:mediumint(8) unsigned;not null;comment:目录ID"`  // 目录ID
	UserID      uint32 `gorm:"column:idx_clt_uid;type:mediumint(8) unsigned;not null;comment:用户UID"` // 用户UID
	CreatedTime uint32 `gorm:"column:idx_clt_dateline;type:int(10) unsigned;not null"`
}

// TableName IndexCollect's table name
func (*IndexCollect) TableName() string {
	return TableNameIndexCollect
}
