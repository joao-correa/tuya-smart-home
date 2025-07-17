package api

type Scene struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Dps       Dps      `json:"dps"`
	DeviceIds []string `json:"deviceIds"`
}

type Scenes struct {
	Scenes map[string]Scene `json:"scenes"`
}

type Dps struct {
	Switch     bool   `json:"switch"`
	Mode       string `json:"mode"`
	Brightness int    `json:"brightness"`
	ColourTemp int    `json:"colourtemp"`
	Colour     string `json:"colour"`
	Scene      string `json:"scene"`
	SceneData  string `json:"sceneData"`
	Timer      string `json:"timer"`
	Music      string `json:"music"`
	ValueMin   string `json:"valueMin"`
	ValueMax   string `json:"valueMax"`
	ValueHex   string `json:"valueHex"`
}
