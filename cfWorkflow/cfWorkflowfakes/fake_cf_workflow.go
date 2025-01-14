// Code generated by counterfeiter. DO NOT EDIT.
package cfWorkflowfakes

import (
	"context"
	"sync"

	"github.com/cloudfoundry/uptimer/cfCmdGenerator"
	"github.com/cloudfoundry/uptimer/cfWorkflow"
	"github.com/cloudfoundry/uptimer/cmdStartWaiter"
)

type FakeCfWorkflow struct {
	AppUrlStub        func() string
	appUrlMutex       sync.RWMutex
	appUrlArgsForCall []struct {
	}
	appUrlReturns struct {
		result1 string
	}
	appUrlReturnsOnCall map[int]struct {
		result1 string
	}
	CreateAndBindSyslogDrainServiceStub        func(cfCmdGenerator.CfCmdGenerator, string) []cmdStartWaiter.CmdStartWaiter
	createAndBindSyslogDrainServiceMutex       sync.RWMutex
	createAndBindSyslogDrainServiceArgsForCall []struct {
		arg1 cfCmdGenerator.CfCmdGenerator
		arg2 string
	}
	createAndBindSyslogDrainServiceReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	createAndBindSyslogDrainServiceReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	DeleteStub        func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}
	deleteReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	deleteReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	MapRouteStub        func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter
	mapRouteMutex       sync.RWMutex
	mapRouteArgsForCall []struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}
	mapRouteReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	mapRouteReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	OrgStub        func() string
	orgMutex       sync.RWMutex
	orgArgsForCall []struct {
	}
	orgReturns struct {
		result1 string
	}
	orgReturnsOnCall map[int]struct {
		result1 string
	}
	PushStub        func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter
	pushMutex       sync.RWMutex
	pushArgsForCall []struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}
	pushReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	pushReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	QuotaStub        func() string
	quotaMutex       sync.RWMutex
	quotaArgsForCall []struct {
	}
	quotaReturns struct {
		result1 string
	}
	quotaReturnsOnCall map[int]struct {
		result1 string
	}
	RecentLogsStub        func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter
	recentLogsMutex       sync.RWMutex
	recentLogsArgsForCall []struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}
	recentLogsReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	recentLogsReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	SetupStub        func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter
	setupMutex       sync.RWMutex
	setupArgsForCall []struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}
	setupReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	setupReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	SpaceStub        func() string
	spaceMutex       sync.RWMutex
	spaceArgsForCall []struct {
	}
	spaceReturns struct {
		result1 string
	}
	spaceReturnsOnCall map[int]struct {
		result1 string
	}
	StreamLogsStub        func(context.Context, cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter
	streamLogsMutex       sync.RWMutex
	streamLogsArgsForCall []struct {
		arg1 context.Context
		arg2 cfCmdGenerator.CfCmdGenerator
	}
	streamLogsReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	streamLogsReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	TearDownStub        func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter
	tearDownMutex       sync.RWMutex
	tearDownArgsForCall []struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}
	tearDownReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	tearDownReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCfWorkflow) AppUrl() string {
	fake.appUrlMutex.Lock()
	ret, specificReturn := fake.appUrlReturnsOnCall[len(fake.appUrlArgsForCall)]
	fake.appUrlArgsForCall = append(fake.appUrlArgsForCall, struct {
	}{})
	stub := fake.AppUrlStub
	fakeReturns := fake.appUrlReturns
	fake.recordInvocation("AppUrl", []interface{}{})
	fake.appUrlMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCfWorkflow) AppUrlCallCount() int {
	fake.appUrlMutex.RLock()
	defer fake.appUrlMutex.RUnlock()
	return len(fake.appUrlArgsForCall)
}

func (fake *FakeCfWorkflow) AppUrlCalls(stub func() string) {
	fake.appUrlMutex.Lock()
	defer fake.appUrlMutex.Unlock()
	fake.AppUrlStub = stub
}

