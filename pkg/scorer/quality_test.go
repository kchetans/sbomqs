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
	"testing"

	"github.com/kchetans/sbomqs/pkg/cpe"
	"github.com/kchetans/sbomqs/pkg/purl"
	"github.com/kchetans/sbomqs/pkg/sbom"
	"github.com/kchetans/sbomqs/pkg/sbom/sbomfakes"
)

func sampleDocs() []sbom.Document {
	var fakeLicRestrictive = &sbomfakes.FakeLicense{}
	var fakeRestComp = &sbomfakes.FakeComponent{}
	var fakeDoc2 = &sbomfakes.FakeDocument{}
	fakeLicRestrictive.NameReturns("GPL-3.0")
	fakeRestComp.LicensesReturns([]sbom.License{fakeLicRestrictive})
	fakeDoc2.ComponentsReturns([]sbom.Component{fakeRestComp})

	var fakeLicNonRestrictive = &sbomfakes.FakeLicense{}
	var fakeComp = &sbomfakes.FakeComponent{}
	var fakeDoc = &sbomfakes.FakeDocument{}
	fakeLicNonRestrictive.NameReturns("MIT")
	fakeComp.LicensesReturns([]sbom.License{fakeLicNonRestrictive})
	fakeDoc.ComponentsReturns([]sbom.Component{fakeComp})

	var fakeComp3 = &sbomfakes.FakeComponent{}
	var fakeDoc3 = &sbomfakes.FakeDocument{}
	fakeComp3.LicensesReturns([]sbom.License{})
	fakeDoc3.ComponentsReturns([]sbom.Component{fakeComp3})

	var fakeLicDep = &sbomfakes.FakeLicense{}
	var fakeComp4 = &sbomfakes.FakeComponent{}
	var fakeDoc4 = &sbomfakes.FakeDocument{}
	fakeLicDep.DeprecatedReturns(true)
	fakeComp4.LicensesReturns([]sbom.License{fakeLicDep})
	fakeDoc4.ComponentsReturns([]sbom.Component{fakeComp4})

	return []sbom.Document{fakeDoc, fakeDoc2, fakeDoc3, fakeDoc4}
}

func Test_compWithRestrictedLicensesScore(t *testing.T) {
	testDocs := sampleDocs()

	type args struct {
		d sbom.Document
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Doc has no restrictive licenses", args{d: testDocs[0]}, 10.0},
		{"Doc has restrictive licenses", args{d: testDocs[1]}, 0.0},
		{"Doc has no licenses", args{d: testDocs[2]}, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compWithRestrictedLicensesScore(tt.args.d); got.score != tt.want {
				t.Errorf("compWithRestrictedLicensesScore() = %v, want %v", got.score, tt.want)
			}
		})
	}
}

func Test_compWithNoDepLicensesScore(t *testing.T) {
	testDocs := sampleDocs()
	type args struct {
		d sbom.Document
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Doc has no deprecated licenses", args{d: testDocs[0]}, 10.0},
		{"Doc has deprecated licenses", args{d: testDocs[3]}, 0.0},
		{"Doc has no licenses", args{d: testDocs[2]}, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compWithNoDepLicensesScore(tt.args.d); got.score != tt.want {
				t.Errorf("compWithNoDepLicensesScore() = %v, want %v", got.score, tt.want)
			}
		})
	}
}

func Test_compWithAnyLookupIdScore(t *testing.T) {
	testDocs := func() []sbom.Document {
		doc := &sbomfakes.FakeDocument{}
		comp := &sbomfakes.FakeComponent{}
		comp.CpesReturns([]cpe.CPE{})
		comp.PurlsReturns([]purl.PURL{})
		doc.ComponentsReturns([]sbom.Component{comp})

		docWithCPE := &sbomfakes.FakeDocument{}
		comp2 := &sbomfakes.FakeComponent{}
		comp2.CpesReturns([]cpe.CPE{cpe.NewCPE("cpe:2.3:a:spdx:tools-golang:v0.4.0:*:*:*:*:*:*:*")})
		comp2.PurlsReturns([]purl.PURL{})
		docWithCPE.ComponentsReturns([]sbom.Component{comp2})

		docWithPURL := &sbomfakes.FakeDocument{}
		comp3 := &sbomfakes.FakeComponent{}
		comp3.CpesReturns([]cpe.CPE{})
		comp3.PurlsReturns([]purl.PURL{purl.NewPURL("pkg:golang/github.com/spf13/pflag@1.0.5")})
		docWithPURL.ComponentsReturns([]sbom.Component{comp3})

		docWithBoth := &sbomfakes.FakeDocument{}
		comp5 := &sbomfakes.FakeComponent{}
		comp5.CpesReturns([]cpe.CPE{cpe.NewCPE("cpe:2.3:a:spdx:tools-golang:v0.4.0:*:*:*:*:*:*:*")})
		comp5.PurlsReturns([]purl.PURL{purl.NewPURL("pkg:golang/github.com/spf13/pflag@1.0.5")})
		docWithBoth.ComponentsReturns([]sbom.Component{comp5})

		return []sbom.Document{doc, docWithCPE, docWithPURL, docWithBoth}

	}()

	type args struct {
		d sbom.Document
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Doc with no vuln lookup ids", args{d: testDocs[0]}, 0.0},
		{"Doc with cpe only vuln lookup ids", args{d: testDocs[1]}, 10.0},
		{"Doc with purl only vuln lookup ids", args{d: testDocs[2]}, 10.0},
		{"Doc with both cpe & purl vuln lookup ids", args{d: testDocs[3]}, 10.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compWithAnyLookupIdScore(tt.args.d); got.score != tt.want {
				t.Errorf("compWithAnyLookupIdScore() = %v, want %v", got.score, tt.want)
			}
		})
	}
}

