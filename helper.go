//: Copyright Verizon Media
//: Licensed under the terms of the Apache 2.0 License. See LICENSE file in the project root for terms.

package vssh

import (
	"golang.org/x/crypto/ssh"
)

// GetConfigPEM returns SSH configuration that uses the given private key.
// the keyfile should be unencrypted PEM-encoded private key file.
func GetConfigPEM(user, keyFile string) (*ssh.ClientConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetConfigUserPass returns SSH configuration that uses the given
// username and password.
func GetConfigUserPass(user, password string) *ssh.ClientConfig {
	_ = "STUB: not implemented"
	return nil
}

func getConfig(user string, auths ...ssh.AuthMethod) *ssh.ClientConfig {
	_ = "STUB: not implemented"
	return nil
}
