package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/wolf00/golang_lms/campaign/client"
	"github.com/wolf00/golang_lms/campaign/handler"

	campaign "github.com/wolf00/golang_lms/campaign/proto/campaign"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.campaign"),
		micro.Version("0.1"),
	)

	// Initialise service
	service.Init(
		// micro.WrapHandler(client.UserWrapper(service))
		micro.WrapHandler(client.LeadTemplateWrapper(service)),
	)

	// Register Handler
	campaign.RegisterCampaignHandler(service.Server(), new(handler.CampaignServiceHandler))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
