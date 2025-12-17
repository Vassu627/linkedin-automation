package auth

import (
	"errors"
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"

	"linkedin-automation/stealth"
)

// Login performs a proof-of-concept LinkedIn login flow
// This does NOT bypass captchas or 2FA
func Login(page *rod.Page, email string, password string) error {
	fmt.Println("Navigating to LinkedIn login page")

	page.MustNavigate("https://www.linkedin.com/login")
	page.MustWaitLoad()

	stealth.Think(1200, 2500)

	// locate fields
	emailInput, err := page.Element("#username")
	if err != nil {
		return errors.New("email input not found")
	}

	passwordInput, err := page.Element("#password")
	if err != nil {
		return errors.New("password input not found")
	}

	// type credentials
	stealth.HumanType(emailInput, email)
	stealth.Think(400, 800)
	stealth.HumanType(passwordInput, password)

	stealth.Think(500, 1000)

	// click submit
	loginBtn, err := page.Element("button[type=submit]")
	if err != nil {
		return errors.New("login button not found")
	}

	loginBtn.Click(proto.InputMouseButtonLeft, 1)

	stealth.Think(2500, 3500)

	// checkpoint detection (PoC)
	if page.MustHas("input[name=pin]") || page.MustHas("iframe") {
		return errors.New("security checkpoint detected (2FA / captcha)")
	}

	fmt.Println("Login attempt completed (PoC)")
	return nil
}
