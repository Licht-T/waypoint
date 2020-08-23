package state

import (
	"strings"

	"github.com/hashicorp/go-memdb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/hashicorp/waypoint/internal/server/gen"
)

func init() {
	schemas = append(schemas, workspaceIndexSchema)
}

// WorkspaceList lists all the workspaces.
func (s *State) WorkspaceList() ([]*pb.Workspace, error) {
	memTxn := s.inmem.Txn(false)
	defer memTxn.Abort()

	iter, err := memTxn.Get(
		workspaceIndexTableName,
		workspaceIndexIdIndexName+"_prefix",
		"",
	)
	if err != nil {
		return nil, err
	}

	wsMap := map[string]*pb.Workspace{}
	for {
		raw := iter.Next()
		if raw == nil {
			break
		}

		idx := raw.(*workspaceIndexRecord)

		// Get our workspace
		ws, ok := wsMap[idx.Name]
		if !ok {
			ws = &pb.Workspace{Name: idx.Name}
			wsMap[idx.Name] = ws
		}

		// Append our apps
		ws.Applications = append(ws.Applications, &pb.Ref_Application{
			Project:     idx.Project,
			Application: idx.App,
		})
	}

	result := make([]*pb.Workspace, 0, len(wsMap))
	for _, v := range wsMap {
		result = append(result, v)
	}

	return result, nil
}

// WorkspaceGet gets a workspace with a specific name. If it doesn't exist,
// this will return an error with codes.NotFound.
func (s *State) WorkspaceGet(n string) (*pb.Workspace, error) {
	// We implement this in terms of list for now.
	wsList, err := s.WorkspaceList()
	if err != nil {
		return nil, err
	}

	for _, ws := range wsList {
		if strings.EqualFold(ws.Name, n) {
			return ws, nil
		}
	}

	return nil, status.Errorf(codes.NotFound,
		"not found for name: %q", n)
}

// workspaceInit creates an initial record for the given workspace or
// returns one if it already exists.
func (s *State) workspaceInit(
	memTxn *memdb.Txn,
	ref *pb.Ref_Workspace,
	app *pb.Ref_Application,
) (*workspaceIndexRecord, error) {
	rec, err := s.workspaceGet(memTxn, ref, app)
	if err != nil {
		return nil, err
	}
	if rec != nil {
		return rec, nil
	}

	rec = &workspaceIndexRecord{
		Name:    ref.Workspace,
		Project: app.Project,
		App:     app.Application,
	}
	return rec, s.workspacePut(memTxn, rec)
}

// workspacePut updates the workspace record.
func (s *State) workspacePut(
	memTxn *memdb.Txn,
	rec *workspaceIndexRecord,
) error {
	return memTxn.Insert(workspaceIndexTableName, rec)
}

func (s *State) workspaceGet(
	memTxn *memdb.Txn,
	ref *pb.Ref_Workspace,
	app *pb.Ref_Application,
) (*workspaceIndexRecord, error) {
	raw, err := memTxn.First(
		workspaceIndexTableName,
		workspaceIndexIdIndexName,
		ref.Workspace,
		app.Project,
		app.Application,
	)
	if err != nil {
		return nil, err
	}
	if raw == nil {
		return nil, nil
	}

	return raw.(*workspaceIndexRecord), nil
}

func workspaceIndexSchema() *memdb.TableSchema {
	return &memdb.TableSchema{
		Name: workspaceIndexTableName,
		Indexes: map[string]*memdb.IndexSchema{
			workspaceIndexIdIndexName: &memdb.IndexSchema{
				Name:         workspaceIndexIdIndexName,
				AllowMissing: false,
				Unique:       true,
				Indexer: &memdb.CompoundIndex{
					Indexes: []memdb.Indexer{
						&memdb.StringFieldIndex{
							Field:     "Name",
							Lowercase: true,
						},

						&memdb.StringFieldIndex{
							Field:     "Project",
							Lowercase: true,
						},

						&memdb.StringFieldIndex{
							Field:     "App",
							Lowercase: true,
						},
					},
				},
			},
		},
	}
}

const (
	workspaceIndexTableName   = "workspace-index"
	workspaceIndexIdIndexName = "id"
)

type workspaceIndexRecord struct {
	Name    string
	Project string
	App     string
}