func (fake *FakeCfWorkflow) AppUrlReturns(result1 string) {
	fake.appUrlMutex.Lock()
	defer fake.appUrlMutex.Unlock()
	fake.AppUrlStub = nil
	fake.appUrlReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCfWorkflow) AppUrlReturnsOnCall(i int, result1 string) {
	fake.appUrlMutex.Lock()
	defer fake.appUrlMutex.Unlock()
	fake.AppUrlStub = nil
	if fake.appUrlReturnsOnCall == nil {
		fake.appUrlReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.appUrlReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCfWorkflow) CreateAndBindSyslogDrainService(arg1 cfCmdGenerator.CfCmdGenerator, arg2 string) []cmdStartWaiter.CmdStartWaiter {
	fake.createAndBindSyslogDrainServiceMutex.Lock()
	ret, specificReturn := fake.createAndBindSyslogDrainServiceReturnsOnCall[len(fake.createAndBindSyslogDrainServiceArgsForCall)]
	fake.createAndBindSyslogDrainServiceArgsForCall = append(fake.createAndBindSyslogDrainServiceArgsForCall, struct {
		arg1 cfCmdGenerator.CfCmdGenerator
		arg2 string
	}{arg1, arg2})
	stub := fake.CreateAndBindSyslogDrainServiceStub
	fakeReturns := fake.createAndBindSyslogDrainServiceReturns
	fake.recordInvocation("CreateAndBindSyslogDrainService", []interface{}{arg1, arg2})
	fake.createAndBindSyslogDrainServiceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCfWorkflow) CreateAndBindSyslogDrainServiceCallCount() int {
	fake.createAndBindSyslogDrainServiceMutex.RLock()
	defer fake.createAndBindSyslogDrainServiceMutex.RUnlock()
	return len(fake.createAndBindSyslogDrainServiceArgsForCall)
}

func (fake *FakeCfWorkflow) CreateAndBindSyslogDrainServiceCalls(stub func(cfCmdGenerator.CfCmdGenerator, string) []cmdStartWaiter.CmdStartWaiter) {
	fake.createAndBindSyslogDrainServiceMutex.Lock()
	defer fake.createAndBindSyslogDrainServiceMutex.Unlock()
	fake.CreateAndBindSyslogDrainServiceStub = stub
}

func (fake *FakeCfWorkflow) CreateAndBindSyslogDrainServiceArgsForCall(i int) (cfCmdGenerator.CfCmdGenerator, string) {
	fake.createAndBindSyslogDrainServiceMutex.RLock()
	defer fake.createAndBindSyslogDrainServiceMutex.RUnlock()
	argsForCall := fake.createAndBindSyslogDrainServiceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCfWorkflow) CreateAndBindSyslogDrainServiceReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.createAndBindSyslogDrainServiceMutex.Lock()
	defer fake.createAndBindSyslogDrainServiceMutex.Unlock()
	fake.CreateAndBindSyslogDrainServiceStub = nil
	fake.createAndBindSyslogDrainServiceReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) CreateAndBindSyslogDrainServiceReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.createAndBindSyslogDrainServiceMutex.Lock()
	defer fake.createAndBindSyslogDrainServiceMutex.Unlock()
	fake.CreateAndBindSyslogDrainServiceStub = nil
	if fake.createAndBindSyslogDrainServiceReturnsOnCall == nil {
		fake.createAndBindSyslogDrainServiceReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.createAndBindSyslogDrainServiceReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) Delete(arg1 cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}{arg1})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCfWorkflow) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeCfWorkflow) DeleteCalls(stub func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeCfWorkflow) DeleteArgsForCall(i int) cfCmdGenerator.CfCmdGenerator {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCfWorkflow) DeleteReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) DeleteReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) MapRoute(arg1 cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter {
	fake.mapRouteMutex.Lock()
	ret, specificReturn := fake.mapRouteReturnsOnCall[len(fake.mapRouteArgsForCall)]
	fake.mapRouteArgsForCall = append(fake.mapRouteArgsForCall, struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}{arg1})
	stub := fake.MapRouteStub
	fakeReturns := fake.mapRouteReturns
	fake.recordInvocation("MapRoute", []interface{}{arg1})
	fake.mapRouteMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCfWorkflow) MapRouteCallCount() int {
	fake.mapRouteMutex.RLock()
	defer fake.mapRouteMutex.RUnlock()
	return len(fake.mapRouteArgsForCall)
}

