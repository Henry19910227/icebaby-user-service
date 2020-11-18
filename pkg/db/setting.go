package db

import (
	"github.com/spf13/viper"
)

// MysqlViperSetting ...
type MysqlViperSetting struct {
	vp   *viper.Viper
	mode string
}

// NewMysqlSetting ...
func NewMysqlSetting(viperTool *viper.Viper) *MysqlViperSetting {
	return &MysqlViperSetting{viperTool, viperTool.GetString("Server.RunMode")}
}

// GetUserName ...
func (setting *MysqlViperSetting) GetUserName() string {
	if setting.mode == "debug" {
		return setting.vp.GetString("Database.Debug.UserName")
	}
	return setting.vp.GetString("Database.Release.UserName")
}

// GetPassword ...
func (setting *MysqlViperSetting) GetPassword() string {
	if setting.mode == "debug" {
		return setting.vp.GetString("Database.Debug.Password")
	}
	return setting.vp.GetString("Database.Release.Password")
}

// GetHost ...
func (setting *MysqlViperSetting) GetHost() string {
	if setting.mode == "debug" {
		return setting.vp.GetString("Database.Debug.Host")
	}
	return setting.vp.GetString("Database.Release.Host")
}

// GetDatabase ...
func (setting *MysqlViperSetting) GetDatabase() string {
	if setting.mode == "debug" {
		return setting.vp.GetString("Database.Debug.DBName")
	}
	return setting.vp.GetString("Database.Release.DBName")
}
