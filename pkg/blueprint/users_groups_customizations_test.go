package blueprint

import (
	"testing"

	"github.com/osbuild/blueprint/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestSSHKey(t *testing.T) {
	expectedSSHKeys := []SSHKeyCustomization{
		{
			User: "test-user",
			Key:  "test-key",
		},
	}
	TestCustomizations := Customizations{
		SSHKey: expectedSSHKeys,
	}

	retUser := TestCustomizations.GetUsers()[0].Name
	retKey := *TestCustomizations.GetUsers()[0].Key

	assert.Equal(t, expectedSSHKeys[0].User, retUser)
	assert.Equal(t, expectedSSHKeys[0].Key, retKey)
}

func TestGetUsers(t *testing.T) {
	Desc := "Test descritpion"
	Pass := "testpass"
	Key := "testkey"
	Home := "Home"
	Shell := "Shell"
	Groups := []string{
		"Group",
	}
	UID := 123
	GID := 321
	ExpireDate := 12345
	ForcePasswordReset := true

	expectedUsers := []UserCustomization{
		{
			Name:               "John",
			Description:        &Desc,
			Password:           &Pass,
			Key:                &Key,
			Home:               &Home,
			Shell:              &Shell,
			Groups:             Groups,
			UID:                &UID,
			GID:                &GID,
			ExpireDate:         &ExpireDate,
			ForcePasswordReset: &ForcePasswordReset,
		},
	}

	TestCustomizations := Customizations{
		User: expectedUsers,
	}

	retUsers := TestCustomizations.GetUsers()

	assert.ElementsMatch(t, expectedUsers, retUsers)
}

func TestGetGroups(t *testing.T) {
	type testCase struct {
		groups               []GroupCustomization
		expectedErrorMessage string
	}

	testCases := map[string]testCase{
		"nil": {
			groups: nil,
		},
		"none": {
			groups: []GroupCustomization{},
		},
		"single": {
			groups: []GroupCustomization{
				{
					Name: "TestGroup",
					GID:  common.ToPtr(1234),
				},
			},
		},
		"multi": {
			groups: []GroupCustomization{
				{
					Name: "TestGroup",
					GID:  common.ToPtr(1234),
				},
				{
					Name: "sysgrp",
					GID:  common.ToPtr(998),
				},
				{
					Name: "wheel",
					GID:  common.ToPtr(42),
				},
			},
		},
		"duplicate-names": {
			groups: []GroupCustomization{
				{
					Name: "TestGroup",
					GID:  common.ToPtr(1234),
				},
				{
					Name: "sysgrp",
					GID:  common.ToPtr(998),
				},
				{
					Name: "wheel",
					GID:  common.ToPtr(42),
				},
				{
					Name: "wheel",
					GID:  common.ToPtr(43),
				},
			},
			expectedErrorMessage: "invalid group customizations:\nduplicate group name: wheel",
		},
		"duplicate-gids": {
			groups: []GroupCustomization{
				{
					Name: "TestGroup",
					GID:  common.ToPtr(1234),
				},
				{
					Name: "sysgrp",
					GID:  common.ToPtr(42),
				},
				{
					Name: "wheel",
					GID:  common.ToPtr(42),
				},
			},
			expectedErrorMessage: "invalid group customizations:\nduplicate group ID: 42",
		},
		"duplicate-both": {
			groups: []GroupCustomization{
				{
					Name: "TestGroup",
					GID:  common.ToPtr(1234),
				},
				{
					Name: "wheel",
					GID:  common.ToPtr(42),
				},
				{
					Name: "wheel",
					GID:  common.ToPtr(42),
				},
			},
			expectedErrorMessage: "invalid group customizations:\nduplicate group name: wheel\nduplicate group ID: 42",
		},
		"duplicate-multi": {
			groups: []GroupCustomization{
				{
					Name: "test",
					GID:  common.ToPtr(1234),
				},
				{
					Name: "wheel",
					GID:  common.ToPtr(42),
				},
				{
					Name: "wheel",
					GID:  common.ToPtr(42),
				},
				{
					Name: "user",
					GID:  common.ToPtr(1234),
				},
				{
					Name: "test",
					GID:  common.ToPtr(4321),
				},
			},
			expectedErrorMessage: "invalid group customizations:\nduplicate group name: wheel\nduplicate group ID: 42\nduplicate group ID: 1234\nduplicate group name: test",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			c := Customizations{
				Group: tc.groups,
			}

			groups, err := c.GetGroups()
			if tc.expectedErrorMessage != "" {
				assert.EqualError(err, tc.expectedErrorMessage)
			} else {
				assert.NoError(err)
				assert.ElementsMatch(tc.groups, groups)
			}
		})
	}
}
