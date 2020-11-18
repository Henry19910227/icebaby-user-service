package upload

import (
	"github.com/spf13/viper"
)

// GPUploadSetting ...
type GPUploadSetting struct {
	vp *viper.Viper
}

// NewUploadSetting ...
func NewUploadSetting(filename string) (*GPUploadSetting, error) {
	vp := viper.New()
	vp.SetConfigFile(filename)
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &GPUploadSetting{vp}, nil
}

// GetUploadSavePath ...
func (setting *GPUploadSetting) GetUploadSavePath() string {
	return setting.vp.GetString("App.UploadSavePath")
}

// GetUploadImageAllowExts ...
func (setting *GPUploadSetting) GetUploadImageAllowExts() []string {
	return setting.vp.GetStringSlice("App.UploadImageAllowExt")
}

// GetUploadImageMaxSize ...
func (setting *GPUploadSetting) GetUploadImageMaxSize() int {
	return setting.vp.GetInt("App.UploadImageMaxSize")
}
