// Auto-generated by avdl-compiler v1.3.24 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/config.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type GetCurrentStatusRes struct {
	Configured     bool  `codec:"configured" json:"configured"`
	Registered     bool  `codec:"registered" json:"registered"`
	LoggedIn       bool  `codec:"loggedIn" json:"loggedIn"`
	SessionIsValid bool  `codec:"sessionIsValid" json:"sessionIsValid"`
	User           *User `codec:"user,omitempty" json:"user,omitempty"`
}

func (o GetCurrentStatusRes) DeepCopy() GetCurrentStatusRes {
	return GetCurrentStatusRes{
		Configured:     o.Configured,
		Registered:     o.Registered,
		LoggedIn:       o.LoggedIn,
		SessionIsValid: o.SessionIsValid,
		User: (func(x *User) *User {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.User),
	}
}

type SessionStatus struct {
	SessionFor string `codec:"SessionFor" json:"SessionFor"`
	Loaded     bool   `codec:"Loaded" json:"Loaded"`
	Cleared    bool   `codec:"Cleared" json:"Cleared"`
	SaltOnly   bool   `codec:"SaltOnly" json:"SaltOnly"`
	Expired    bool   `codec:"Expired" json:"Expired"`
}

func (o SessionStatus) DeepCopy() SessionStatus {
	return SessionStatus{
		SessionFor: o.SessionFor,
		Loaded:     o.Loaded,
		Cleared:    o.Cleared,
		SaltOnly:   o.SaltOnly,
		Expired:    o.Expired,
	}
}

type ClientDetails struct {
	Pid        int        `codec:"pid" json:"pid"`
	ClientType ClientType `codec:"clientType" json:"clientType"`
	Argv       []string   `codec:"argv" json:"argv"`
	Desc       string     `codec:"desc" json:"desc"`
	Version    string     `codec:"version" json:"version"`
}

func (o ClientDetails) DeepCopy() ClientDetails {
	return ClientDetails{
		Pid:        o.Pid,
		ClientType: o.ClientType.DeepCopy(),
		Argv: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.Argv),
		Desc:    o.Desc,
		Version: o.Version,
	}
}

type PlatformInfo struct {
	Os        string `codec:"os" json:"os"`
	OsVersion string `codec:"osVersion" json:"osVersion"`
	Arch      string `codec:"arch" json:"arch"`
	GoVersion string `codec:"goVersion" json:"goVersion"`
}

func (o PlatformInfo) DeepCopy() PlatformInfo {
	return PlatformInfo{
		Os:        o.Os,
		OsVersion: o.OsVersion,
		Arch:      o.Arch,
		GoVersion: o.GoVersion,
	}
}

type LoadDeviceErr struct {
	Where string `codec:"where" json:"where"`
	Desc  string `codec:"desc" json:"desc"`
}

func (o LoadDeviceErr) DeepCopy() LoadDeviceErr {
	return LoadDeviceErr{
		Where: o.Where,
		Desc:  o.Desc,
	}
}

type ExtendedStatus struct {
	Standalone             bool            `codec:"standalone" json:"standalone"`
	PassphraseStreamCached bool            `codec:"passphraseStreamCached" json:"passphraseStreamCached"`
	TsecCached             bool            `codec:"tsecCached" json:"tsecCached"`
	DeviceSigKeyCached     bool            `codec:"deviceSigKeyCached" json:"deviceSigKeyCached"`
	DeviceEncKeyCached     bool            `codec:"deviceEncKeyCached" json:"deviceEncKeyCached"`
	PaperSigKeyCached      bool            `codec:"paperSigKeyCached" json:"paperSigKeyCached"`
	PaperEncKeyCached      bool            `codec:"paperEncKeyCached" json:"paperEncKeyCached"`
	StoredSecret           bool            `codec:"storedSecret" json:"storedSecret"`
	SecretPromptSkip       bool            `codec:"secretPromptSkip" json:"secretPromptSkip"`
	RememberPassphrase     bool            `codec:"rememberPassphrase" json:"rememberPassphrase"`
	Device                 *Device         `codec:"device,omitempty" json:"device,omitempty"`
	DeviceErr              *LoadDeviceErr  `codec:"deviceErr,omitempty" json:"deviceErr,omitempty"`
	LogDir                 string          `codec:"logDir" json:"logDir"`
	Session                *SessionStatus  `codec:"session,omitempty" json:"session,omitempty"`
	DefaultUsername        string          `codec:"defaultUsername" json:"defaultUsername"`
	ProvisionedUsernames   []string        `codec:"provisionedUsernames" json:"provisionedUsernames"`
	Clients                []ClientDetails `codec:"Clients" json:"Clients"`
	DeviceEkNames          []string        `codec:"deviceEkNames" json:"deviceEkNames"`
	PlatformInfo           PlatformInfo    `codec:"platformInfo" json:"platformInfo"`
	DefaultDeviceID        DeviceID        `codec:"defaultDeviceID" json:"defaultDeviceID"`
}

