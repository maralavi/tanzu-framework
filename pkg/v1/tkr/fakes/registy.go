// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkr/pkg/registry"
)

type Registry struct {
	GetFileStub        func(string, string, string) ([]byte, error)
	getFileMutex       sync.RWMutex
	getFileArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	getFileReturns struct {
		result1 []byte
		result2 error
	}
	getFileReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetFilesStub        func(string, string) (map[string][]byte, error)
	getFilesMutex       sync.RWMutex
	getFilesArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getFilesReturns struct {
		result1 map[string][]byte
		result2 error
	}
	getFilesReturnsOnCall map[int]struct {
		result1 map[string][]byte
		result2 error
	}
	ListImageTagsStub        func(string) ([]string, error)
	listImageTagsMutex       sync.RWMutex
	listImageTagsArgsForCall []struct {
		arg1 string
	}
	listImageTagsReturns struct {
		result1 []string
		result2 error
	}
	listImageTagsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Registry) GetFile(arg1 string, arg2 string, arg3 string) ([]byte, error) {
	fake.getFileMutex.Lock()
	ret, specificReturn := fake.getFileReturnsOnCall[len(fake.getFileArgsForCall)]
	fake.getFileArgsForCall = append(fake.getFileArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetFile", []interface{}{arg1, arg2, arg3})
	fake.getFileMutex.Unlock()
	if fake.GetFileStub != nil {
		return fake.GetFileStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getFileReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Registry) GetFileCallCount() int {
	fake.getFileMutex.RLock()
	defer fake.getFileMutex.RUnlock()
	return len(fake.getFileArgsForCall)
}

func (fake *Registry) GetFileCalls(stub func(string, string, string) ([]byte, error)) {
	fake.getFileMutex.Lock()
	defer fake.getFileMutex.Unlock()
	fake.GetFileStub = stub
}

func (fake *Registry) GetFileArgsForCall(i int) (string, string, string) {
	fake.getFileMutex.RLock()
	defer fake.getFileMutex.RUnlock()
	argsForCall := fake.getFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *Registry) GetFileReturns(result1 []byte, result2 error) {
	fake.getFileMutex.Lock()
	defer fake.getFileMutex.Unlock()
	fake.GetFileStub = nil
	fake.getFileReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Registry) GetFileReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getFileMutex.Lock()
	defer fake.getFileMutex.Unlock()
	fake.GetFileStub = nil
	if fake.getFileReturnsOnCall == nil {
		fake.getFileReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getFileReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Registry) GetFiles(arg1 string, arg2 string) (map[string][]byte, error) {
	fake.getFilesMutex.Lock()
	ret, specificReturn := fake.getFilesReturnsOnCall[len(fake.getFilesArgsForCall)]
	fake.getFilesArgsForCall = append(fake.getFilesArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetFiles", []interface{}{arg1, arg2})
	fake.getFilesMutex.Unlock()
	if fake.GetFilesStub != nil {
		return fake.GetFilesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getFilesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Registry) GetFilesCallCount() int {
	fake.getFilesMutex.RLock()
	defer fake.getFilesMutex.RUnlock()
	return len(fake.getFilesArgsForCall)
}

func (fake *Registry) GetFilesCalls(stub func(string, string) (map[string][]byte, error)) {
	fake.getFilesMutex.Lock()
	defer fake.getFilesMutex.Unlock()
	fake.GetFilesStub = stub
}

func (fake *Registry) GetFilesArgsForCall(i int) (string, string) {
	fake.getFilesMutex.RLock()
	defer fake.getFilesMutex.RUnlock()
	argsForCall := fake.getFilesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Registry) GetFilesReturns(result1 map[string][]byte, result2 error) {
	fake.getFilesMutex.Lock()
	defer fake.getFilesMutex.Unlock()
	fake.GetFilesStub = nil
	fake.getFilesReturns = struct {
		result1 map[string][]byte
		result2 error
	}{result1, result2}
}

func (fake *Registry) GetFilesReturnsOnCall(i int, result1 map[string][]byte, result2 error) {
	fake.getFilesMutex.Lock()
	defer fake.getFilesMutex.Unlock()
	fake.GetFilesStub = nil
	if fake.getFilesReturnsOnCall == nil {
		fake.getFilesReturnsOnCall = make(map[int]struct {
			result1 map[string][]byte
			result2 error
		})
	}
	fake.getFilesReturnsOnCall[i] = struct {
		result1 map[string][]byte
		result2 error
	}{result1, result2}
}

func (fake *Registry) ListImageTags(arg1 string) ([]string, error) {
	fake.listImageTagsMutex.Lock()
	ret, specificReturn := fake.listImageTagsReturnsOnCall[len(fake.listImageTagsArgsForCall)]
	fake.listImageTagsArgsForCall = append(fake.listImageTagsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ListImageTags", []interface{}{arg1})
	fake.listImageTagsMutex.Unlock()
	if fake.ListImageTagsStub != nil {
		return fake.ListImageTagsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listImageTagsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Registry) ListImageTagsCallCount() int {
	fake.listImageTagsMutex.RLock()
	defer fake.listImageTagsMutex.RUnlock()
	return len(fake.listImageTagsArgsForCall)
}

func (fake *Registry) ListImageTagsCalls(stub func(string) ([]string, error)) {
	fake.listImageTagsMutex.Lock()
	defer fake.listImageTagsMutex.Unlock()
	fake.ListImageTagsStub = stub
}

func (fake *Registry) ListImageTagsArgsForCall(i int) string {
	fake.listImageTagsMutex.RLock()
	defer fake.listImageTagsMutex.RUnlock()
	argsForCall := fake.listImageTagsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Registry) ListImageTagsReturns(result1 []string, result2 error) {
	fake.listImageTagsMutex.Lock()
	defer fake.listImageTagsMutex.Unlock()
	fake.ListImageTagsStub = nil
	fake.listImageTagsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *Registry) ListImageTagsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listImageTagsMutex.Lock()
	defer fake.listImageTagsMutex.Unlock()
	fake.ListImageTagsStub = nil
	if fake.listImageTagsReturnsOnCall == nil {
		fake.listImageTagsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listImageTagsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *Registry) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getFileMutex.RLock()
	defer fake.getFileMutex.RUnlock()
	fake.getFilesMutex.RLock()
	defer fake.getFilesMutex.RUnlock()
	fake.listImageTagsMutex.RLock()
	defer fake.listImageTagsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Registry) recordInvocation(key string, args []interface{}) {
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

var _ registry.Registry = new(Registry)
