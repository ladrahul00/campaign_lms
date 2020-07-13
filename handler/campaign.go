package handler

import (
	"context"

	campaign "github.com/wolf00/campaign_lms/proto/campaign"
	services "github.com/wolf00/campaign_lms/services"
)

type CampaignServiceHandler struct {
	services.CampaignService
}

// Create is a single request handler called via client.Call or the generated client code
func (e *CampaignServiceHandler) Create(ctx context.Context, req *campaign.NewCampaignRequest, rsp *campaign.Response) error {
	return e.CampaignService.Create(ctx, req, rsp)
}

// Get is a single request handler called via client.Call or the generated client code
func (e *CampaignServiceHandler) Get(ctx context.Context, req *campaign.CampaignByIdRequest, rsp *campaign.CampaignResponse) error {
	campaign, err := e.CampaignService.CampaignByID(ctx, req.CampaignId)
	if err != nil {
		return err
	}
	e.CampaignService.Campaign2CampaignResponse(campaign, rsp)
	return nil
}

// GetByTag is a single request handler called via client.Call or the generated client code
func (e *CampaignServiceHandler) GetByTag(ctx context.Context, req *campaign.CampaignByTagRequest, rsp *campaign.CampaignResponse) error {
	campaign, err := e.CampaignService.CampaignByTag(ctx, req.CampaignTag)
	if err != nil {
		return err
	}
	e.CampaignService.Campaign2CampaignResponse(campaign, rsp)
	return nil
}
