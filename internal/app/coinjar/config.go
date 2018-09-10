package coinjar

type Config struct {
	PersonalToken string `envconfig:"PERSONAL_TOKEN"`
	WebHookSecret string `envconfig:"WEBHOOK_SECRET"`
}
