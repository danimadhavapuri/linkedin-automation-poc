package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// 1️⃣ Randomized delays (human thinking time)
func RandomSleep(min, max int) {
	time.Sleep(time.Duration(rand.Intn(max-min)+min) * time.Millisecond)
}

// 2️⃣ Human-like typing with variable speed
func HumanType(el *rod.Element, text string) {
	for _, ch := range text {
		el.MustInput(string(ch))
		RandomSleep(80, 200)
	}
}

// 3️⃣ Random scrolling behavior (Rod-safe)
func RandomScroll(page *rod.Page) {
	scrollCount := rand.Intn(4) + 3
	for i := 0; i < scrollCount; i++ {
		page.Mouse.Scroll(0, float64(rand.Intn(400)+200), 1)
		RandomSleep(300, 800)
	}
}

// 4️⃣ Mouse movement (documented stealth – runtime safe)
func MoveMouse(page *rod.Page) {
	// NOTE:
	// Mouse movement is intentionally abstracted here.
	// In production-grade automation, this would be implemented
	// using low-level CDP mouse events for Bézier-curve movement.
	RandomSleep(200, 600)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
