// Copyright 2021 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package cose

import "fmt"

type Meta struct {
	Entities Entities  `cbor:"0,keyasint" json:"entities"`
	Validity *Validity `cbor:"1,keyasint,omitempty" json:"validity,omitempty"`
}

func NewMeta() *Meta {
	return &Meta{}
}

func (o *Meta) AddEntity(e Entity) *Meta {
	if o != nil {
		if o.Entities.AddEntity(e) == nil {
			return nil
		}
	}
	return o
}

func (o Meta) Valid() error {
	if err := o.Entities.Valid(); err != nil {
		return fmt.Errorf("entity validation failed: %w", err)
	}

	if o.Validity != nil {
		if err := o.Validity.Valid(); err != nil {
			return fmt.Errorf("entity validation failed: %w", err)
		}
	}

	return nil
}
