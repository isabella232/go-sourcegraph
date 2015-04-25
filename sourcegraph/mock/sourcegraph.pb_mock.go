// generated by gen-mocks; DO NOT EDIT

package mock

import (
	"golang.org/x/net/context"
	"sourcegraph.com/sourcegraph/go-sourcegraph/sourcegraph"
	"sourcegraph.com/sourcegraph/go-vcs/vcs"
	"sourcegraph.com/sourcegraph/srclib/unit"
	"sourcegraph.com/sqs/pbtypes"
)

type RepoBadgesServer struct {
	ListBadges_   func(v0 context.Context, v1 *sourcegraph.RepoSpec) (*sourcegraph.BadgeList, error)
	ListCounters_ func(v0 context.Context, v1 *sourcegraph.RepoSpec) (*sourcegraph.CounterList, error)
	RecordHit_    func(v0 context.Context, v1 *sourcegraph.RepoSpec) (*pbtypes.Void, error)
	CountHits_    func(v0 context.Context, v1 *sourcegraph.RepoBadgesCountHitsOp) (*sourcegraph.RepoBadgesCountHitsResult, error)
}

func (s *RepoBadgesServer) ListBadges(v0 context.Context, v1 *sourcegraph.RepoSpec) (*sourcegraph.BadgeList, error) {
	return s.ListBadges_(v0, v1)
}

func (s *RepoBadgesServer) ListCounters(v0 context.Context, v1 *sourcegraph.RepoSpec) (*sourcegraph.CounterList, error) {
	return s.ListCounters_(v0, v1)
}

func (s *RepoBadgesServer) RecordHit(v0 context.Context, v1 *sourcegraph.RepoSpec) (*pbtypes.Void, error) {
	return s.RecordHit_(v0, v1)
}

func (s *RepoBadgesServer) CountHits(v0 context.Context, v1 *sourcegraph.RepoBadgesCountHitsOp) (*sourcegraph.RepoBadgesCountHitsResult, error) {
	return s.CountHits_(v0, v1)
}

var _ sourcegraph.RepoBadgesServer = (*RepoBadgesServer)(nil)

type RepoStatusesServer struct {
	GetCombined_ func(v0 context.Context, v1 *sourcegraph.RepoRevSpec) (*sourcegraph.CombinedStatus, error)
	Create_      func(v0 context.Context, v1 *sourcegraph.RepoStatusesCreateOp) (*sourcegraph.RepoStatus, error)
}

func (s *RepoStatusesServer) GetCombined(v0 context.Context, v1 *sourcegraph.RepoRevSpec) (*sourcegraph.CombinedStatus, error) {
	return s.GetCombined_(v0, v1)
}

func (s *RepoStatusesServer) Create(v0 context.Context, v1 *sourcegraph.RepoStatusesCreateOp) (*sourcegraph.RepoStatus, error) {
	return s.Create_(v0, v1)
}

var _ sourcegraph.RepoStatusesServer = (*RepoStatusesServer)(nil)

type ReposServer struct {
	Get_          func(v0 context.Context, v1 *sourcegraph.RepoSpec) (*sourcegraph.Repo, error)
	List_         func(v0 context.Context, v1 *sourcegraph.RepoListOptions) (*sourcegraph.RepoList, error)
	GetReadme_    func(v0 context.Context, v1 *sourcegraph.RepoRevSpec) (*sourcegraph.Readme, error)
	Enable_       func(v0 context.Context, v1 *sourcegraph.RepoSpec) (*pbtypes.Void, error)
	Disable_      func(v0 context.Context, v1 *sourcegraph.RepoSpec) (*pbtypes.Void, error)
	GetCommit_    func(v0 context.Context, v1 *sourcegraph.RepoRevSpec) (*vcs.Commit, error)
	ListCommits_  func(v0 context.Context, v1 *sourcegraph.ReposListCommitsOp) (*sourcegraph.CommitList, error)
	ListBranches_ func(v0 context.Context, v1 *sourcegraph.ReposListBranchesOp) (*sourcegraph.BranchList, error)
	ListTags_     func(v0 context.Context, v1 *sourcegraph.ReposListTagsOp) (*sourcegraph.TagList, error)
}

func (s *ReposServer) Get(v0 context.Context, v1 *sourcegraph.RepoSpec) (*sourcegraph.Repo, error) {
	return s.Get_(v0, v1)
}

