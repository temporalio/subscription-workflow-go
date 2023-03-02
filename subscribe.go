package subscribe_emails

import "time"

// EmailInfo is the data that the SendContentEmail uses to send the message.
type EmailInfo struct {
	FirstName    string
	LastName     string
	EmailAddress string
}

// Campaign is the info about the email campaign.
type Campaign struct {
	WelcomeEmail     string
	UnsubscribeEmail string
	Mail            string
}

// Periods contains duration info for trial and billing periods
type Periods struct {
	TrialPeriod  time.Duration
	BillingPeriod time.Duration
	MaxBillingPeriods int
	BillingPeriodCharge int
}

// Subscription is the user email and the campaign they'll receive.
type Subscription struct {
	EmailInfo    EmailInfo
	Campaign     Campaign
	Periods      Periods
}
