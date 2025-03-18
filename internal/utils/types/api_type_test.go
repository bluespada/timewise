package types

import (
	"testing"
)

func TestApiResult(t *testing.T) {
	apiResult := NewApiResponse()
	apiResult.Message = "Success"

	if apiResult.Error != false {
		t.Errorf("Expected false, got %v", apiResult.Error)
	}

	if apiResult.Message != "Success" {
		t.Errorf("Expected Success, got %v", apiResult.Message)
	}

	if apiResult.Data != nil {
		t.Errorf("Expected nil, got %v", apiResult.Data)
	}
}

func TestApiError(t *testing.T) {
	apiResult := NewApiResponse()
	apiResult.Error = true

	if apiResult.Error != true {
		t.Errorf("Expected true, got %v", apiResult.Error)
	}
}

func TestApiCustomMessage(t *testing.T) {
	apiResult := NewApiResponse()
	apiResult.Message = "Custom Error Message"

	if apiResult.Message != "Custom Error Message" {
		t.Errorf("Expected 'Custom Error Message', got %v", apiResult.Message)
	}
}

func TestApiCustomData(t *testing.T) {
	apiResult := NewApiResponse()
	customData := map[string]interface{}{"key": "value"}
	apiResult.Data = customData

	if apiResult.Data == nil {
		t.Errorf("Expected custom data, got nil")
	}

	if data, ok := apiResult.Data.(map[string]interface{}); !ok || data["key"] != "value" {
		t.Errorf("Expected custom data with key 'key' and value 'value', got %v", apiResult.Data)
	}
}
