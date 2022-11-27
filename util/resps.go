package util

type respTemplate struct {
	Status int `json:"status"`
	Info string `json:"info"`
}

var OK=respTemplate{
	Status:200,
	Info:"success",
}
