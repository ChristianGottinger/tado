package tado

import (
	"context"
	"github.com/gonzolino/gotado/v2"
	"github.com/procyon-projects/chrono"
	"github.com/sirupsen/logrus"
	"tado/pkg/common"
	"time"
)

type TadoAdapter struct {
	home    *gotado.Home
	context context.Context
}

const (
	clientID     = "tado-web-app"
	clientSecret = "wZaRN7rpjn3FoNyF5IFuxg9uMzYJcvOoQ8QWiIqS3hfk6gLhVlG57j5YNoZL2Rtc"
)

func Init(homename, username, password string) TadoAdapter {
	ctx := context.Background()
	client := gotado.New(clientID, clientSecret)
	user, err := client.Me(ctx, username, password)
	common.PanicOnError(err)

	home, err := user.GetHome(ctx, homename)
	common.PanicOnError(err)

	return TadoAdapter{
		home:    home,
		context: ctx,
	}
}

func (t *TadoAdapter) CheckOpenWindows(taskScheduler chrono.TaskScheduler) {
	//logrus.Infof("Open Window detection triggered")

	zones, err := t.home.GetZones(t.context)
	common.PanicOnError(err)

	for _, zone := range zones {
		//logrus.Infof("   Checking zone %s", zone.Name)
		zonestate, err := zone.GetState(t.context)
		common.PanicOnError(err)

		if zonestate.OpenWindow != nil {
			logrus.Infof("Open window detected in %s - Heating turned off", zone.Name)
			zone.SetHeatingOff(t.context)

			taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
				t.scheduleHeatingOn(zone.Name)
			}, time.Duration(zonestate.OpenWindow.RemainingTimeInSeconds))

			logrus.Infof("Heating scheduled to be reactivated in %d seconds", zonestate.OpenWindow.RemainingTimeInSeconds)
		}
	}
}

func (t *TadoAdapter) scheduleHeatingOn(zonename string) {
	zone, err := t.home.GetZone(t.context, zonename)
	common.PanicOnError(err)
	zone.ResumeSchedule(t.context)
	logrus.Infof("Heating schedule reactivated for zone %s", zonename)
}
