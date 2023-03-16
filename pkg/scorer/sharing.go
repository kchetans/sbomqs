// Copyright 2023 Interlynk.io
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scorer

import (
	"fmt"

	"github.com/kchetans/sbomqs/pkg/licenses"
	"github.com/kchetans/sbomqs/pkg/sbom"
	"github.com/samber/lo"
)

func sharableLicenseScore(d sbom.Document) score {
	s := newScore(CategorySharing, string(docShareLicense))

	lics := d.Spec().Licenses()

	freeLics := lo.CountBy(lics, func(l sbom.License) bool {
		return licenses.FreeForAnyUse(l.Short())
	})

	if len(lics) > 0 && freeLics == len(lics) {
		s.setScore(10.0)
	}

	s.setDesc(fmt.Sprintf("doc has a sharable license free %d :: of %d", freeLics, len(lics)))
	return *s
}
