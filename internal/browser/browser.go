package browser

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func LaunchBrowser() (*rod.Browser, *rod.Page) {
	u := launcher.New().
		Headless(false).
		Leakless(false). // ✅ DISABLE leakless.exe
		Set("disable-blink-features", "AutomationControlled").
		MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	width := rand.Intn(400) + 1200
	height := rand.Intn(300) + 800

	page := browser.MustPage("")
	page.MustSetViewport(width, height, 1, false)

	return browser, page
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
