package contextmanager_test

import (
	"errors"
	"testing"

	contextmanager "github.com/snoekiede/gocontextmanager"
)

func TestWithResource(t *testing.T) {
	// Test case 1: Successful execution
	value := 42
	expectedResult := "success"
	expectedError := errors.New("error")

	result, err := contextmanager.WithResource(value, func(a int) (string, error) {
		// Perform some action with the resource
		return expectedResult, expectedError
	}, func(element int) {
		// Dispose the resource
	})

	if err != expectedError {
		t.Errorf("Expected error: %v, got: %v", expectedError, err)
	}

	if result != expectedResult {
		t.Errorf("Expected result: %v, got: %v", expectedResult, result)
	}

	// Test case 2: Error during execution
	value = 0
	expectedError = errors.New("error during execution")

	result, err = contextmanager.WithResource(value, func(a int) (string, error) {
		// Perform some action with the resource
		return "", expectedError
	}, func(element int) {
		// Dispose the resource
	})

	if err != expectedError {
		t.Errorf("Expected error: %v, got: %v", expectedError, err)
	}

	if result != "" {
		t.Errorf("Expected empty result, got: %v", result)
	}
}

func TestContextManager_Dispose(t *testing.T) {
	// Test case 1: Dispose function is called
	value := 42
	disposeCalled := false

	cm := &contextmanager.ContextManager[int]{Value: value}
	cm.Dispose(func(a int) {
		disposeCalled = true
	})

	if !disposeCalled {
		t.Error("Dispose function was not called")
	}

}
