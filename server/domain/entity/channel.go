package entity

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"

	"private-conda-repo/config"
)

type Channel struct {
	Id        int       `json:"id"`
	Channel   string    `json:"channel"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email"`
	CreatedOn time.Time `json:"created_on"`

	PackageCounts []PackageCount
}

const (
	saltLen = 8
	joinKey = ";;;;"
)

func (c *Channel) HasValidPassword(password string) bool {
	salt := strings.Split(c.Password, joinKey)[1]
	return hashPassword(password, salt) == c.Password
}

func (c *Channel) IsValid() (err error) {
	nameRegex := regexp.MustCompile(`^\w[\w\-]{2,50}$`)

	conf, err := config.New()
	if err != nil {
		return err
	}

	c.Channel = strings.TrimSpace(c.Channel)
	c.Password = strings.TrimSpace(c.Password)
	c.Email = strings.TrimSpace(c.Email)

	if !nameRegex.MatchString(c.Channel) {
		err = multierror.Append(err, errors.New("user/channel name length must be between [2, 50] characters and can only be alphanumeric with dashes"))
	}
	if len(c.Password) < 4 {
		err = multierror.Append(err, errors.New("password must be >= 4 characters"))
	}
	if !conf.UserConfig.ValidateEmail(c.Email) {
		err = multierror.Append(err, errors.New("invalid email address"))
	}

	return
}

// This is the proper way to change the password. Setting the password directly on the struct field
// does not hash it. Meaning that subsequently, when checking if the password is valid, the check
// will fail since the check will hash the incoming password.
func (c *Channel) SetPassword(password string) {
	salt := generateSalt()
	c.Password = hashPassword(password, salt)
}

func NewChannel(name, password, email string) *Channel {
	c := &Channel{
		Channel:   strings.ToLower(strings.TrimSpace(name)),
		Email:     strings.TrimSpace(strings.ToLower(email)),
		CreatedOn: time.Now().UTC(),
	}
	c.SetPassword(password)

	return c
}

func hashPassword(plainPassword, salt string) string {
	h := sha256.New()
	h.Write([]byte(plainPassword + salt))
	p := fmt.Sprintf("%x", h.Sum(nil))[:52]

	return p + joinKey + salt
}

func generateSalt() string {
	b := make([]byte, saltLen)
	rand.Read(b)
	str := fmt.Sprintf("%x", b)
	return str[:saltLen]
}
