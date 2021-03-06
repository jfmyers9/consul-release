// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/hashicorp/consul/api"
)

type FakeconsulAPIAgent struct {
	MembersStub        func(wan bool) ([]*api.AgentMember, error)
	membersMutex       sync.RWMutex
	membersArgsForCall []struct {
		wan bool
	}
	membersReturns struct {
		result1 []*api.AgentMember
		result2 error
	}
	JoinCall struct {
		CallCount int
		Stub      func(member string, wan bool) error
		Receives  struct {
			Members []string
			WAN     bool
		}
		Returns struct {
			Error error
		}
	}
	SelfCall struct {
		CallCount int
		Returns   struct {
			SelfInfo map[string]map[string]interface{}
			Error    error
		}
	}
}

func (fake *FakeconsulAPIAgent) Self() (map[string]map[string]interface{}, error) {
	fake.SelfCall.CallCount++
	return fake.SelfCall.Returns.SelfInfo, fake.SelfCall.Returns.Error
}

func (fake *FakeconsulAPIAgent) Members(wan bool) ([]*api.AgentMember, error) {
	fake.membersMutex.Lock()
	fake.membersArgsForCall = append(fake.membersArgsForCall, struct {
		wan bool
	}{wan})
	fake.membersMutex.Unlock()
	if fake.MembersStub != nil {
		return fake.MembersStub(wan)
	} else {
		return fake.membersReturns.result1, fake.membersReturns.result2
	}
}

func (fake *FakeconsulAPIAgent) MembersCallCount() int {
	fake.membersMutex.RLock()
	defer fake.membersMutex.RUnlock()
	return len(fake.membersArgsForCall)
}

func (fake *FakeconsulAPIAgent) MembersArgsForCall(i int) bool {
	fake.membersMutex.RLock()
	defer fake.membersMutex.RUnlock()
	return fake.membersArgsForCall[i].wan
}

func (fake *FakeconsulAPIAgent) MembersReturns(result1 []*api.AgentMember, result2 error) {
	fake.MembersStub = nil
	fake.membersReturns = struct {
		result1 []*api.AgentMember
		result2 error
	}{result1, result2}
}

func (fake *FakeconsulAPIAgent) Join(member string, wan bool) error {
	fake.JoinCall.CallCount++
	fake.JoinCall.Receives.Members = append(fake.JoinCall.Receives.Members, member)
	fake.JoinCall.Receives.WAN = wan
	if fake.JoinCall.Stub != nil {
		return fake.JoinCall.Stub(member, wan)
	}
	return fake.JoinCall.Returns.Error
}

// var _ confab.consulAPIAgent = new(FakeconsulAPIAgent)
