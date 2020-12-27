package models

type BlockInfo struct {
	Hash string `json:"hash"`
	Size float64 `json:"size"`
	Height float64 `json:"height"`
	Version float64 `json:"version"`
}
