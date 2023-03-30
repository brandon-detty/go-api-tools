package entity

type SampleEntityId int8

type SampleEntity struct {
	id   *SampleEntityId
	Name string
}

func NewSampleEntity() *SampleEntity {
	return &SampleEntity{
		id: new(SampleEntityId),
	}
}

func (e SampleEntity) Id() SampleEntityId {
	return *e.id
}

func (e SampleEntity) SetId(id SampleEntityId) {
	*e.id = id
}
