// Code generated by protoc-gen-go-bnet. DO NOT EDIT.

package challenge_service

import (
	"context"
	"strings"
)

// ChallengeServiceHash is the hash of the service name
const ChallengeServiceHash uint32 = 1908107315 // bgs.protocol.challenge.v1.ChallengeService
// ChallengeServiceOriginalFullyQualifiedDescriptorName is the original fully qualified descriptor name
const ChallengeServiceOriginalFullyQualifiedDescriptorName string = "bnet.protocol.challenge.ChallengeService"

// ChallengeServiceOriginalFullyQualifiedDescriptorNameHash is the hash of the original fully qualified descriptor name
const ChallengeServiceOriginalFullyQualifiedDescriptorNameHash uint32 = 3686756121 // bnet.protocol.challenge.ChallengeService

// ChallengeService is the client API for bgs.protocol.challenge.v1.ChallengeService
type ChallengeService interface {
	// ChallengePicked has a method id of 1 and a method name of bgs.protocol.challenge.v1.ChallengeService.ChallengePicked
	ChallengePicked(context.Context, *ChallengePickedRequest, interface{}) error
	// ChallengeAnswered has a method id of 2 and a method name of bgs.protocol.challenge.v1.ChallengeService.ChallengeAnswered
	ChallengeAnswered(context.Context, *ChallengeAnsweredRequest, interface{}) error
	// ChallengeCancelled has a method id of 3 and a method name of bgs.protocol.challenge.v1.ChallengeService.ChallengeCancelled
	ChallengeCancelled(context.Context, *ChallengeCancelledRequest, interface{}) error
	// SendChallengeToUser has a method id of 4 and a method name of bgs.protocol.challenge.v1.ChallengeService.SendChallengeToUser
	SendChallengeToUser(context.Context, *SendChallengeToUserRequest, interface{}) error
}

// ChallengeServiceMethods is a map of method ids to method names
var ChallengeServiceMethods = map[uint32]string{
	1: "bgs.protocol.challenge.v1.ChallengeService.ChallengePicked",
	2: "bgs.protocol.challenge.v1.ChallengeService.ChallengeAnswered",
	3: "bgs.protocol.challenge.v1.ChallengeService.ChallengeCancelled",
	4: "bgs.protocol.challenge.v1.ChallengeService.SendChallengeToUser",
}

// ChallengeServiceMethod is the method id for bgs.protocol.challenge.v1.ChallengeService
type ChallengeServiceMethod uint32

// Enum value maps for ChallengeServiceMethod.
const (
	// ChallengeServiceMethod_ChallengePicked ChallengeServiceMethod = 1 // bgs.protocol.challenge.v1.ChallengeService.ChallengePicked
	ChallengeServiceMethod_ChallengePicked ChallengeServiceMethod = 1 // bgs.protocol.challenge.v1.ChallengeService.ChallengePicked
	// ChallengeServiceMethod_ChallengeAnswered ChallengeServiceMethod = 2 // bgs.protocol.challenge.v1.ChallengeService.ChallengeAnswered
	ChallengeServiceMethod_ChallengeAnswered ChallengeServiceMethod = 2 // bgs.protocol.challenge.v1.ChallengeService.ChallengeAnswered
	// ChallengeServiceMethod_ChallengeCancelled ChallengeServiceMethod = 3 // bgs.protocol.challenge.v1.ChallengeService.ChallengeCancelled
	ChallengeServiceMethod_ChallengeCancelled ChallengeServiceMethod = 3 // bgs.protocol.challenge.v1.ChallengeService.ChallengeCancelled
	// ChallengeServiceMethod_SendChallengeToUser ChallengeServiceMethod = 4 // bgs.protocol.challenge.v1.ChallengeService.SendChallengeToUser
	ChallengeServiceMethod_SendChallengeToUser ChallengeServiceMethod = 4 // bgs.protocol.challenge.v1.ChallengeService.SendChallengeToUser
)

// ToUint32 converts a ChallengeServiceMethod enum to a uint32
func (m ChallengeServiceMethod) ToUint32() uint32 {
	return uint32(m)
}

// GetMethodFullName returns the method full name for the given method id
func (m ChallengeServiceMethod) GetMethodFullName() string {
	return ChallengeServiceMethods[uint32(m)]
}

// GetMethodName returns the method name for the given method id
func (m ChallengeServiceMethod) GetMethodName() string {
	if _, ok := ChallengeServiceMethods[uint32(m)]; !ok {
		return "No Method Exist"
	}
	return ChallengeServiceMethods[uint32(m)][strings.LastIndex(ChallengeServiceMethods[uint32(m)], ".")+1:]
}

