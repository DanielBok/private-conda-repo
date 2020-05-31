package config

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type userConfig struct {
	EmailDomain string         `mapstructure:"email_domain"`
	domainRegex *regexp.Regexp `mapstructure:"-"`
}

func (u *userConfig) Init() error {
	u.EmailDomain = strings.TrimSpace(u.EmailDomain)
	if u.EmailDomain == "" {
		u.domainRegex = nil
		return nil
	}

	re, err := regexp.Compile(u.EmailDomain)
	if err != nil {
		return errors.Errorf("could not compile '%s' as a valid regex", u.EmailDomain)
	}
	u.domainRegex = re
	return nil
}

// Returns true if email is valid. Otherwise false
func (u *userConfig) ValidateEmail(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	if u.domainRegex == nil {
		return true
	}

	return u.domainRegex.MatchString(parts[1])
}
