package sourcegraph

import "sourcegraph.com/sourcegraph/go-vcs/vcs"

type MockRepoTreeService struct {
	Get_    func(entry TreeEntrySpec, opt *RepoTreeGetOptions) (*TreeEntry, Response, error)
	Search_ func(rev RepoRevSpec, opt *RepoTreeSearchOptions) ([]*vcs.SearchResult, Response, error)
}

func (s *MockRepoTreeService) Get(entry TreeEntrySpec, opt *RepoTreeGetOptions) (*TreeEntry, Response, error) {
	return s.Get_(entry, opt)
}

func (s *MockRepoTreeService) Search(rev RepoRevSpec, opt *RepoTreeSearchOptions) ([]*vcs.SearchResult, Response, error) {
	return s.Search_(rev, opt)
}
