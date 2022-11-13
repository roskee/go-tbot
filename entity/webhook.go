package entity

import (
	"encoding/json"
	"log"
)

type Webhook struct {
	URL       string `json:"url"`
	IPAddress string `json:"ip_address"`
}

// FromJSONBody modifies the current Webhook with matching values in body
func (u *Webhook) FromJSONBody(body []byte) {
	err := json.Unmarshal(body, u)
	if err != nil {
		log.Printf("error while unmarshalling webhook: %+v", err)
	}
}
