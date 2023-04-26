// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"context"
	"el_porvenir.com/cloudfunction/bigqueryPorvenir"
	"el_porvenir.com/cloudfunction/siigo"
	siigo_request "el_porvenir.com/cloudfunction/siigo/request"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

// ElPorvenirSiigo consumes a Pub/Sub message.
func ElPorvenirSiigo(ctx context.Context, m PubSubMessage) error {
	token := siigo_request.GetAuthToken()
	siigoData := siigo.GetSiigoData(token)
	bigQueryPorvenir := bigqueryPorvenir.NewBigQueryPorvenir(ctx, siigoData)
	bigQueryPorvenir.SiigoElPorvenir()
	return nil
}
