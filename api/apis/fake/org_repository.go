// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"code.cloudfoundry.org/cf-k8s-controllers/api/apis"
	"code.cloudfoundry.org/cf-k8s-controllers/api/authorization"
	"code.cloudfoundry.org/cf-k8s-controllers/api/repositories"
)

type OrgRepository struct {
	CreateOrgStub        func(context.Context, authorization.Info, repositories.CreateOrgMessage) (repositories.OrgRecord, error)
	createOrgMutex       sync.RWMutex
	createOrgArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.CreateOrgMessage
	}
	createOrgReturns struct {
		result1 repositories.OrgRecord
		result2 error
	}
	createOrgReturnsOnCall map[int]struct {
		result1 repositories.OrgRecord
		result2 error
	}
	ListOrgsStub        func(context.Context, authorization.Info, []string) ([]repositories.OrgRecord, error)
	listOrgsMutex       sync.RWMutex
	listOrgsArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 []string
	}
	listOrgsReturns struct {
		result1 []repositories.OrgRecord
		result2 error
	}
	listOrgsReturnsOnCall map[int]struct {
		result1 []repositories.OrgRecord
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *OrgRepository) CreateOrg(arg1 context.Context, arg2 authorization.Info, arg3 repositories.CreateOrgMessage) (repositories.OrgRecord, error) {
	fake.createOrgMutex.Lock()
	ret, specificReturn := fake.createOrgReturnsOnCall[len(fake.createOrgArgsForCall)]
	fake.createOrgArgsForCall = append(fake.createOrgArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.CreateOrgMessage
	}{arg1, arg2, arg3})
	stub := fake.CreateOrgStub
	fakeReturns := fake.createOrgReturns
	fake.recordInvocation("CreateOrg", []interface{}{arg1, arg2, arg3})
	fake.createOrgMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *OrgRepository) CreateOrgCallCount() int {
	fake.createOrgMutex.RLock()
	defer fake.createOrgMutex.RUnlock()
	return len(fake.createOrgArgsForCall)
}

func (fake *OrgRepository) CreateOrgCalls(stub func(context.Context, authorization.Info, repositories.CreateOrgMessage) (repositories.OrgRecord, error)) {
	fake.createOrgMutex.Lock()
	defer fake.createOrgMutex.Unlock()
	fake.CreateOrgStub = stub
}

func (fake *OrgRepository) CreateOrgArgsForCall(i int) (context.Context, authorization.Info, repositories.CreateOrgMessage) {
	fake.createOrgMutex.RLock()
	defer fake.createOrgMutex.RUnlock()
	argsForCall := fake.createOrgArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *OrgRepository) CreateOrgReturns(result1 repositories.OrgRecord, result2 error) {
	fake.createOrgMutex.Lock()
	defer fake.createOrgMutex.Unlock()
	fake.CreateOrgStub = nil
	fake.createOrgReturns = struct {
		result1 repositories.OrgRecord
		result2 error
	}{result1, result2}
}

func (fake *OrgRepository) CreateOrgReturnsOnCall(i int, result1 repositories.OrgRecord, result2 error) {
	fake.createOrgMutex.Lock()
	defer fake.createOrgMutex.Unlock()
	fake.CreateOrgStub = nil
	if fake.createOrgReturnsOnCall == nil {
		fake.createOrgReturnsOnCall = make(map[int]struct {
			result1 repositories.OrgRecord
			result2 error
		})
	}
	fake.createOrgReturnsOnCall[i] = struct {
		result1 repositories.OrgRecord
		result2 error
	}{result1, result2}
}

func (fake *OrgRepository) ListOrgs(arg1 context.Context, arg2 authorization.Info, arg3 []string) ([]repositories.OrgRecord, error) {
	var arg3Copy []string
	if arg3 != nil {
		arg3Copy = make([]string, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.listOrgsMutex.Lock()
	ret, specificReturn := fake.listOrgsReturnsOnCall[len(fake.listOrgsArgsForCall)]
	fake.listOrgsArgsForCall = append(fake.listOrgsArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 []string
	}{arg1, arg2, arg3Copy})
	stub := fake.ListOrgsStub
	fakeReturns := fake.listOrgsReturns
	fake.recordInvocation("ListOrgs", []interface{}{arg1, arg2, arg3Copy})
	fake.listOrgsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *OrgRepository) ListOrgsCallCount() int {
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	return len(fake.listOrgsArgsForCall)
}

func (fake *OrgRepository) ListOrgsCalls(stub func(context.Context, authorization.Info, []string) ([]repositories.OrgRecord, error)) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = stub
}

func (fake *OrgRepository) ListOrgsArgsForCall(i int) (context.Context, authorization.Info, []string) {
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	argsForCall := fake.listOrgsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *OrgRepository) ListOrgsReturns(result1 []repositories.OrgRecord, result2 error) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = nil
	fake.listOrgsReturns = struct {
		result1 []repositories.OrgRecord
		result2 error
	}{result1, result2}
}

func (fake *OrgRepository) ListOrgsReturnsOnCall(i int, result1 []repositories.OrgRecord, result2 error) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = nil
	if fake.listOrgsReturnsOnCall == nil {
		fake.listOrgsReturnsOnCall = make(map[int]struct {
			result1 []repositories.OrgRecord
			result2 error
		})
	}
	fake.listOrgsReturnsOnCall[i] = struct {
		result1 []repositories.OrgRecord
		result2 error
	}{result1, result2}
}

func (fake *OrgRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createOrgMutex.RLock()
	defer fake.createOrgMutex.RUnlock()
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *OrgRepository) recordInvocation(key string, args []interface{}) {
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

var _ apis.OrgRepository = new(OrgRepository)
