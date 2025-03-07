// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"

	github "github.com/google/go-github/v49/github"
)

func AnyPtrToGithubIssueCommentEvent() *github.IssueCommentEvent {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*github.IssueCommentEvent))(nil)).Elem()))
	var nullValue *github.IssueCommentEvent
	return nullValue
}

func EqPtrToGithubIssueCommentEvent(value *github.IssueCommentEvent) *github.IssueCommentEvent {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *github.IssueCommentEvent
	return nullValue
}

func NotEqPtrToGithubIssueCommentEvent(value *github.IssueCommentEvent) *github.IssueCommentEvent {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue *github.IssueCommentEvent
	return nullValue
}

func PtrToGithubIssueCommentEventThat(matcher pegomock.ArgumentMatcher) *github.IssueCommentEvent {
	pegomock.RegisterMatcher(matcher)
	var nullValue *github.IssueCommentEvent
	return nullValue
}