// ChallengeListenerHash is the hash of the service name
const ChallengeListenerHash uint32 = 3336112824 // bgs.protocol.challenge.v1.ChallengeListener
// ChallengeListenerOriginalFullyQualifiedDescriptorName is the original fully qualified descriptor name
const ChallengeListenerOriginalFullyQualifiedDescriptorName string = "bnet.protocol.challenge.ChallengeNotify"

// ChallengeListenerOriginalFullyQualifiedDescriptorNameHash is the hash of the original fully qualified descriptor name
const ChallengeListenerOriginalFullyQualifiedDescriptorNameHash uint32 = 3151632159 // bnet.protocol.challenge.ChallengeNotify

// ChallengeListener is the client API for bgs.protocol.challenge.v1.ChallengeListener
type ChallengeListener interface {
	// OnChallengeUser has a method id of 1 and a method name of bgs.protocol.challenge.v1.ChallengeListener.OnChallengeUser
	OnChallengeUser(context.Context, *ChallengeUserRequest, interface{}) error
	// OnChallengeResult has a method id of 2 and a method name of bgs.protocol.challenge.v1.ChallengeListener.OnChallengeResult
	OnChallengeResult(context.Context, *ChallengeResultRequest, interface{}) error
	// OnExternalChallenge has a method id of 3 and a method name of bgs.protocol.challenge.v1.ChallengeListener.OnExternalChallenge
	OnExternalChallenge(context.Context, *ChallengeExternalRequest, interface{}) error
	// OnExternalChallengeResult has a method id of 4 and a method name of bgs.protocol.challenge.v1.ChallengeListener.OnExternalChallengeResult
	OnExternalChallengeResult(context.Context, *ChallengeExternalResult, interface{}) error
}

// ChallengeListenerMethods is a map of method ids to method names
var ChallengeListenerMethods = map[uint32]string{
	1: "bgs.protocol.challenge.v1.ChallengeListener.OnChallengeUser",
	2: "bgs.protocol.challenge.v1.ChallengeListener.OnChallengeResult",
	3: "bgs.protocol.challenge.v1.ChallengeListener.OnExternalChallenge",
	4: "bgs.protocol.challenge.v1.ChallengeListener.OnExternalChallengeResult",
}

// ChallengeListenerMethod is the method id for bgs.protocol.challenge.v1.ChallengeListener
type ChallengeListenerMethod uint32

// Enum value maps for ChallengeListenerMethod.
const (
	// ChallengeListenerMethod_OnChallengeUser ChallengeListenerMethod = 1 // bgs.protocol.challenge.v1.ChallengeListener.OnChallengeUser
	ChallengeListenerMethod_OnChallengeUser ChallengeListenerMethod = 1 // bgs.protocol.challenge.v1.ChallengeListener.OnChallengeUser
	// ChallengeListenerMethod_OnChallengeResult ChallengeListenerMethod = 2 // bgs.protocol.challenge.v1.ChallengeListener.OnChallengeResult
	ChallengeListenerMethod_OnChallengeResult ChallengeListenerMethod = 2 // bgs.protocol.challenge.v1.ChallengeListener.OnChallengeResult
	// ChallengeListenerMethod_OnExternalChallenge ChallengeListenerMethod = 3 // bgs.protocol.challenge.v1.ChallengeListener.OnExternalChallenge
	ChallengeListenerMethod_OnExternalChallenge ChallengeListenerMethod = 3 // bgs.protocol.challenge.v1.ChallengeListener.OnExternalChallenge
	// ChallengeListenerMethod_OnExternalChallengeResult ChallengeListenerMethod = 4 // bgs.protocol.challenge.v1.ChallengeListener.OnExternalChallengeResult
	ChallengeListenerMethod_OnExternalChallengeResult ChallengeListenerMethod = 4 // bgs.protocol.challenge.v1.ChallengeListener.OnExternalChallengeResult
)

// ToUint32 converts a ChallengeListenerMethod enum to a uint32
func (m ChallengeListenerMethod) ToUint32() uint32 {
	return uint32(m)
}

// GetMethodFullName returns the method full name for the given method id
func (m ChallengeListenerMethod) GetMethodFullName() string {
	return ChallengeListenerMethods[uint32(m)]
}

// GetMethodName returns the method name for the given method id
func (m ChallengeListenerMethod) GetMethodName() string {
	if _, ok := ChallengeListenerMethods[uint32(m)]; !ok {
		return "No Method Exist"
	}
	return ChallengeListenerMethods[uint32(m)][strings.LastIndex(ChallengeListenerMethods[uint32(m)], ".")+1:]
}
