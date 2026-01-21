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
		groups []GroupCustomization
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
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := Customizations{
				Group: tc.groups,
			}

			groups, err := c.GetGroups()
			assert.NoError(t, err)
			assert.ElementsMatch(t, tc.groups, groups)
		})
	}
}
