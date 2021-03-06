// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"
)

type Platform struct {
	GetDeploymentPayloadStub        func(string) ([]byte, error)
	getDeploymentPayloadMutex       sync.RWMutex
	getDeploymentPayloadArgsForCall []struct {
		arg1 string
	}
	getDeploymentPayloadReturns struct {
		result1 []byte
		result2 error
	}
	getDeploymentPayloadReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	ValidateCodePackageStub        func([]byte) error
	validateCodePackageMutex       sync.RWMutex
	validateCodePackageArgsForCall []struct {
		arg1 []byte
	}
	validateCodePackageReturns struct {
		result1 error
	}
	validateCodePackageReturnsOnCall map[int]struct {
		result1 error
	}
	ValidatePathStub        func(string) error
	validatePathMutex       sync.RWMutex
	validatePathArgsForCall []struct {
		arg1 string
	}
	validatePathReturns struct {
		result1 error
	}
	validatePathReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Platform) GetDeploymentPayload(arg1 string) ([]byte, error) {
	fake.getDeploymentPayloadMutex.Lock()
	ret, specificReturn := fake.getDeploymentPayloadReturnsOnCall[len(fake.getDeploymentPayloadArgsForCall)]
	fake.getDeploymentPayloadArgsForCall = append(fake.getDeploymentPayloadArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetDeploymentPayloadStub
	fakeReturns := fake.getDeploymentPayloadReturns
	fake.recordInvocation("GetDeploymentPayload", []interface{}{arg1})
	fake.getDeploymentPayloadMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Platform) GetDeploymentPayloadCallCount() int {
	fake.getDeploymentPayloadMutex.RLock()
	defer fake.getDeploymentPayloadMutex.RUnlock()
	return len(fake.getDeploymentPayloadArgsForCall)
}

func (fake *Platform) GetDeploymentPayloadCalls(stub func(string) ([]byte, error)) {
	fake.getDeploymentPayloadMutex.Lock()
	defer fake.getDeploymentPayloadMutex.Unlock()
	fake.GetDeploymentPayloadStub = stub
}

func (fake *Platform) GetDeploymentPayloadArgsForCall(i int) string {
	fake.getDeploymentPayloadMutex.RLock()
	defer fake.getDeploymentPayloadMutex.RUnlock()
	argsForCall := fake.getDeploymentPayloadArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Platform) GetDeploymentPayloadReturns(result1 []byte, result2 error) {
	fake.getDeploymentPayloadMutex.Lock()
	defer fake.getDeploymentPayloadMutex.Unlock()
	fake.GetDeploymentPayloadStub = nil
	fake.getDeploymentPayloadReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Platform) GetDeploymentPayloadReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getDeploymentPayloadMutex.Lock()
	defer fake.getDeploymentPayloadMutex.Unlock()
	fake.GetDeploymentPayloadStub = nil
	if fake.getDeploymentPayloadReturnsOnCall == nil {
		fake.getDeploymentPayloadReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getDeploymentPayloadReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Platform) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	stub := fake.NameStub
	fakeReturns := fake.nameReturns
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Platform) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *Platform) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *Platform) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *Platform) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *Platform) ValidateCodePackage(arg1 []byte) error {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.validateCodePackageMutex.Lock()
	ret, specificReturn := fake.validateCodePackageReturnsOnCall[len(fake.validateCodePackageArgsForCall)]
	fake.validateCodePackageArgsForCall = append(fake.validateCodePackageArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.ValidateCodePackageStub
	fakeReturns := fake.validateCodePackageReturns
	fake.recordInvocation("ValidateCodePackage", []interface{}{arg1Copy})
	fake.validateCodePackageMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Platform) ValidateCodePackageCallCount() int {
	fake.validateCodePackageMutex.RLock()
	defer fake.validateCodePackageMutex.RUnlock()
	return len(fake.validateCodePackageArgsForCall)
}

func (fake *Platform) ValidateCodePackageCalls(stub func([]byte) error) {
	fake.validateCodePackageMutex.Lock()
	defer fake.validateCodePackageMutex.Unlock()
	fake.ValidateCodePackageStub = stub
}

func (fake *Platform) ValidateCodePackageArgsForCall(i int) []byte {
	fake.validateCodePackageMutex.RLock()
	defer fake.validateCodePackageMutex.RUnlock()
	argsForCall := fake.validateCodePackageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Platform) ValidateCodePackageReturns(result1 error) {
	fake.validateCodePackageMutex.Lock()
	defer fake.validateCodePackageMutex.Unlock()
	fake.ValidateCodePackageStub = nil
	fake.validateCodePackageReturns = struct {
		result1 error
	}{result1}
}

func (fake *Platform) ValidateCodePackageReturnsOnCall(i int, result1 error) {
	fake.validateCodePackageMutex.Lock()
	defer fake.validateCodePackageMutex.Unlock()
	fake.ValidateCodePackageStub = nil
	if fake.validateCodePackageReturnsOnCall == nil {
		fake.validateCodePackageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateCodePackageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Platform) ValidatePath(arg1 string) error {
	fake.validatePathMutex.Lock()
	ret, specificReturn := fake.validatePathReturnsOnCall[len(fake.validatePathArgsForCall)]
	fake.validatePathArgsForCall = append(fake.validatePathArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ValidatePathStub
	fakeReturns := fake.validatePathReturns
	fake.recordInvocation("ValidatePath", []interface{}{arg1})
	fake.validatePathMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Platform) ValidatePathCallCount() int {
	fake.validatePathMutex.RLock()
	defer fake.validatePathMutex.RUnlock()
	return len(fake.validatePathArgsForCall)
}

func (fake *Platform) ValidatePathCalls(stub func(string) error) {
	fake.validatePathMutex.Lock()
	defer fake.validatePathMutex.Unlock()
	fake.ValidatePathStub = stub
}

func (fake *Platform) ValidatePathArgsForCall(i int) string {
	fake.validatePathMutex.RLock()
	defer fake.validatePathMutex.RUnlock()
	argsForCall := fake.validatePathArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Platform) ValidatePathReturns(result1 error) {
	fake.validatePathMutex.Lock()
	defer fake.validatePathMutex.Unlock()
	fake.ValidatePathStub = nil
	fake.validatePathReturns = struct {
		result1 error
	}{result1}
}

func (fake *Platform) ValidatePathReturnsOnCall(i int, result1 error) {
	fake.validatePathMutex.Lock()
	defer fake.validatePathMutex.Unlock()
	fake.ValidatePathStub = nil
	if fake.validatePathReturnsOnCall == nil {
		fake.validatePathReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validatePathReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Platform) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getDeploymentPayloadMutex.RLock()
	defer fake.getDeploymentPayloadMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.validateCodePackageMutex.RLock()
	defer fake.validateCodePackageMutex.RUnlock()
	fake.validatePathMutex.RLock()
	defer fake.validatePathMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Platform) recordInvocation(key string, args []interface{}) {
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
