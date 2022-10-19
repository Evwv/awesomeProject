package models

import (
	"fmt"
	"sync"
	"time"
)

const (
	Created    = "created"
	InProgress = "in progress"
	Failed     = "failed"
	Finished   = "finished"
)

type Build struct {
	Id                int
	Name              string
	Status            string
	CreatedDate       time.Time
	StatusChangedDate time.Time
	FinishDate        time.Time
}

type BuildStore struct {
	sync.Mutex
	builds map[int]Build
	nextId int
}

func New() *BuildStore {
	bs := &BuildStore{}
	bs.builds = make(map[int]Build)
	bs.nextId = 0
	return bs
}

func (bs *BuildStore) Save(name string) int {
	bs.Lock()
	defer bs.Unlock()

	build := Build{
		Id:                bs.nextId,
		Name:              name,
		CreatedDate:       time.Now(),
		StatusChangedDate: time.Now(),
		FinishDate:        time.Now(),
		Status:            Created,
	}
	bs.builds[bs.nextId] = build
	bs.nextId++
	return build.Id
}

func (bs *BuildStore) GetBuildById(id int) (Build, error) {
	bs.Lock()
	defer bs.Unlock()

	build, ok := bs.builds[id]
	if ok {
		return build, nil
	} else {
		return Build{}, fmt.Errorf("build with id=%d not found", id)
	}
}

func (bs *BuildStore) GetAllBuildsFromQuery() []Build {
	bs.Lock()
	defer bs.Unlock()

	allBuilds := make([]Build, 0, len(bs.builds))
	for _, build := range bs.builds {
		allBuilds = append(allBuilds, build)
	}
	return allBuilds
}
