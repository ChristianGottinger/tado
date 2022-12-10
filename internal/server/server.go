package server

import (
	"context"
	"github.com/procyon-projects/chrono"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"tado/internal/cli_flags"
	"tado/internal/tado"
	"tado/pkg/common"
	"time"
)

func Start(viper *viper.Viper) {

	logrus.Infof("Starting tado open windows detection")
	taskScheduler := chrono.NewDefaultTaskScheduler()
	homename := viper.GetString(cli_flags.Homename)
	username := viper.GetString(cli_flags.Username)
	password := viper.GetString(cli_flags.Password)
	adapter := tado.Init(homename, username, password)

	_, err := taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
		openWindowError := adapter.CheckOpenWindows()
		if openWindowError != nil {
			logrus.Errorf("Error during open windows check: %s", openWindowError.Error())
		}
	}, 10*time.Second)

	if err == nil {
		logrus.Infof("Task has been scheduled successfully")
	}

	common.Wait()
}
