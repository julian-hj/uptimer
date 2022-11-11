// Code generated by counterfeiter. DO NOT EDIT.
package measurementfakes

import (
	"sync"
	"time"

	"github.com/cloudfoundry/uptimer/measurement"
)

type FakeResultSet struct {
	FailedStub        func() int
	failedMutex       sync.RWMutex
	failedArgsForCall []struct {
	}
	failedReturns struct {
		result1 int
	}
	failedReturnsOnCall map[int]struct {
		result1 int
	}
	RecordFailureStub        func()
	recordFailureMutex       sync.RWMutex
	recordFailureArgsForCall []struct {
	}
	RecordSuccessStub        func()
	recordSuccessMutex       sync.RWMutex
	recordSuccessArgsForCall []struct {
	}
	SuccessesSinceLastFailureStub        func() (int, time.Time)
	successesSinceLastFailureMutex       sync.RWMutex
	successesSinceLastFailureArgsForCall []struct {
	}
	successesSinceLastFailureReturns struct {
		result1 int
		result2 time.Time
	}
	successesSinceLastFailureReturnsOnCall map[int]struct {
		result1 int
		result2 time.Time
	}
	SuccessfulStub        func() int
	successfulMutex       sync.RWMutex
	successfulArgsForCall []struct {
	}
	successfulReturns struct {
		result1 int
	}
	successfulReturnsOnCall map[int]struct {
		result1 int
	}
	TotalStub        func() int
	totalMutex       sync.RWMutex
	totalArgsForCall []struct {
	}
	totalReturns struct {
		result1 int
	}
	totalReturnsOnCall map[int]struct {
		result1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResultSet) Failed() int {
	fake.failedMutex.Lock()
	ret, specificReturn := fake.failedReturnsOnCall[len(fake.failedArgsForCall)]
	fake.failedArgsForCall = append(fake.failedArgsForCall, struct {
	}{})
	stub := fake.FailedStub
	fakeReturns := fake.failedReturns
	fake.recordInvocation("Failed", []interface{}{})
	fake.failedMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeResultSet) FailedCallCount() int {
	fake.failedMutex.RLock()
	defer fake.failedMutex.RUnlock()
	return len(fake.failedArgsForCall)
}

func (fake *FakeResultSet) FailedCalls(stub func() int) {
	fake.failedMutex.Lock()
	defer fake.failedMutex.Unlock()
	fake.FailedStub = stub
}

func (fake *FakeResultSet) FailedReturns(result1 int) {
	fake.failedMutex.Lock()
	defer fake.failedMutex.Unlock()
	fake.FailedStub = nil
	fake.failedReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeResultSet) FailedReturnsOnCall(i int, result1 int) {
	fake.failedMutex.Lock()
	defer fake.failedMutex.Unlock()
	fake.FailedStub = nil
	if fake.failedReturnsOnCall == nil {
		fake.failedReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.failedReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeResultSet) RecordFailure() {
	fake.recordFailureMutex.Lock()
	fake.recordFailureArgsForCall = append(fake.recordFailureArgsForCall, struct {
	}{})
	stub := fake.RecordFailureStub
	fake.recordInvocation("RecordFailure", []interface{}{})
	fake.recordFailureMutex.Unlock()
	if stub != nil {
		fake.RecordFailureStub()
	}
}

func (fake *FakeResultSet) RecordFailureCallCount() int {
	fake.recordFailureMutex.RLock()
	defer fake.recordFailureMutex.RUnlock()
	return len(fake.recordFailureArgsForCall)
}

func (fake *FakeResultSet) RecordFailureCalls(stub func()) {
	fake.recordFailureMutex.Lock()
	defer fake.recordFailureMutex.Unlock()
	fake.RecordFailureStub = stub
}

func (fake *FakeResultSet) RecordSuccess() {
	fake.recordSuccessMutex.Lock()
	fake.recordSuccessArgsForCall = append(fake.recordSuccessArgsForCall, struct {
	}{})
	stub := fake.RecordSuccessStub
	fake.recordInvocation("RecordSuccess", []interface{}{})
	fake.recordSuccessMutex.Unlock()
	if stub != nil {
		fake.RecordSuccessStub()
	}
}

func (fake *FakeResultSet) RecordSuccessCallCount() int {
	fake.recordSuccessMutex.RLock()
	defer fake.recordSuccessMutex.RUnlock()
	return len(fake.recordSuccessArgsForCall)
}

func (fake *FakeResultSet) RecordSuccessCalls(stub func()) {
	fake.recordSuccessMutex.Lock()
	defer fake.recordSuccessMutex.Unlock()
	fake.RecordSuccessStub = stub
}

func (fake *FakeResultSet) SuccessesSinceLastFailure() (int, time.Time) {
	fake.successesSinceLastFailureMutex.Lock()
	ret, specificReturn := fake.successesSinceLastFailureReturnsOnCall[len(fake.successesSinceLastFailureArgsForCall)]
	fake.successesSinceLastFailureArgsForCall = append(fake.successesSinceLastFailureArgsForCall, struct {
	}{})
	stub := fake.SuccessesSinceLastFailureStub
	fakeReturns := fake.successesSinceLastFailureReturns
	fake.recordInvocation("SuccessesSinceLastFailure", []interface{}{})
	fake.successesSinceLastFailureMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResultSet) SuccessesSinceLastFailureCallCount() int {
	fake.successesSinceLastFailureMutex.RLock()
	defer fake.successesSinceLastFailureMutex.RUnlock()
	return len(fake.successesSinceLastFailureArgsForCall)
}

func (fake *FakeResultSet) SuccessesSinceLastFailureCalls(stub func() (int, time.Time)) {
	fake.successesSinceLastFailureMutex.Lock()
	defer fake.successesSinceLastFailureMutex.Unlock()
	fake.SuccessesSinceLastFailureStub = stub
}

func (fake *FakeResultSet) SuccessesSinceLastFailureReturns(result1 int, result2 time.Time) {
	fake.successesSinceLastFailureMutex.Lock()
	defer fake.successesSinceLastFailureMutex.Unlock()
	fake.SuccessesSinceLastFailureStub = nil
	fake.successesSinceLastFailureReturns = struct {
		result1 int
		result2 time.Time
	}{result1, result2}
}

func (fake *FakeResultSet) SuccessesSinceLastFailureReturnsOnCall(i int, result1 int, result2 time.Time) {
	fake.successesSinceLastFailureMutex.Lock()
	defer fake.successesSinceLastFailureMutex.Unlock()
	fake.SuccessesSinceLastFailureStub = nil
	if fake.successesSinceLastFailureReturnsOnCall == nil {
		fake.successesSinceLastFailureReturnsOnCall = make(map[int]struct {
			result1 int
			result2 time.Time
		})
	}
	fake.successesSinceLastFailureReturnsOnCall[i] = struct {
		result1 int
		result2 time.Time
	}{result1, result2}
}

func (fake *FakeResultSet) Successful() int {
	fake.successfulMutex.Lock()
	ret, specificReturn := fake.successfulReturnsOnCall[len(fake.successfulArgsForCall)]
	fake.successfulArgsForCall = append(fake.successfulArgsForCall, struct {
	}{})
	stub := fake.SuccessfulStub
	fakeReturns := fake.successfulReturns
	fake.recordInvocation("Successful", []interface{}{})
	fake.successfulMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeResultSet) SuccessfulCallCount() int {
	fake.successfulMutex.RLock()
	defer fake.successfulMutex.RUnlock()
	return len(fake.successfulArgsForCall)
}

func (fake *FakeResultSet) SuccessfulCalls(stub func() int) {
	fake.successfulMutex.Lock()
	defer fake.successfulMutex.Unlock()
	fake.SuccessfulStub = stub
}

func (fake *FakeResultSet) SuccessfulReturns(result1 int) {
	fake.successfulMutex.Lock()
	defer fake.successfulMutex.Unlock()
	fake.SuccessfulStub = nil
	fake.successfulReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeResultSet) SuccessfulReturnsOnCall(i int, result1 int) {
	fake.successfulMutex.Lock()
	defer fake.successfulMutex.Unlock()
	fake.SuccessfulStub = nil
	if fake.successfulReturnsOnCall == nil {
		fake.successfulReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.successfulReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeResultSet) Total() int {
	fake.totalMutex.Lock()
	ret, specificReturn := fake.totalReturnsOnCall[len(fake.totalArgsForCall)]
	fake.totalArgsForCall = append(fake.totalArgsForCall, struct {
	}{})
	stub := fake.TotalStub
	fakeReturns := fake.totalReturns
	fake.recordInvocation("Total", []interface{}{})
	fake.totalMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeResultSet) TotalCallCount() int {
	fake.totalMutex.RLock()
	defer fake.totalMutex.RUnlock()
	return len(fake.totalArgsForCall)
}

func (fake *FakeResultSet) TotalCalls(stub func() int) {
	fake.totalMutex.Lock()
	defer fake.totalMutex.Unlock()
	fake.TotalStub = stub
}

func (fake *FakeResultSet) TotalReturns(result1 int) {
	fake.totalMutex.Lock()
	defer fake.totalMutex.Unlock()
	fake.TotalStub = nil
	fake.totalReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeResultSet) TotalReturnsOnCall(i int, result1 int) {
	fake.totalMutex.Lock()
	defer fake.totalMutex.Unlock()
	fake.TotalStub = nil
	if fake.totalReturnsOnCall == nil {
		fake.totalReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.totalReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeResultSet) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.failedMutex.RLock()
	defer fake.failedMutex.RUnlock()
	fake.recordFailureMutex.RLock()
	defer fake.recordFailureMutex.RUnlock()
	fake.recordSuccessMutex.RLock()
	defer fake.recordSuccessMutex.RUnlock()
	fake.successesSinceLastFailureMutex.RLock()
	defer fake.successesSinceLastFailureMutex.RUnlock()
	fake.successfulMutex.RLock()
	defer fake.successfulMutex.RUnlock()
	fake.totalMutex.RLock()
	defer fake.totalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResultSet) recordInvocation(key string, args []interface{}) {
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

var _ measurement.ResultSet = new(FakeResultSet)
