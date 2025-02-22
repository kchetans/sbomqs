// Code generated by counterfeiter. DO NOT EDIT.
package sbomfakes

import (
	"sync"

	"github.com/kchetans/sbomqs/pkg/sbom"
)

type FakeChecksum struct {
	AlgoStub        func() string
	algoMutex       sync.RWMutex
	algoArgsForCall []struct {
	}
	algoReturns struct {
		result1 string
	}
	algoReturnsOnCall map[int]struct {
		result1 string
	}
	ContentStub        func() string
	contentMutex       sync.RWMutex
	contentArgsForCall []struct {
	}
	contentReturns struct {
		result1 string
	}
	contentReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeChecksum) Algo() string {
	fake.algoMutex.Lock()
	ret, specificReturn := fake.algoReturnsOnCall[len(fake.algoArgsForCall)]
	fake.algoArgsForCall = append(fake.algoArgsForCall, struct {
	}{})
	stub := fake.AlgoStub
	fakeReturns := fake.algoReturns
	fake.recordInvocation("Algo", []interface{}{})
	fake.algoMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeChecksum) AlgoCallCount() int {
	fake.algoMutex.RLock()
	defer fake.algoMutex.RUnlock()
	return len(fake.algoArgsForCall)
}

func (fake *FakeChecksum) AlgoCalls(stub func() string) {
	fake.algoMutex.Lock()
	defer fake.algoMutex.Unlock()
	fake.AlgoStub = stub
}

func (fake *FakeChecksum) AlgoReturns(result1 string) {
	fake.algoMutex.Lock()
	defer fake.algoMutex.Unlock()
	fake.AlgoStub = nil
	fake.algoReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeChecksum) AlgoReturnsOnCall(i int, result1 string) {
	fake.algoMutex.Lock()
	defer fake.algoMutex.Unlock()
	fake.AlgoStub = nil
	if fake.algoReturnsOnCall == nil {
		fake.algoReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.algoReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeChecksum) Content() string {
	fake.contentMutex.Lock()
	ret, specificReturn := fake.contentReturnsOnCall[len(fake.contentArgsForCall)]
	fake.contentArgsForCall = append(fake.contentArgsForCall, struct {
	}{})
	stub := fake.ContentStub
	fakeReturns := fake.contentReturns
	fake.recordInvocation("Content", []interface{}{})
	fake.contentMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeChecksum) ContentCallCount() int {
	fake.contentMutex.RLock()
	defer fake.contentMutex.RUnlock()
	return len(fake.contentArgsForCall)
}

func (fake *FakeChecksum) ContentCalls(stub func() string) {
	fake.contentMutex.Lock()
	defer fake.contentMutex.Unlock()
	fake.ContentStub = stub
}

func (fake *FakeChecksum) ContentReturns(result1 string) {
	fake.contentMutex.Lock()
	defer fake.contentMutex.Unlock()
	fake.ContentStub = nil
	fake.contentReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeChecksum) ContentReturnsOnCall(i int, result1 string) {
	fake.contentMutex.Lock()
	defer fake.contentMutex.Unlock()
	fake.ContentStub = nil
	if fake.contentReturnsOnCall == nil {
		fake.contentReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.contentReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeChecksum) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.algoMutex.RLock()
	defer fake.algoMutex.RUnlock()
	fake.contentMutex.RLock()
	defer fake.contentMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeChecksum) recordInvocation(key string, args []interface{}) {
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

var _ sbom.Checksum = new(FakeChecksum)
