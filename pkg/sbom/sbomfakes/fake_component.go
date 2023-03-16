// Code generated by counterfeiter. DO NOT EDIT.
package sbomfakes

import (
	"sync"

	"github.com/kchetans/sbomqs/pkg/cpe"
	"github.com/kchetans/sbomqs/pkg/purl"
	"github.com/kchetans/sbomqs/pkg/sbom"
)

type FakeComponent struct {
	ChecksumsStub        func() []sbom.Checksum
	checksumsMutex       sync.RWMutex
	checksumsArgsForCall []struct {
	}
	checksumsReturns struct {
		result1 []sbom.Checksum
	}
	checksumsReturnsOnCall map[int]struct {
		result1 []sbom.Checksum
	}
	CpesStub        func() []cpe.CPE
	cpesMutex       sync.RWMutex
	cpesArgsForCall []struct {
	}
	cpesReturns struct {
		result1 []cpe.CPE
	}
	cpesReturnsOnCall map[int]struct {
		result1 []cpe.CPE
	}
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 string
	}
	iDReturnsOnCall map[int]struct {
		result1 string
	}
	LicensesStub        func() []sbom.License
	licensesMutex       sync.RWMutex
	licensesArgsForCall []struct {
	}
	licensesReturns struct {
		result1 []sbom.License
	}
	licensesReturnsOnCall map[int]struct {
		result1 []sbom.License
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	PrimaryPurposeStub        func() string
	primaryPurposeMutex       sync.RWMutex
	primaryPurposeArgsForCall []struct {
	}
	primaryPurposeReturns struct {
		result1 string
	}
	primaryPurposeReturnsOnCall map[int]struct {
		result1 string
	}
	PurlsStub        func() []purl.PURL
	purlsMutex       sync.RWMutex
	purlsArgsForCall []struct {
	}
	purlsReturns struct {
		result1 []purl.PURL
	}
	purlsReturnsOnCall map[int]struct {
		result1 []purl.PURL
	}
	RequiredFieldsStub        func() bool
	requiredFieldsMutex       sync.RWMutex
	requiredFieldsArgsForCall []struct {
	}
	requiredFieldsReturns struct {
		result1 bool
	}
	requiredFieldsReturnsOnCall map[int]struct {
		result1 bool
	}
	SupplierNameStub        func() string
	supplierNameMutex       sync.RWMutex
	supplierNameArgsForCall []struct {
	}
	supplierNameReturns struct {
		result1 string
	}
	supplierNameReturnsOnCall map[int]struct {
		result1 string
	}
	VersionStub        func() string
	versionMutex       sync.RWMutex
	versionArgsForCall []struct {
	}
	versionReturns struct {
		result1 string
	}
	versionReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeComponent) Checksums() []sbom.Checksum {
	fake.checksumsMutex.Lock()
	ret, specificReturn := fake.checksumsReturnsOnCall[len(fake.checksumsArgsForCall)]
	fake.checksumsArgsForCall = append(fake.checksumsArgsForCall, struct {
	}{})
	stub := fake.ChecksumsStub
	fakeReturns := fake.checksumsReturns
	fake.recordInvocation("Checksums", []interface{}{})
	fake.checksumsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) ChecksumsCallCount() int {
	fake.checksumsMutex.RLock()
	defer fake.checksumsMutex.RUnlock()
	return len(fake.checksumsArgsForCall)
}

func (fake *FakeComponent) ChecksumsCalls(stub func() []sbom.Checksum) {
	fake.checksumsMutex.Lock()
	defer fake.checksumsMutex.Unlock()
	fake.ChecksumsStub = stub
}

func (fake *FakeComponent) ChecksumsReturns(result1 []sbom.Checksum) {
	fake.checksumsMutex.Lock()
	defer fake.checksumsMutex.Unlock()
	fake.ChecksumsStub = nil
	fake.checksumsReturns = struct {
		result1 []sbom.Checksum
	}{result1}
}

func (fake *FakeComponent) ChecksumsReturnsOnCall(i int, result1 []sbom.Checksum) {
	fake.checksumsMutex.Lock()
	defer fake.checksumsMutex.Unlock()
	fake.ChecksumsStub = nil
	if fake.checksumsReturnsOnCall == nil {
		fake.checksumsReturnsOnCall = make(map[int]struct {
			result1 []sbom.Checksum
		})
	}
	fake.checksumsReturnsOnCall[i] = struct {
		result1 []sbom.Checksum
	}{result1}
}

func (fake *FakeComponent) Cpes() []cpe.CPE {
	fake.cpesMutex.Lock()
	ret, specificReturn := fake.cpesReturnsOnCall[len(fake.cpesArgsForCall)]
	fake.cpesArgsForCall = append(fake.cpesArgsForCall, struct {
	}{})
	stub := fake.CpesStub
	fakeReturns := fake.cpesReturns
	fake.recordInvocation("Cpes", []interface{}{})
	fake.cpesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) CpesCallCount() int {
	fake.cpesMutex.RLock()
	defer fake.cpesMutex.RUnlock()
	return len(fake.cpesArgsForCall)
}

