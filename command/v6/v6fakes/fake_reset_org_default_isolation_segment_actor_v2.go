// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v2action "code.cloudfoundry.org/cli/actor/v2action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeResetOrgDefaultIsolationSegmentActorV2 struct {
	GetOrganizationByNameStub        func(string) (v2action.Organization, v2action.Warnings, error)
	getOrganizationByNameMutex       sync.RWMutex
	getOrganizationByNameArgsForCall []struct {
		arg1 string
	}
	getOrganizationByNameReturns struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}
	getOrganizationByNameReturnsOnCall map[int]struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResetOrgDefaultIsolationSegmentActorV2) GetOrganizationByName(arg1 string) (v2action.Organization, v2action.Warnings, error) {
	fake.getOrganizationByNameMutex.Lock()
	ret, specificReturn := fake.getOrganizationByNameReturnsOnCall[len(fake.getOrganizationByNameArgsForCall)]
	fake.getOrganizationByNameArgsForCall = append(fake.getOrganizationByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrganizationByName", []interface{}{arg1})
	fake.getOrganizationByNameMutex.Unlock()
	if fake.GetOrganizationByNameStub != nil {
		return fake.GetOrganizationByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationByNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeResetOrgDefaultIsolationSegmentActorV2) GetOrganizationByNameCallCount() int {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return len(fake.getOrganizationByNameArgsForCall)
}

func (fake *FakeResetOrgDefaultIsolationSegmentActorV2) GetOrganizationByNameCalls(stub func(string) (v2action.Organization, v2action.Warnings, error)) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = stub
}

func (fake *FakeResetOrgDefaultIsolationSegmentActorV2) GetOrganizationByNameArgsForCall(i int) string {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	argsForCall := fake.getOrganizationByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeResetOrgDefaultIsolationSegmentActorV2) GetOrganizationByNameReturns(result1 v2action.Organization, result2 v2action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	fake.getOrganizationByNameReturns = struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResetOrgDefaultIsolationSegmentActorV2) GetOrganizationByNameReturnsOnCall(i int, result1 v2action.Organization, result2 v2action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	if fake.getOrganizationByNameReturnsOnCall == nil {
		fake.getOrganizationByNameReturnsOnCall = make(map[int]struct {
			result1 v2action.Organization
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getOrganizationByNameReturnsOnCall[i] = struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResetOrgDefaultIsolationSegmentActorV2) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResetOrgDefaultIsolationSegmentActorV2) recordInvocation(key string, args []interface{}) {
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

var _ v6.ResetOrgDefaultIsolationSegmentActorV2 = new(FakeResetOrgDefaultIsolationSegmentActorV2)