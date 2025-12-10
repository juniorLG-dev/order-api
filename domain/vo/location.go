package vo

type Location struct {
	Country    string `json:"country"`
	State      string `json:"state"`
	City       string `json:"city"`
	Complement string `json:"complement"`
	CEP        CEP    `json:"cep"`
}
