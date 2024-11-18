// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/reguardprotection/warrantysphere/internal/utils"
	"time"
)

type StaffAggregateStatus struct {
}

type StaffAggregate struct {
	ID          string `json:"id"`
	ReferenceID string `json:"referenceId"`
	// Datetime when the object was created.
	Created time.Time `json:"created"`
	// Datetime when the object was last modified.
	Modified       time.Time            `json:"modified"`
	Status         StaffAggregateStatus `json:"status"`
	Email          string               `json:"email"`
	Name           string               `json:"name"`
	Phone          string               `json:"phone"`
	UserID         string               `json:"userId"`
	User           UserAggregate        `json:"user"`
	ServiceAccount bool                 `json:"serviceAccount"`
	Roles          []string             `json:"roles"`
}

func (s StaffAggregate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *StaffAggregate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *StaffAggregate) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *StaffAggregate) GetReferenceID() string {
	if o == nil {
		return ""
	}
	return o.ReferenceID
}

func (o *StaffAggregate) GetCreated() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Created
}

func (o *StaffAggregate) GetModified() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Modified
}

func (o *StaffAggregate) GetStatus() StaffAggregateStatus {
	if o == nil {
		return StaffAggregateStatus{}
	}
	return o.Status
}

func (o *StaffAggregate) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *StaffAggregate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *StaffAggregate) GetPhone() string {
	if o == nil {
		return ""
	}
	return o.Phone
}

func (o *StaffAggregate) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *StaffAggregate) GetUser() UserAggregate {
	if o == nil {
		return UserAggregate{}
	}
	return o.User
}

func (o *StaffAggregate) GetServiceAccount() bool {
	if o == nil {
		return false
	}
	return o.ServiceAccount
}

func (o *StaffAggregate) GetRoles() []string {
	if o == nil {
		return []string{}
	}
	return o.Roles
}