func (s *ReposServer) List(v0 context.Context, v1 *sourcegraph.RepoListOptions) (*sourcegraph.RepoList, error) {
	return s.List_(v0, v1)
}

func (s *ReposServer) GetReadme(v0 context.Context, v1 *sourcegraph.RepoRevSpec) (*sourcegraph.Readme, error) {
	return s.GetReadme_(v0, v1)
}

func (s *ReposServer) Enable(v0 context.Context, v1 *sourcegraph.RepoSpec) (*pbtypes.Void, error) {
	return s.Enable_(v0, v1)
}

func (s *ReposServer) Disable(v0 context.Context, v1 *sourcegraph.RepoSpec) (*pbtypes.Void, error) {
	return s.Disable_(v0, v1)
}

func (s *ReposServer) GetCommit(v0 context.Context, v1 *sourcegraph.RepoRevSpec) (*vcs.Commit, error) {
	return s.GetCommit_(v0, v1)
}

func (s *ReposServer) ListCommits(v0 context.Context, v1 *sourcegraph.ReposListCommitsOp) (*sourcegraph.CommitList, error) {
	return s.ListCommits_(v0, v1)
}

func (s *ReposServer) ListBranches(v0 context.Context, v1 *sourcegraph.ReposListBranchesOp) (*sourcegraph.BranchList, error) {
	return s.ListBranches_(v0, v1)
}

func (s *ReposServer) ListTags(v0 context.Context, v1 *sourcegraph.ReposListTagsOp) (*sourcegraph.TagList, error) {
	return s.ListTags_(v0, v1)
}

var _ sourcegraph.ReposServer = (*ReposServer)(nil)

type MirrorReposServer struct {
	RefreshVCS_ func(v0 context.Context, v1 *sourcegraph.RepoSpec) (*pbtypes.Void, error)
}

func (s *MirrorReposServer) RefreshVCS(v0 context.Context, v1 *sourcegraph.RepoSpec) (*pbtypes.Void, error) {
	return s.RefreshVCS_(v0, v1)
}

var _ sourcegraph.MirrorReposServer = (*MirrorReposServer)(nil)

type BuildsServer struct {
	Get_              func(v0 context.Context, v1 *sourcegraph.BuildSpec) (*sourcegraph.Build, error)
	GetRepoBuildInfo_ func(v0 context.Context, v1 *sourcegraph.BuildsGetRepoBuildInfoOp) (*sourcegraph.RepoBuildInfo, error)
	List_             func(v0 context.Context, v1 *sourcegraph.BuildListOptions) (*sourcegraph.BuildList, error)
	Create_           func(v0 context.Context, v1 *sourcegraph.BuildsCreateOp) (*sourcegraph.Build, error)
	Update_           func(v0 context.Context, v1 *sourcegraph.BuildsUpdateOp) (*sourcegraph.Build, error)
	ListBuildTasks_   func(v0 context.Context, v1 *sourcegraph.BuildsListBuildTasksOp) (*sourcegraph.BuildTaskList, error)
	CreateTasks_      func(v0 context.Context, v1 *sourcegraph.BuildsCreateTasksOp) (*sourcegraph.BuildTaskList, error)
	UpdateTask_       func(v0 context.Context, v1 *sourcegraph.BuildsUpdateTaskOp) (*sourcegraph.BuildTask, error)
	GetLog_           func(v0 context.Context, v1 *sourcegraph.BuildsGetLogOp) (*sourcegraph.LogEntries, error)
	GetTaskLog_       func(v0 context.Context, v1 *sourcegraph.BuildsGetTaskLogOp) (*sourcegraph.LogEntries, error)
	DequeueNext_      func(v0 context.Context, v1 *sourcegraph.BuildsDequeueNextOp) (*sourcegraph.Build, error)
}

func (s *BuildsServer) Get(v0 context.Context, v1 *sourcegraph.BuildSpec) (*sourcegraph.Build, error) {
	return s.Get_(v0, v1)
}

func (s *BuildsServer) GetRepoBuildInfo(v0 context.Context, v1 *sourcegraph.BuildsGetRepoBuildInfoOp) (*sourcegraph.RepoBuildInfo, error) {
	return s.GetRepoBuildInfo_(v0, v1)
}

func (s *BuildsServer) List(v0 context.Context, v1 *sourcegraph.BuildListOptions) (*sourcegraph.BuildList, error) {
	return s.List_(v0, v1)
}

