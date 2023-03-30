package entity

// an Entity can be stored in a Repo and is identified by a value of type I
type Entity[I comparable] interface {
	Id() I
	SetId(id I)
}
