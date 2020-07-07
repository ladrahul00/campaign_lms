package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Status constants
const (
	ACTIVE     = "ACTIVE"
	DRAFT      = "DRAFT"
	INACTIVE   = "INACTIVE"
	COMPLETE   = "COMPLETE"
	INCOMPLETE = "INCOMPLETE"
)

// Campaign type definition
type Campaign struct {
	ID            primitive.ObjectID  `bson:"_id,omitempty"`
	CreatedBy     primitive.ObjectID  `bson:"createdBy"`
	Purpose       string              `bson:"purpose"`
	Description   string              `bson:"description,omitempty"`
	CampaignTag   string              `bson:"campaignTag"`
	CampaignName  string              `bson:"campaignName"`
	Status        string              `bson:"status"`
	TemplateID    primitive.ObjectID  `bson:"templateId"`
	CreatedOn     primitive.Timestamp `bson:"createdOn,omitempty"`
	StartDateTime time.Time           `bson:"startDateTime,omitempty"`
	EndDateTime   time.Time           `bson:"endDateTime,omitempty"`
}
