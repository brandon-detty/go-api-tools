package repo

import "testing"

func TestMapRepoSatisfiesRepo(t *testing.T) {
	var r Repo[interface{}, int8] = NewMapRepo[interface{}, int8]()
	_ = r
}

type name struct {
	first string
	last  string
}

func TestMapRepoSetAndGet(t *testing.T) {
	var r Repo[name, int8] = NewMapRepo[name, int8]()

	var id int8 = 42
	in := &name{"Alan", "Atkins"}
	if _, err := r.Set(id, in); err != nil {
		t.Log("Set() failed")
		t.FailNow()
	}

	out, err := r.Get(id)
	if err != nil {
		t.Log("Get(id) failed after Set(id, item)")
		t.FailNow()
	}

	if in != out {
		t.Log("Get(id) didn't return item after Set(id, item)")
		t.FailNow()
	}
}
