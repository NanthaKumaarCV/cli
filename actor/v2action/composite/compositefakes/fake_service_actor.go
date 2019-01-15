// Code generated by counterfeiter. DO NOT EDIT.
package compositefakes

import (
	sync "sync"

	v2action "code.cloudfoundry.org/cli/actor/v2action"
	composite "code.cloudfoundry.org/cli/actor/v2action/composite"
)

type FakeServiceActor struct {
	GetServicesWithPlansForBrokerStub        func(string) (v2action.ServicesWithPlans, v2action.Warnings, error)
	getServicesWithPlansForBrokerMutex       sync.RWMutex
	getServicesWithPlansForBrokerArgsForCall []struct {
		arg1 string
	}
	getServicesWithPlansForBrokerReturns struct {
		result1 v2action.ServicesWithPlans
		result2 v2action.Warnings
		result3 error
	}
	getServicesWithPlansForBrokerReturnsOnCall map[int]struct {
		result1 v2action.ServicesWithPlans
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceActor) GetServicesWithPlansForBroker(arg1 string) (v2action.ServicesWithPlans, v2action.Warnings, error) {
	fake.getServicesWithPlansForBrokerMutex.Lock()
	ret, specificReturn := fake.getServicesWithPlansForBrokerReturnsOnCall[len(fake.getServicesWithPlansForBrokerArgsForCall)]
	fake.getServicesWithPlansForBrokerArgsForCall = append(fake.getServicesWithPlansForBrokerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetServicesWithPlansForBroker", []interface{}{arg1})
	fake.getServicesWithPlansForBrokerMutex.Unlock()
	if fake.GetServicesWithPlansForBrokerStub != nil {
		return fake.GetServicesWithPlansForBrokerStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getServicesWithPlansForBrokerReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeServiceActor) GetServicesWithPlansForBrokerCallCount() int {
	fake.getServicesWithPlansForBrokerMutex.RLock()
	defer fake.getServicesWithPlansForBrokerMutex.RUnlock()
	return len(fake.getServicesWithPlansForBrokerArgsForCall)
}

func (fake *FakeServiceActor) GetServicesWithPlansForBrokerCalls(stub func(string) (v2action.ServicesWithPlans, v2action.Warnings, error)) {
	fake.getServicesWithPlansForBrokerMutex.Lock()
	defer fake.getServicesWithPlansForBrokerMutex.Unlock()
	fake.GetServicesWithPlansForBrokerStub = stub
}

func (fake *FakeServiceActor) GetServicesWithPlansForBrokerArgsForCall(i int) string {
	fake.getServicesWithPlansForBrokerMutex.RLock()
	defer fake.getServicesWithPlansForBrokerMutex.RUnlock()
	argsForCall := fake.getServicesWithPlansForBrokerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceActor) GetServicesWithPlansForBrokerReturns(result1 v2action.ServicesWithPlans, result2 v2action.Warnings, result3 error) {
	fake.getServicesWithPlansForBrokerMutex.Lock()
	defer fake.getServicesWithPlansForBrokerMutex.Unlock()
	fake.GetServicesWithPlansForBrokerStub = nil
	fake.getServicesWithPlansForBrokerReturns = struct {
		result1 v2action.ServicesWithPlans
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServiceActor) GetServicesWithPlansForBrokerReturnsOnCall(i int, result1 v2action.ServicesWithPlans, result2 v2action.Warnings, result3 error) {
	fake.getServicesWithPlansForBrokerMutex.Lock()
	defer fake.getServicesWithPlansForBrokerMutex.Unlock()
	fake.GetServicesWithPlansForBrokerStub = nil
	if fake.getServicesWithPlansForBrokerReturnsOnCall == nil {
		fake.getServicesWithPlansForBrokerReturnsOnCall = make(map[int]struct {
			result1 v2action.ServicesWithPlans
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServicesWithPlansForBrokerReturnsOnCall[i] = struct {
		result1 v2action.ServicesWithPlans
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServiceActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getServicesWithPlansForBrokerMutex.RLock()
	defer fake.getServicesWithPlansForBrokerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceActor) recordInvocation(key string, args []interface{}) {
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

var _ composite.ServiceActor = new(FakeServiceActor)
