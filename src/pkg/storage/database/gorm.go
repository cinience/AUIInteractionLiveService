package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"ApsaraLive/pkg/models"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormDb struct {
	db *gorm.DB
}

func NewGormDb(driver string, dsn string) (*GormDb, error) {
	g := &GormDb{}
	var err error

	slowLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			//设定慢查询时间阈值为100ms
			SlowThreshold: 100 * time.Millisecond,
			//设置日志级别，只有Warn和Info级别会输出慢查询日志
			LogLevel: logger.Warn,
		},
	)

	if driver == "sqlite3" {
		g.db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	} else {
		g.db, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn,   // DSN data source name
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{Logger: slowLogger})
	}

	if err != nil {
		return nil, err
	}

	err = g.db.AutoMigrate(&models.RoomInfo{})

	return g, err
}

func (g *GormDb) GetDB() *gorm.DB {
	return g.db.Debug()
}

func (g *GormDb) CreateRoom(r *models.RoomInfo) error {
	rst := g.GetDB().Create(r)

	return rst.Error
}

func (g *GormDb) GetRoomList(pageSize int, pageNum int, status int) (int64, []string, error) {
	offset := (pageNum - 1) * pageSize
	var rooms []models.RoomInfo
	var err error

	var totalCount int64
	if status != models.LiveStatusAll {
		err = g.GetDB().Where("status = ?", status).Model(&models.RoomInfo{}).Count(&totalCount).Error
	} else {
		err = g.GetDB().Model(&models.RoomInfo{}).Count(&totalCount).Error
	}

	if status != models.LiveStatusAll {
		err = g.GetDB().Order("created_at desc").Limit(pageSize).Offset(offset).Where("status = ?", status).Find(&rooms).Error
	} else {
		err = g.GetDB().Order("created_at desc").Limit(pageSize).Offset(offset).Find(&rooms).Error
	}

	if err != nil {
		return totalCount, nil, fmt.Errorf("list data failed. err: %w", err)
	}

	var ids []string
	for _, item := range rooms {
		ids = append(ids, item.ID)
	}

	return totalCount, ids, nil
}

func (g *GormDb) GetRoom(id string) (*models.RoomInfo, error) {
	var r models.RoomInfo
	rst := g.GetDB().Where("id = ?", id).First(&r)
	if rst.Error != nil {
		return &r, fmt.Errorf("get record failed. err: %w", rst.Error)
	}

	return &r, nil
}

func (g *GormDb) UpdateRoom(id string, r *models.RoomInfo) error {
	rst := g.GetDB().Where("id = ?", id).Updates(r)

	return rst.Error
}

func (g *GormDb) DeleteRoom(id string) error {
	return g.GetDB().Where("id = ?", id).Delete(&models.RoomInfo{}).Error
}
