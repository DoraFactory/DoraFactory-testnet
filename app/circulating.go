package app

	
import (
	"github.com/cosmos/cosmos-sdk/client"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
	"encoding/json"
	"cosmossdk.io/math"
)

type SupplyResponse struct {
	Amount struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"amount"`
}

type BalanceResponse struct {
	Balance struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"balance"`
}

func (app *App) RegisterCirculatingAPI(ctx client.Context, rtr *mux.Router) error {

	rtr.HandleFunc("/circulating_supply", func(w http.ResponseWriter, r *http.Request) {
	   app.handleCirculatingSupplyQuery(w, r, ctx)
   }).Methods("GET")

   return nil
}

func (app *App) handleCirculatingSupplyQuery(w http.ResponseWriter, r *http.Request, ctx client.Context) {
   circulatingSupply, err := app.getSomeCirculatingSupply(ctx)
   if err != nil {
	   http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	   return
   }

   w.Header().Set("Content-Type", "application/json")
   json.NewEncoder(w).Encode(map[string]interface{}{
	   "circulating_supply": circulatingSupply,
   })
}


func (app *App) getSomeCirculatingSupply(ctx client.Context) (string, error) {

   supplyUrl := "https://vota-rest.dorafactory.org/cosmos/bank/v1beta1/supply/by_denom?denom=peaka"
   multisigUrl := "https://vota-rest.dorafactory.org/cosmos/bank/v1beta1/balances/dora1z47xcnkqtmu4pq9mvxgqjrrm7mev3z2c8tfwlv/by_denom?denom=peaka"

   response, err := http.Get(supplyUrl)
   if err != nil {
	   return "", err
   }
   defer response.Body.Close()

   body, err := ioutil.ReadAll(response.Body)
   if err != nil {
	   return "", err
   }

   var supplyResponse SupplyResponse
	err = json.Unmarshal(body, &supplyResponse)
	if err != nil {
		return "", err
	}

	// get total supply
	amount, _ := math.NewIntFromString(supplyResponse.Amount.Amount)

	response_multisig, err := http.Get(multisigUrl)
	if err != nil {
		return "", err
	}
	defer response_multisig.Body.Close()
 
	// get multisig account balance
	body_multisig, err := ioutil.ReadAll(response_multisig.Body)
	if err != nil {
		return "", err
	}

	var balanceResponse BalanceResponse
	err = json.Unmarshal(body_multisig, &balanceResponse)
	if err != nil {
		return "", err
	}

	multisig_balance, _ := math.NewIntFromString(balanceResponse.Balance.Amount)

	// get circulating supply
	res := amount.Sub(multisig_balance).String()

   return res , nil
}