// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegistryServiceConfig RegistryServiceConfig stores daemon registry services configuration.
//
//
// swagger:model RegistryServiceConfig
type RegistryServiceConfig struct {

	// List of IP ranges to which nondistributable artifacts can be pushed,
	// using the CIDR syntax [RFC 4632](https://tools.ietf.org/html/4632).
	//
	// Some images (for example, Windows base images) contain artifacts
	// whose distribution is restricted by license. When these images are
	// pushed to a registry, restricted artifacts are not included.
	//
	// This configuration override this behavior, and enables the daemon to
	// push nondistributable artifacts to all registries whose resolved IP
	// address is within the subnet described by the CIDR syntax.
	//
	// This option is useful when pushing images containing
	// nondistributable artifacts to a registry on an air-gapped network so
	// hosts on that network can pull the images without connecting to
	// another server.
	//
	// > **Warning**: Nondistributable artifacts typically have restrictions
	// > on how and where they can be distributed and shared. Only use this
	// > feature to push artifacts to private registries and ensure that you
	// > are in compliance with any terms that cover redistributing
	// > nondistributable artifacts.
	//
	// Example: ["::1/128","127.0.0.0/8"]
	AllowNondistributableArtifactsCIDRs []string `json:"AllowNondistributableArtifactsCIDRs"`

	// List of registry hostnames to which nondistributable artifacts can be
	// pushed, using the format `<hostname>[:<port>]` or `<IP address>[:<port>]`.
	//
	// Some images (for example, Windows base images) contain artifacts
	// whose distribution is restricted by license. When these images are
	// pushed to a registry, restricted artifacts are not included.
	//
	// This configuration override this behavior for the specified
	// registries.
	//
	// This option is useful when pushing images containing
	// nondistributable artifacts to a registry on an air-gapped network so
	// hosts on that network can pull the images without connecting to
	// another server.
	//
	// > **Warning**: Nondistributable artifacts typically have restrictions
	// > on how and where they can be distributed and shared. Only use this
	// > feature to push artifacts to private registries and ensure that you
	// > are in compliance with any terms that cover redistributing
	// > nondistributable artifacts.
	//
	// Example: ["registry.internal.corp.example.com:3000","[2001:db8:a0b:12f0::1]:443"]
	AllowNondistributableArtifactsHostnames []string `json:"AllowNondistributableArtifactsHostnames"`

	// index configs
	// Example: {"127.0.0.1:5000":{"Mirrors":[],"Name":"127.0.0.1:5000","Official":false,"Secure":false},"[2001:db8:a0b:12f0::1]:80":{"Mirrors":[],"Name":"[2001:db8:a0b:12f0::1]:80","Official":false,"Secure":false},"docker.io":{"Mirrors":["https://hub-mirror.corp.example.com:5000/"],"Name":"docker.io","Official":true,"Secure":true},"registry.internal.corp.example.com:3000":{"Mirrors":[],"Name":"registry.internal.corp.example.com:3000","Official":false,"Secure":false}}
	IndexConfigs map[string]IndexInfo `json:"IndexConfigs,omitempty"`

	// List of IP ranges of insecure registries, using the CIDR syntax
	// ([RFC 4632](https://tools.ietf.org/html/4632)). Insecure registries
	// accept un-encrypted (HTTP) and/or untrusted (HTTPS with certificates
	// from unknown CAs) communication.
	//
	// By default, local registries (`127.0.0.0/8`) are configured as
	// insecure. All other registries are secure. Communicating with an
	// insecure registry is not possible if the daemon assumes that registry
	// is secure.
	//
	// This configuration override this behavior, insecure communication with
	// registries whose resolved IP address is within the subnet described by
	// the CIDR syntax.
	//
	// Registries can also be marked insecure by hostname. Those registries
	// are listed under `IndexConfigs` and have their `Secure` field set to
	// `false`.
	//
	// > **Warning**: Using this option can be useful when running a local
	// > registry, but introduces security vulnerabilities. This option
	// > should therefore ONLY be used for testing purposes. For increased
	// > security, users should add their CA to their system's list of trusted
	// > CAs instead of enabling this option.
	//
	// Example: ["::1/128","127.0.0.0/8"]
	InsecureRegistryCIDRs []string `json:"InsecureRegistryCIDRs"`

	// List of registry URLs that act as a mirror for the official
	// (`docker.io`) registry.
	//
	// Example: ["https://hub-mirror.corp.example.com:5000/","https://[2001:db8:a0b:12f0::1]/"]
	Mirrors []string `json:"Mirrors"`
}

// Validate validates this registry service config
func (m *RegistryServiceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndexConfigs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistryServiceConfig) validateIndexConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.IndexConfigs) { // not required
		return nil
	}

	for k := range m.IndexConfigs {

		if err := validate.Required("IndexConfigs"+"."+k, "body", m.IndexConfigs[k]); err != nil {
			return err
		}
		if val, ok := m.IndexConfigs[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("IndexConfigs" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("IndexConfigs" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this registry service config based on the context it is used
func (m *RegistryServiceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIndexConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistryServiceConfig) contextValidateIndexConfigs(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.IndexConfigs {

		if val, ok := m.IndexConfigs[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistryServiceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistryServiceConfig) UnmarshalBinary(b []byte) error {
	var res RegistryServiceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
