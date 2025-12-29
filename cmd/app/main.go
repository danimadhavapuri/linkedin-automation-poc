package main

import (
	"github.com/sirupsen/logrus"

	"linkedin-automation-poc/internal/auth"
	"linkedin-automation-poc/internal/browser"
	"linkedin-automation-poc/internal/search"
)

func main() {
	logrus.Info("Starting LinkedIn Automation POC")

	_, page := browser.LaunchBrowser()

	logrus.Info("Browser launched")

	auth.Login(page)
	search.RunSearch(page)

	logrus.Info("Execution finished")
}
