package entity

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type Channel struct {
	Id        int       `json:"id"`
	Channel   string    `json:"channel"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email"`
	CreatedOn time.Time `json:"createdOn"`

	PackageCounts []PackageCount `json:"package_counts"`
}

const (
	saltLen = 8
	joinKey = ";;;;"
)

func (c *Channel) HasValidPassword(password string) bool {
	salt := strings.Split(c.Password, joinKey)[1]
	return hashPassword(password, salt) == c.Password
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
	_, err := rand.Read(b)
	if err != nil {
		panic(errors.Wrap(err, "This should not happen!"))
	}

	str := fmt.Sprintf("%x", b)
	return str[:saltLen]
}
