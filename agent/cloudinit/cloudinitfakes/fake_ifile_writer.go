// Code generated by counterfeiter. DO NOT EDIT.
package cloudinitfakes

import (
	"sync"

	"github.com/vmware-tanzu/cluster-api-provider-bringyourownhost/agent/cloudinit"
)

type FakeIFileWriter struct {
	MkdirIfNotExistsStub        func(string) error
	mkdirIfNotExistsMutex       sync.RWMutex
	mkdirIfNotExistsArgsForCall []struct {
		arg1 string
	}
	mkdirIfNotExistsReturns struct {
		result1 error
	}
	mkdirIfNotExistsReturnsOnCall map[int]struct {
		result1 error
	}
	WriteToFileStub        func(*cloudinit.Files) error
	writeToFileMutex       sync.RWMutex
	writeToFileArgsForCall []struct {
		arg1 *cloudinit.Files
	}
	writeToFileReturns struct {
		result1 error
	}
	writeToFileReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIFileWriter) MkdirIfNotExists(arg1 string) error {
	fake.mkdirIfNotExistsMutex.Lock()
	ret, specificReturn := fake.mkdirIfNotExistsReturnsOnCall[len(fake.mkdirIfNotExistsArgsForCall)]
	fake.mkdirIfNotExistsArgsForCall = append(fake.mkdirIfNotExistsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.MkdirIfNotExistsStub
	fakeReturns := fake.mkdirIfNotExistsReturns
	fake.recordInvocation("MkdirIfNotExists", []interface{}{arg1})
	fake.mkdirIfNotExistsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIFileWriter) MkdirIfNotExistsCallCount() int {
	fake.mkdirIfNotExistsMutex.RLock()
	defer fake.mkdirIfNotExistsMutex.RUnlock()
	return len(fake.mkdirIfNotExistsArgsForCall)
}

func (fake *FakeIFileWriter) MkdirIfNotExistsCalls(stub func(string) error) {
	fake.mkdirIfNotExistsMutex.Lock()
	defer fake.mkdirIfNotExistsMutex.Unlock()
	fake.MkdirIfNotExistsStub = stub
}

func (fake *FakeIFileWriter) MkdirIfNotExistsArgsForCall(i int) string {
	fake.mkdirIfNotExistsMutex.RLock()
	defer fake.mkdirIfNotExistsMutex.RUnlock()
	argsForCall := fake.mkdirIfNotExistsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIFileWriter) MkdirIfNotExistsReturns(result1 error) {
	fake.mkdirIfNotExistsMutex.Lock()
	defer fake.mkdirIfNotExistsMutex.Unlock()
	fake.MkdirIfNotExistsStub = nil
	fake.mkdirIfNotExistsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIFileWriter) MkdirIfNotExistsReturnsOnCall(i int, result1 error) {
	fake.mkdirIfNotExistsMutex.Lock()
	defer fake.mkdirIfNotExistsMutex.Unlock()
	fake.MkdirIfNotExistsStub = nil
	if fake.mkdirIfNotExistsReturnsOnCall == nil {
		fake.mkdirIfNotExistsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mkdirIfNotExistsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIFileWriter) WriteToFile(arg1 *cloudinit.Files) error {
	fake.writeToFileMutex.Lock()
	ret, specificReturn := fake.writeToFileReturnsOnCall[len(fake.writeToFileArgsForCall)]
	fake.writeToFileArgsForCall = append(fake.writeToFileArgsForCall, struct {
		arg1 *cloudinit.Files
	}{arg1})
	stub := fake.WriteToFileStub
	fakeReturns := fake.writeToFileReturns
	fake.recordInvocation("WriteToFile", []interface{}{arg1})
	fake.writeToFileMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIFileWriter) WriteToFileCallCount() int {
	fake.writeToFileMutex.RLock()
	defer fake.writeToFileMutex.RUnlock()
	return len(fake.writeToFileArgsForCall)
}

func (fake *FakeIFileWriter) WriteToFileCalls(stub func(*cloudinit.Files) error) {
	fake.writeToFileMutex.Lock()
	defer fake.writeToFileMutex.Unlock()
	fake.WriteToFileStub = stub
}

func (fake *FakeIFileWriter) WriteToFileArgsForCall(i int) *cloudinit.Files {
	fake.writeToFileMutex.RLock()
	defer fake.writeToFileMutex.RUnlock()
	argsForCall := fake.writeToFileArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIFileWriter) WriteToFileReturns(result1 error) {
	fake.writeToFileMutex.Lock()
	defer fake.writeToFileMutex.Unlock()
	fake.WriteToFileStub = nil
	fake.writeToFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIFileWriter) WriteToFileReturnsOnCall(i int, result1 error) {
	fake.writeToFileMutex.Lock()
	defer fake.writeToFileMutex.Unlock()
	fake.WriteToFileStub = nil
	if fake.writeToFileReturnsOnCall == nil {
		fake.writeToFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeToFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIFileWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mkdirIfNotExistsMutex.RLock()
	defer fake.mkdirIfNotExistsMutex.RUnlock()
	fake.writeToFileMutex.RLock()
	defer fake.writeToFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIFileWriter) recordInvocation(key string, args []interface{}) {
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

var _ cloudinit.IFileWriter = new(FakeIFileWriter)