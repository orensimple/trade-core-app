package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/google/uuid"
	"github.com/orensimple/trade-core-app/internal/app/domain"
	"github.com/prometheus/common/log"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
)

const DefaultCurrencyCode = "RU"

// BillingCreateResponse is response of billing
type BillingCreateResponse struct {
	ID            uuid.UUID
	CurrencyCode  string
	Name          string
	Balance       decimal.Decimal
	BlockedAmount decimal.Decimal
}

// billingCreateRequest is request of billing create
type billingCreateRequest struct {
	CurrencyCode string `json:"currency_code"`
	Name         string `json:"name"`
}

// Billing is billing app
type Billing struct{}

func (b Billing) Create(u *domain.User) (*BillingCreateResponse, error) {
	host := viper.Get("billing_host")
	url := fmt.Sprintf("%v/api/account", host)

	data, err := json.Marshal(billingCreateRequest{CurrencyCode: DefaultCurrencyCode, Name: u.ID.String()})
	if err != nil {
		return nil, err
	}

	log.Info("Sending request")
	res, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	log.Infof("Dump response with code, '%d'", res.StatusCode)
	dump, err := httputil.DumpResponse(res, true)
	if err == nil {
		log.Debugf("account response '%q", dump)
	}
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("something wrong")
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp *BillingCreateResponse
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b Billing) Get(u *domain.Account) (*BillingCreateResponse, error) {
	host := viper.Get("billing_host")
	url := fmt.Sprintf("%v/api/account/%s", host, u.AccountID)

	log.Info("Sending request")
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	log.Infof("Dump response with code, '%d'", res.StatusCode)
	dump, err := httputil.DumpResponse(res, true)
	if err == nil {
		log.Debugf("account response '%q", dump)
	}
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("something wrong")
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp *BillingCreateResponse
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
