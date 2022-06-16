package permission

import (
	"context"
	"fmt"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"sync"
)

type memoryRepository struct {
	locker *sync.RWMutex
	data   map[uint]model.Permission
}

func newMemoryRepository() *memoryRepository {
	return &memoryRepository{
		locker: &sync.RWMutex{},
		data:   make(map[uint]model.Permission),
	}
}

func (repo *memoryRepository) Sets(ctx context.Context, mos model.Permissions) error {
	if mos.Len() == 0 {
		return nil
	}

	if mos.Len() == 1 {
		return repo.Set(ctx, *mos[0])
	}

	for idx, _ := range mos {
		mo := mos[idx]
		repo.data[mo.GetKey()] = *mo
	}
	return nil
}

func (repo *memoryRepository) Set(ctx context.Context, mo model.Permission) error {
	repo.locker.Lock()
	defer repo.locker.Unlock()

	repo.data[mo.GetKey()] = mo
	return nil
}

func (repo *memoryRepository) Get(ctx context.Context, key uint) (mo model.Permission, err error) {
	repo.locker.RLock()
	defer repo.locker.RUnlock()

	val, ok := repo.data[key]
	if !ok {
		err = fmt.Errorf("permission with key %d not found in memory", key)
	}
	return val, err
}

func (repo *memoryRepository) Gets(ctx context.Context, keys ...uint) (mos model.Permissions, err error) {
	repo.locker.RLock()
	defer repo.locker.RUnlock()

	for _, key := range keys {
		val, ok := repo.data[key]
		if !ok {
			err = fmt.Errorf("permission with key %d not found in memory", key)
			return nil, err
		}
		mos = append(mos, &val)
	}

	return mos, err
}

func (repo *memoryRepository) Del(ctx context.Context, keys ...uint) error {
	repo.locker.Lock()
	defer repo.locker.Unlock()

	for _, key := range keys {
		delete(repo.data, key)
	}
	return nil
}

func (repo *memoryRepository) Data(ctx context.Context) map[uint]model.Permission {
	repo.locker.RLock()
	defer repo.locker.RUnlock()

	return repo.data
}
