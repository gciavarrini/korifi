// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"code.cloudfoundry.org/korifi/api/authorization"
	"code.cloudfoundry.org/korifi/api/handlers"
	"code.cloudfoundry.org/korifi/api/repositories"
)

type CFServiceBrokerRepository struct {
	CreateServiceBrokerStub        func(context.Context, authorization.Info, repositories.CreateServiceBrokerMessage) (repositories.ServiceBrokerResource, error)
	createServiceBrokerMutex       sync.RWMutex
	createServiceBrokerArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.CreateServiceBrokerMessage
	}
	createServiceBrokerReturns struct {
		result1 repositories.ServiceBrokerResource
		result2 error
	}
	createServiceBrokerReturnsOnCall map[int]struct {
		result1 repositories.ServiceBrokerResource
		result2 error
	}
	ListServiceBrokersStub        func(context.Context, authorization.Info) ([]repositories.ServiceBrokerResource, error)
	listServiceBrokersMutex       sync.RWMutex
	listServiceBrokersArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
	}
	listServiceBrokersReturns struct {
		result1 []repositories.ServiceBrokerResource
		result2 error
	}
	listServiceBrokersReturnsOnCall map[int]struct {
		result1 []repositories.ServiceBrokerResource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CFServiceBrokerRepository) CreateServiceBroker(arg1 context.Context, arg2 authorization.Info, arg3 repositories.CreateServiceBrokerMessage) (repositories.ServiceBrokerResource, error) {
	fake.createServiceBrokerMutex.Lock()
	ret, specificReturn := fake.createServiceBrokerReturnsOnCall[len(fake.createServiceBrokerArgsForCall)]
	fake.createServiceBrokerArgsForCall = append(fake.createServiceBrokerArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.CreateServiceBrokerMessage
	}{arg1, arg2, arg3})
	stub := fake.CreateServiceBrokerStub
	fakeReturns := fake.createServiceBrokerReturns
	fake.recordInvocation("CreateServiceBroker", []interface{}{arg1, arg2, arg3})
	fake.createServiceBrokerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFServiceBrokerRepository) CreateServiceBrokerCallCount() int {
	fake.createServiceBrokerMutex.RLock()
	defer fake.createServiceBrokerMutex.RUnlock()
	return len(fake.createServiceBrokerArgsForCall)
}

func (fake *CFServiceBrokerRepository) CreateServiceBrokerCalls(stub func(context.Context, authorization.Info, repositories.CreateServiceBrokerMessage) (repositories.ServiceBrokerResource, error)) {
	fake.createServiceBrokerMutex.Lock()
	defer fake.createServiceBrokerMutex.Unlock()
	fake.CreateServiceBrokerStub = stub
}

func (fake *CFServiceBrokerRepository) CreateServiceBrokerArgsForCall(i int) (context.Context, authorization.Info, repositories.CreateServiceBrokerMessage) {
	fake.createServiceBrokerMutex.RLock()
	defer fake.createServiceBrokerMutex.RUnlock()
	argsForCall := fake.createServiceBrokerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFServiceBrokerRepository) CreateServiceBrokerReturns(result1 repositories.ServiceBrokerResource, result2 error) {
	fake.createServiceBrokerMutex.Lock()
	defer fake.createServiceBrokerMutex.Unlock()
	fake.CreateServiceBrokerStub = nil
	fake.createServiceBrokerReturns = struct {
		result1 repositories.ServiceBrokerResource
		result2 error
	}{result1, result2}
}

func (fake *CFServiceBrokerRepository) CreateServiceBrokerReturnsOnCall(i int, result1 repositories.ServiceBrokerResource, result2 error) {
	fake.createServiceBrokerMutex.Lock()
	defer fake.createServiceBrokerMutex.Unlock()
	fake.CreateServiceBrokerStub = nil
	if fake.createServiceBrokerReturnsOnCall == nil {
		fake.createServiceBrokerReturnsOnCall = make(map[int]struct {
			result1 repositories.ServiceBrokerResource
			result2 error
		})
	}
	fake.createServiceBrokerReturnsOnCall[i] = struct {
		result1 repositories.ServiceBrokerResource
		result2 error
	}{result1, result2}
}

func (fake *CFServiceBrokerRepository) ListServiceBrokers(arg1 context.Context, arg2 authorization.Info) ([]repositories.ServiceBrokerResource, error) {
	fake.listServiceBrokersMutex.Lock()
	ret, specificReturn := fake.listServiceBrokersReturnsOnCall[len(fake.listServiceBrokersArgsForCall)]
	fake.listServiceBrokersArgsForCall = append(fake.listServiceBrokersArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
	}{arg1, arg2})
	stub := fake.ListServiceBrokersStub
	fakeReturns := fake.listServiceBrokersReturns
	fake.recordInvocation("ListServiceBrokers", []interface{}{arg1, arg2})
	fake.listServiceBrokersMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFServiceBrokerRepository) ListServiceBrokersCallCount() int {
	fake.listServiceBrokersMutex.RLock()
	defer fake.listServiceBrokersMutex.RUnlock()
	return len(fake.listServiceBrokersArgsForCall)
}

func (fake *CFServiceBrokerRepository) ListServiceBrokersCalls(stub func(context.Context, authorization.Info) ([]repositories.ServiceBrokerResource, error)) {
	fake.listServiceBrokersMutex.Lock()
	defer fake.listServiceBrokersMutex.Unlock()
	fake.ListServiceBrokersStub = stub
}

func (fake *CFServiceBrokerRepository) ListServiceBrokersArgsForCall(i int) (context.Context, authorization.Info) {
	fake.listServiceBrokersMutex.RLock()
	defer fake.listServiceBrokersMutex.RUnlock()
	argsForCall := fake.listServiceBrokersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CFServiceBrokerRepository) ListServiceBrokersReturns(result1 []repositories.ServiceBrokerResource, result2 error) {
	fake.listServiceBrokersMutex.Lock()
	defer fake.listServiceBrokersMutex.Unlock()
	fake.ListServiceBrokersStub = nil
	fake.listServiceBrokersReturns = struct {
		result1 []repositories.ServiceBrokerResource
		result2 error
	}{result1, result2}
}

func (fake *CFServiceBrokerRepository) ListServiceBrokersReturnsOnCall(i int, result1 []repositories.ServiceBrokerResource, result2 error) {
	fake.listServiceBrokersMutex.Lock()
	defer fake.listServiceBrokersMutex.Unlock()
	fake.ListServiceBrokersStub = nil
	if fake.listServiceBrokersReturnsOnCall == nil {
		fake.listServiceBrokersReturnsOnCall = make(map[int]struct {
			result1 []repositories.ServiceBrokerResource
			result2 error
		})
	}
	fake.listServiceBrokersReturnsOnCall[i] = struct {
		result1 []repositories.ServiceBrokerResource
		result2 error
	}{result1, result2}
}

func (fake *CFServiceBrokerRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createServiceBrokerMutex.RLock()
	defer fake.createServiceBrokerMutex.RUnlock()
	fake.listServiceBrokersMutex.RLock()
	defer fake.listServiceBrokersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CFServiceBrokerRepository) recordInvocation(key string, args []interface{}) {
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

var _ handlers.CFServiceBrokerRepository = new(CFServiceBrokerRepository)
