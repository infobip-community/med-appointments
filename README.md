# med-appointments

A CLI demo showcasing a simple medical appointment system with notifications and reminders.

This use case covers a scenario in which patients can schedule a medical appointment and get an appointment
reminder over SMS or email using the [Infobip API Go SDK](https://github.com/infobip-community/infobip-api-go-sdk). The use case also showcases how to schedule
messages for follow up notifications.

For a more in-depth exploration of the project, check our [DevHub post](infobip.com/developers).

## Prerequisites

* Working [Go installation](https://go.dev/doc/install)
* [Infobip Account](https://www.infobip.com/signup)
* A terminal
* [Git installation](https://git-scm.com/downloads)

## Setup

Before running the project, you must set a few constants with your account credentials, `IB_BASE_URL`
and `IB_API_KEY`.
You can get your credentials by logging into your [Infobip account](https://portal.infobip.com/login/).
Once you configured the variables, you can run this project with the following commands:

```bash
cd med-appointments
go get "github.com/infobip-community/infobip-api-go-sdk/v3"
go run main.go prompts.go
```

This will move the terminal to the cloned folder, install the SDK, and run the project.