func (o ExtendedStatus) DeepCopy() ExtendedStatus {
	return ExtendedStatus{
		Standalone:             o.Standalone,
		PassphraseStreamCached: o.PassphraseStreamCached,
		TsecCached:             o.TsecCached,
		DeviceSigKeyCached:     o.DeviceSigKeyCached,
		DeviceEncKeyCached:     o.DeviceEncKeyCached,
		PaperSigKeyCached:      o.PaperSigKeyCached,
		PaperEncKeyCached:      o.PaperEncKeyCached,
		StoredSecret:           o.StoredSecret,
		SecretPromptSkip:       o.SecretPromptSkip,
		RememberPassphrase:     o.RememberPassphrase,
		Device: (func(x *Device) *Device {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Device),
		DeviceErr: (func(x *LoadDeviceErr) *LoadDeviceErr {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.DeviceErr),
		LogDir: o.LogDir,
		Session: (func(x *SessionStatus) *SessionStatus {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Session),
		DefaultUsername: o.DefaultUsername,
		ProvisionedUsernames: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.ProvisionedUsernames),
		Clients: (func(x []ClientDetails) []ClientDetails {
			if x == nil {
				return nil
			}
			ret := make([]ClientDetails, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Clients),
		DeviceEkNames: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.DeviceEkNames),
		PlatformInfo:    o.PlatformInfo.DeepCopy(),
		DefaultDeviceID: o.DefaultDeviceID.DeepCopy(),
	}
}

type AllProvisionedUsernames struct {
	DefaultUsername      string   `codec:"defaultUsername" json:"defaultUsername"`
	ProvisionedUsernames []string `codec:"provisionedUsernames" json:"provisionedUsernames"`
}

func (o AllProvisionedUsernames) DeepCopy() AllProvisionedUsernames {
	return AllProvisionedUsernames{
		DefaultUsername: o.DefaultUsername,
		ProvisionedUsernames: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.ProvisionedUsernames),
	}
}

type ForkType int

const (
	ForkType_NONE     ForkType = 0
	ForkType_AUTO     ForkType = 1
	ForkType_WATCHDOG ForkType = 2
	ForkType_LAUNCHD  ForkType = 3
	ForkType_SYSTEMD  ForkType = 4
)

func (o ForkType) DeepCopy() ForkType { return o }

var ForkTypeMap = map[string]ForkType{
	"NONE":     0,
	"AUTO":     1,
	"WATCHDOG": 2,
	"LAUNCHD":  3,
	"SYSTEMD":  4,
}

var ForkTypeRevMap = map[ForkType]string{
	0: "NONE",
	1: "AUTO",
	2: "WATCHDOG",
	3: "LAUNCHD",
	4: "SYSTEMD",
}

func (e ForkType) String() string {
	if v, ok := ForkTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type Config struct {
	ServerURI      string   `codec:"serverURI" json:"serverURI"`
	SocketFile     string   `codec:"socketFile" json:"socketFile"`
	Label          string   `codec:"label" json:"label"`
	RunMode        string   `codec:"runMode" json:"runMode"`
	GpgExists      bool     `codec:"gpgExists" json:"gpgExists"`
	GpgPath        string   `codec:"gpgPath" json:"gpgPath"`
	Version        string   `codec:"version" json:"version"`
	Path           string   `codec:"path" json:"path"`
	BinaryRealpath string   `codec:"binaryRealpath" json:"binaryRealpath"`
	ConfigPath     string   `codec:"configPath" json:"configPath"`
	VersionShort   string   `codec:"versionShort" json:"versionShort"`
	VersionFull    string   `codec:"versionFull" json:"versionFull"`
	IsAutoForked   bool     `codec:"isAutoForked" json:"isAutoForked"`
	ForkType       ForkType `codec:"forkType" json:"forkType"`
}

func (o Config) DeepCopy() Config {
	return Config{
		ServerURI:      o.ServerURI,
		SocketFile:     o.SocketFile,
		Label:          o.Label,
		RunMode:        o.RunMode,
		GpgExists:      o.GpgExists,
		GpgPath:        o.GpgPath,
		Version:        o.Version,
		Path:           o.Path,
		BinaryRealpath: o.BinaryRealpath,
		ConfigPath:     o.ConfigPath,
		VersionShort:   o.VersionShort,
		VersionFull:    o.VersionFull,
		IsAutoForked:   o.IsAutoForked,
		ForkType:       o.ForkType.DeepCopy(),
	}
}

type ConfigValue struct {
	IsNull bool    `codec:"isNull" json:"isNull"`
	B      *bool   `codec:"b,omitempty" json:"b,omitempty"`
	I      *int    `codec:"i,omitempty" json:"i,omitempty"`
	S      *string `codec:"s,omitempty" json:"s,omitempty"`
	O      *string `codec:"o,omitempty" json:"o,omitempty"`
}

func (o ConfigValue) DeepCopy() ConfigValue {
	return ConfigValue{
		IsNull: o.IsNull,
		B: (func(x *bool) *bool {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.B),
		I: (func(x *int) *int {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.I),
		S: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.S),
		O: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.O),
	}
}

type OutOfDateInfo struct {
	UpgradeTo         string `codec:"upgradeTo" json:"upgradeTo"`
	UpgradeURI        string `codec:"upgradeURI" json:"upgradeURI"`
	CustomMessage     string `codec:"customMessage" json:"customMessage"`
	CriticalClockSkew int64  `codec:"criticalClockSkew" json:"criticalClockSkew"`
}

func (o OutOfDateInfo) DeepCopy() OutOfDateInfo {
	return OutOfDateInfo{
		UpgradeTo:         o.UpgradeTo,
		UpgradeURI:        o.UpgradeURI,
		CustomMessage:     o.CustomMessage,
		CriticalClockSkew: o.CriticalClockSkew,
	}
}

type BootstrapStatus struct {
	Registered bool     `codec:"registered" json:"registered"`
	LoggedIn   bool     `codec:"loggedIn" json:"loggedIn"`
	Uid        UID      `codec:"uid" json:"uid"`
	Username   string   `codec:"username" json:"username"`
	DeviceID   DeviceID `codec:"deviceID" json:"deviceID"`
	DeviceName string   `codec:"deviceName" json:"deviceName"`
	Following  []string `codec:"following" json:"following"`
	Followers  []string `codec:"followers" json:"followers"`
}

func (o BootstrapStatus) DeepCopy() BootstrapStatus {
	return BootstrapStatus{
		Registered: o.Registered,
		LoggedIn:   o.LoggedIn,
		Uid:        o.Uid.DeepCopy(),
		Username:   o.Username,
		DeviceID:   o.DeviceID.DeepCopy(),
		DeviceName: o.DeviceName,
		Following: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.Following),
		Followers: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.Followers),
	}
}

type GetCurrentStatusArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type GetExtendedStatusArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type GetAllProvisionedUsernamesArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type GetConfigArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type SetUserConfigArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Username  string `codec:"username" json:"username"`
	Key       string `codec:"key" json:"key"`
	Value     string `codec:"value" json:"value"`
}

type SetPathArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Path      string `codec:"path" json:"path"`
}

type HelloIAmArg struct {
	Details ClientDetails `codec:"details" json:"details"`
}

type SetValueArg struct {
	Path  string      `codec:"path" json:"path"`
	Value ConfigValue `codec:"value" json:"value"`
}

type ClearValueArg struct {
	Path string `codec:"path" json:"path"`
}

type GetValueArg struct {
	Path string `codec:"path" json:"path"`
}

type CheckAPIServerOutOfDateWarningArg struct {
}

type WaitForClientArg struct {
	ClientType ClientType  `codec:"clientType" json:"clientType"`
	Timeout    DurationSec `codec:"timeout" json:"timeout"`
}

type GetBootstrapStatusArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type GetRememberPassphraseArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type SetRememberPassphraseArg struct {
	SessionID int  `codec:"sessionID" json:"sessionID"`
	Remember  bool `codec:"remember" json:"remember"`
}

type ConfigInterface interface {
	GetCurrentStatus(context.Context, int) (GetCurrentStatusRes, error)
	GetExtendedStatus(context.Context, int) (ExtendedStatus, error)
	GetAllProvisionedUsernames(context.Context, int) (AllProvisionedUsernames, error)
	GetConfig(context.Context, int) (Config, error)
	// Change user config.
	// For example, to update primary picture source:
	// key=picture.source, value=twitter (or github)
	SetUserConfig(context.Context, SetUserConfigArg) error
	SetPath(context.Context, SetPathArg) error
	HelloIAm(context.Context, ClientDetails) error
	SetValue(context.Context, SetValueArg) error
	ClearValue(context.Context, string) error
	GetValue(context.Context, string) (ConfigValue, error)
	// Check whether the API server has told us we're out of date.
	CheckAPIServerOutOfDateWarning(context.Context) (OutOfDateInfo, error)
	// Wait for client type to connect to service.
	WaitForClient(context.Context, WaitForClientArg) (bool, error)
	GetBootstrapStatus(context.Context, int) (BootstrapStatus, error)
	GetRememberPassphrase(context.Context, int) (bool, error)
	SetRememberPassphrase(context.Context, SetRememberPassphraseArg) error
}

