package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOneOptions struct {
	*options.FindOneOptions
	readPreference bool
}

// FindOne creates a new FindOneOptions instance.
func OptionsFindOne() *FindOneOptions {
	return &FindOneOptions{FindOneOptions: options.FindOne()}
}

func (s FindOneOptions) ReadPreference() bool {
	return s.readPreference
}

func (s *FindOneOptions) SetReadPreference(flag bool) *FindOneOptions {
	s.readPreference = flag
	return s
}

type FindOptions struct {
	*options.FindOptions
	readPreference bool
}

// Find creates a new FindOptions instance.
func OptionsFind() *FindOptions {
	return &FindOptions{FindOptions: options.Find()}
}

func (s FindOptions) ReadPreference() bool {
	return s.readPreference
}

func (s *FindOptions) SetReadPreference(flag bool) *FindOptions {
	s.readPreference = flag
	return s
}

type AggregateOptions struct {
	*options.AggregateOptions
	page           *int64
	limit          *int64
	noNeedCount    *bool
	sort           bson.D
	readPreference bool
	facetsPipeline []bson.M
}

// Aggregate creates a new AggregateOptions instance.
func OptionsAggregate() *AggregateOptions {
	return &AggregateOptions{AggregateOptions: options.Aggregate()}
}

func (s AggregateOptions) ReadPreference() bool {
	return s.readPreference
}

func (s *AggregateOptions) SetReadPreference(flag bool) *AggregateOptions {
	s.readPreference = flag
	return s
}

// Use for method AggregationFacets
func (s *AggregateOptions) SetSort(sort bson.D) *AggregateOptions {
	s.sort = sort
	return s
}

// Use for method AggregationFacets
func (s *AggregateOptions) SetPage(page *int64) *AggregateOptions {
	s.page = page
	return s
}

// Use for method AggregationFacets
func (s *AggregateOptions) SetLimit(limit *int64) *AggregateOptions {
	s.limit = limit
	return s
}

// Use for method AggregationFacets
func (s *AggregateOptions) SetNoNeedCount(flag *bool) *AggregateOptions {
	s.noNeedCount = flag
	return s
}

// Use for method AggregationFacets
func (s *AggregateOptions) SetFacetsPipeline(facetsPipeline ...bson.M) *AggregateOptions {
	s.facetsPipeline = append(s.facetsPipeline, facetsPipeline...)
	return s
}

type CountOptions struct {
	*options.CountOptions
	readPreference bool
}

// Count creates a new CountOptions instance.
func OptionsCount() *CountOptions {
	return &CountOptions{CountOptions: options.Count()}
}

func (s CountOptions) ReadPreference() bool {
	return s.readPreference
}

func (s *CountOptions) SetReadPreference(flag bool) *CountOptions {
	s.readPreference = flag
	return s
}
