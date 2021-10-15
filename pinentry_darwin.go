// +build darwin

package pinentry

import "github.com/gopasspw/pinentry/gpgconf"

// GetBinary always returns pinentry-mac
func GetBinary() string {
	return "pinentry-mac"
}
