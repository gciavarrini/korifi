// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"code.cloudfoundry.org/cf-k8s-api/apis"
	"code.cloudfoundry.org/cf-k8s-api/payloads"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ApplyManifestAction struct {
	Stub        func(context.Context, client.Client, string, payloads.SpaceManifestApply) error
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 context.Context
		arg2 client.Client
		arg3 string
		arg4 payloads.SpaceManifestApply
	}
	returns struct {
		result1 error
	}
	returnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ApplyManifestAction) Spy(arg1 context.Context, arg2 client.Client, arg3 string, arg4 payloads.SpaceManifestApply) error {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 context.Context
		arg2 client.Client
		arg3 string
		arg4 payloads.SpaceManifestApply
	}{arg1, arg2, arg3, arg4})
	stub := fake.Stub
	returns := fake.returns
	fake.recordInvocation("ApplyManifestAction", []interface{}{arg1, arg2, arg3, arg4})
	fake.mutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return returns.result1
}

func (fake *ApplyManifestAction) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *ApplyManifestAction) Calls(stub func(context.Context, client.Client, string, payloads.SpaceManifestApply) error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

func (fake *ApplyManifestAction) ArgsForCall(i int) (context.Context, client.Client, string, payloads.SpaceManifestApply) {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1, fake.argsForCall[i].arg2, fake.argsForCall[i].arg3, fake.argsForCall[i].arg4
}

func (fake *ApplyManifestAction) Returns(result1 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		result1 error
	}{result1}
}

func (fake *ApplyManifestAction) ReturnsOnCall(i int, result1 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ApplyManifestAction) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ApplyManifestAction) recordInvocation(key string, args []interface{}) {
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

var _ apis.ApplyManifestAction = new(ApplyManifestAction).Spy
