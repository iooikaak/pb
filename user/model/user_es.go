package model

import (
	"github.com/olivere/elastic"
	"time"
)

type ESUser struct {
	UserID        int64             `json:"-"`
	OpenID        int64             `json:"openid"`
	Gender        byte              `json:"gender"`
	Age           int32             `json:"age"`
	Birthday      time.Time         `json:"birthday"`
	Birthday2     string            `json:"birthday2"`
	Games         []int32           `json:"games"`
	RegisterTime  time.Time         `json:"register_time"`
	LastLoginTime time.Time         `json:"last_login_time"`
	AppVersion    string            `json:"app_version"`
	OS            string            `json:"os"`
	DeviceModel   string            `json:"device_model"`
	TokenType     byte              `json:"token_type"`
	DeviceToken   string            `json:"device_token"`
	Location      *elastic.GeoPoint `json:"location"`
	LocationDesc  string            `json:"location_desc"`
	Tags          []string          `json:"tags"`
	AppID         string            `json:"app_id"`
}
