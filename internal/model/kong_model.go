package model

import "github.com/google/uuid"

type KongConsumer struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type KongConsumerRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type KongConsumerResponse struct {
	ID           uuid.UUID    `json:"uuid"`
	Secret       string       `json:"secret"`
	Consumer     KongConsumer `json:"consumer"`
	CreatedAt    int          `json:"created_at"`
	Key          string       `json:"key"`
	RsaPublicKey string       `json:"rsa_public_key"`
}
