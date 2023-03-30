package entity

import "testing"

func TestSampleEntitySatisfiesEntity(t *testing.T) {
	var e Entity[SampleEntityId] = SampleEntity{}
	e = NewSampleEntity()
	_ = e
}

func TestSampleEntitySetAndGetId(t *testing.T) {
	var e Entity[SampleEntityId] = NewSampleEntity()
	var id SampleEntityId = 7
	e.SetId(id)
	if id != e.Id() {
		t.FailNow()
	}
}
