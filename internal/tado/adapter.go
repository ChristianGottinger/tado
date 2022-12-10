package tado

import (
	"context"
	"github.com/gonzolino/gotado/v2"
	"github.com/sirupsen/logrus"
	"tado/pkg/common"
)

type TadoAdapter struct {
	context  context.Context
	user     gotado.User
	homename string
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

	return TadoAdapter{
		homename: homename,
		user:     *user,
		context:  ctx,
	}
}

func (t *TadoAdapter) CheckOpenWindows() error {
	logrus.Infof("Open Window detection triggered")

	home, err := t.user.GetHome(t.context, t.homename)
	if err != nil {
		return err
	}

	zones, err := home.GetZones(t.context)
	if err != nil {
		return err
	}

	for _, zone := range zones {
		logrus.Infof("   Checking zone %s", zone.Name)
		zonestate, err := zone.GetState(t.context)
		if err != nil {
			return err
		}

		if zonestate.OpenWindowDetected {
			logrus.Infof("Open window detected in %s - Heating turned off", zone.Name)
			err := zone.OpenWindow(t.context)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
