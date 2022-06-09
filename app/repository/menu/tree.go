package menu

import (
	"context"
	"fmt"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"sort"
	"sync"
)

type treeRepository struct {
	data     map[uint]*model.Menu
	treeData *model.Menus
	locker   *sync.RWMutex
}

func newTreeRepository() *treeRepository {
	return &treeRepository{
		data:     make(map[uint]*model.Menu, 0),
		treeData: new(model.Menus),
		locker:   &sync.RWMutex{},
	}
}

func (repo *treeRepository) Set(ctx context.Context, menus ...*model.Menu) error {
	repo.locker.Lock()
	defer repo.locker.Unlock()

	for _, menu := range menus {
		repo.data[menu.GetKey()] = menu
	}
	repo.makeTreeData()
	return nil
}

func (repo *treeRepository) Get(ctx context.Context, key uint) (*model.Menu, error) {
	repo.locker.RLock()
	defer repo.locker.RUnlock()

	val, ok := repo.data[key]
	if !ok {
		return nil, fmt.Errorf("not found menu data from cache")
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

func (repo *treeRepository) Data(ctx context.Context) map[uint]*model.Menu {
	repo.locker.RLock()
	defer repo.locker.RUnlock()

	return repo.data
}

func (repo *treeRepository) TreeData() model.Menus {
	return *repo.treeData
}

func (repo *treeRepository) makeTreeData() {
	menus := repo.data
	refer := make(map[uint]*model.Menu, 0)
	tree := make(model.Menus, 0)
	for idx, menu := range menus {
		val := *(menus)[idx]
		refer[menu.ID] = &val
	}

	for idx, menu := range menus {
		pid := *menu.ParentId
		if pid == 0 {
			tree = append(tree, (refer)[idx])
		} else {
			if _, ok := refer[pid]; ok {
				if refer[pid].Children == nil {
					refer[pid].Children = &model.Menus{}
				}
				*refer[pid].Children = append(*refer[pid].Children, (refer)[idx])
			}
		}
	}
	repo.sortTreeData(&tree)
	repo.treeData = &tree
}

func (repo *treeRepository) sortTreeData(tree *model.Menus) {
	sort.SliceStable(*tree, func(i, j int) bool {
		return *((*tree)[i].SortNum) > *((*tree)[j].SortNum) || (*tree)[i].ID > (*tree)[j].ID
	})

	for _, menu := range *tree {
		if menu.Children != nil && len(*menu.Children) > 0 {
			repo.sortTreeData(menu.Children)
		}
	}
}
