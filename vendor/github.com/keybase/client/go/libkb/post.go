// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package libkb

import (
	"fmt"
	"runtime/debug"

	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	jsonw "github.com/keybase/go-jsonw"
)

type PostProofRes struct {
	Text     string
	ID       string
	Metadata *jsonw.Wrapper
}

type PostProofArg struct {
	Sig            string
	SigInner       []byte
	ID             keybase1.SigID
	RemoteUsername string
	ProofType      string
	Supersede      bool
	RemoteKey      string
	SigningKey     GenericKey
}

func PostProof(m MetaContext, arg PostProofArg) (*PostProofRes, error) {
	hargs := HTTPArgs{
		"sig_id_base":     S{arg.ID.ToString(false)},
		"sig_id_short":    S{arg.ID.ToShortID()},
		"sig":             S{arg.Sig},
		"is_remote_proof": B{true},
		"supersede":       B{arg.Supersede},
		"signing_kid":     S{arg.SigningKey.GetKID().String()},
		"type":            S{arg.ProofType},
	}
	if len(arg.SigInner) > 0 {
		hargs["sig_inner"] = S{string(arg.SigInner)}
	}
	hargs.Add(arg.RemoteKey, S{arg.RemoteUsername})

	res, err := m.G().API.Post(APIArg{
		Endpoint:    "sig/post",
		SessionType: APISessionTypeREQUIRED,
		Args:        hargs,
		MetaContext: m,
	})

	if err != nil {
		return nil, err
	}
	var tmp PostProofRes
	res.Body.AtKey("proof_text").GetStringVoid(&tmp.Text, &err)
	res.Body.AtKey("proof_id").GetStringVoid(&tmp.ID, &err)
	tmp.Metadata = res.Body.AtKey("proof_metadata")

	var ret *PostProofRes
	if err == nil {
		ret = &tmp
	}
	return ret, err
}

type PostAuthProofArg struct {
	uid keybase1.UID
	sig string
	key GenericKey
}

type PostAuthProofRes struct {
	SessionID string `json:"session"`
	AuthID    string `json:"auth_id"`
	CSRFToken string `json:"csrf_token"`
	UIDHex    string `json:"uid"`
	Username  string `json:"username"`
	PPGen     int    `json:"passphrase_generation"`
}

func PostAuthProof(m MetaContext, arg PostAuthProofArg) (*PostAuthProofRes, error) {
	hargs := HTTPArgs{
		"uid":         UIDArg(arg.uid),
		"sig":         S{arg.sig},
		"signing_kid": S{arg.key.GetKID().String()},
	}
	res, err := m.G().API.Post(APIArg{
		Endpoint:    "sig/post_auth",
		SessionType: APISessionTypeNONE,
		Args:        hargs,
		MetaContext: m,
	})
	if err != nil {
		return nil, err
	}
	var ret *PostAuthProofRes
	var tmp PostAuthProofRes
	if err = res.Body.UnmarshalAgain(&tmp); err == nil {
		ret = &tmp
	}
	return ret, err
}

type InviteRequestArg struct {
	Email    string
	Fullname string
	Notes    string
}

func PostInviteRequest(m MetaContext, arg InviteRequestArg) (err error) {
	_, err = m.G().API.Post(APIArg{
		Endpoint: "invitation_request",
		Args: HTTPArgs{
			"email":     S{arg.Email},
			"full_name": S{arg.Fullname},
			"notes":     S{arg.Notes},
		},
		MetaContext: m,
	})
	return err
}

func DeletePrimary(m MetaContext) (err error) {
	_, err = m.G().API.Post(APIArg{
		Endpoint:    "key/revoke",
		SessionType: APISessionTypeREQUIRED,
		Args: HTTPArgs{
			"revoke_primary":  I{1},
			"revocation_type": I{RevSimpleDelete},
		},
		MetaContext: m,
	})
	return
}

