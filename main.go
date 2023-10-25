package main

import (
	"context"
	"fmt"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
	"time"
)

// Change these according to your Infobip account details.
const (
	EMAIL_FROM  = "<your-email-sender>"
	IB_BASE_URL = "<your-base-url>"
	IB_API_KEY  = "<your-api-key>"
)

const (
	N_DATE_OPTS   = 4
	CHANNEL_EMAIL = 1
	CHANNEL_SMS   = 2
)

type Patient struct {
	Name  string
	Email string
	Phone string
}

// This function simulates a call to the calendar service or database. It returns a list of available dates.
func getAvailableDates(nDates int) []time.Time {
	now := time.Now()
	var dates []time.Time
	for i := 1; i <= nDates; i++ {
		dates = append(dates, now.AddDate(0, 0, i+1).Round(time.Hour*24))
	}

	return dates
}

func sendEmail(client infobip.Client, from string, to string, subject string, text string, sendAt time.Time) int {
	mail := models.EmailMsg{
		From:    from,
		To:      to,
		Subject: subject,
		Text:    text,
		BulkID:  fmt.Sprintf("appointments-%d", time.Now().Unix()),
		SendAt:  sendAt.Format(time.RFC3339),
	}

	resp, respDetails, _ := client.Email.Send(context.Background(), mail)

	fmt.Printf("%+v\n", resp)
	fmt.Println(respDetails)

	return respDetails.HTTPResponse.StatusCode
}

func sendSMS(client infobip.Client, to string, text string, sendAt time.Time) int {
	sms := models.SMSMsg{
		Destinations: []models.SMSDestination{
			{To: to},
		},
		Text:   text,
		SendAt: sendAt.Format(time.RFC3339),
	}
	request := models.SendSMSRequest{
		BulkID:   fmt.Sprintf("appointments-%d", time.Now().Unix()),
		Messages: []models.SMSMsg{sms},
	}

	resp, respDetails, _ := client.SMS.Send(context.Background(), request)

	fmt.Printf("%+v\n", resp)
	fmt.Println(respDetails)

	return respDetails.HTTPResponse.StatusCode
}

func greet() {
	fmt.Println("==================================")
	fmt.Println("Welcome to MedAppointments System!")
	fmt.Println("==================================")
	fmt.Println()
}

func bye() {
	fmt.Println()
	fmt.Println("===============================================")
	fmt.Println("Thank you for letting us keep you healthy! Bye!")
	fmt.Println("===============================================")
}

func sendAppointmentReminder(client infobip.Client, channel int, patient Patient, appointment time.Time) {
	subject := "Dr. Bip Appointment Reminder"
	text := fmt.Sprintf("%s, you have an appointment with Dr. Bip tomorrow!", patient.Name)
	date := appointment.AddDate(0, 0, -1)

	if channel == CHANNEL_EMAIL {
		sendEmail(client, EMAIL_FROM, patient.Email, subject, text, date)
	} else {
		sendSMS(client, patient.Phone, text, date)
	}
}

func sendFollowUpReminder(client infobip.Client, channel int, patient Patient, appointment time.Time) {
	fmt.Println("Sending follow-up reminder!")
	subject := "Dr. Bip Follow-Up Reminder"
	text := fmt.Sprintf("%s, please schedule a follow-up appointment with Dr. Bip soon.", patient.Name)
	date := appointment.AddDate(0, 0, 25)
	if channel == CHANNEL_EMAIL {
		sendEmail(client, EMAIL_FROM, patient.Email, subject, text, date)
	} else {
		sendSMS(client, patient.Phone, text, date)
	}
}

func getInfobipClient() infobip.Client {
	client, _ := infobip.NewClient(
		IB_BASE_URL,
		IB_API_KEY,
	)

	return client
}

func main() {
	greet()

	patient := Patient{}
	patient.Name = promptText("Enter the patient's name: ")

	appointment := promptDate("Choose a date for your appointment:", getAvailableDates(N_DATE_OPTS), N_DATE_OPTS)

	channel := promptChannel("How would you like to be reminded?\n1) Email\n2) SMS")

	if channel == CHANNEL_EMAIL {
		patient.Email = promptText("Enter your email address: ")
	} else {
		patient.Phone = promptText("Enter your cell phone number (with country code): ")
	}

	fmt.Println("All set! You'll be reminded 24 hours before the appointment.")

	client := getInfobipClient()

	sendAppointmentReminder(client, channel, patient, appointment)
	fmt.Println(appointment)

	if promptYesNo("Do you want a reminder for a follow-up after the recommended period? (one month)") {
		sendFollowUpReminder(client, channel, patient, appointment)
	}

	bye()
}
