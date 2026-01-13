package fetcher

import (
	"github.com/bashmills/gevm/internal/platform"
	"github.com/bashmills/gevm/internal/repository"
	"github.com/bashmills/gevm/semver"
)

type Fetcher interface {
	FetchAsset(platform platform.Platform, semver semver.Semver) (*repository.Asset, error)
	FetchDownloads(mono bool) ([]repository.Download, error)
}
