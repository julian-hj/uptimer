// Code generated by counterfeiter. DO NOT EDIT.
package appLogValidatorfakes

import (
	"sync"

	"github.com/cloudfoundry/uptimer/appLogValidator"
)

type FakeAppLogValidator struct {
	IsNewerStub        func(string) (bool, error)
	isNewerMutex       sync.RWMutex
	isNewerArgsForCall []struct {
		arg1 string
	}
	isNewerReturns struct {
		result1 bool
		result2 error
	}
	isNewerReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAppLogValidator) IsNewer(arg1 string) (bool, error) {
	fake.isNewerMutex.Lock()
	ret, specificReturn := fake.isNewerReturnsOnCall[len(fake.isNewerArgsForCall)]
	fake.isNewerArgsForCall = append(fake.isNewerArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.IsNewerStub
	fakeReturns := fake.isNewerReturns
	fake.recordInvocation("IsNewer", []interface{}{arg1})
	fake.isNewerMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAppLogValidator) IsNewerCallCount() int {
	fake.isNewerMutex.RLock()
	defer fake.isNewerMutex.RUnlock()
	return len(fake.isNewerArgsForCall)
}

func (fake *FakeAppLogValidator) IsNewerCalls(stub func(string) (bool, error)) {
	fake.isNewerMutex.Lock()
	defer fake.isNewerMutex.Unlock()
	fake.IsNewerStub = stub
}

func (fake *FakeAppLogValidator) IsNewerArgsForCall(i int) string {
	fake.isNewerMutex.RLock()
	defer fake.isNewerMutex.RUnlock()
	argsForCall := fake.isNewerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAppLogValidator) IsNewerReturns(result1 bool, result2 error) {
	fake.isNewerMutex.Lock()
	defer fake.isNewerMutex.Unlock()
	fake.IsNewerStub = nil
	fake.isNewerReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeAppLogValidator) IsNewerReturnsOnCall(i int, result1 bool, result2 error) {
	fake.isNewerMutex.Lock()
	defer fake.isNewerMutex.Unlock()
	fake.IsNewerStub = nil
	if fake.isNewerReturnsOnCall == nil {
		fake.isNewerReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isNewerReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeAppLogValidator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isNewerMutex.RLock()
	defer fake.isNewerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAppLogValidator) recordInvocation(key string, args []interface{}) {
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

var _ appLogValidator.AppLogValidator = new(FakeAppLogValidator)
