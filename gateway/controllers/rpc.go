package controllers

import (
	"context"
	"vcs-kafka-learning-go-gateway/models"
	"vcs-kafka-learning-go-gateway/proto"
)

func (s *responseServer) SendResponse(ctx context.Context, in *proto.ResponseBody) (*proto.NoMessage, error) {

	responseChannel <- models.Response{
		Message:   in.Message,
		IsSucceed: in.IsSucceed,
	}

	return &proto.NoMessage{}, nil
}
