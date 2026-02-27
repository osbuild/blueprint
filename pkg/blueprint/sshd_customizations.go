package blueprint

type SshdCustomization struct {
	PasswordAuthentication          *bool `json:"password_authentication,omitempty" toml:"password_authentication,omitempty"`
	ChallengeResponseAuthentication *bool `json:"challenge_response_authentication,omitempty" toml:"challenge_response_authentication,omitempty"`
	ClientAliveInterval             *int  `json:"client_alive_interval,omitempty" toml:"client_alive_interval,omitempty"`
	PermitRootLogin                 any   `json:"permit_root_login,omitempty" toml:"permit_root_login,omitempty"`
}