func ConfigProtocol(i ConfigInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.config",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getCurrentStatus": {
				MakeArg: func() interface{} {
					ret := make([]GetCurrentStatusArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetCurrentStatusArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetCurrentStatusArg)(nil), args)
						return
					}
					ret, err = i.GetCurrentStatus(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getExtendedStatus": {
				MakeArg: func() interface{} {
					ret := make([]GetExtendedStatusArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetExtendedStatusArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetExtendedStatusArg)(nil), args)
						return
					}
					ret, err = i.GetExtendedStatus(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getAllProvisionedUsernames": {
				MakeArg: func() interface{} {
					ret := make([]GetAllProvisionedUsernamesArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetAllProvisionedUsernamesArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetAllProvisionedUsernamesArg)(nil), args)
						return
					}
					ret, err = i.GetAllProvisionedUsernames(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getConfig": {
				MakeArg: func() interface{} {
					ret := make([]GetConfigArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetConfigArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetConfigArg)(nil), args)
						return
					}
					ret, err = i.GetConfig(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"setUserConfig": {
				MakeArg: func() interface{} {
					ret := make([]SetUserConfigArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SetUserConfigArg)
					if !ok {
						err = rpc.NewTypeError((*[]SetUserConfigArg)(nil), args)
						return
					}
					err = i.SetUserConfig(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"setPath": {
				MakeArg: func() interface{} {
					ret := make([]SetPathArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SetPathArg)
					if !ok {
						err = rpc.NewTypeError((*[]SetPathArg)(nil), args)
						return
					}
					err = i.SetPath(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"helloIAm": {
				MakeArg: func() interface{} {
					ret := make([]HelloIAmArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]HelloIAmArg)
					if !ok {
						err = rpc.NewTypeError((*[]HelloIAmArg)(nil), args)
						return
					}
					err = i.HelloIAm(ctx, (*typedArgs)[0].Details)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"setValue": {
				MakeArg: func() interface{} {
					ret := make([]SetValueArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SetValueArg)
					if !ok {
						err = rpc.NewTypeError((*[]SetValueArg)(nil), args)
						return
					}
					err = i.SetValue(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"clearValue": {
				MakeArg: func() interface{} {
					ret := make([]ClearValueArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ClearValueArg)
					if !ok {
						err = rpc.NewTypeError((*[]ClearValueArg)(nil), args)
						return
					}
					err = i.ClearValue(ctx, (*typedArgs)[0].Path)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getValue": {
				MakeArg: func() interface{} {
					ret := make([]GetValueArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetValueArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetValueArg)(nil), args)
						return
					}
					ret, err = i.GetValue(ctx, (*typedArgs)[0].Path)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"checkAPIServerOutOfDateWarning": {
				MakeArg: func() interface{} {
					ret := make([]CheckAPIServerOutOfDateWarningArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.CheckAPIServerOutOfDateWarning(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"waitForClient": {
				MakeArg: func() interface{} {
					ret := make([]WaitForClientArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]WaitForClientArg)
					if !ok {
						err = rpc.NewTypeError((*[]WaitForClientArg)(nil), args)
						return
					}
					ret, err = i.WaitForClient(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getBootstrapStatus": {
				MakeArg: func() interface{} {
					ret := make([]GetBootstrapStatusArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetBootstrapStatusArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetBootstrapStatusArg)(nil), args)
						return
					}
					ret, err = i.GetBootstrapStatus(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getRememberPassphrase": {
				MakeArg: func() interface{} {
					ret := make([]GetRememberPassphraseArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetRememberPassphraseArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetRememberPassphraseArg)(nil), args)
						return
					}
					ret, err = i.GetRememberPassphrase(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"setRememberPassphrase": {
				MakeArg: func() interface{} {
					ret := make([]SetRememberPassphraseArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SetRememberPassphraseArg)
					if !ok {
						err = rpc.NewTypeError((*[]SetRememberPassphraseArg)(nil), args)
						return
					}
					err = i.SetRememberPassphrase(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type ConfigClient struct {
	Cli rpc.GenericClient
}

func (c ConfigClient) GetCurrentStatus(ctx context.Context, sessionID int) (res GetCurrentStatusRes, err error) {
	__arg := GetCurrentStatusArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.config.getCurrentStatus", []interface{}{__arg}, &res)
	return
}

func (c ConfigClient) GetExtendedStatus(ctx context.Context, sessionID int) (res ExtendedStatus, err error) {
	__arg := GetExtendedStatusArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.config.getExtendedStatus", []interface{}{__arg}, &res)
	return
}

func (c ConfigClient) GetAllProvisionedUsernames(ctx context.Context, sessionID int) (res AllProvisionedUsernames, err error) {
	__arg := GetAllProvisionedUsernamesArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.config.getAllProvisionedUsernames", []interface{}{__arg}, &res)
	return
}

func (c ConfigClient) GetConfig(ctx context.Context, sessionID int) (res Config, err error) {
	__arg := GetConfigArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.config.getConfig", []interface{}{__arg}, &res)
	return
}

// Change user config.
// For example, to update primary picture source:
// key=picture.source, value=twitter (or github)
func (c ConfigClient) SetUserConfig(ctx context.Context, __arg SetUserConfigArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.config.setUserConfig", []interface{}{__arg}, nil)
	return
}

func (c ConfigClient) SetPath(ctx context.Context, __arg SetPathArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.config.setPath", []interface{}{__arg}, nil)
	return
}

func (c ConfigClient) HelloIAm(ctx context.Context, details ClientDetails) (err error) {
	__arg := HelloIAmArg{Details: details}
	err = c.Cli.Call(ctx, "keybase.1.config.helloIAm", []interface{}{__arg}, nil)
	return
}

func (c ConfigClient) SetValue(ctx context.Context, __arg SetValueArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.config.setValue", []interface{}{__arg}, nil)
	return
}

func (c ConfigClient) ClearValue(ctx context.Context, path string) (err error) {
	__arg := ClearValueArg{Path: path}
	err = c.Cli.Call(ctx, "keybase.1.config.clearValue", []interface{}{__arg}, nil)
	return
}

func (c ConfigClient) GetValue(ctx context.Context, path string) (res ConfigValue, err error) {
	__arg := GetValueArg{Path: path}
	err = c.Cli.Call(ctx, "keybase.1.config.getValue", []interface{}{__arg}, &res)
	return
}

// Check whether the API server has told us we're out of date.
func (c ConfigClient) CheckAPIServerOutOfDateWarning(ctx context.Context) (res OutOfDateInfo, err error) {
	err = c.Cli.Call(ctx, "keybase.1.config.checkAPIServerOutOfDateWarning", []interface{}{CheckAPIServerOutOfDateWarningArg{}}, &res)
	return
}

// Wait for client type to connect to service.
func (c ConfigClient) WaitForClient(ctx context.Context, __arg WaitForClientArg) (res bool, err error) {
	err = c.Cli.Call(ctx, "keybase.1.config.waitForClient", []interface{}{__arg}, &res)
	return
}

func (c ConfigClient) GetBootstrapStatus(ctx context.Context, sessionID int) (res BootstrapStatus, err error) {
	__arg := GetBootstrapStatusArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.config.getBootstrapStatus", []interface{}{__arg}, &res)
	return
}

func (c ConfigClient) GetRememberPassphrase(ctx context.Context, sessionID int) (res bool, err error) {
	__arg := GetRememberPassphraseArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.config.getRememberPassphrase", []interface{}{__arg}, &res)
	return
}

func (c ConfigClient) SetRememberPassphrase(ctx context.Context, __arg SetRememberPassphraseArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.config.setRememberPassphrase", []interface{}{__arg}, nil)
	return
}
