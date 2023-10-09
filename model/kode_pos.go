package model

var BaseUrl = "https://kodepos-2d475.firebaseio.com/kota_kab/k69.json"

type KodePos struct {
	Kecamatan string
	Kelurahan string
	KodePos   string
}