func (fake *FakeCfWorkflow) MapRouteCalls(stub func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter) {
	fake.mapRouteMutex.Lock()
	defer fake.mapRouteMutex.Unlock()
	fake.MapRouteStub = stub
}

func (fake *FakeCfWorkflow) MapRouteArgsForCall(i int) cfCmdGenerator.CfCmdGenerator {
	fake.mapRouteMutex.RLock()
	defer fake.mapRouteMutex.RUnlock()
	argsForCall := fake.mapRouteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCfWorkflow) MapRouteReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.mapRouteMutex.Lock()
	defer fake.mapRouteMutex.Unlock()
	fake.MapRouteStub = nil
	fake.mapRouteReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) MapRouteReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.mapRouteMutex.Lock()
	defer fake.mapRouteMutex.Unlock()
	fake.MapRouteStub = nil
	if fake.mapRouteReturnsOnCall == nil {
		fake.mapRouteReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.mapRouteReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) Org() string {
	fake.orgMutex.Lock()
	ret, specificReturn := fake.orgReturnsOnCall[len(fake.orgArgsForCall)]
	fake.orgArgsForCall = append(fake.orgArgsForCall, struct {
	}{})
	stub := fake.OrgStub
	fakeReturns := fake.orgReturns
	fake.recordInvocation("Org", []interface{}{})
	fake.orgMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCfWorkflow) OrgCallCount() int {
	fake.orgMutex.RLock()
	defer fake.orgMutex.RUnlock()
	return len(fake.orgArgsForCall)
}

func (fake *FakeCfWorkflow) OrgCalls(stub func() string) {
	fake.orgMutex.Lock()
	defer fake.orgMutex.Unlock()
	fake.OrgStub = stub
}

func (fake *FakeCfWorkflow) OrgReturns(result1 string) {
	fake.orgMutex.Lock()
	defer fake.orgMutex.Unlock()
	fake.OrgStub = nil
	fake.orgReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCfWorkflow) OrgReturnsOnCall(i int, result1 string) {
	fake.orgMutex.Lock()
	defer fake.orgMutex.Unlock()
	fake.OrgStub = nil
	if fake.orgReturnsOnCall == nil {
		fake.orgReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.orgReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCfWorkflow) Push(arg1 cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter {
	fake.pushMutex.Lock()
	ret, specificReturn := fake.pushReturnsOnCall[len(fake.pushArgsForCall)]
	fake.pushArgsForCall = append(fake.pushArgsForCall, struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}{arg1})
	stub := fake.PushStub
	fakeReturns := fake.pushReturns
	fake.recordInvocation("Push", []interface{}{arg1})
	fake.pushMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCfWorkflow) PushCallCount() int {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	return len(fake.pushArgsForCall)
}

func (fake *FakeCfWorkflow) PushCalls(stub func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter) {
	fake.pushMutex.Lock()
	defer fake.pushMutex.Unlock()
	fake.PushStub = stub
}

func (fake *FakeCfWorkflow) PushArgsForCall(i int) cfCmdGenerator.CfCmdGenerator {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	argsForCall := fake.pushArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCfWorkflow) PushReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.pushMutex.Lock()
	defer fake.pushMutex.Unlock()
	fake.PushStub = nil
	fake.pushReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) PushReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.pushMutex.Lock()
	defer fake.pushMutex.Unlock()
	fake.PushStub = nil
	if fake.pushReturnsOnCall == nil {
		fake.pushReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.pushReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) Quota() string {
	fake.quotaMutex.Lock()
	ret, specificReturn := fake.quotaReturnsOnCall[len(fake.quotaArgsForCall)]
	fake.quotaArgsForCall = append(fake.quotaArgsForCall, struct {
	}{})
	stub := fake.QuotaStub
	fakeReturns := fake.quotaReturns
	fake.recordInvocation("Quota", []interface{}{})
	fake.quotaMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCfWorkflow) QuotaCallCount() int {
	fake.quotaMutex.RLock()
	defer fake.quotaMutex.RUnlock()
	return len(fake.quotaArgsForCall)
}

func (fake *FakeCfWorkflow) QuotaCalls(stub func() string) {
	fake.quotaMutex.Lock()
	defer fake.quotaMutex.Unlock()
	fake.QuotaStub = stub
}

