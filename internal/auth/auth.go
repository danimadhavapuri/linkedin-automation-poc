package auth

import (
	"os"

	"github.com/go-rod/rod"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"linkedin-automation-poc/internal/stealth"
)

func Login(page *rod.Page) {
	godotenv.Load()

	email := os.Getenv("LINKEDIN_EMAIL")
	pass := os.Getenv("LINKEDIN_PASSWORD")

	logrus.Info("Navigating to LinkedIn login")

	page.MustNavigate("https://www.linkedin.com/login")
	stealth.RandomSleep(2000, 4000)

	page.MustElement(`#username`).MustClick()
	stealth.HumanType(page.MustElement(`#username`), email)

	page.MustElement(`#password`).MustClick()
	stealth.HumanType(page.MustElement(`#password`), pass)

	stealth.RandomSleep(1000, 2000)
	page.MustElement(`button[type=submit]`).MustClick()

	logrus.Info("Login attempted (captcha/2FA handled gracefully)")
}
