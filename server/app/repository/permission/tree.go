package permission

import (
	"context"
	"fmt"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"sort"
	"sync"
)

type treeRepository struct {
	data     map[uint]*model.Permission
	treeData *model.Permissions
	locker   *sync.RWMutex
}

func newTreeRepository() *treeRepository {
	return &treeRepository{
		data:     make(map[uint]*model.Permission, 0),
		treeData: new(model.Permissions),
		locker:   &sync.RWMutex{},
	}
}

func (repo *treeRepository) Set(ctx context.Context, permissions ...*model.Permission) error {
	repo.locker.Lock()
	defer repo.locker.Unlock()

	for _, permission := range permissions {
		repo.data[permission.GetKey()] = permission
	}
	repo.makeTreeData()
	return nil
}

func (repo *treeRepository) Get(ctx context.Context, key uint) (*model.Permission, error) {
	repo.locker.RLock()
	defer repo.locker.RUnlock()

	val, ok := repo.data[key]
	if !ok {
		return nil, fmt.Errorf("not found Permissions data from cache")
	}
	return val, nil
}

func (repo *treeRepository) Del(ctx context.Context, keys ...uint) error {
	repo.locker.Lock()
	defer repo.locker.Unlock()

	for _, key := range keys {
		delete(repo.data, key)
	}
	repo.makeTreeData()
	return nil
}

func (repo *treeRepository) Data(ctx context.Context) map[uint]*model.Permission {
	repo.locker.RLock()
	defer repo.locker.RUnlock()

	return repo.data
}

func (repo *treeRepository) TreeData() model.Permissions {
	return *repo.treeData
}

func (repo *treeRepository) makeTreeData() {
	permissions := repo.data
	refer := make(map[uint]*model.Permission, 0)
	tree := make(model.Permissions, 0)
	for idx, Permissions := range permissions {
		val := *(permissions)[idx]
		refer[Permissions.ID] = &val
	}

	for idx, permission := range refer {
		pid := *permission.ParentId
		if pid == 0 {
			tree = append(tree, (refer)[idx])
		} else {
			if _, ok := refer[pid]; ok {
				if refer[pid].Children == nil {
					refer[pid].Children = &model.Permissions{}
				}
				*refer[pid].Children = append(*refer[pid].Children, (refer)[idx])
			}
		}
	}
	repo.sortTreeData(&tree)
	repo.treeData = &tree
}

func (repo *treeRepository) sortTreeData(tree *model.Permissions) {
	sort.SliceStable(*tree, func(i, j int) bool {
		return *((*tree)[i].SortNum) > *((*tree)[j].SortNum) || (*tree)[i].ID > (*tree)[j].ID
	})

	for _, Permissions := range *tree {
		if Permissions.Children != nil && len(*Permissions.Children) > 0 {
			repo.sortTreeData(Permissions.Children)
		}
	}
}