func CheckPosted(m MetaContext, proofID string) (found bool, status keybase1.ProofStatus, state keybase1.ProofState, err error) {
	res, e2 := m.G().API.Post(APIArg{
		Endpoint:    "sig/posted",
		SessionType: APISessionTypeREQUIRED,
		Args: HTTPArgs{
			"proof_id": S{proofID},
		},
		MetaContext: m,
	})
	if e2 != nil {
		err = e2
		return
	}
	var (
		rfound  bool
		rstatus int
		rstate  int
		rerr    error
	)
	res.Body.AtKey("proof_ok").GetBoolVoid(&rfound, &rerr)
	res.Body.AtPath("proof_res.status").GetIntVoid(&rstatus, &rerr)
	res.Body.AtPath("proof_res.state").GetIntVoid(&rstate, &rerr)
	return rfound, keybase1.ProofStatus(rstatus), keybase1.ProofState(rstate), rerr
}

func CheckPostedViaSigID(m MetaContext, sigID keybase1.SigID) (found bool, status keybase1.ProofStatus, state keybase1.ProofState, err error) {
	res, e2 := m.G().API.Post(APIArg{
		Endpoint:    "sig/posted",
		SessionType: APISessionTypeREQUIRED,
		Args: HTTPArgs{
			"sig_id": S{sigID.ToString(true)},
		},
		MetaContext: m,
	})
	if e2 != nil {
		err = e2
		return
	}

	var (
		rfound  bool
		rstatus int
		rstate  int
		rerr    error
	)
	res.Body.AtKey("proof_ok").GetBoolVoid(&rfound, &rerr)
	res.Body.AtPath("proof_res.status").GetIntVoid(&rstatus, &rerr)
	res.Body.AtPath("proof_res.state").GetIntVoid(&rstate, &rerr)
	return rfound, keybase1.ProofStatus(rstatus), keybase1.ProofState(rstate), rerr
}

func PostDeviceLKS(m MetaContext, deviceID keybase1.DeviceID, deviceType string, serverHalf LKSecServerHalf,
	ppGen PassphraseGeneration,
	clientHalfRecovery string, clientHalfRecoveryKID keybase1.KID) error {
	m.CDebugf("| PostDeviceLKS: %s", deviceID)
	if serverHalf.IsNil() {
		return fmt.Errorf("PostDeviceLKS: called with empty serverHalf")
	}
	if ppGen < 1 {
		m.CWarningf("PostDeviceLKS: ppGen < 1 (%d)", ppGen)
		debug.PrintStack()
	}
	arg := APIArg{
		Endpoint:    "device/update",
		SessionType: APISessionTypeREQUIRED,
		Args: HTTPArgs{
			"device_id":       S{Val: deviceID.String()},
			"type":            S{Val: deviceType},
			"lks_server_half": S{Val: serverHalf.EncodeToHex()},
			"ppgen":           I{Val: int(ppGen)},
			"lks_client_half": S{Val: clientHalfRecovery},
			"kid":             S{Val: clientHalfRecoveryKID.String()},
			"platform":        S{Val: GetPlatformString()},
		},
		RetryCount:  10,
		MetaContext: m,
	}
	_, err := m.G().API.Post(arg)
	if err != nil {
		m.CInfof("device/update(%+v) failed: %s", arg.Args, err)
	}
	return err
}

func CheckInvitationCode(m MetaContext, code string) error {
	arg := APIArg{
		Endpoint:    "invitation/check",
		SessionType: APISessionTypeNONE,
		Args: HTTPArgs{
			"invitation_id": S{Val: code},
		},
		MetaContext: m,
	}
	_, err := m.G().API.Get(arg)
	return err
}

func GetInvitationCode(m MetaContext) (string, error) {
	arg := APIArg{
		Endpoint:    "invitation_bypass_request",
		SessionType: APISessionTypeNONE,
		MetaContext: m,
	}
	res, err := m.G().API.Get(arg)
	var invitationID string
	if err == nil {
		invitationID, err = res.Body.AtKey("invitation_id").GetString()
	}
	return invitationID, err
}
