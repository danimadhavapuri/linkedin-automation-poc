package search

import (
	"github.com/go-rod/rod"
	"github.com/sirupsen/logrus"

	"linkedin-automation-poc/internal/stealth"
)

func RunSearch(page *rod.Page) {
	logrus.Info("Running sample search")
	page.MustNavigate("https://www.linkedin.com/search/results/people/")
	stealth.RandomScroll(page)
}
