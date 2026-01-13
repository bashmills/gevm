package repository

import (
	"github.com/bashmills/gevm/internal/platform"
	"github.com/bashmills/gevm/semver"
)

type Download struct {
	Assets map[platform.Platform]Asset
	Relver semver.Relver
}

func (d Download) HasAsset(platform platform.Platform) bool {
	return d.Assets[platform].IsValid()
}

type Asset struct {
	DownloadURL string
	Name        string
}

func (a Asset) IsValid() bool {
	return len(a.DownloadURL) > 0 && len(a.Name) > 0
}
