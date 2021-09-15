// Copyright 2021 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package cose

import (
	"errors"
	"fmt"
)

type Role int64

/*
$corim-role-type-choice /= corim.manifest-creator
$corim-role-type-choice /= corim.manifest-signer

; $entity-name-type-choice values
;
; Person or organization responsible for creating the CoRIM
corim.manifest-creator = 1

; Person or organization responsible for signing the CoRIM
corim.manifest-signer = 2
*/

const (
	RoleManifestCreator Role = iota + 1
	RoleManifestSigner
)

type Roles []Role

/*
func NewRoles() *Roles {
	return new(Roles)
}
*/

func (o *Roles) Add(roles ...Role) *Roles {
	if o != nil {
		*o = append(*o, roles...)
	}

	return o
}

func isRole(r Role) bool {
	if r != RoleManifestCreator && r != RoleManifestSigner {
		return false
	}
	return true
}

func (o Roles) Valid() error {
	if len(o) == 0 {
		return errors.New("empty roles")
	}

	for i, r := range o {
		if !isRole(r) {
			return fmt.Errorf("unknown role %d at index %d", r, i)
		}
	}

	return nil
}

/*
func (o Roles) ToCBOR() ([]byte, error) {
	if err := o.Valid(); err != nil {
		return nil, err
	}

	data, err := em.Marshal(o)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (o *Roles) FromCBOR(data []byte) error {
	err := dm.Unmarshal(data, o)
	if err != nil {
		return err
	}

	return o.Valid()
}

func (o *Roles) UnmarshalJSON(data []byte) error {
	var a []string

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	if len(a) == 0 {
		return fmt.Errorf("no roles found")
	}

	var r Role

	for _, s := range a {
		switch s {
		case "creator":
			r = RoleManifestCreator
		case "signer":
			r = RoleManifestSigner
		default:
			return fmt.Errorf("unknown role '%s'", s)
		}
		o = o.Add(r)
	}

	return nil
}
*/
