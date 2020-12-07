// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"
)

type ACLProvider struct {
	CheckACLStub        func(string, string, interface{}) error
	checkACLMutex       sync.RWMutex
	checkACLArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 interface{}
	}
	checkACLReturns struct {
		result1 error
	}
	checkACLReturnsOnCall map[int]struct {
		result1 error
	}
	CheckACLNoChannelStub        func(string, interface{}) error
	checkACLNoChannelMutex       sync.RWMutex
	checkACLNoChannelArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	checkACLNoChannelReturns struct {
		result1 error
	}
	checkACLNoChannelReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ACLProvider) CheckACL(arg1 string, arg2 string, arg3 interface{}) error {
	fake.checkACLMutex.Lock()
	ret, specificReturn := fake.checkACLReturnsOnCall[len(fake.checkACLArgsForCall)]
	fake.checkACLArgsForCall = append(fake.checkACLArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 interface{}
	}{arg1, arg2, arg3})
	stub := fake.CheckACLStub
	fakeReturns := fake.checkACLReturns
	fake.recordInvocation("CheckACL", []interface{}{arg1, arg2, arg3})
	fake.checkACLMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ACLProvider) CheckACLCallCount() int {
	fake.checkACLMutex.RLock()
	defer fake.checkACLMutex.RUnlock()
	return len(fake.checkACLArgsForCall)
}

func (fake *ACLProvider) CheckACLCalls(stub func(string, string, interface{}) error) {
	fake.checkACLMutex.Lock()
	defer fake.checkACLMutex.Unlock()
	fake.CheckACLStub = stub
}

func (fake *ACLProvider) CheckACLArgsForCall(i int) (string, string, interface{}) {
	fake.checkACLMutex.RLock()
	defer fake.checkACLMutex.RUnlock()
	argsForCall := fake.checkACLArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ACLProvider) CheckACLReturns(result1 error) {
	fake.checkACLMutex.Lock()
	defer fake.checkACLMutex.Unlock()
	fake.CheckACLStub = nil
	fake.checkACLReturns = struct {
		result1 error
	}{result1}
}

func (fake *ACLProvider) CheckACLReturnsOnCall(i int, result1 error) {
	fake.checkACLMutex.Lock()
	defer fake.checkACLMutex.Unlock()
	fake.CheckACLStub = nil
	if fake.checkACLReturnsOnCall == nil {
		fake.checkACLReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkACLReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ACLProvider) CheckACLNoChannel(arg1 string, arg2 interface{}) error {
	fake.checkACLNoChannelMutex.Lock()
	ret, specificReturn := fake.checkACLNoChannelReturnsOnCall[len(fake.checkACLNoChannelArgsForCall)]
	fake.checkACLNoChannelArgsForCall = append(fake.checkACLNoChannelArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	stub := fake.CheckACLNoChannelStub
	fakeReturns := fake.checkACLNoChannelReturns
	fake.recordInvocation("CheckACLNoChannel", []interface{}{arg1, arg2})
	fake.checkACLNoChannelMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ACLProvider) CheckACLNoChannelCallCount() int {
	fake.checkACLNoChannelMutex.RLock()
	defer fake.checkACLNoChannelMutex.RUnlock()
	return len(fake.checkACLNoChannelArgsForCall)
}

func (fake *ACLProvider) CheckACLNoChannelCalls(stub func(string, interface{}) error) {
	fake.checkACLNoChannelMutex.Lock()
	defer fake.checkACLNoChannelMutex.Unlock()
	fake.CheckACLNoChannelStub = stub
}

func (fake *ACLProvider) CheckACLNoChannelArgsForCall(i int) (string, interface{}) {
	fake.checkACLNoChannelMutex.RLock()
	defer fake.checkACLNoChannelMutex.RUnlock()
	argsForCall := fake.checkACLNoChannelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ACLProvider) CheckACLNoChannelReturns(result1 error) {
	fake.checkACLNoChannelMutex.Lock()
	defer fake.checkACLNoChannelMutex.Unlock()
	fake.CheckACLNoChannelStub = nil
	fake.checkACLNoChannelReturns = struct {
		result1 error
	}{result1}
}

func (fake *ACLProvider) CheckACLNoChannelReturnsOnCall(i int, result1 error) {
	fake.checkACLNoChannelMutex.Lock()
	defer fake.checkACLNoChannelMutex.Unlock()
	fake.CheckACLNoChannelStub = nil
	if fake.checkACLNoChannelReturnsOnCall == nil {
		fake.checkACLNoChannelReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkACLNoChannelReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ACLProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkACLMutex.RLock()
	defer fake.checkACLMutex.RUnlock()
	fake.checkACLNoChannelMutex.RLock()
	defer fake.checkACLNoChannelMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ACLProvider) recordInvocation(key string, args []interface{}) {
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