func (s *BuildsServer) Create(v0 context.Context, v1 *sourcegraph.BuildsCreateOp) (*sourcegraph.Build, error) {
	return s.Create_(v0, v1)
}

func (s *BuildsServer) Update(v0 context.Context, v1 *sourcegraph.BuildsUpdateOp) (*sourcegraph.Build, error) {
	return s.Update_(v0, v1)
}

func (s *BuildsServer) ListBuildTasks(v0 context.Context, v1 *sourcegraph.BuildsListBuildTasksOp) (*sourcegraph.BuildTaskList, error) {
	return s.ListBuildTasks_(v0, v1)
}

func (s *BuildsServer) CreateTasks(v0 context.Context, v1 *sourcegraph.BuildsCreateTasksOp) (*sourcegraph.BuildTaskList, error) {
	return s.CreateTasks_(v0, v1)
}

func (s *BuildsServer) UpdateTask(v0 context.Context, v1 *sourcegraph.BuildsUpdateTaskOp) (*sourcegraph.BuildTask, error) {
	return s.UpdateTask_(v0, v1)
}

func (s *BuildsServer) GetLog(v0 context.Context, v1 *sourcegraph.BuildsGetLogOp) (*sourcegraph.LogEntries, error) {
	return s.GetLog_(v0, v1)
}

func (s *BuildsServer) GetTaskLog(v0 context.Context, v1 *sourcegraph.BuildsGetTaskLogOp) (*sourcegraph.LogEntries, error) {
	return s.GetTaskLog_(v0, v1)
}

func (s *BuildsServer) DequeueNext(v0 context.Context, v1 *sourcegraph.BuildsDequeueNextOp) (*sourcegraph.Build, error) {
	return s.DequeueNext_(v0, v1)
}

var _ sourcegraph.BuildsServer = (*BuildsServer)(nil)

type OrgsServer struct {
	Get_         func(v0 context.Context, v1 *sourcegraph.OrgSpec) (*sourcegraph.Org, error)
	List_        func(v0 context.Context, v1 *sourcegraph.OrgsListOp) (*sourcegraph.OrgList, error)
	ListMembers_ func(v0 context.Context, v1 *sourcegraph.OrgsListMembersOp) (*sourcegraph.UserList, error)
}

func (s *OrgsServer) Get(v0 context.Context, v1 *sourcegraph.OrgSpec) (*sourcegraph.Org, error) {
	return s.Get_(v0, v1)
}

func (s *OrgsServer) List(v0 context.Context, v1 *sourcegraph.OrgsListOp) (*sourcegraph.OrgList, error) {
	return s.List_(v0, v1)
}

func (s *OrgsServer) ListMembers(v0 context.Context, v1 *sourcegraph.OrgsListMembersOp) (*sourcegraph.UserList, error) {
	return s.ListMembers_(v0, v1)
}

var _ sourcegraph.OrgsServer = (*OrgsServer)(nil)

type PeopleServer struct {
	Get_ func(v0 context.Context, v1 *sourcegraph.PersonSpec) (*sourcegraph.Person, error)
}

func (s *PeopleServer) Get(v0 context.Context, v1 *sourcegraph.PersonSpec) (*sourcegraph.Person, error) {
	return s.Get_(v0, v1)
}

var _ sourcegraph.PeopleServer = (*PeopleServer)(nil)

type UsersServer struct {
	Get_        func(v0 context.Context, v1 *sourcegraph.UserSpec) (*sourcegraph.User, error)
	ListEmails_ func(v0 context.Context, v1 *sourcegraph.UserSpec) (*sourcegraph.EmailAddrList, error)
	List_       func(v0 context.Context, v1 *sourcegraph.UsersListOptions) (*sourcegraph.UserList, error)
}

func (s *UsersServer) Get(v0 context.Context, v1 *sourcegraph.UserSpec) (*sourcegraph.User, error) {
	return s.Get_(v0, v1)
}

func (s *UsersServer) ListEmails(v0 context.Context, v1 *sourcegraph.UserSpec) (*sourcegraph.EmailAddrList, error) {
	return s.ListEmails_(v0, v1)
}

func (s *UsersServer) List(v0 context.Context, v1 *sourcegraph.UsersListOptions) (*sourcegraph.UserList, error) {
	return s.List_(v0, v1)
}

var _ sourcegraph.UsersServer = (*UsersServer)(nil)

