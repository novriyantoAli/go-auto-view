package model

/**
 */
type ChannelCodeResponse struct {
	ChannelCode string `json:"channelCode"`
}

/**
* ProfileDetail
 */
type ProfileDetail struct {
	ProfileName string `json:"profileName"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	ChannelCode string `json:"channelCode"`
}

/**
* Profile
 */
type Profile struct {
	Prefix string          `json:"prefix"`
	Split  string          `json:"split"`
	Detail []ProfileDetail `json:"detail"`
}

/**
* openBrowserCommand
 */
type OpenBrowserCommand struct {
	ApplicationName string   `json:"applicationName"`
	Arguments       []string `json:"arguments"`
}

/**
* InstallationBrowserCommand
 */
type InstallationBrowserCommand struct {
	ApplicationName string   `json:"applicationName"`
	Arguments       []string `json:"arguments"`
}

/**
* InstallationBrowserCommand
 */
type KillBrowserCommand struct {
	ApplicationName string   `json:"applicationName"`
	Arguments       []string `json:"arguments"`
}

/**
*
 */
type Step struct {
	WaitElement string `json:"waitElement"`
	Url         string `json:"url"`
}

/**
* JavascriptConfig
 */
type JavascriptConfig struct {
	Mode         int    `json:"mode"`
	Keyword      string `json:"keyword"`
	ChannelName  string `json:"channelName"`
	ChannelCode  string `json:"channelCode"`
	PlaylistName string `json:"playlistName"`
	PlaylistCode string `json:"playlistCode"`
	Step         []Step `json:"step"`
}

/**
* WifiConfig
 */

type WifiConfig struct {
	Active         bool   `json:"active"`
	ConnectionName string `json:"connectionName"`
	Password       string `json:"password"`
}

type Config struct {
	WiCon                      WifiConfig                 `json:"wifiConfig"`
	SourceProfile              string                     `json:"sourceProfile"`
	DestinationProfile         string                     `json:"destinationProfile"`
	Profile                    Profile                    `json:"profile"`
	ProxyList                  []string                   `json:"proxyList"`
	OpenBrowserCommand         OpenBrowserCommand         `json:"openBrowserCommand"`
	InstallationBrowserCommand InstallationBrowserCommand `json:"installationBrowserCommand"`
	KillBrowserCommand         KillBrowserCommand         `json:"killBrowserCommand"`
	JobName                    string                     `json:"jobName"`
	MaxTime                    int64                      `json:"maxTime"`
	Port                       string                     `json:"port"`
	CloneBrowser               int                        `json:"cloneBrowser"`
	JavascriptConfig           JavascriptConfig           `json:"javascriptConfig"`
}

type Credentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type Stp struct {
	Username string `json:"username"`
	Command  string `json:"command"`
}

type Process struct {
	ProfileDetail ProfileDetail
	ProcessID     int
}
