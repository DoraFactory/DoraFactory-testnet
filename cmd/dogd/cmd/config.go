package cmd

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/DoraFactory/doravota/app"
)

const (
	// set token decimal
	HumanCoinUnit = "dear"
	BaseCoinUnit  = "udear"
	DoraExponent  = 6

	DefaultBondDenom = BaseCoinUnit
)

func initSDKConfig() {
	// Set prefixes
	accountPubKeyPrefix := app.AccountAddressPrefix + "pub"
	validatorAddressPrefix := app.AccountAddressPrefix + "valoper"
	validatorPubKeyPrefix := app.AccountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := app.AccountAddressPrefix + "valcons"
	consNodePubKeyPrefix := app.AccountAddressPrefix + "valconspub"

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(app.AccountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)

	err := sdk.RegisterDenom(HumanCoinUnit, sdk.OneDec())
	if err != nil {
		panic(err)
	}
	err = sdk.RegisterDenom(BaseCoinUnit, sdk.NewDecWithPrec(1, DoraExponent))
	if err != nil {
		panic(err)
	}

	// set the genesis default denom
	sdk.DefaultBondDenom = DefaultBondDenom

	config.Seal()
}


const (
    Version   = "0.1.0"
)

func setVersionInfo() {
	// version.Name = Name
	// version.AppName = AppName
	version.Version = Version
	// version.Commit = Commit
	// version.BuildTags = BuildTags
}
