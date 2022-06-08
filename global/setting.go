package global

import (
	"github.com/qiuqiu1999/go-blog/pkg/logger"
	"github.com/qiuqiu1999/go-blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	Logger *logger.Logger

	EmailSetting *setting.EmailSettingS

	JWTSetting *setting.JWTSettingS
)
