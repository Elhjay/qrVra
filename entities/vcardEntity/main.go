package vcardentity

type VcardEntity struct {
	Pid        string `json:"pid"`
	Name       string `json:"name"`
	Cell       string `json:"cell"`
	Work       string `json:"work"`
	Email      string `json:"email"`
	Department string `json:"department"`
	Section    string `json:"section"`
}

type FinalUsers struct {
	Pid        string `json:"pid"`
	Name       string `json:"name"`
	Cell       string `json:"cell"`
	Work       string `json:"work"`
	Email      string `json:"email"`
	Department string `json:"department"`
	Section    string `json:"section"`
	Link       string `json:"link"`
	QrLink     string `json:"qrlink"`
}