func (fake *FakeCfWorkflow) QuotaReturns(result1 string) {
	fake.quotaMutex.Lock()
	defer fake.quotaMutex.Unlock()
	fake.QuotaStub = nil
	fake.quotaReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCfWorkflow) QuotaReturnsOnCall(i int, result1 string) {
	fake.quotaMutex.Lock()
	defer fake.quotaMutex.Unlock()
	fake.QuotaStub = nil
	if fake.quotaReturnsOnCall == nil {
		fake.quotaReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.quotaReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCfWorkflow) RecentLogs(arg1 cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter {
	fake.recentLogsMutex.Lock()
	ret, specificReturn := fake.recentLogsReturnsOnCall[len(fake.recentLogsArgsForCall)]
	fake.recentLogsArgsForCall = append(fake.recentLogsArgsForCall, struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}{arg1})
	stub := fake.RecentLogsStub
	fakeReturns := fake.recentLogsReturns
	fake.recordInvocation("RecentLogs", []interface{}{arg1})
	fake.recentLogsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCfWorkflow) RecentLogsCallCount() int {
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	return len(fake.recentLogsArgsForCall)
}

func (fake *FakeCfWorkflow) RecentLogsCalls(stub func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter) {
	fake.recentLogsMutex.Lock()
	defer fake.recentLogsMutex.Unlock()
	fake.RecentLogsStub = stub
}

func (fake *FakeCfWorkflow) RecentLogsArgsForCall(i int) cfCmdGenerator.CfCmdGenerator {
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	argsForCall := fake.recentLogsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCfWorkflow) RecentLogsReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.recentLogsMutex.Lock()
	defer fake.recentLogsMutex.Unlock()
	fake.RecentLogsStub = nil
	fake.recentLogsReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) RecentLogsReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.recentLogsMutex.Lock()
	defer fake.recentLogsMutex.Unlock()
	fake.RecentLogsStub = nil
	if fake.recentLogsReturnsOnCall == nil {
		fake.recentLogsReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.recentLogsReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) Setup(arg1 cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter {
	fake.setupMutex.Lock()
	ret, specificReturn := fake.setupReturnsOnCall[len(fake.setupArgsForCall)]
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}{arg1})
	stub := fake.SetupStub
	fakeReturns := fake.setupReturns
	fake.recordInvocation("Setup", []interface{}{arg1})
	fake.setupMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCfWorkflow) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *FakeCfWorkflow) SetupCalls(stub func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = stub
}

func (fake *FakeCfWorkflow) SetupArgsForCall(i int) cfCmdGenerator.CfCmdGenerator {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	argsForCall := fake.setupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCfWorkflow) SetupReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) SetupReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	if fake.setupReturnsOnCall == nil {
		fake.setupReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.setupReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) Space() string {
	fake.spaceMutex.Lock()
	ret, specificReturn := fake.spaceReturnsOnCall[len(fake.spaceArgsForCall)]
	fake.spaceArgsForCall = append(fake.spaceArgsForCall, struct {
	}{})
	stub := fake.SpaceStub
	fakeReturns := fake.spaceReturns
	fake.recordInvocation("Space", []interface{}{})
	fake.spaceMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCfWorkflow) SpaceCallCount() int {
	fake.spaceMutex.RLock()
	defer fake.spaceMutex.RUnlock()
	return len(fake.spaceArgsForCall)
}

func (fake *FakeCfWorkflow) SpaceCalls(stub func() string) {
	fake.spaceMutex.Lock()
	defer fake.spaceMutex.Unlock()
	fake.SpaceStub = stub
}

