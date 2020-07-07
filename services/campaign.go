package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/micro/go-micro/v2/util/log"
	"github.com/wolf00/golang_lms/campaign/db"
	"github.com/wolf00/golang_lms/campaign/db/models"
	"github.com/wolf00/golang_lms/campaign/utilities"

	campaign "github.com/wolf00/golang_lms/campaign/proto/campaign"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CampaignService hanldes
type CampaignService struct {
}

// Create service
func (c *CampaignService) Create(ctx context.Context, req *campaign.NewCampaignRequest, rsp *campaign.Response) error {
	var err error
	validationErrors := c.validateCampaignFields(req)
	if len(validationErrors) > 0 {
		rsp.Message = strings.Join(validationErrors, ", ")
		rsp.Status = false
		return nil
	}
	newCampaign := models.Campaign{}
	newCampaign.CampaignName = req.CampaignName
	newCampaign.Purpose = req.Purpose
	newCampaign.Description = req.Description
	newCampaign.CampaignTag = req.CampaignTag
	// TO_DO: Validate user id
	newCampaign.CreatedBy, err = primitive.ObjectIDFromHex(req.CreatedBy)
	if err != nil {
		return fmt.Errorf("invalid userid")
	}
	_, err = leadTemplateByID(ctx, req.TemplateId)
	if err != nil {
		return err
	}
	newCampaign.TemplateID, err = primitive.ObjectIDFromHex(req.TemplateId)
	if err != nil {
		return fmt.Errorf("invalid template id")
	}
	newCampaign.StartDateTime, err = utilities.DatabaseTimestamp(req.StartDateTime)
	if err != nil {
		return err
	}
	newCampaign.EndDateTime, err = utilities.DatabaseTimestamp(req.EndDateTime)
	if err != nil {
		return err
	}
	if newCampaign.StartDateTime.Before(time.Now()) {
		return fmt.Errorf("campaign start date cannot be before current date")
	}
	if newCampaign.StartDateTime.After(newCampaign.EndDateTime) {
		return fmt.Errorf("campaign start date cannot be after campaign end date")
	}
	_, err = createCampaign(ctx, newCampaign)
	if err != nil {
		log.Error(err)
		rsp.Message = "failed to create campaign"
		rsp.Status = false
		return err
	}
	rsp.Message = "campaign created succesfully"
	rsp.Status = true
	return nil
}

// CampaignByID ads
func (c *CampaignService) CampaignByID(ctx context.Context, campaignID string) (models.Campaign, error) {
	var campaign models.Campaign
	id, err := primitive.ObjectIDFromHex(campaignID)
	if err != nil {
		log.Error(err)
		return campaign, err
	}
	filter := bson.M{"_id": id}
	campaign, err = getByFilter(ctx, filter)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

// CampaignByTag ads
func (c *CampaignService) CampaignByTag(ctx context.Context, campaignTag string) (models.Campaign, error) {
	var campaign models.Campaign
	filter := bson.M{"campaignTag": campaignTag}
	campaign, err := getByFilter(ctx, filter)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

// Campaign2CampaignResponse asd
func (c *CampaignService) Campaign2CampaignResponse(campaign models.Campaign, rsp *campaign.CampaignResponse) {
	rsp.CampaignName = campaign.CampaignName
	rsp.Purpose = campaign.Purpose
	rsp.CampaignTag = campaign.CampaignTag
	rsp.CreatedBy = campaign.CreatedBy.Hex()
	rsp.Description = campaign.Description
	rsp.EndDateTime = campaign.EndDateTime.Format(utilities.TIME_FORMAT)
	rsp.Id = campaign.ID.Hex()
	rsp.StartDateTime = campaign.StartDateTime.Format(utilities.TIME_FORMAT)
	rsp.TemplateId = campaign.TemplateID.Hex()
}

func (c *CampaignService) validateCampaignFields(req *campaign.NewCampaignRequest) []string {
	validationErrors := []string{}
	if !utilities.ValidateString(req.Purpose) {
		validationErrors = append(validationErrors, "campaign purpose cannot be empty")
	}
	if !utilities.ValidateString(req.CampaignName) {
		validationErrors = append(validationErrors, "campaign name cannot be empty")
	}
	if !utilities.ValidateString(req.StartDateTime) {
		validationErrors = append(validationErrors, "campaign start date cannot be empty")
	}
	if !utilities.ValidateString(req.EndDateTime) {
		validationErrors = append(validationErrors, "campaign end date cannot be empty")
	}
	if !utilities.ValidateString(req.EndDateTime) {
		validationErrors = append(validationErrors, "template cannot be empty")
	}
	if !utilities.ValidateString(req.EndDateTime) {
		validationErrors = append(validationErrors, "user id cannot be empty")
	}
	if !utilities.ValidateString(req.CampaignTag) {
		validationErrors = append(validationErrors, "campaign tag cannot be empty")
	} else {
		err := utilities.ValidateTag(req.CampaignTag)
		if err != nil {
			validationErrors = append(validationErrors, err.Error())
		}
	}
	return validationErrors
}

func getByFilter(ctx context.Context, filter interface{}) (models.Campaign, error) {
	helper := db.Campaigns(ctx)
	var campaign models.Campaign

	err := helper.FindOne(ctx, filter).Decode(&campaign)
	if err != nil {
		log.Error(err)
		return campaign, err
	}

	return campaign, nil
}

func createCampaign(ctx context.Context, newCampaign models.Campaign) (*mongo.InsertOneResult, error) {
	helper := db.Campaigns(ctx)
	// TO_DO: Update the status dynamically after we have approval system and versioning in place
	newCampaign.Status = models.ACTIVE
	newCampaign.CreatedOn = utilities.CurrentDateTime()
	return helper.InsertOne(ctx, newCampaign)
}
