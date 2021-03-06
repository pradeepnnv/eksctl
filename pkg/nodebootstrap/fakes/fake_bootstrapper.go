// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/weaveworks/eksctl/pkg/nodebootstrap"
)

type FakeBootstrapper struct {
	UserDataStub        func() (string, error)
	userDataMutex       sync.RWMutex
	userDataArgsForCall []struct {
	}
	userDataReturns struct {
		result1 string
		result2 error
	}
	userDataReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBootstrapper) UserData() (string, error) {
	fake.userDataMutex.Lock()
	ret, specificReturn := fake.userDataReturnsOnCall[len(fake.userDataArgsForCall)]
	fake.userDataArgsForCall = append(fake.userDataArgsForCall, struct {
	}{})
	stub := fake.UserDataStub
	fakeReturns := fake.userDataReturns
	fake.recordInvocation("UserData", []interface{}{})
	fake.userDataMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBootstrapper) UserDataCallCount() int {
	fake.userDataMutex.RLock()
	defer fake.userDataMutex.RUnlock()
	return len(fake.userDataArgsForCall)
}

func (fake *FakeBootstrapper) UserDataCalls(stub func() (string, error)) {
	fake.userDataMutex.Lock()
	defer fake.userDataMutex.Unlock()
	fake.UserDataStub = stub
}

func (fake *FakeBootstrapper) UserDataReturns(result1 string, result2 error) {
	fake.userDataMutex.Lock()
	defer fake.userDataMutex.Unlock()
	fake.UserDataStub = nil
	fake.userDataReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBootstrapper) UserDataReturnsOnCall(i int, result1 string, result2 error) {
	fake.userDataMutex.Lock()
	defer fake.userDataMutex.Unlock()
	fake.UserDataStub = nil
	if fake.userDataReturnsOnCall == nil {
		fake.userDataReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.userDataReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBootstrapper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.userDataMutex.RLock()
	defer fake.userDataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBootstrapper) recordInvocation(key string, args []interface{}) {
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

var _ nodebootstrap.Bootstrapper = new(FakeBootstrapper)
