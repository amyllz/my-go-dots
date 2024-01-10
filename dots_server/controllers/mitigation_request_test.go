package controllers_test

import (
	"testing"

	"github.com/nttdots/go-dots/dots_common"
	"github.com/nttdots/go-dots/dots_common/messages"
	"github.com/nttdots/go-dots/dots_server/controllers"
	"github.com/nttdots/go-dots/dots_server/models"
)

func TestMitigationRequestPut(t *testing.T) {
	mitigationRequest := controllers.MitigationRequest{}

	scopes := []messages.Scope{
		{
			MitigationId:    1234567,
			TargetPrefix:    []string{"192.168.1.10/24", "192.168.1.11/32"},
			TargetPortRange: []messages.TargetPortRange{{LowerPort: 8989, UpperPort: 9999}},
			TargetProtocol:  []int{1, 2},
			FQDN:            []string{"golang.org"},
			URI:             []string{""},
			AliasName:       []string{""},
			Lifetime:        0,
		},
		{
			MitigationId:    2345678,
			TargetPrefix:    []string{"192.168.1.110/24", "192.168.1.111/24"},
			TargetPortRange: []messages.TargetPortRange{{LowerPort: 8989, UpperPort: 9999}},
			TargetProtocol:  []int{1, 2},
			FQDN:            []string{"golang.org"},
			URI:             []string{""},
			AliasName:       []string{""},
			Lifetime:        0,
		},
	}
	request := messages.MitigationRequest{}
	request.MitigationScope = messages.MitigationScope{}
	request.MitigationScope.Scopes = scopes
	customer, err := models.GetCustomerById(123)
	if err != nil {
		t.Errorf("get customer data error: %s", err.Error())
		return
	}
	response, err := mitigationRequest.HandlePut(controllers.Request{ Body: &request }, customer)

	if err != nil {
		t.Errorf("post method return error: %s", err.Error())
		return
	}
	switch code := response.Code; code {
	case dots_common.BadRequest:
		t.Errorf("got %s, want %s", code.String(), dots_common.Created.String())
	case dots_common.Created:
	default:
		t.Errorf("invalid result code: %s", code.String())
	}
}
