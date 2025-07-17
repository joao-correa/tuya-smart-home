package api

type Device struct {
	ID         string `json:"id"`
	IP         string `json:"ip"`
	ProductKey string `json:"productKey"`
	Key        string `json:"key"`
	MAC        string `json:"mac"`
	Version    string `json:"ver"`
}

type Devices struct {
	Devices []Device `json:"devices"`
}
