package sourcegraph

import (
	"sourcegraph.com/sourcegraph/go-sourcegraph/router"
	"sourcegraph.com/sourcegraph/srclib/graph"
)

type DocPagesService interface {
	Get(docPage DocPageSpec, opt *DocPageGetOptions) (*graph.DocPage, Response, error)
}

type DocPageSpec struct {
	RepoRev RepoRevSpec
	Path    string

	// TODO(new-arch): what is the primary key for a doc page? update it here
	// when we figure out the best way to set a primary key for doc pages.
}

func (s *DocPageSpec) RouteVars() map[string]string {
	m := s.RepoRev.RouteVars()
	m["Path"] = s.Path
	return m
}

type docPagesService struct {
	client *Client
}

var _ DocPagesService = &docPagesService{}

type DocPageGetOptions struct{}

func (s *docPagesService) Get(docPage DocPageSpec, opt *DocPageGetOptions) (*graph.DocPage, Response, error) {
	url, err := s.client.url(router.RepoDocPage, docPage.RouteVars(), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	var docPage_ *graph.DocPage
	resp, err := s.client.Do(req, &docPage_)
	if err != nil {
		return nil, resp, err
	}

	return docPage_, resp, nil
}

type MockDocPagesService struct {
	Get_ func(docPage DocPageSpec, opt *DocPageGetOptions) (*graph.DocPage, Response, error)
}

var _ DocPagesService = MockDocPagesService{}

func (s MockDocPagesService) Get(docPage DocPageSpec, opt *DocPageGetOptions) (*graph.DocPage, Response, error) {
	if s.Get_ == nil {
		return nil, &HTTPResponse{}, nil
	}
	return s.Get_(docPage, opt)
}