func (fake *FakeComponent) CpesCalls(stub func() []cpe.CPE) {
	fake.cpesMutex.Lock()
	defer fake.cpesMutex.Unlock()
	fake.CpesStub = stub
}

func (fake *FakeComponent) CpesReturns(result1 []cpe.CPE) {
	fake.cpesMutex.Lock()
	defer fake.cpesMutex.Unlock()
	fake.CpesStub = nil
	fake.cpesReturns = struct {
		result1 []cpe.CPE
	}{result1}
}

func (fake *FakeComponent) CpesReturnsOnCall(i int, result1 []cpe.CPE) {
	fake.cpesMutex.Lock()
	defer fake.cpesMutex.Unlock()
	fake.CpesStub = nil
	if fake.cpesReturnsOnCall == nil {
		fake.cpesReturnsOnCall = make(map[int]struct {
			result1 []cpe.CPE
		})
	}
	fake.cpesReturnsOnCall[i] = struct {
		result1 []cpe.CPE
	}{result1}
}

func (fake *FakeComponent) ID() string {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	stub := fake.IDStub
	fakeReturns := fake.iDReturns
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeComponent) IDCalls(stub func() string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeComponent) IDReturns(result1 string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeComponent) IDReturnsOnCall(i int, result1 string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeComponent) Licenses() []sbom.License {
	fake.licensesMutex.Lock()
	ret, specificReturn := fake.licensesReturnsOnCall[len(fake.licensesArgsForCall)]
	fake.licensesArgsForCall = append(fake.licensesArgsForCall, struct {
	}{})
	stub := fake.LicensesStub
	fakeReturns := fake.licensesReturns
	fake.recordInvocation("Licenses", []interface{}{})
	fake.licensesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) LicensesCallCount() int {
	fake.licensesMutex.RLock()
	defer fake.licensesMutex.RUnlock()
	return len(fake.licensesArgsForCall)
}

func (fake *FakeComponent) LicensesCalls(stub func() []sbom.License) {
	fake.licensesMutex.Lock()
	defer fake.licensesMutex.Unlock()
	fake.LicensesStub = stub
}

func (fake *FakeComponent) LicensesReturns(result1 []sbom.License) {
	fake.licensesMutex.Lock()
	defer fake.licensesMutex.Unlock()
	fake.LicensesStub = nil
	fake.licensesReturns = struct {
		result1 []sbom.License
	}{result1}
}

func (fake *FakeComponent) LicensesReturnsOnCall(i int, result1 []sbom.License) {
	fake.licensesMutex.Lock()
	defer fake.licensesMutex.Unlock()
	fake.LicensesStub = nil
	if fake.licensesReturnsOnCall == nil {
		fake.licensesReturnsOnCall = make(map[int]struct {
			result1 []sbom.License
		})
	}
	fake.licensesReturnsOnCall[i] = struct {
		result1 []sbom.License
	}{result1}
}

func (fake *FakeComponent) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	stub := fake.NameStub
	fakeReturns := fake.nameReturns
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeComponent) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeComponent) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeComponent) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeComponent) PrimaryPurpose() string {
	fake.primaryPurposeMutex.Lock()
	ret, specificReturn := fake.primaryPurposeReturnsOnCall[len(fake.primaryPurposeArgsForCall)]
	fake.primaryPurposeArgsForCall = append(fake.primaryPurposeArgsForCall, struct {
	}{})
	stub := fake.PrimaryPurposeStub
	fakeReturns := fake.primaryPurposeReturns
	fake.recordInvocation("PrimaryPurpose", []interface{}{})
	fake.primaryPurposeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) PrimaryPurposeCallCount() int {
	fake.primaryPurposeMutex.RLock()
	defer fake.primaryPurposeMutex.RUnlock()
	return len(fake.primaryPurposeArgsForCall)
}

func (fake *FakeComponent) PrimaryPurposeCalls(stub func() string) {
	fake.primaryPurposeMutex.Lock()
	defer fake.primaryPurposeMutex.Unlock()
	fake.PrimaryPurposeStub = stub
}

