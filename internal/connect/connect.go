package connect

import "github.com/sirupsen/logrus"

func SendConnection() {
	logrus.Info("Connection request simulated (rate-limited)")
}
