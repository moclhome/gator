package main

import (
	"bootdev/go/gator/internal/config"
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	testConfig := config.Config{
		Db_url: "test_db",
	}
	theState := state{
		config: &testConfig,
	}
	theCommand := command{
		name: "login",
	}
	cases := map[string]struct {
		input          []string
		expectedConfig config.Config
		expectedError  error
	}{
		"correct input": {
			input: []string{"MyName"},
			expectedConfig: config.Config{
				Db_url:            "test_db",
				Current_user_name: "MyName",
			},
			expectedError: nil,
		},
		"no parameter": {
			input: []string{},
			expectedConfig: config.Config{
				Db_url:            "test_db",
				Current_user_name: "",
			},
			expectedError: fmt.Errorf("Usage: login <username>"),
		},
		"two parameters": {
			input: []string{"param1", "param2"},
			expectedConfig: config.Config{
				Db_url:            "test_db",
				Current_user_name: "",
			},
			expectedError: fmt.Errorf("Usage: login <username>"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			theCommand.arguments = c.input
			err := handlerLogin(&theState, theCommand)
			if err != nil && c.expectedError != nil && err.Error() != c.expectedError.Error() {
				t.Fatalf("Error during login: %v, expected error: %v", err, c.expectedError)
			}
			if err == nil && c.expectedError != nil {
				t.Fatalf("No error during login. Expected error: %v", c.expectedError)
			}
			if err != nil && c.expectedError == nil {
				t.Fatalf("Error during login: %v, no expected error.", err)
			}

			errorText, ok := testConfigData("Test if config is set", *theState.config, c.expectedConfig)
			if !ok {
				t.Fatalf("%s", errorText)
			}
			newConfig, err := config.Read()
			errorText, ok = testConfigData("Test if file is written", newConfig, c.expectedConfig)
		})
		// remove the user from file for next test
		theState.config.SetUser("")
	}
}

func testConfigData(testComment string, cfg config.Config, expectedCfg config.Config) (errorText string, ok bool) {
	ok = true
	errorText = ""
	if cfg.Current_user_name != expectedCfg.Current_user_name {
		errorText = fmt.Sprintf("%s: User name: %s; expected user name: %s", testComment, cfg.Current_user_name, expectedCfg.Current_user_name)
		ok = false
	}
	if cfg.Db_url != expectedCfg.Db_url {
		errorText = fmt.Sprintf("%s: DB URL: %s; expected DB URL: %s", testComment, cfg.Db_url, expectedCfg.Db_url)
		ok = false
	}
	return
}
