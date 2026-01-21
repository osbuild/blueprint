package blueprint

import "strings"

func (c *Customizations) GetUsers() []UserCustomization {
	if c == nil || (c.User == nil && c.SSHKey == nil) {
		return nil
	}

	var users []UserCustomization

	// prepend sshkey for backwards compat (overridden by users)
	if len(c.SSHKey) > 0 {
		for _, k := range c.SSHKey {
			key := k.Key
			users = append(users, UserCustomization{
				Name: k.User,
				Key:  &key,
			})
		}
	}

	users = append(users, c.User...)

	// sanitize user home directory in blueprint: if it has a trailing slash,
	// it might lead to the directory not getting the correct selinux labels
	for idx := range users {
		u := users[idx]
		if u.Home != nil {
			homedir := strings.TrimRight(*u.Home, "/")
			u.Home = &homedir
			users[idx] = u
		}
	}
	return users
}

func (c *Customizations) GetGroups() []GroupCustomization {
	if c == nil {
		return nil
	}

	return c.Group
}
