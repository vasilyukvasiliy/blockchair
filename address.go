// Copyright 2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchair

import (
	"encoding/json"
	"fmt"
)

type AddressResponse struct {
	Data   []Address `json:"data"`
	Rows   uint      `json:"rows"`
	Limit  int64     `json:"limit"`
	Time   float64   `json:"time"`
	Cache  int       `json:"cache"`
	Source string    `json:"source"`
}

type Address struct {
	SumValue            json.Number `json:"sum_value"`
	SumValueUsd         json.Number `json:"sum_value_usd"`
	SumSpendingValueUsd json.Number `json:"sum_spending_value_usd"`
	MaxTimeReceiving    string      `json:"max_time_receiving"`
	MaxTimeSpending     string      `json:"max_time_spending"`
	MinTimeReceiving    string      `json:"min_time_receiving"`
	CountTotal          json.Number `json:"count_total"`
	Rate                json.Number `json:"rate"`
	SumValueUnspent     json.Number `json:"sum_value_unspent"`
	SumValueUnspentUsd  json.Number `json:"sum_value_unspent_usd"`
	CountUnspent        json.Number `json:"count_unspent"`
	PluUsd              json.Number `json:"plu_usd"`
	MinTimeSpending     string      `json:"min_time_spending"`
	PlUsd               json.Number `json:"pl_usd"`
	ReceivingActivity   []Activity  `json:"receiving_activity"`
	SpendingActivity    []Activity  `json:"spending_activity"`
}

type Activity struct {
	Year  int    `json:"year"`
	Month int    `json:"month"`
	Value string `json:"value"`
}

// GetAddress
// https://api.blockchair.com/bitcoin/dashboards/address/{address}
func (c *Client) GetAddress(address string) (a *Address, e error) {
	response, e := c.GetAddressRaw(address)

	if len(response.Data) == 1 {
		a = &response.Data[0]
	} else {
		if len(response.Data) > 1 {
			e = fmt.Errorf("Unexpected response from the server")
		}
	}

	return
}

// GetAddressRaw
// https://api.blockchair.com/bitcoin/dashboards/address/{address}
func (c *Client) GetAddressRaw(address string) (response *AddressResponse, e error) {
	response = &AddressResponse{}
	e = c.DoRequest("/dashboards/address/"+address, response)

	return
}
