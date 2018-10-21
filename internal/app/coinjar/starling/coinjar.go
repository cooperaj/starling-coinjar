package starling

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/billglover/starling"
	"github.com/cooperaj/starling-coinjar/internal/app/coinjar"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

var (
	client *starling.Client
	ctx    context.Context
)

type StarlingCoinJar struct {
	Name           string
	Currency       string
	RoundTo        int8
	SavingsGoalUID string
}

func NewCoinJar(config coinjar.Config) coinjar.CoinJar {
	var coinJar = StarlingCoinJar{Name: config.CoinJarName}
	coinJar.Currency = "GBP"
	coinJar.RoundTo = int8(config.RoundTo)

	client = coinJar.starlingClient(config.PersonalToken)

	savingsGoalUID, err := coinJar.ensureStarlingSavingsGoal(config.CoinJarName)
	if err != nil {
		panic(fmt.Sprintf("Unable to ensure a Starling Savings Goal exists: %s", err.Error()))
	}
	coinJar.SavingsGoalUID = savingsGoalUID

	return &coinJar
}

func (cj *StarlingCoinJar) AddFunds(amount int8) error {
	if amount == 0 {
		return nil
	}

	var change = starling.Amount{
		Currency:   cj.Currency,
		MinorUnits: int64(amount),
	}

	_, _, err := client.TransferToSavingsGoal(ctx, cj.SavingsGoalUID, change)
	if err != nil {
		panic(fmt.Sprintf("Cannot add funds to the \"%s\" coin jar: %s", cj.Name, err))
	}

	return nil
}

func (cj *StarlingCoinJar) GetRoundTo() int8 {
	return cj.RoundTo
}

func (cj *StarlingCoinJar) ensureStarlingSavingsGoal(name string) (savingsGoalUID string, err error) {
	// get list of savings goals
	savingsGoals, _, err := client.SavingsGoals(ctx)
	if err != nil {
		return "", err
	}

	// check ours is in it
	for _, savingsGoal := range savingsGoals {
		if savingsGoal.Name == name {
			return savingsGoal.UID, nil
		}
	}

	// if not, make it
	uuid := uuid.New()
	err = cj.makeSavingsGoal(uuid, name)
	if err != nil {
		return "", err
	}

	fmt.Printf("New savings goal %s created...\n", name)

	return uuid.String(), nil
}

func (cj *StarlingCoinJar) makeSavingsGoal(uuid uuid.UUID, name string) error {
	image, _ := coinjar.Asset("assets/coins.jpg")
	request := starling.SavingsGoalRequest{
		Name:               name,
		Currency:           cj.Currency,
		Base64EncodedPhoto: base64.StdEncoding.EncodeToString(image),
	}

	_, err := client.CreateSavingsGoal(ctx, uuid.String(), request)
	if err != nil {
		return err
	}

	return nil
}

func (cj *StarlingCoinJar) starlingClient(accessToken string) *starling.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	ctx = context.Background()
	tc := oauth2.NewClient(ctx, ts)

	baseURL, _ := url.Parse(starling.ProdURL)
	return starling.NewClientWithOptions(tc, starling.ClientOptions{BaseURL: baseURL})
}