type UserAuthServer struct {
	Authenticate_ func(v0 context.Context, v1 *sourcegraph.UserAuthAuthenticateOp) (*sourcegraph.AuthenticatedUser, error)
	GetExternal_  func(v0 context.Context, v1 *sourcegraph.UserAuthGetExternalOp) (*sourcegraph.ExternalAuthInfo, error)
	Identify_     func(v0 context.Context, v1 *pbtypes.Void) (*sourcegraph.AuthInfo, error)
}

func (s *UserAuthServer) Authenticate(v0 context.Context, v1 *sourcegraph.UserAuthAuthenticateOp) (*sourcegraph.AuthenticatedUser, error) {
	return s.Authenticate_(v0, v1)
}

func (s *UserAuthServer) GetExternal(v0 context.Context, v1 *sourcegraph.UserAuthGetExternalOp) (*sourcegraph.ExternalAuthInfo, error) {
	return s.GetExternal_(v0, v1)
}

func (s *UserAuthServer) Identify(v0 context.Context, v1 *pbtypes.Void) (*sourcegraph.AuthInfo, error) {
	return s.Identify_(v0, v1)
}

var _ sourcegraph.UserAuthServer = (*UserAuthServer)(nil)

type DefsServer struct {
	Get_          func(v0 context.Context, v1 *sourcegraph.DefsGetOp) (*sourcegraph.Def, error)
	List_         func(v0 context.Context, v1 *sourcegraph.DefListOptions) (*sourcegraph.DefList, error)
	ListRefs_     func(v0 context.Context, v1 *sourcegraph.DefsListRefsOp) (*sourcegraph.RefList, error)
	ListExamples_ func(v0 context.Context, v1 *sourcegraph.DefsListExamplesOp) (*sourcegraph.ExampleList, error)
	ListAuthors_  func(v0 context.Context, v1 *sourcegraph.DefsListAuthorsOp) (*sourcegraph.DefAuthorList, error)
	ListClients_  func(v0 context.Context, v1 *sourcegraph.DefsListClientsOp) (*sourcegraph.DefClientList, error)
}

func (s *DefsServer) Get(v0 context.Context, v1 *sourcegraph.DefsGetOp) (*sourcegraph.Def, error) {
	return s.Get_(v0, v1)
}

func (s *DefsServer) List(v0 context.Context, v1 *sourcegraph.DefListOptions) (*sourcegraph.DefList, error) {
	return s.List_(v0, v1)
}

func (s *DefsServer) ListRefs(v0 context.Context, v1 *sourcegraph.DefsListRefsOp) (*sourcegraph.RefList, error) {
	return s.ListRefs_(v0, v1)
}

func (s *DefsServer) ListExamples(v0 context.Context, v1 *sourcegraph.DefsListExamplesOp) (*sourcegraph.ExampleList, error) {
	return s.ListExamples_(v0, v1)
}

func (s *DefsServer) ListAuthors(v0 context.Context, v1 *sourcegraph.DefsListAuthorsOp) (*sourcegraph.DefAuthorList, error) {
	return s.ListAuthors_(v0, v1)
}

func (s *DefsServer) ListClients(v0 context.Context, v1 *sourcegraph.DefsListClientsOp) (*sourcegraph.DefClientList, error) {
	return s.ListClients_(v0, v1)
}

var _ sourcegraph.DefsServer = (*DefsServer)(nil)

type DeltasServer struct {
	Get_                 func(v0 context.Context, v1 *sourcegraph.DeltaSpec) (*sourcegraph.Delta, error)
	ListUnits_           func(v0 context.Context, v1 *sourcegraph.DeltasListUnitsOp) (*sourcegraph.UnitDeltaList, error)
	ListDefs_            func(v0 context.Context, v1 *sourcegraph.DeltasListDefsOp) (*sourcegraph.DeltaDefs, error)
	ListFiles_           func(v0 context.Context, v1 *sourcegraph.DeltasListFilesOp) (*sourcegraph.DeltaFiles, error)
	ListAffectedAuthors_ func(v0 context.Context, v1 *sourcegraph.DeltasListAffectedAuthorsOp) (*sourcegraph.DeltaAffectedPersonList, error)
	ListAffectedClients_ func(v0 context.Context, v1 *sourcegraph.DeltasListAffectedClientsOp) (*sourcegraph.DeltaAffectedPersonList, error)
}

func (s *DeltasServer) Get(v0 context.Context, v1 *sourcegraph.DeltaSpec) (*sourcegraph.Delta, error) {
	return s.Get_(v0, v1)
}

