package registry

import "testing"

func TestRegistyIsPopulatedAsExpected(t *testing.T) {
	expectedRegistyEntries := []string{"virtualbox"}

	if len(Registry) != len(expectedRegistyEntries) {
		t.Errorf("Unexpected number of drivers registered. Expected %d, got: %d", len(expectedRegistyEntries), len(Registry))
	}
	for _, key := range expectedRegistyEntries {
		if _, exists := Registry[key]; !exists {
			t.Errorf("Expected driver '%s' to be registered.", key)
		}
	}
}

func TestValidDriversIsPopulatedAsExpected(t *testing.T) {
	expectedDrivers := []string{"virtualbox"}

	if len(ValidDrivers) != len(expectedDrivers) {
		t.Errorf("Unexpected number of drivers registered. Expected %d, got: %d", len(expectedDrivers), len(ValidDrivers))
	}
	for idx := range expectedDrivers {
		if ValidDrivers[idx] != expectedDrivers[idx] {
			t.Errorf("Expected driver '%s' to be registered at index %d.", expectedDrivers[idx], idx)
		}
	}
}
