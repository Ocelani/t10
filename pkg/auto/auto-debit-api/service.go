package main

import "github.com/Ocelani/t10/pkg/auto"

type Service struct {
	Repository auto.Repository
}

// NewMongoService is used to create a single instance of the Repository.
func NewMongoService(mongoURL, collection string) auto.Service {
	return &Service{Repository: NewMongoRepository(mongoURL, collection)}
}

// Insert auto.Entity.
func (s *Service) Insert(auto auto.Entity) (auto.Entity, error) {
	return s.Repository.Create(auto)
}

// FindAll auto.Entity data.
func (s *Service) FindAll() ([]auto.Entity, error) {
	return s.Repository.ReadAll()
}

// FindOneWithID auto.Entity.
func (s *Service) FindAllWithStatus(status string) ([]auto.Entity, error) {
	return s.Repository.ReadAllWithStatus(status)
}

// FindOneWithID auto.Entity.
func (s *Service) FindOneWithID(id string) (auto.Entity, error) {
	return s.Repository.ReadOneWithID(id)
}

// FindOneWithName auto.Entity.
func (s *Service) FindOneWithName(name string) (auto.Entity, error) {
	return s.Repository.ReadOneWithName(name)
}

// Remove auto.Entity.
func (s *Service) Remove(id string) error {
	return s.Repository.Delete(id)
}
