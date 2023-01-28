package repo

import (
	"log"
	"reflect"
)

type IRepository interface {
	SetupRepo()
}

var Context = &repoContext{}

type repoContext struct {
	repos []IRepository
}

func (ctx *repoContext) RegisterRepo(repo IRepository) {
	ctx.repos = append(ctx.repos, repo)
}

func (ctx *repoContext) Setup() {
	for _, repo := range ctx.repos {
		repo.SetupRepo()
		repoName := reflect.TypeOf(repo).Elem().Name()
		log.Printf("[Context.SetupRepo]: %s repo registered successfully\n", repoName)
	}
}
