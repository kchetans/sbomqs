// Code generated by counterfeiter. DO NOT EDIT.
package sbomfakes

import (
	"sync"

	"github.com/kchetans/sbomqs/pkg/sbom"
)

type FakeDocument struct {
	AuthorsStub        func() []sbom.Author
	authorsMutex       sync.RWMutex
	authorsArgsForCall []struct {
	}
	authorsReturns struct {
		result1 []sbom.Author
	}
	authorsReturnsOnCall map[int]struct {
		result1 []sbom.Author
	}
	ComponentsStub        func() []sbom.Component
	componentsMutex       sync.RWMutex
	componentsArgsForCall []struct {
	}
	componentsReturns struct {
		result1 []sbom.Component
	}
	componentsReturnsOnCall map[int]struct {
		result1 []sbom.Component
	}
	LogsStub        func() []string
	logsMutex       sync.RWMutex
	logsArgsForCall []struct {
	}
	logsReturns struct {
		result1 []string
	}
	logsReturnsOnCall map[int]struct {
		result1 []string
	}
	RelationsStub        func() []sbom.Relation
	relationsMutex       sync.RWMutex
	relationsArgsForCall []struct {
	}
	relationsReturns struct {
		result1 []sbom.Relation
	}
	relationsReturnsOnCall map[int]struct {
		result1 []sbom.Relation
	}
	SpecStub        func() sbom.Spec
	specMutex       sync.RWMutex
	specArgsForCall []struct {
	}
	specReturns struct {
		result1 sbom.Spec
	}
	specReturnsOnCall map[int]struct {
		result1 sbom.Spec
	}
	ToolsStub        func() []sbom.Tool
	toolsMutex       sync.RWMutex
	toolsArgsForCall []struct {
	}
	toolsReturns struct {
		result1 []sbom.Tool
	}
	toolsReturnsOnCall map[int]struct {
		result1 []sbom.Tool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDocument) Authors() []sbom.Author {
	fake.authorsMutex.Lock()
	ret, specificReturn := fake.authorsReturnsOnCall[len(fake.authorsArgsForCall)]
	fake.authorsArgsForCall = append(fake.authorsArgsForCall, struct {
	}{})
	stub := fake.AuthorsStub
	fakeReturns := fake.authorsReturns
	fake.recordInvocation("Authors", []interface{}{})
	fake.authorsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDocument) AuthorsCallCount() int {
	fake.authorsMutex.RLock()
	defer fake.authorsMutex.RUnlock()
	return len(fake.authorsArgsForCall)
}

func (fake *FakeDocument) AuthorsCalls(stub func() []sbom.Author) {
	fake.authorsMutex.Lock()
	defer fake.authorsMutex.Unlock()
	fake.AuthorsStub = stub
}

func (fake *FakeDocument) AuthorsReturns(result1 []sbom.Author) {
	fake.authorsMutex.Lock()
	defer fake.authorsMutex.Unlock()
	fake.AuthorsStub = nil
	fake.authorsReturns = struct {
		result1 []sbom.Author
	}{result1}
}

func (fake *FakeDocument) AuthorsReturnsOnCall(i int, result1 []sbom.Author) {
	fake.authorsMutex.Lock()
	defer fake.authorsMutex.Unlock()
	fake.AuthorsStub = nil
	if fake.authorsReturnsOnCall == nil {
		fake.authorsReturnsOnCall = make(map[int]struct {
			result1 []sbom.Author
		})
	}
	fake.authorsReturnsOnCall[i] = struct {
		result1 []sbom.Author
	}{result1}
}

func (fake *FakeDocument) Components() []sbom.Component {
	fake.componentsMutex.Lock()
	ret, specificReturn := fake.componentsReturnsOnCall[len(fake.componentsArgsForCall)]
	fake.componentsArgsForCall = append(fake.componentsArgsForCall, struct {
	}{})
	stub := fake.ComponentsStub
	fakeReturns := fake.componentsReturns
	fake.recordInvocation("Components", []interface{}{})
	fake.componentsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDocument) ComponentsCallCount() int {
	fake.componentsMutex.RLock()
	defer fake.componentsMutex.RUnlock()
	return len(fake.componentsArgsForCall)
}

func (fake *FakeDocument) ComponentsCalls(stub func() []sbom.Component) {
	fake.componentsMutex.Lock()
	defer fake.componentsMutex.Unlock()
	fake.ComponentsStub = stub
}

func (fake *FakeDocument) ComponentsReturns(result1 []sbom.Component) {
	fake.componentsMutex.Lock()
	defer fake.componentsMutex.Unlock()
	fake.ComponentsStub = nil
	fake.componentsReturns = struct {
		result1 []sbom.Component
	}{result1}
}

func (fake *FakeDocument) ComponentsReturnsOnCall(i int, result1 []sbom.Component) {
	fake.componentsMutex.Lock()
	defer fake.componentsMutex.Unlock()
	fake.ComponentsStub = nil
	if fake.componentsReturnsOnCall == nil {
		fake.componentsReturnsOnCall = make(map[int]struct {
			result1 []sbom.Component
		})
	}
	fake.componentsReturnsOnCall[i] = struct {
		result1 []sbom.Component
	}{result1}
}

func (fake *FakeDocument) Logs() []string {
	fake.logsMutex.Lock()
	ret, specificReturn := fake.logsReturnsOnCall[len(fake.logsArgsForCall)]
	fake.logsArgsForCall = append(fake.logsArgsForCall, struct {
	}{})
	stub := fake.LogsStub
	fakeReturns := fake.logsReturns
	fake.recordInvocation("Logs", []interface{}{})
	fake.logsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDocument) LogsCallCount() int {
	fake.logsMutex.RLock()
	defer fake.logsMutex.RUnlock()
	return len(fake.logsArgsForCall)
}

func (fake *FakeDocument) LogsCalls(stub func() []string) {
	fake.logsMutex.Lock()
	defer fake.logsMutex.Unlock()
	fake.LogsStub = stub
}

func (fake *FakeDocument) LogsReturns(result1 []string) {
	fake.logsMutex.Lock()
	defer fake.logsMutex.Unlock()
	fake.LogsStub = nil
	fake.logsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeDocument) LogsReturnsOnCall(i int, result1 []string) {
	fake.logsMutex.Lock()
	defer fake.logsMutex.Unlock()
	fake.LogsStub = nil
	if fake.logsReturnsOnCall == nil {
		fake.logsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.logsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeDocument) Relations() []sbom.Relation {
	fake.relationsMutex.Lock()
	ret, specificReturn := fake.relationsReturnsOnCall[len(fake.relationsArgsForCall)]
	fake.relationsArgsForCall = append(fake.relationsArgsForCall, struct {
	}{})
	stub := fake.RelationsStub
	fakeReturns := fake.relationsReturns
	fake.recordInvocation("Relations", []interface{}{})
	fake.relationsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDocument) RelationsCallCount() int {
	fake.relationsMutex.RLock()
	defer fake.relationsMutex.RUnlock()
	return len(fake.relationsArgsForCall)
}

func (fake *FakeDocument) RelationsCalls(stub func() []sbom.Relation) {
	fake.relationsMutex.Lock()
	defer fake.relationsMutex.Unlock()
	fake.RelationsStub = stub
}

func (fake *FakeDocument) RelationsReturns(result1 []sbom.Relation) {
	fake.relationsMutex.Lock()
	defer fake.relationsMutex.Unlock()
	fake.RelationsStub = nil
	fake.relationsReturns = struct {
		result1 []sbom.Relation
	}{result1}
}

func (fake *FakeDocument) RelationsReturnsOnCall(i int, result1 []sbom.Relation) {
	fake.relationsMutex.Lock()
	defer fake.relationsMutex.Unlock()
	fake.RelationsStub = nil
	if fake.relationsReturnsOnCall == nil {
		fake.relationsReturnsOnCall = make(map[int]struct {
			result1 []sbom.Relation
		})
	}
	fake.relationsReturnsOnCall[i] = struct {
		result1 []sbom.Relation
	}{result1}
}

func (fake *FakeDocument) Spec() sbom.Spec {
	fake.specMutex.Lock()
	ret, specificReturn := fake.specReturnsOnCall[len(fake.specArgsForCall)]
	fake.specArgsForCall = append(fake.specArgsForCall, struct {
	}{})
	stub := fake.SpecStub
	fakeReturns := fake.specReturns
	fake.recordInvocation("Spec", []interface{}{})
	fake.specMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDocument) SpecCallCount() int {
	fake.specMutex.RLock()
	defer fake.specMutex.RUnlock()
	return len(fake.specArgsForCall)
}

func (fake *FakeDocument) SpecCalls(stub func() sbom.Spec) {
	fake.specMutex.Lock()
	defer fake.specMutex.Unlock()
	fake.SpecStub = stub
}

func (fake *FakeDocument) SpecReturns(result1 sbom.Spec) {
	fake.specMutex.Lock()
	defer fake.specMutex.Unlock()
	fake.SpecStub = nil
	fake.specReturns = struct {
		result1 sbom.Spec
	}{result1}
}

func (fake *FakeDocument) SpecReturnsOnCall(i int, result1 sbom.Spec) {
	fake.specMutex.Lock()
	defer fake.specMutex.Unlock()
	fake.SpecStub = nil
	if fake.specReturnsOnCall == nil {
		fake.specReturnsOnCall = make(map[int]struct {
			result1 sbom.Spec
		})
	}
	fake.specReturnsOnCall[i] = struct {
		result1 sbom.Spec
	}{result1}
}

func (fake *FakeDocument) Tools() []sbom.Tool {
	fake.toolsMutex.Lock()
	ret, specificReturn := fake.toolsReturnsOnCall[len(fake.toolsArgsForCall)]
	fake.toolsArgsForCall = append(fake.toolsArgsForCall, struct {
	}{})
	stub := fake.ToolsStub
	fakeReturns := fake.toolsReturns
	fake.recordInvocation("Tools", []interface{}{})
	fake.toolsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDocument) ToolsCallCount() int {
	fake.toolsMutex.RLock()
	defer fake.toolsMutex.RUnlock()
	return len(fake.toolsArgsForCall)
}

func (fake *FakeDocument) ToolsCalls(stub func() []sbom.Tool) {
	fake.toolsMutex.Lock()
	defer fake.toolsMutex.Unlock()
	fake.ToolsStub = stub
}

func (fake *FakeDocument) ToolsReturns(result1 []sbom.Tool) {
	fake.toolsMutex.Lock()
	defer fake.toolsMutex.Unlock()
	fake.ToolsStub = nil
	fake.toolsReturns = struct {
		result1 []sbom.Tool
	}{result1}
}

func (fake *FakeDocument) ToolsReturnsOnCall(i int, result1 []sbom.Tool) {
	fake.toolsMutex.Lock()
	defer fake.toolsMutex.Unlock()
	fake.ToolsStub = nil
	if fake.toolsReturnsOnCall == nil {
		fake.toolsReturnsOnCall = make(map[int]struct {
			result1 []sbom.Tool
		})
	}
	fake.toolsReturnsOnCall[i] = struct {
		result1 []sbom.Tool
	}{result1}
}

func (fake *FakeDocument) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.authorsMutex.RLock()
	defer fake.authorsMutex.RUnlock()
	fake.componentsMutex.RLock()
	defer fake.componentsMutex.RUnlock()
	fake.logsMutex.RLock()
	defer fake.logsMutex.RUnlock()
	fake.relationsMutex.RLock()
	defer fake.relationsMutex.RUnlock()
	fake.specMutex.RLock()
	defer fake.specMutex.RUnlock()
	fake.toolsMutex.RLock()
	defer fake.toolsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDocument) recordInvocation(key string, args []interface{}) {
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

var _ sbom.Document = new(FakeDocument)
