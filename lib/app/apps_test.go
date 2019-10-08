/*
Copyright 2018 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package app

import (
	"testing"

	"github.com/gravitational/gravity/lib/loc"
	"github.com/gravitational/gravity/lib/pack"
	"github.com/gravitational/gravity/lib/schema"

	. "gopkg.in/check.v1"
)

func TestAppUtils(t *testing.T) { TestingT(t) }

type AppUtilsSuite struct{}

var _ = Suite(&AppUtilsSuite{})

func (s *AppUtilsSuite) TestUpdatedDependencies(c *C) {
	manifest1, err := schema.ParseManifestYAMLNoValidate([]byte(app1Manifest))
	c.Assert(err, IsNil)
	app1 := Application{
		Package: loc.MustParseLocator("repo/app:1.0.0"),
		PackageEnvelope: pack.PackageEnvelope{
			Manifest: []byte(app1Manifest),
		},
		Manifest: *manifest1,
	}
	manifest2, err := schema.ParseManifestYAMLNoValidate([]byte(app2Manifest))
	c.Assert(err, IsNil)
	app2 := Application{
		Package: loc.MustParseLocator("repo/app:2.0.0"),
		PackageEnvelope: pack.PackageEnvelope{
			Manifest: []byte(app2Manifest),
		},
		Manifest: *manifest2,
	}

	updates, err := GetUpdatedDependencies(app1, app2, *manifest1, *manifest2)
	c.Assert(err, IsNil)
	c.Assert(updates, DeepEquals, []loc.Locator{
		loc.MustParseLocator("repo/dep-2:2.0.0"),
		loc.MustParseLocator("repo/app:2.0.0"),
	})

	updates, err = GetUpdatedDependencies(app1, app1, *manifest1, *manifest1)
	c.Assert(err, IsNil)
	c.Assert(updates, DeepEquals, []loc.Locator(nil))
}

const app1Manifest = `apiVersion: bundle.gravitational.io/v2
kind: Bundle
metadata:
  name: app
  resourceVersion: 1.0.0
dependencies:
  apps:
    - repo/dep-1:1.0.0
    - repo/dep-2:1.0.0`

const app2Manifest = `apiVersion: bundle.gravitational.io/v2
kind: Bundle
metadata:
  name: app
  resourceVersion: 2.0.0
dependencies:
  apps:
    - repo/dep-1:1.0.0
    - repo/dep-2:2.0.0`
