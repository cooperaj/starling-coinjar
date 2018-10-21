package coinjar

// Config Provides the configuration options for the CoinJar service
type Config struct {
	PersonalToken string         `envconfig:"PERSONAL_TOKEN" required:"true"`
	WebHookSecret string         `envconfig:"WEBHOOK_SECRET" required:"true"`
	CoinJarName   string         `envconfig:"COINJAR_NAME" default:"Coin Jar"`
	RoundTo       RoundToDecoder `envconfig:"ROUND_TO" default:"pound"`
}

type RoundToDecoder int8

func (rvd *RoundToDecoder) Decode(value string) error {
	switch value {
	case "fifty":
		*rvd = ChangeToFiftyPence
	case "twenty":
		*rvd = ChangeToTwentyPence
	case "ten":
		*rvd = ChangeToTenPence
	case "five":
		*rvd = ChangeToFivePence
	default:
		*rvd = ChangeToAPound
	}

	return nil
}
