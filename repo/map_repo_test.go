package repo

import (
	"testing"

	e "github.com/brandon-detty/go-api-tools/entity"
)

type SampleRepo Repo[e.SampleEntity, e.SampleEntityId]

func newSampleRepo() SampleRepo {
	return NewMapRepo[e.SampleEntity, e.SampleEntityId]()
}

func TestMapRepoSatisfiesRepo(t *testing.T) {
	var r SampleRepo = newSampleRepo()
	_ = r
}

func TestMapRepoSaveAndGet(t *testing.T) {
	var r SampleRepo = newSampleRepo()

	var id e.SampleEntityId = 42
	in := e.NewSampleEntity()
	in.SetId(id)
	if _, err := r.Save(in); err != nil {
		t.Log("Save() failed")
		t.FailNow()
	}

	out, err := r.Get(id)
	if err != nil {
		t.Log("Get(itemId) failed after Save(item)")
		t.FailNow()
	}

	if in != out {
		t.Log("Get(itemId) didn't return item after Save(item)")
		t.FailNow()
	}
}

func TestMapRepoDelete(t *testing.T) {
	var r SampleRepo = newSampleRepo()

	var id e.SampleEntityId = 42
	in := e.NewSampleEntity()
	in.SetId(id)
	if _, err := r.Save(in); err != nil {
		t.Log("Save() failed")
		t.FailNow()
	}

	r.Delete(id)

	_, err := r.Get(id)
	if err == nil {
		t.Log("Delete(itemId) failed to remove item from repo")
		t.FailNow()
	}
}
