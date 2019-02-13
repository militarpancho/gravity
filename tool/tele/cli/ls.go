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

package cli

import (
	"github.com/gravitational/gravity/lib/catalog"
	"github.com/gravitational/gravity/lib/constants"
	"github.com/gravitational/gravity/lib/localenv"

	"github.com/gravitational/trace"
)

func list(env localenv.LocalEnvironment, all bool, format constants.Format) error {
	lister, err := catalog.NewLister()
	if err != nil {
		return trace.Wrap(err)
	}
	err = catalog.List(lister, all, format)
	if err != nil {
		return trace.Wrap(err)
	}
	return nil
}
