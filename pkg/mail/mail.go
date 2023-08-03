package mail

import (
	"encoding/base64"
	"fmt"
	"news-api/internal/user/domain"
	"news-api/pkg/errs"
	"news-api/pkg/jwthelper"
	"os"
	"strconv"
	"time"

	"github.com/skip2/go-qrcode"
	"gopkg.in/gomail.v2"
)

func GenerateWIB(t time.Time) time.Time {
	wib, err := time.LoadLocation("Asia/Jakarta")
	if err == nil {
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), wib)
	}
	return t
}

func Verify_mail(form *domain.User) errs.Error {
	var link, text string
	port, _ := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	expires := form.CreatedAt.Add(time.Hour * 24)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("CONFIG_SENDER_NAME"))
	mailer.SetHeader("To", form.Email)
	mailer.SetHeader("Subject", "Code Verification")
	link = os.Getenv("RUNNING_HOST") + "authentication/verify/" + form.Id.String()
	text = "Please, click this link before " + expires.Format(time.RFC1123)
	mailer.SetBody("text/html", "Hello, "+form.Name+"<br>have a nice day, here's your credentials : <br> Email : "+form.Email+"<br> Password : "+form.Password+"<br><br>"+"Here's your verification link, click link<br><a href='"+link+"'>Click here</a><br>"+text)

	dialer := gomail.NewDialer(
		os.Getenv("CONFIG_SMTP_HOST"),
		port,
		os.Getenv("CONFIG_AUTH_EMAIL"),
		os.Getenv("CONFIG_AUTH_PASSWORD"),
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return errs.Wrap(err)
	}

	return nil
}

func Upgrade_mail(form *jwthelper.JWTClaim) errs.Error {
	var link string
	port, _ := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("CONFIG_SENDER_NAME"))
	mailer.SetHeader("To", form.Email)
	mailer.SetHeader("Subject", "Code Verification")
	link = os.Getenv("RUNNING_HOST") + "authentication/upgrade/" + form.Id.String()
	var png []byte
	png, _ = qrcode.Encode(link, qrcode.Medium, 256)
	out := base64.StdEncoding.EncodeToString(png)
	linkimage := fmt.Sprintf("<img src=\"data:image/png;base64,%s\"/>", out)
	mailer.SetBody("text/html", "Here's your qris payment <br> For Test click this link<br><a href='"+link+"'>Click here</a><br>"+linkimage)
	dialer := gomail.NewDialer(
		os.Getenv("CONFIG_SMTP_HOST"),
		port,
		os.Getenv("CONFIG_AUTH_EMAIL"),
		os.Getenv("CONFIG_AUTH_PASSWORD"),
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return errs.Wrap(err)
	}

	return nil
}

// func Forgot_password(form domain.User) {
// 	var link, text string
// 	mailer := gomail.NewMessage()
// 	mailer.SetHeader("From", CONFIG_SENDER_NAME)
// 	mailer.SetHeader("To", form.Email, "voucherlabs.official@gmail.com")
// 	mailer.SetAddressHeader("Cc", "voucherlabs.official@gmail.com", "Tra Lala La")
// 	mailer.SetHeader("Subject", "Code Verification")
// 	link = os.Getenv("RUNNING_HOST") + "reset-password/" + form.Id.String()

// 	mailer.SetBody("text/html", "Please click this link to reset your password<br><a href='"+link+"'>Click here</a><br>"+text)

// 	dialer := gomail.NewDialer(
// 		CONFIG_SMTP_HOST,
// 		port,
// 		CONFIG_AUTH_EMAIL,
// 		CONFIG_AUTH_PASSWORD,
// 	)

// 	err := dialer.DialAndSend(mailer)
// 	if err != nil {
// 		log.Print(dialer)
// 		log.Fatal(err.Error())
// 	}

// 	log.Println("Mail sent!")
// }
