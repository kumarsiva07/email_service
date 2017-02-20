package email

import (
	"os"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/go-gomail/gomail"
)

type Email struct {
	Message      *gomail.Message
	ResponseChan chan EmailResponse
}

type EmailResponse struct {
	Error error
}

var (
	Sender     gomail.SendCloser
	Open       = false
	EmailQueue = make(chan Email)
	CloseTime  = time.Duration(30)
)

func Start() {
	if os.Getenv("SMTP_HOST") == "" {
		os.Setenv("SMTP_HOST", "localhost")
	}

	if os.Getenv("SMTP_PORT") == "" {
		os.Setenv("SMTP_PORT", "25")
	}

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatal("Could not parse SMTP_PORT into an integer")
	}

	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"))

	for {
		select {
		case m := <-EmailQueue:
			if !Open {
				if Sender, err = d.Dial(); err != nil {
					log.Error(err)
					m.ResponseChan <- EmailResponse{Error: err}
					break
				}

				log.Debug("Connected to SMTP server made")
				Open = true
			}

			t := time.Now()
			if err := gomail.Send(Sender, m.Message); err != nil {
				log.Error(err)
				m.ResponseChan <- EmailResponse{Error: err}
				break
			}

			m.ResponseChan <- EmailResponse{}
			log.Debugf("Time = %+v\n", time.Since(t))

		// Close the connection to the SMTP server if no email was sent in
		// the last 30 seconds.
		case <-time.After(CloseTime * time.Second):
			if Open {
				if err := Sender.Close(); err != nil {
					log.Error(err)
				}
				Open = false
			}
		}
	}
}
