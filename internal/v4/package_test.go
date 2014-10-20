// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package v4_test

import (
	"testing"

	jujutesting "github.com/juju/testing"

	"github.com/juju/charmstore/internal/storetesting"
)

func TestPackage(t *testing.T) {
	storetesting.ElasticSearchTestPackage(t, func(t2 *testing.T) {
		jujutesting.MgoTestPackage(t2, nil)
	})
}