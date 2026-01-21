package blueprint

import (
	"testing"

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
	GID := 1234
	expectedGroups := []GroupCustomization{
		{
			Name: "TestGroup",
			GID:  &GID,
		},
	}

	TestCustomizations := Customizations{
		Group: expectedGroups,
	}

	retGroups := TestCustomizations.GetGroups()

	assert.ElementsMatch(t, expectedGroups, retGroups)
}
