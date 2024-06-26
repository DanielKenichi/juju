// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package user

import (
	"encoding/asn1"
	"encoding/base64"

	"github.com/juju/errors"
	"gopkg.in/yaml.v2"

	"github.com/juju/juju/jujuclient"
)

// ControllerCommand defines a subset methods of the modelcmd.ControllerCommandBase that the generateUserControllerAccessToken function uses.
type ControllerCommand interface {
	ControllerName() (string, error)
	ClientStore() jujuclient.ClientStore
}

func generateUserControllerAccessToken(command ControllerCommand, username string, secretKey []byte) (string, error) {
	controllerName, err := command.ControllerName()
	if err != nil {
		return "", errors.Trace(err)
	}

	// Generate the base64-encoded string for the user to pass to
	// "juju register". We marshal the information using ASN.1
	// to keep the size down, since we need to encode binary data.
	controllerDetails, err := command.ClientStore().ControllerByName(controllerName)
	if err != nil {
		return "", errors.Trace(err)
	}
	registrationInfo := jujuclient.RegistrationInfo{
		User:           username,
		Addrs:          controllerDetails.APIEndpoints,
		SecretKey:      secretKey,
		ControllerName: controllerName,
	}
	if controllerDetails.Proxy != nil {
		controllerDetails.Proxy.Proxier.Insecure()
		proxyConfig, err := yaml.Marshal(controllerDetails.Proxy)
		if err != nil {
			return "", errors.Trace(err)
		}
		if len(proxyConfig) > 0 {
			registrationInfo.ProxyConfig = string(proxyConfig)
		}
	}
	registrationData, err := asn1.Marshal(registrationInfo)
	if err != nil {
		return "", errors.Trace(err)
	}

	// Use URLEncoding so we don't get + or / in the string,
	// and pad with zero bytes so we don't get =; this all
	// makes it easier to copy & paste in a terminal.
	//
	// The embedded ASN.1 data is length-encoded, so the
	// padding will not complicate decoding.
	remainder := len(registrationData) % 3
	if remainder != 0 {
		var pad [3]byte
		registrationData = append(registrationData, pad[:3-remainder]...)
	}
	return base64.URLEncoding.EncodeToString(registrationData), nil
}
