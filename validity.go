// Copyright 2021 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package cose

import (
	"fmt"
	"time"
)

type Validity struct {
	NotBefore *time.Time `cbor:"0,keyasint,omitempty" json:"not-before,omitempty"`
	NotAfter  time.Time  `cbor:"1,keyasint" json:"not-after"`
}

func (o Validity) Valid() error {
	if o.NotBefore != nil {
		if delta := o.NotAfter.Sub(*o.NotBefore); delta < 0 {
			return fmt.Errorf("invalid not-before / not-after: negative delta (%d)", delta)
		}
	}
	return nil
}
