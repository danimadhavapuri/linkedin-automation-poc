# LinkedIn Automation – Stealth Browser Automation (Proof of Concept)

## ⚠️ Disclaimer (Important)

**Educational Purpose Only**

This project is a technical proof-of-concept created solely for learning and evaluation purposes.  
It demonstrates browser automation, stealth techniques, and clean Go architecture.

**Terms of Service Warning**

Automating LinkedIn violates LinkedIn’s Terms of Service.  
This project must NOT be used on real accounts, deployed in production, or used for real automation.

---

## 📌 Project Overview

This project demonstrates a Go-based browser automation system using the Rod library.  
The primary focus is on implementing human-like behavior patterns, stealth mechanisms, and modular software architecture rather than full LinkedIn automation.

The automation intentionally stops at safe interaction points (login/search) to avoid misuse while still proving technical capability.

---

## 🎯 Objectives

- Demonstrate browser automation using Go
- Implement stealth and anti-detection techniques
- Follow clean, modular Go architecture
- Handle real-world automation challenges safely

---

## 🧱 Project Structure

```text
linkedin-automation-poc/
│
├── cmd/app/
│   └── main.go              # Application entry point
│
├── internal/
│   ├── browser/             # Browser launch & fingerprint control
│   ├── stealth/             # Human-like behavior & anti-detection
│   ├── auth/                # Login flow handling
│   ├── search/              # Search flow (demonstration)
│   ├── connect/             # Connection logic (stubbed)
│   ├── messaging/           # Messaging logic (stubbed)
│   ├── config/              # Centralized configuration
│   └── storage/             # State persistence (future-ready)
│
├── configs/
│   └── config.yaml
│
├── .env.example
├── README.md
└── go.mod
```

---

## 🕵️ Stealth & Anti-Bot Techniques

### Implemented in Code
1. Randomized action delays (human thinking time)
2. Human-like typing with variable speed
3. Random scrolling behavior
4. Browser automation flag disabling
5. Randomized viewport sizes

### Designed / Documented (Proof-of-Concept)
6. Mouse movement abstraction (Bezier-curve approach planned)
7. Rate limiting and action quotas
8. Business-hour activity scheduling

> Modern bot detection relies on behavioral patterns.  
> This project demonstrates how multiple subtle techniques together reduce detection probability.

---

## 🔐 Authentication Strategy

- Credentials loaded securely using environment variables
- Slow, human-like input for login fields
- Graceful handling of captcha or 2FA checkpoints
- Automation stops safely when manual intervention is required

---

## ⚙️ Configuration Management

- Centralized configuration via Go struct
- Environment-based secrets (no hardcoding)
- Safe defaults for demo execution
- Easily extendable to YAML or JSON configuration

---

## 🧪 Execution Flow

1. Launch Chromium with stealth configuration
2. Apply browser fingerprint masking
3. Navigate to LinkedIn login page
4. Perform human-like login interaction
5. Demonstrate navigation/search behavior
6. Stop execution safely (proof-of-concept boundary)

---

## 🚀 How to Run

### 1️⃣ Set environment variables
```env
LINKEDIN_EMAIL=your_email_here
LINKEDIN_PASSWORD=your_password_here
```

### 2️⃣ Run the application
```bash
go run cmd/app/main.go
```

> On first run, Rod will automatically download Chromium.

---

## 🛑 Known Limitations (Intentional)

- No real connection requests are sent
- No captcha or 2FA bypassing
- Messaging and connection modules are stubbed
- Designed strictly as a technical demonstration

These limitations are intentional to ensure ethical usage.

---

## 🧠 Key Engineering Takeaways

- Stealth automation is primarily behavioral, not just technical
- Modular architecture improves maintainability
- Automation must balance realism and safety
- OS-level security (e.g., antivirus) is a real-world challenge

---

## 🎥 Demo Video

The demo video shows:
- Project structure overview
- Stealth configuration
- Application execution
- Browser launch and login flow

---

## 📌 Final Note

This project demonstrates advanced browser automation concepts, anti-detection awareness, and clean Go engineering practices.  
It is intended strictly as a proof-of-concept for technical evaluation.
