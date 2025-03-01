package env

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetGinReleaseMode(thisService string) {
	s, _ := Env.GetUserService(thisService)

	level := s.GetParamString("debug_level")
	if level == "debug" {
		zap.L().Info("setting gin debug level to debug")
		gin.SetMode(gin.DebugMode)
		os.Setenv("GIN_MODE", "debug")
	} else {
		zap.L().Info("setting gin debug level to release")
		gin.SetMode(gin.ReleaseMode)
		os.Setenv("GIN_MODE", "release")
	}
}
