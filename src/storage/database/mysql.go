package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"skyExplorerBack/src/constant/config"
	"skyExplorerBack/src/utils/logger"
)

var Client *gorm.DB

func init() {
	var err error
	var cfg gorm.Config
	cfg = gorm.Config{
		PrepareStmt: true,
		Logger:      logger.NewCustomLogger(gormLogger.Info),
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix: "test",
		//},
		ConnPool: nil,
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.EnvCfg.MysqlUser, config.EnvCfg.MysqlPassword,
		config.EnvCfg.MysqlHost, config.EnvCfg.MysqlPort, config.EnvCfg.MysqlDatabase)
	if Client, err = gorm.Open(mysql.Open(dsn), &cfg); err != nil {
		panic(err)
	}

	TableAutoMigrate()
}

func TableAutoMigrate() {
	if !config.EnvCfg.AutoMigrate {
		logger.Info("未启用迁移数据库")
		return
	}

	// 自定义连接表，包括反向引用
	if err := Client.SetupJoinTable(&dbmodels.User{},
		"WorkInfo", &dbmodels.UserWorkHistory{}); err != nil {
		panic(err)
	}

	// 反向引用
	if err := Client.SetupJoinTable(&dbmodels.WorkInfo{},
		"Users", &dbmodels.UserWorkHistory{}); err != nil {
		panic(err)
	}

	if err := Client.AutoMigrate(&dbmodels.User{},
		&dbmodels.WorkInfo{}, &dbmodels.UserInfo{},
		&dbmodels.WithdrawInfo{}, &dbmodels.CommunityInfo{},
		&dbmodels.CompanyInfo{}, &dbmodels.Advertisement{}, &dbmodels.Admin{},
		&dbmodels.WorkTemplate{}, &dbmodels.UserBalanceChangeLog{}); err != nil {
		panic(err)
		return
	}
}
