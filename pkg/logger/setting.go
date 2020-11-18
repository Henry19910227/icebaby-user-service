package logger

import (
	"github.com/spf13/viper"
)

// GPLogSetting ...
type GPLogSetting struct {
	vp   *viper.Viper
	mode string
}

// NewGPLogSetting ...
func NewGPLogSetting(viperTool *viper.Viper) *GPLogSetting {
	return &GPLogSetting{viperTool, viperTool.GetString("Server.RunMode")}
}

// GetLogFilePath ...
func (setting *GPLogSetting) GetLogFilePath() string {
	if setting.mode == "debug" {
		return setting.vp.GetString("Log.Debug.Path")
	}
	return setting.vp.GetString("Log.Release.Path")

}

// GetLogFileName ...
func (setting *GPLogSetting) GetLogFileName() string {
	return setting.vp.GetString("Log.FileName")
}

// GetLogFileExt ...
func (setting *GPLogSetting) GetLogFileExt() string {
	return setting.vp.GetString("Log.FileExt")
}

// GetRunMode ...
func (setting *GPLogSetting) GetRunMode() string {
	return setting.vp.GetString("Server.RunMode")
}
