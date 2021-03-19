package auto

import "go.mongodb.org/mongo-driver/bson/primitive"

// * AUTO-DEBIT * //

type Entity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `json:"name" binding:"required,min=2" bson:"name"`
	Amount    string             `json:"amount" binding:"required,min=2" bson:"amount"`
	Frequency string             `json:"frequency" binding:"required,min=2" bson:"frequency"`
	Status    string             `json:"status" binding:"required,min=2" bson:"status"`
}

// Repository interface allows us to access the CRUD Operations of the database.
type Repository interface {
	Create(Entity) (Entity, error)
	ReadAll() ([]Entity, error)
	ReadAllWithStatus(status string) ([]Entity, error)
	ReadOneWithID(id string) (Entity, error)
	ReadOneWithName(name string) (Entity, error)
	Delete(id string) error
}

// Service is an interface from which our api module can access our repository.
type Service interface {
	Insert(Entity) (Entity, error)
	FindAll() ([]Entity, error)
	FindAllWithStatus(status string) ([]Entity, error)
	FindOneWithID(id string) (Entity, error)
	FindOneWithName(name string) (Entity, error)
	Remove(id string) error
}
