// Code generated by counterfeiter. DO NOT EDIT.
package v2actionfakes

import (
	sync "sync"

	v2action "code.cloudfoundry.org/cli/actor/v2action"
	uaa "code.cloudfoundry.org/cli/api/uaa"
	constant "code.cloudfoundry.org/cli/api/uaa/constant"
)

type FakeUAAClient struct {
	APIVersionStub        func() string
	aPIVersionMutex       sync.RWMutex
	aPIVersionArgsForCall []struct {
	}
	aPIVersionReturns struct {
		result1 string
	}
	aPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	AuthenticateStub        func(string, string, string, constant.GrantType) (string, string, error)
	authenticateMutex       sync.RWMutex
	authenticateArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 constant.GrantType
	}
	authenticateReturns struct {
		result1 string
		result2 string
		result3 error
	}
	authenticateReturnsOnCall map[int]struct {
		result1 string
		result2 string
		result3 error
	}
	CreateUserStub        func(string, string, string) (uaa.User, error)
	createUserMutex       sync.RWMutex
	createUserArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	createUserReturns struct {
		result1 uaa.User
		result2 error
	}
	createUserReturnsOnCall map[int]struct {
		result1 uaa.User
		result2 error
	}
	GetSSHPasscodeStub        func(string, string) (string, error)
	getSSHPasscodeMutex       sync.RWMutex
	getSSHPasscodeArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSSHPasscodeReturns struct {
		result1 string
		result2 error
	}
	getSSHPasscodeReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	RefreshAccessTokenStub        func(string) (uaa.RefreshedTokens, error)
	refreshAccessTokenMutex       sync.RWMutex
	refreshAccessTokenArgsForCall []struct {
		arg1 string
	}
	refreshAccessTokenReturns struct {
		result1 uaa.RefreshedTokens
		result2 error
	}
	refreshAccessTokenReturnsOnCall map[int]struct {
		result1 uaa.RefreshedTokens
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUAAClient) APIVersion() string {
	fake.aPIVersionMutex.Lock()
	ret, specificReturn := fake.aPIVersionReturnsOnCall[len(fake.aPIVersionArgsForCall)]
	fake.aPIVersionArgsForCall = append(fake.aPIVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("APIVersion", []interface{}{})
	fake.aPIVersionMutex.Unlock()
	if fake.APIVersionStub != nil {
		return fake.APIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.aPIVersionReturns
	return fakeReturns.result1
}

func (fake *FakeUAAClient) APIVersionCallCount() int {
	fake.aPIVersionMutex.RLock()
	defer fake.aPIVersionMutex.RUnlock()
	return len(fake.aPIVersionArgsForCall)
}

func (fake *FakeUAAClient) APIVersionCalls(stub func() string) {
	fake.aPIVersionMutex.Lock()
	defer fake.aPIVersionMutex.Unlock()
	fake.APIVersionStub = stub
}

func (fake *FakeUAAClient) APIVersionReturns(result1 string) {
	fake.aPIVersionMutex.Lock()
	defer fake.aPIVersionMutex.Unlock()
	fake.APIVersionStub = nil
	fake.aPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeUAAClient) APIVersionReturnsOnCall(i int, result1 string) {
	fake.aPIVersionMutex.Lock()
	defer fake.aPIVersionMutex.Unlock()
	fake.APIVersionStub = nil
	if fake.aPIVersionReturnsOnCall == nil {
		fake.aPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.aPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeUAAClient) Authenticate(arg1 string, arg2 string, arg3 string, arg4 constant.GrantType) (string, string, error) {
	fake.authenticateMutex.Lock()
	ret, specificReturn := fake.authenticateReturnsOnCall[len(fake.authenticateArgsForCall)]
	fake.authenticateArgsForCall = append(fake.authenticateArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 constant.GrantType
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Authenticate", []interface{}{arg1, arg2, arg3, arg4})
	fake.authenticateMutex.Unlock()
	if fake.AuthenticateStub != nil {
		return fake.AuthenticateStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.authenticateReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeUAAClient) AuthenticateCallCount() int {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	return len(fake.authenticateArgsForCall)
}

func (fake *FakeUAAClient) AuthenticateCalls(stub func(string, string, string, constant.GrantType) (string, string, error)) {
	fake.authenticateMutex.Lock()
	defer fake.authenticateMutex.Unlock()
	fake.AuthenticateStub = stub
}

func (fake *FakeUAAClient) AuthenticateArgsForCall(i int) (string, string, string, constant.GrantType) {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	argsForCall := fake.authenticateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeUAAClient) AuthenticateReturns(result1 string, result2 string, result3 error) {
	fake.authenticateMutex.Lock()
	defer fake.authenticateMutex.Unlock()
	fake.AuthenticateStub = nil
	fake.authenticateReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUAAClient) AuthenticateReturnsOnCall(i int, result1 string, result2 string, result3 error) {
	fake.authenticateMutex.Lock()
	defer fake.authenticateMutex.Unlock()
	fake.AuthenticateStub = nil
	if fake.authenticateReturnsOnCall == nil {
		fake.authenticateReturnsOnCall = make(map[int]struct {
			result1 string
			result2 string
			result3 error
		})
	}
	fake.authenticateReturnsOnCall[i] = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUAAClient) CreateUser(arg1 string, arg2 string, arg3 string) (uaa.User, error) {
	fake.createUserMutex.Lock()
	ret, specificReturn := fake.createUserReturnsOnCall[len(fake.createUserArgsForCall)]
	fake.createUserArgsForCall = append(fake.createUserArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateUser", []interface{}{arg1, arg2, arg3})
	fake.createUserMutex.Unlock()
	if fake.CreateUserStub != nil {
		return fake.CreateUserStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) CreateUserCallCount() int {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return len(fake.createUserArgsForCall)
}

func (fake *FakeUAAClient) CreateUserCalls(stub func(string, string, string) (uaa.User, error)) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = stub
}

func (fake *FakeUAAClient) CreateUserArgsForCall(i int) (string, string, string) {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	argsForCall := fake.createUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUAAClient) CreateUserReturns(result1 uaa.User, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	fake.createUserReturns = struct {
		result1 uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) CreateUserReturnsOnCall(i int, result1 uaa.User, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	if fake.createUserReturnsOnCall == nil {
		fake.createUserReturnsOnCall = make(map[int]struct {
			result1 uaa.User
			result2 error
		})
	}
	fake.createUserReturnsOnCall[i] = struct {
		result1 uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) GetSSHPasscode(arg1 string, arg2 string) (string, error) {
	fake.getSSHPasscodeMutex.Lock()
	ret, specificReturn := fake.getSSHPasscodeReturnsOnCall[len(fake.getSSHPasscodeArgsForCall)]
	fake.getSSHPasscodeArgsForCall = append(fake.getSSHPasscodeArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetSSHPasscode", []interface{}{arg1, arg2})
	fake.getSSHPasscodeMutex.Unlock()
	if fake.GetSSHPasscodeStub != nil {
		return fake.GetSSHPasscodeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getSSHPasscodeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) GetSSHPasscodeCallCount() int {
	fake.getSSHPasscodeMutex.RLock()
	defer fake.getSSHPasscodeMutex.RUnlock()
	return len(fake.getSSHPasscodeArgsForCall)
}

func (fake *FakeUAAClient) GetSSHPasscodeCalls(stub func(string, string) (string, error)) {
	fake.getSSHPasscodeMutex.Lock()
	defer fake.getSSHPasscodeMutex.Unlock()
	fake.GetSSHPasscodeStub = stub
}

func (fake *FakeUAAClient) GetSSHPasscodeArgsForCall(i int) (string, string) {
	fake.getSSHPasscodeMutex.RLock()
	defer fake.getSSHPasscodeMutex.RUnlock()
	argsForCall := fake.getSSHPasscodeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUAAClient) GetSSHPasscodeReturns(result1 string, result2 error) {
	fake.getSSHPasscodeMutex.Lock()
	defer fake.getSSHPasscodeMutex.Unlock()
	fake.GetSSHPasscodeStub = nil
	fake.getSSHPasscodeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) GetSSHPasscodeReturnsOnCall(i int, result1 string, result2 error) {
	fake.getSSHPasscodeMutex.Lock()
	defer fake.getSSHPasscodeMutex.Unlock()
	fake.GetSSHPasscodeStub = nil
	if fake.getSSHPasscodeReturnsOnCall == nil {
		fake.getSSHPasscodeReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getSSHPasscodeReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) RefreshAccessToken(arg1 string) (uaa.RefreshedTokens, error) {
	fake.refreshAccessTokenMutex.Lock()
	ret, specificReturn := fake.refreshAccessTokenReturnsOnCall[len(fake.refreshAccessTokenArgsForCall)]
	fake.refreshAccessTokenArgsForCall = append(fake.refreshAccessTokenArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RefreshAccessToken", []interface{}{arg1})
	fake.refreshAccessTokenMutex.Unlock()
	if fake.RefreshAccessTokenStub != nil {
		return fake.RefreshAccessTokenStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.refreshAccessTokenReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) RefreshAccessTokenCallCount() int {
	fake.refreshAccessTokenMutex.RLock()
	defer fake.refreshAccessTokenMutex.RUnlock()
	return len(fake.refreshAccessTokenArgsForCall)
}

func (fake *FakeUAAClient) RefreshAccessTokenCalls(stub func(string) (uaa.RefreshedTokens, error)) {
	fake.refreshAccessTokenMutex.Lock()
	defer fake.refreshAccessTokenMutex.Unlock()
	fake.RefreshAccessTokenStub = stub
}

func (fake *FakeUAAClient) RefreshAccessTokenArgsForCall(i int) string {
	fake.refreshAccessTokenMutex.RLock()
	defer fake.refreshAccessTokenMutex.RUnlock()
	argsForCall := fake.refreshAccessTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUAAClient) RefreshAccessTokenReturns(result1 uaa.RefreshedTokens, result2 error) {
	fake.refreshAccessTokenMutex.Lock()
	defer fake.refreshAccessTokenMutex.Unlock()
	fake.RefreshAccessTokenStub = nil
	fake.refreshAccessTokenReturns = struct {
		result1 uaa.RefreshedTokens
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) RefreshAccessTokenReturnsOnCall(i int, result1 uaa.RefreshedTokens, result2 error) {
	fake.refreshAccessTokenMutex.Lock()
	defer fake.refreshAccessTokenMutex.Unlock()
	fake.RefreshAccessTokenStub = nil
	if fake.refreshAccessTokenReturnsOnCall == nil {
		fake.refreshAccessTokenReturnsOnCall = make(map[int]struct {
			result1 uaa.RefreshedTokens
			result2 error
		})
	}
	fake.refreshAccessTokenReturnsOnCall[i] = struct {
		result1 uaa.RefreshedTokens
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.aPIVersionMutex.RLock()
	defer fake.aPIVersionMutex.RUnlock()
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	fake.getSSHPasscodeMutex.RLock()
	defer fake.getSSHPasscodeMutex.RUnlock()
	fake.refreshAccessTokenMutex.RLock()
	defer fake.refreshAccessTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUAAClient) recordInvocation(key string, args []interface{}) {
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

var _ v2action.UAAClient = new(FakeUAAClient)
