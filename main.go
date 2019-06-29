package main

import (
	"context"
	"encoding/base64"
	"fmt"

	"google.golang.org/api/pubsub/v1"
)

func ReceivedMessage(ctx context.Context, m *pubsub.PubsubMessage) error {
	decoded, err := base64.StdEncoding.DecodeString(m.Data)
	if err != nil {
		return err
	}

	fmt.Println(string(decoded))

	return nil
}