func Test_compWithMultipleIdScore(t *testing.T) {
	testDocs := func() []sbom.Document {
		doc := &sbomfakes.FakeDocument{}
		comp := &sbomfakes.FakeComponent{}
		comp.CpesReturns([]cpe.CPE{})
		comp.PurlsReturns([]purl.PURL{})
		doc.ComponentsReturns([]sbom.Component{comp})

		docWithCPE := &sbomfakes.FakeDocument{}
		comp2 := &sbomfakes.FakeComponent{}
		comp2.CpesReturns([]cpe.CPE{cpe.NewCPE("cpe:2.3:a:spdx:tools-golang:v0.4.0:*:*:*:*:*:*:*")})
		comp2.PurlsReturns([]purl.PURL{})
		docWithCPE.ComponentsReturns([]sbom.Component{comp2})

		docWithPURL := &sbomfakes.FakeDocument{}
		comp3 := &sbomfakes.FakeComponent{}
		comp3.CpesReturns([]cpe.CPE{})
		comp3.PurlsReturns([]purl.PURL{purl.NewPURL("pkg:golang/github.com/spf13/pflag@1.0.5")})
		docWithPURL.ComponentsReturns([]sbom.Component{comp3})

		docWithBoth := &sbomfakes.FakeDocument{}
		comp5 := &sbomfakes.FakeComponent{}
		comp5.CpesReturns([]cpe.CPE{cpe.NewCPE("cpe:2.3:a:spdx:tools-golang:v0.4.0:*:*:*:*:*:*:*")})
		comp5.PurlsReturns([]purl.PURL{purl.NewPURL("pkg:golang/github.com/spf13/pflag@1.0.5")})
		docWithBoth.ComponentsReturns([]sbom.Component{comp5})

		return []sbom.Document{doc, docWithCPE, docWithPURL, docWithBoth}

	}()
	type args struct {
		d sbom.Document
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Doc with no vuln lookup ids", args{d: testDocs[0]}, 0.0},
		{"Doc with cpe only vuln lookup ids", args{d: testDocs[1]}, 0.0},
		{"Doc with purl only vuln lookup ids", args{d: testDocs[2]}, 0.0},
		{"Doc with both cpe & purl vuln lookup ids", args{d: testDocs[3]}, 10.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compWithMultipleIdScore(tt.args.d); got.score != tt.want {
				t.Errorf("compWithMultipleIdScore() = %v, want %v", got.score, tt.want)
			}
		})
	}
}

type tool struct {
	name    string
	version string
}

func (t tool) Name() string {
	return t.name
}
func (t tool) Version() string {
	return t.version
}

func Test_docWithCreatorScore(t *testing.T) {
	testDocs := func() []sbom.Document {
		doc := &sbomfakes.FakeDocument{}
		doc.ToolsReturns([]sbom.Tool{})

		docWithCreator := &sbomfakes.FakeDocument{}
		docWithCreator.ToolsReturns([]sbom.Tool{tool{name: "cyclonedx-gomod", version: "v0.1.0"}})

		docWithTool := &sbomfakes.FakeDocument{}
		docWithTool.ToolsReturns([]sbom.Tool{tool{name: "cyclonedx-gomod", version: ""}})

		docWithVer := &sbomfakes.FakeDocument{}
		docWithVer.ToolsReturns([]sbom.Tool{tool{name: "", version: "v0.1.0"}})
		return []sbom.Document{doc, docWithCreator, docWithTool, docWithVer}
	}
	type args struct {
		d sbom.Document
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"doc with no tool info", args{d: testDocs()[0]}, 0.0},
		{"doc with tool info", args{d: testDocs()[1]}, 10.0},
		{"doc with tool and no version", args{d: testDocs()[2]}, 0.0},
		{"doc with no tool has version", args{d: testDocs()[3]}, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := docWithCreatorScore(tt.args.d); got.score != tt.want {
				t.Errorf("docWithCreatorScore() = %v, want %v", got.score, tt.want)
			}
		})
	}
}
