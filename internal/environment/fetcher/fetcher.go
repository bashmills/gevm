package fetcher

import (
	"github.com/bashidogames/gevm/internal/platform"
	"github.com/bashidogames/gevm/internal/repository"
	"github.com/bashidogames/gevm/semver"
)

type Fetcher interface {
	FetchAsset(platform platform.Platform, semver semver.Semver) (*repository.Asset, error)
	FetchDownloads(mono bool) ([]repository.Download, error)
}
