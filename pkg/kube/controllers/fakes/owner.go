// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/cf-operator/pkg/kube/apis"
	"code.cloudfoundry.org/cf-operator/pkg/kube/controllers/extendedjob"
	v1 "k8s.io/api/core/v1"
)

type FakeOwner struct {
	ListConfigsOwnedByStub        func(context.Context, apis.Object) ([]apis.Object, error)
	listConfigsOwnedByMutex       sync.RWMutex
	listConfigsOwnedByArgsForCall []struct {
		arg1 context.Context
		arg2 apis.Object
	}
	listConfigsOwnedByReturns struct {
		result1 []apis.Object
		result2 error
	}
	listConfigsOwnedByReturnsOnCall map[int]struct {
		result1 []apis.Object
		result2 error
	}
	RemoveOwnerReferencesStub        func(context.Context, apis.Object, []apis.Object) error
	removeOwnerReferencesMutex       sync.RWMutex
	removeOwnerReferencesArgsForCall []struct {
		arg1 context.Context
		arg2 apis.Object
		arg3 []apis.Object
	}
	removeOwnerReferencesReturns struct {
		result1 error
	}
	removeOwnerReferencesReturnsOnCall map[int]struct {
		result1 error
	}
	SyncStub        func(context.Context, apis.Object, v1.PodSpec) error
	syncMutex       sync.RWMutex
	syncArgsForCall []struct {
		arg1 context.Context
		arg2 apis.Object
		arg3 v1.PodSpec
	}
	syncReturns struct {
		result1 error
	}
	syncReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOwner) ListConfigsOwnedBy(arg1 context.Context, arg2 apis.Object) ([]apis.Object, error) {
	fake.listConfigsOwnedByMutex.Lock()
	ret, specificReturn := fake.listConfigsOwnedByReturnsOnCall[len(fake.listConfigsOwnedByArgsForCall)]
	fake.listConfigsOwnedByArgsForCall = append(fake.listConfigsOwnedByArgsForCall, struct {
		arg1 context.Context
		arg2 apis.Object
	}{arg1, arg2})
	fake.recordInvocation("ListConfigsOwnedBy", []interface{}{arg1, arg2})
	fake.listConfigsOwnedByMutex.Unlock()
	if fake.ListConfigsOwnedByStub != nil {
		return fake.ListConfigsOwnedByStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listConfigsOwnedByReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOwner) ListConfigsOwnedByCallCount() int {
	fake.listConfigsOwnedByMutex.RLock()
	defer fake.listConfigsOwnedByMutex.RUnlock()
	return len(fake.listConfigsOwnedByArgsForCall)
}

func (fake *FakeOwner) ListConfigsOwnedByCalls(stub func(context.Context, apis.Object) ([]apis.Object, error)) {
	fake.listConfigsOwnedByMutex.Lock()
	defer fake.listConfigsOwnedByMutex.Unlock()
	fake.ListConfigsOwnedByStub = stub
}

func (fake *FakeOwner) ListConfigsOwnedByArgsForCall(i int) (context.Context, apis.Object) {
	fake.listConfigsOwnedByMutex.RLock()
	defer fake.listConfigsOwnedByMutex.RUnlock()
	argsForCall := fake.listConfigsOwnedByArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOwner) ListConfigsOwnedByReturns(result1 []apis.Object, result2 error) {
	fake.listConfigsOwnedByMutex.Lock()
	defer fake.listConfigsOwnedByMutex.Unlock()
	fake.ListConfigsOwnedByStub = nil
	fake.listConfigsOwnedByReturns = struct {
		result1 []apis.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeOwner) ListConfigsOwnedByReturnsOnCall(i int, result1 []apis.Object, result2 error) {
	fake.listConfigsOwnedByMutex.Lock()
	defer fake.listConfigsOwnedByMutex.Unlock()
	fake.ListConfigsOwnedByStub = nil
	if fake.listConfigsOwnedByReturnsOnCall == nil {
		fake.listConfigsOwnedByReturnsOnCall = make(map[int]struct {
			result1 []apis.Object
			result2 error
		})
	}
	fake.listConfigsOwnedByReturnsOnCall[i] = struct {
		result1 []apis.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeOwner) RemoveOwnerReferences(arg1 context.Context, arg2 apis.Object, arg3 []apis.Object) error {
	var arg3Copy []apis.Object
	if arg3 != nil {
		arg3Copy = make([]apis.Object, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.removeOwnerReferencesMutex.Lock()
	ret, specificReturn := fake.removeOwnerReferencesReturnsOnCall[len(fake.removeOwnerReferencesArgsForCall)]
	fake.removeOwnerReferencesArgsForCall = append(fake.removeOwnerReferencesArgsForCall, struct {
		arg1 context.Context
		arg2 apis.Object
		arg3 []apis.Object
	}{arg1, arg2, arg3Copy})
	fake.recordInvocation("RemoveOwnerReferences", []interface{}{arg1, arg2, arg3Copy})
	fake.removeOwnerReferencesMutex.Unlock()
	if fake.RemoveOwnerReferencesStub != nil {
		return fake.RemoveOwnerReferencesStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.removeOwnerReferencesReturns
	return fakeReturns.result1
}

func (fake *FakeOwner) RemoveOwnerReferencesCallCount() int {
	fake.removeOwnerReferencesMutex.RLock()
	defer fake.removeOwnerReferencesMutex.RUnlock()
	return len(fake.removeOwnerReferencesArgsForCall)
}

func (fake *FakeOwner) RemoveOwnerReferencesCalls(stub func(context.Context, apis.Object, []apis.Object) error) {
	fake.removeOwnerReferencesMutex.Lock()
	defer fake.removeOwnerReferencesMutex.Unlock()
	fake.RemoveOwnerReferencesStub = stub
}

func (fake *FakeOwner) RemoveOwnerReferencesArgsForCall(i int) (context.Context, apis.Object, []apis.Object) {
	fake.removeOwnerReferencesMutex.RLock()
	defer fake.removeOwnerReferencesMutex.RUnlock()
	argsForCall := fake.removeOwnerReferencesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeOwner) RemoveOwnerReferencesReturns(result1 error) {
	fake.removeOwnerReferencesMutex.Lock()
	defer fake.removeOwnerReferencesMutex.Unlock()
	fake.RemoveOwnerReferencesStub = nil
	fake.removeOwnerReferencesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOwner) RemoveOwnerReferencesReturnsOnCall(i int, result1 error) {
	fake.removeOwnerReferencesMutex.Lock()
	defer fake.removeOwnerReferencesMutex.Unlock()
	fake.RemoveOwnerReferencesStub = nil
	if fake.removeOwnerReferencesReturnsOnCall == nil {
		fake.removeOwnerReferencesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeOwnerReferencesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOwner) Sync(arg1 context.Context, arg2 apis.Object, arg3 v1.PodSpec) error {
	fake.syncMutex.Lock()
	ret, specificReturn := fake.syncReturnsOnCall[len(fake.syncArgsForCall)]
	fake.syncArgsForCall = append(fake.syncArgsForCall, struct {
		arg1 context.Context
		arg2 apis.Object
		arg3 v1.PodSpec
	}{arg1, arg2, arg3})
	fake.recordInvocation("Sync", []interface{}{arg1, arg2, arg3})
	fake.syncMutex.Unlock()
	if fake.SyncStub != nil {
		return fake.SyncStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.syncReturns
	return fakeReturns.result1
}

func (fake *FakeOwner) SyncCallCount() int {
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	return len(fake.syncArgsForCall)
}

func (fake *FakeOwner) SyncCalls(stub func(context.Context, apis.Object, v1.PodSpec) error) {
	fake.syncMutex.Lock()
	defer fake.syncMutex.Unlock()
	fake.SyncStub = stub
}

func (fake *FakeOwner) SyncArgsForCall(i int) (context.Context, apis.Object, v1.PodSpec) {
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	argsForCall := fake.syncArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeOwner) SyncReturns(result1 error) {
	fake.syncMutex.Lock()
	defer fake.syncMutex.Unlock()
	fake.SyncStub = nil
	fake.syncReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOwner) SyncReturnsOnCall(i int, result1 error) {
	fake.syncMutex.Lock()
	defer fake.syncMutex.Unlock()
	fake.SyncStub = nil
	if fake.syncReturnsOnCall == nil {
		fake.syncReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.syncReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOwner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listConfigsOwnedByMutex.RLock()
	defer fake.listConfigsOwnedByMutex.RUnlock()
	fake.removeOwnerReferencesMutex.RLock()
	defer fake.removeOwnerReferencesMutex.RUnlock()
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOwner) recordInvocation(key string, args []interface{}) {
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

var _ extendedjob.Owner = new(FakeOwner)