func (fake *FakeComponent) PrimaryPurposeReturns(result1 string) {
	fake.primaryPurposeMutex.Lock()
	defer fake.primaryPurposeMutex.Unlock()
	fake.PrimaryPurposeStub = nil
	fake.primaryPurposeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeComponent) PrimaryPurposeReturnsOnCall(i int, result1 string) {
	fake.primaryPurposeMutex.Lock()
	defer fake.primaryPurposeMutex.Unlock()
	fake.PrimaryPurposeStub = nil
	if fake.primaryPurposeReturnsOnCall == nil {
		fake.primaryPurposeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.primaryPurposeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeComponent) Purls() []purl.PURL {
	fake.purlsMutex.Lock()
	ret, specificReturn := fake.purlsReturnsOnCall[len(fake.purlsArgsForCall)]
	fake.purlsArgsForCall = append(fake.purlsArgsForCall, struct {
	}{})
	stub := fake.PurlsStub
	fakeReturns := fake.purlsReturns
	fake.recordInvocation("Purls", []interface{}{})
	fake.purlsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) PurlsCallCount() int {
	fake.purlsMutex.RLock()
	defer fake.purlsMutex.RUnlock()
	return len(fake.purlsArgsForCall)
}

func (fake *FakeComponent) PurlsCalls(stub func() []purl.PURL) {
	fake.purlsMutex.Lock()
	defer fake.purlsMutex.Unlock()
	fake.PurlsStub = stub
}

func (fake *FakeComponent) PurlsReturns(result1 []purl.PURL) {
	fake.purlsMutex.Lock()
	defer fake.purlsMutex.Unlock()
	fake.PurlsStub = nil
	fake.purlsReturns = struct {
		result1 []purl.PURL
	}{result1}
}

func (fake *FakeComponent) PurlsReturnsOnCall(i int, result1 []purl.PURL) {
	fake.purlsMutex.Lock()
	defer fake.purlsMutex.Unlock()
	fake.PurlsStub = nil
	if fake.purlsReturnsOnCall == nil {
		fake.purlsReturnsOnCall = make(map[int]struct {
			result1 []purl.PURL
		})
	}
	fake.purlsReturnsOnCall[i] = struct {
		result1 []purl.PURL
	}{result1}
}

func (fake *FakeComponent) RequiredFields() bool {
	fake.requiredFieldsMutex.Lock()
	ret, specificReturn := fake.requiredFieldsReturnsOnCall[len(fake.requiredFieldsArgsForCall)]
	fake.requiredFieldsArgsForCall = append(fake.requiredFieldsArgsForCall, struct {
	}{})
	stub := fake.RequiredFieldsStub
	fakeReturns := fake.requiredFieldsReturns
	fake.recordInvocation("RequiredFields", []interface{}{})
	fake.requiredFieldsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) RequiredFieldsCallCount() int {
	fake.requiredFieldsMutex.RLock()
	defer fake.requiredFieldsMutex.RUnlock()
	return len(fake.requiredFieldsArgsForCall)
}

func (fake *FakeComponent) RequiredFieldsCalls(stub func() bool) {
	fake.requiredFieldsMutex.Lock()
	defer fake.requiredFieldsMutex.Unlock()
	fake.RequiredFieldsStub = stub
}

func (fake *FakeComponent) RequiredFieldsReturns(result1 bool) {
	fake.requiredFieldsMutex.Lock()
	defer fake.requiredFieldsMutex.Unlock()
	fake.RequiredFieldsStub = nil
	fake.requiredFieldsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeComponent) RequiredFieldsReturnsOnCall(i int, result1 bool) {
	fake.requiredFieldsMutex.Lock()
	defer fake.requiredFieldsMutex.Unlock()
	fake.RequiredFieldsStub = nil
	if fake.requiredFieldsReturnsOnCall == nil {
		fake.requiredFieldsReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.requiredFieldsReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeComponent) SupplierName() string {
	fake.supplierNameMutex.Lock()
	ret, specificReturn := fake.supplierNameReturnsOnCall[len(fake.supplierNameArgsForCall)]
	fake.supplierNameArgsForCall = append(fake.supplierNameArgsForCall, struct {
	}{})
	stub := fake.SupplierNameStub
	fakeReturns := fake.supplierNameReturns
	fake.recordInvocation("SupplierName", []interface{}{})
	fake.supplierNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) SupplierNameCallCount() int {
	fake.supplierNameMutex.RLock()
	defer fake.supplierNameMutex.RUnlock()
	return len(fake.supplierNameArgsForCall)
}

func (fake *FakeComponent) SupplierNameCalls(stub func() string) {
	fake.supplierNameMutex.Lock()
	defer fake.supplierNameMutex.Unlock()
	fake.SupplierNameStub = stub
}

func (fake *FakeComponent) SupplierNameReturns(result1 string) {
	fake.supplierNameMutex.Lock()
	defer fake.supplierNameMutex.Unlock()
	fake.SupplierNameStub = nil
	fake.supplierNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeComponent) SupplierNameReturnsOnCall(i int, result1 string) {
	fake.supplierNameMutex.Lock()
	defer fake.supplierNameMutex.Unlock()
	fake.SupplierNameStub = nil
	if fake.supplierNameReturnsOnCall == nil {
		fake.supplierNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.supplierNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeComponent) Version() string {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct {
	}{})
	stub := fake.VersionStub
	fakeReturns := fake.versionReturns
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeComponent) VersionCalls(stub func() string) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = stub
}

func (fake *FakeComponent) VersionReturns(result1 string) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeComponent) VersionReturnsOnCall(i int, result1 string) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeComponent) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checksumsMutex.RLock()
	defer fake.checksumsMutex.RUnlock()
	fake.cpesMutex.RLock()
	defer fake.cpesMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.licensesMutex.RLock()
	defer fake.licensesMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.primaryPurposeMutex.RLock()
	defer fake.primaryPurposeMutex.RUnlock()
	fake.purlsMutex.RLock()
	defer fake.purlsMutex.RUnlock()
	fake.requiredFieldsMutex.RLock()
	defer fake.requiredFieldsMutex.RUnlock()
	fake.supplierNameMutex.RLock()
	defer fake.supplierNameMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeComponent) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ sbom.Component = new(FakeComponent)
