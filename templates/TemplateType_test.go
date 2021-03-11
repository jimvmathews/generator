package templates

import "testing"

func TestString(t *testing.T) {
	input := EntityTemplate
	expected := "Entity"
	result := input.String()
	if result != expected {
		t.Errorf("TestString failed. Expected %v, got %v", expected, result)
	}
}

func TestGetTemplateType_validString(t *testing.T) {
	input := "Entity"
	expected := EntityTemplate
	result, err := GetTemplateType(input)
	if err != nil {
		t.Errorf("Did not expect an error %v", err)
	}
	if result != expected {
		t.Errorf("TestString failed. Expected %v, got %v", expected, result)
	}
}

func TestGetTemplateType_invalidString(t *testing.T) {
	input := "NotValidTemplate"
	expected := 0
	result, err := GetTemplateType(input)
	if err == nil {
		t.Errorf("TestString failed. Expected an error and expected %v got result %v", expected, result)
	}
}
