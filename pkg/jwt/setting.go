package jwt

import (
	"time"

	"github.com/spf13/viper"
)

// GPJWTSetting ...
type GPJWTSetting struct {
	vp *viper.Viper
}

// NewJWTSetting ...
func NewJWTSetting(filename string) (*GPJWTSetting, error) {
	vp := viper.New()
	vp.SetConfigFile(filename)
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &GPJWTSetting{vp}, nil
}

// GetTokenSecret ...
func (setting *GPJWTSetting) GetTokenSecret() string {
	return setting.vp.GetString("JWT.Secret")
}

// GetIssuer ...
func (setting *GPJWTSetting) GetIssuer() string {
	return setting.vp.GetString("JWT.Issuer")
}

// GetExpire ...
func (setting *GPJWTSetting) GetExpire() time.Duration {
	return setting.vp.GetDuration("JWT.Expire") * time.Hour
}