func (fake *FakeCfWorkflow) SpaceReturns(result1 string) {
	fake.spaceMutex.Lock()
	defer fake.spaceMutex.Unlock()
	fake.SpaceStub = nil
	fake.spaceReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCfWorkflow) SpaceReturnsOnCall(i int, result1 string) {
	fake.spaceMutex.Lock()
	defer fake.spaceMutex.Unlock()
	fake.SpaceStub = nil
	if fake.spaceReturnsOnCall == nil {
		fake.spaceReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.spaceReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCfWorkflow) StreamLogs(arg1 context.Context, arg2 cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter {
	fake.streamLogsMutex.Lock()
	ret, specificReturn := fake.streamLogsReturnsOnCall[len(fake.streamLogsArgsForCall)]
	fake.streamLogsArgsForCall = append(fake.streamLogsArgsForCall, struct {
		arg1 context.Context
		arg2 cfCmdGenerator.CfCmdGenerator
	}{arg1, arg2})
	stub := fake.StreamLogsStub
	fakeReturns := fake.streamLogsReturns
	fake.recordInvocation("StreamLogs", []interface{}{arg1, arg2})
	fake.streamLogsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCfWorkflow) StreamLogsCallCount() int {
	fake.streamLogsMutex.RLock()
	defer fake.streamLogsMutex.RUnlock()
	return len(fake.streamLogsArgsForCall)
}

func (fake *FakeCfWorkflow) StreamLogsCalls(stub func(context.Context, cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter) {
	fake.streamLogsMutex.Lock()
	defer fake.streamLogsMutex.Unlock()
	fake.StreamLogsStub = stub
}

func (fake *FakeCfWorkflow) StreamLogsArgsForCall(i int) (context.Context, cfCmdGenerator.CfCmdGenerator) {
	fake.streamLogsMutex.RLock()
	defer fake.streamLogsMutex.RUnlock()
	argsForCall := fake.streamLogsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCfWorkflow) StreamLogsReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.streamLogsMutex.Lock()
	defer fake.streamLogsMutex.Unlock()
	fake.StreamLogsStub = nil
	fake.streamLogsReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) StreamLogsReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.streamLogsMutex.Lock()
	defer fake.streamLogsMutex.Unlock()
	fake.StreamLogsStub = nil
	if fake.streamLogsReturnsOnCall == nil {
		fake.streamLogsReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.streamLogsReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) TearDown(arg1 cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter {
	fake.tearDownMutex.Lock()
	ret, specificReturn := fake.tearDownReturnsOnCall[len(fake.tearDownArgsForCall)]
	fake.tearDownArgsForCall = append(fake.tearDownArgsForCall, struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}{arg1})
	stub := fake.TearDownStub
	fakeReturns := fake.tearDownReturns
	fake.recordInvocation("TearDown", []interface{}{arg1})
	fake.tearDownMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCfWorkflow) TearDownCallCount() int {
	fake.tearDownMutex.RLock()
	defer fake.tearDownMutex.RUnlock()
	return len(fake.tearDownArgsForCall)
}

func (fake *FakeCfWorkflow) TearDownCalls(stub func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter) {
	fake.tearDownMutex.Lock()
	defer fake.tearDownMutex.Unlock()
	fake.TearDownStub = stub
}

func (fake *FakeCfWorkflow) TearDownArgsForCall(i int) cfCmdGenerator.CfCmdGenerator {
	fake.tearDownMutex.RLock()
	defer fake.tearDownMutex.RUnlock()
	argsForCall := fake.tearDownArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCfWorkflow) TearDownReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.tearDownMutex.Lock()
	defer fake.tearDownMutex.Unlock()
	fake.TearDownStub = nil
	fake.tearDownReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) TearDownReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.tearDownMutex.Lock()
	defer fake.tearDownMutex.Unlock()
	fake.TearDownStub = nil
	if fake.tearDownReturnsOnCall == nil {
		fake.tearDownReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.tearDownReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appUrlMutex.RLock()
	defer fake.appUrlMutex.RUnlock()
	fake.createAndBindSyslogDrainServiceMutex.RLock()
	defer fake.createAndBindSyslogDrainServiceMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.mapRouteMutex.RLock()
	defer fake.mapRouteMutex.RUnlock()
	fake.orgMutex.RLock()
	defer fake.orgMutex.RUnlock()
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	fake.quotaMutex.RLock()
	defer fake.quotaMutex.RUnlock()
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	fake.spaceMutex.RLock()
	defer fake.spaceMutex.RUnlock()
	fake.streamLogsMutex.RLock()
	defer fake.streamLogsMutex.RUnlock()
	fake.tearDownMutex.RLock()
	defer fake.tearDownMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCfWorkflow) recordInvocation(key string, args []interface{}) {
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

var _ cfWorkflow.CfWorkflow = new(FakeCfWorkflow)
