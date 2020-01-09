package main

import (
	"testing"

	"gitlab.indexexchange.com/app/authentication-api/api/libs/applog"
	"gitlab.indexexchange.com/app/authentication-api/api/settings"

	"github.com/stretchr/testify/assert"
)

var TestSettings = &settings.APISettings{
	CORS: settings.CORSConfig{
		AllowAllOrigins: true,
	},
	Logger: applog.LogConfig{
		LogFile:   "",
		Product:   "golang-testing",
		Component: "Test Suite",
		Stdout:    true,
		Level:     "debug",
	},
}

func TestGreet(t *testing.T) {
	testCases := map[string]struct {
		user          userDetails
		expectedError error
	}{
		"When Name More than 5 characters": {
			user: userDetails{
				userName: "Michael \n",
			},
			expectedError: nil,
		},
		// "When Name less than 5 characters": {
		// 	user: &userDetails{
		// 		userName: "Max\n",
		// 	},
		// 	expectedError: nil,
		// },
	}

	// name, tc := range testCases
	// t.Run(name, func(t *testing.T) {
	// 	defer applog.Initialize(TestSettings.Logger, "Test")()
	// 	err := Greet(tc.user)
	// 	assert.Equal(t, tc.expclearectedError, err)
	// })

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			defer applog.Initialize(TestSettings.Logger, "Test")()
			err := Greet(tc.user)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
