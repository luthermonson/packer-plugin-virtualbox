package common

import (
	"github.com/hashicorp/packer/helper/multistep"
)

func CommHost(host string) func(multistep.StateBag) (string, error) {
	return func(state multistep.StateBag) (string, error) {
		return host, nil
	}
}

func SSHPort(state multistep.StateBag) (uint, error) {
	sshHostPort := state.Get("sshHostPort").(uint)
	return sshHostPort, nil
}
