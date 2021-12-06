package auditmodels

import (
	"testing"

	auditapisdk "github.com/phoenixnap/go-sdk-bmc/auditapi"
	"github.com/stretchr/testify/assert"
)

// Tests
func TestEventFromSdk(test_framework *testing.T) {
	sdkEvent := GenerateEventSdk()
	event := EventFromSdk(sdkEvent)

	assertEqualEvent(test_framework, *event, *sdkEvent)
}

func TestNilEventFromSdk(test_framework *testing.T) {
	event := EventFromSdk(nil)

	assert.Nil(test_framework, event)
}

// assertion functions
func assertEqualEvent(test_framework *testing.T, event Event, sdkEvent auditapisdk.Event) {
	assert.Equal(test_framework, event.Name, sdkEvent.Name)
	assert.Equal(test_framework, event.Timestamp, sdkEvent.Timestamp)

	assertEqualUserInfo(test_framework, event.UserInfo, sdkEvent.UserInfo)
}

func assertEqualUserInfo(test_framework *testing.T, userInfo UserInfo, sdkUserInfo auditapisdk.UserInfo) {
	assert.Equal(test_framework, userInfo.AccountId, sdkUserInfo.AccountId)
	assert.Equal(test_framework, userInfo.ClientId, sdkUserInfo.ClientId)
	assert.Equal(test_framework, userInfo.Username, sdkUserInfo.Username)
}
