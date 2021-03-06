// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// JobResponseBase undocumented
type JobResponseBase struct {
	// Entity is the base model of JobResponseBase
	Entity
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// TenantID undocumented
	TenantID *string `json:"tenantId,omitempty"`
	// CreationDateTime undocumented
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Error undocumented
	Error *ClassificationError `json:"error,omitempty"`
}