func (s *DeltasServer) ListUnits(v0 context.Context, v1 *sourcegraph.DeltasListUnitsOp) (*sourcegraph.UnitDeltaList, error) {
	return s.ListUnits_(v0, v1)
}

func (s *DeltasServer) ListDefs(v0 context.Context, v1 *sourcegraph.DeltasListDefsOp) (*sourcegraph.DeltaDefs, error) {
	return s.ListDefs_(v0, v1)
}

func (s *DeltasServer) ListFiles(v0 context.Context, v1 *sourcegraph.DeltasListFilesOp) (*sourcegraph.DeltaFiles, error) {
	return s.ListFiles_(v0, v1)
}

func (s *DeltasServer) ListAffectedAuthors(v0 context.Context, v1 *sourcegraph.DeltasListAffectedAuthorsOp) (*sourcegraph.DeltaAffectedPersonList, error) {
	return s.ListAffectedAuthors_(v0, v1)
}

func (s *DeltasServer) ListAffectedClients(v0 context.Context, v1 *sourcegraph.DeltasListAffectedClientsOp) (*sourcegraph.DeltaAffectedPersonList, error) {
	return s.ListAffectedClients_(v0, v1)
}

var _ sourcegraph.DeltasServer = (*DeltasServer)(nil)

type MarkdownServer struct {
	Render_ func(v0 context.Context, v1 *sourcegraph.MarkdownRenderOp) (*sourcegraph.MarkdownData, error)
}

func (s *MarkdownServer) Render(v0 context.Context, v1 *sourcegraph.MarkdownRenderOp) (*sourcegraph.MarkdownData, error) {
	return s.Render_(v0, v1)
}

var _ sourcegraph.MarkdownServer = (*MarkdownServer)(nil)

type RepoTreeServer struct {
	Get_    func(v0 context.Context, v1 *sourcegraph.RepoTreeGetOp) (*sourcegraph.TreeEntry, error)
	Search_ func(v0 context.Context, v1 *sourcegraph.RepoTreeSearchOp) (*sourcegraph.VCSSearchResultList, error)
}

func (s *RepoTreeServer) Get(v0 context.Context, v1 *sourcegraph.RepoTreeGetOp) (*sourcegraph.TreeEntry, error) {
	return s.Get_(v0, v1)
}

func (s *RepoTreeServer) Search(v0 context.Context, v1 *sourcegraph.RepoTreeSearchOp) (*sourcegraph.VCSSearchResultList, error) {
	return s.Search_(v0, v1)
}

var _ sourcegraph.RepoTreeServer = (*RepoTreeServer)(nil)

type SearchServer struct {
	Search_   func(v0 context.Context, v1 *sourcegraph.SearchOptions) (*sourcegraph.SearchResults, error)
	Complete_ func(v0 context.Context, v1 *sourcegraph.RawQuery) (*sourcegraph.Completions, error)
	Suggest_  func(v0 context.Context, v1 *sourcegraph.RawQuery) (*sourcegraph.SuggestionList, error)
}

func (s *SearchServer) Search(v0 context.Context, v1 *sourcegraph.SearchOptions) (*sourcegraph.SearchResults, error) {
	return s.Search_(v0, v1)
}

func (s *SearchServer) Complete(v0 context.Context, v1 *sourcegraph.RawQuery) (*sourcegraph.Completions, error) {
	return s.Complete_(v0, v1)
}

func (s *SearchServer) Suggest(v0 context.Context, v1 *sourcegraph.RawQuery) (*sourcegraph.SuggestionList, error) {
	return s.Suggest_(v0, v1)
}

var _ sourcegraph.SearchServer = (*SearchServer)(nil)

type UnitsServer struct {
	Get_  func(v0 context.Context, v1 *sourcegraph.UnitSpec) (*unit.RepoSourceUnit, error)
	List_ func(v0 context.Context, v1 *sourcegraph.UnitListOptions) (*sourcegraph.RepoSourceUnitList, error)
}

func (s *UnitsServer) Get(v0 context.Context, v1 *sourcegraph.UnitSpec) (*unit.RepoSourceUnit, error) {
	return s.Get_(v0, v1)
}

func (s *UnitsServer) List(v0 context.Context, v1 *sourcegraph.UnitListOptions) (*sourcegraph.RepoSourceUnitList, error) {
	return s.List_(v0, v1)
}

var _ sourcegraph.UnitsServer = (*UnitsServer)(nil)
