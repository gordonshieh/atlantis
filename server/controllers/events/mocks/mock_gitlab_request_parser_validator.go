// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/controllers/events (interfaces: GitlabRequestParserValidator)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	http "net/http"
	"reflect"
	"time"
)

type MockGitlabRequestParserValidator struct {
	fail func(message string, callerSkip ...int)
}

func NewMockGitlabRequestParserValidator(options ...pegomock.Option) *MockGitlabRequestParserValidator {
	mock := &MockGitlabRequestParserValidator{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockGitlabRequestParserValidator) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockGitlabRequestParserValidator) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockGitlabRequestParserValidator) ParseAndValidate(_param0 *http.Request, _param1 []byte) (interface{}, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockGitlabRequestParserValidator().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseAndValidate", params, []reflect.Type{reflect.TypeOf((*interface{})(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 interface{}
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(interface{})
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockGitlabRequestParserValidator) VerifyWasCalledOnce() *VerifierMockGitlabRequestParserValidator {
	return &VerifierMockGitlabRequestParserValidator{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockGitlabRequestParserValidator) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockGitlabRequestParserValidator {
	return &VerifierMockGitlabRequestParserValidator{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockGitlabRequestParserValidator) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockGitlabRequestParserValidator {
	return &VerifierMockGitlabRequestParserValidator{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockGitlabRequestParserValidator) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockGitlabRequestParserValidator {
	return &VerifierMockGitlabRequestParserValidator{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockGitlabRequestParserValidator struct {
	mock                   *MockGitlabRequestParserValidator
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockGitlabRequestParserValidator) ParseAndValidate(_param0 *http.Request, _param1 []byte) *MockGitlabRequestParserValidator_ParseAndValidate_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseAndValidate", params, verifier.timeout)
	return &MockGitlabRequestParserValidator_ParseAndValidate_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockGitlabRequestParserValidator_ParseAndValidate_OngoingVerification struct {
	mock              *MockGitlabRequestParserValidator
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockGitlabRequestParserValidator_ParseAndValidate_OngoingVerification) GetCapturedArguments() (*http.Request, []byte) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockGitlabRequestParserValidator_ParseAndValidate_OngoingVerification) GetAllCapturedArguments() (_param0 []*http.Request, _param1 [][]byte) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*http.Request, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(*http.Request)
		}
		_param1 = make([][]byte, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.([]byte)
		}
	}
	return
}
