package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/orensimple/trade-core-app/internal/app/domain"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
)

// Order is order app
type Order struct{}

func (b Order) Create(o *domain.Order) (*domain.Order, error) {
	host := viper.Get("order_host")
	url := fmt.Sprintf("%v/api/order", host)

	data, err := json.Marshal(o)
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
		log.Debugf("order response '%q", dump)
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

	var resp *domain.Order
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b Order) Delete(o *domain.Order) (*domain.Order, error) {
	client := &http.Client{}
	host := viper.Get("order_host")
	url := fmt.Sprintf("%v/api/order/%s", host, o.ID)

	log.Info("Sending request")
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	log.Infof("Dump response with code, '%d'", res.StatusCode)
	dump, err := httputil.DumpResponse(res, true)
	if err == nil {
		log.Debugf("order response '%q", dump)
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

	var resp *domain.Order
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b Order) Find(a *domain.Account) ([]*domain.Order, error) {
	host := viper.Get("order_host")
	url := fmt.Sprintf("%v/api/orders", host)

	f := domain.Order{AccountID: a.AccountID}

	data, err := json.Marshal(f)
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
		log.Debugf("order response '%q", dump)
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

	var resp []*domain.Order
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
