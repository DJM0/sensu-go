// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	errors "errors"
	graphql1 "github.com/graphql-go/graphql"
	graphql "github.com/sensu/sensu-go/graphql"
)

// MutatorIDFieldResolver implement to resolve requests for the Mutator's id field.
type MutatorIDFieldResolver interface {
	// ID implements response to request for id field.
	ID(p graphql.ResolveParams) (string, error)
}

// MutatorNamespaceFieldResolver implement to resolve requests for the Mutator's namespace field.
type MutatorNamespaceFieldResolver interface {
	// Namespace implements response to request for namespace field.
	Namespace(p graphql.ResolveParams) (string, error)
}

// MutatorNameFieldResolver implement to resolve requests for the Mutator's name field.
type MutatorNameFieldResolver interface {
	// Name implements response to request for name field.
	Name(p graphql.ResolveParams) (string, error)
}

// MutatorMetadataFieldResolver implement to resolve requests for the Mutator's metadata field.
type MutatorMetadataFieldResolver interface {
	// Metadata implements response to request for metadata field.
	Metadata(p graphql.ResolveParams) (interface{}, error)
}

// MutatorCommandFieldResolver implement to resolve requests for the Mutator's command field.
type MutatorCommandFieldResolver interface {
	// Command implements response to request for command field.
	Command(p graphql.ResolveParams) (string, error)
}

// MutatorTimeoutFieldResolver implement to resolve requests for the Mutator's timeout field.
type MutatorTimeoutFieldResolver interface {
	// Timeout implements response to request for timeout field.
	Timeout(p graphql.ResolveParams) (int, error)
}

// MutatorEnvVarsFieldResolver implement to resolve requests for the Mutator's envVars field.
type MutatorEnvVarsFieldResolver interface {
	// EnvVars implements response to request for envVars field.
	EnvVars(p graphql.ResolveParams) ([]string, error)
}

// MutatorToJSONFieldResolver implement to resolve requests for the Mutator's toJSON field.
type MutatorToJSONFieldResolver interface {
	// ToJSON implements response to request for toJSON field.
	ToJSON(p graphql.ResolveParams) (interface{}, error)
}

//
// MutatorFieldResolvers represents a collection of methods whose products represent the
// response values of the 'Mutator' type.
//
// == Example SDL
//
//   """
//   Dog's are not hooman.
//   """
//   type Dog implements Pet {
//     "name of this fine beast."
//     name:  String!
//
//     "breed of this silly animal; probably shibe."
//     breed: [Breed]
//   }
//
// == Example generated interface
//
//   // DogResolver ...
//   type DogFieldResolvers interface {
//     DogNameFieldResolver
//     DogBreedFieldResolver
//
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
// == Example implementation ...
//
//   // DogResolver implements DogFieldResolvers interface
//   type DogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *DogResolver) Name(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *DogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
//
type MutatorFieldResolvers interface {
	MutatorIDFieldResolver
	MutatorNamespaceFieldResolver
	MutatorNameFieldResolver
	MutatorMetadataFieldResolver
	MutatorCommandFieldResolver
	MutatorTimeoutFieldResolver
	MutatorEnvVarsFieldResolver
	MutatorToJSONFieldResolver
}

// MutatorAliases implements all methods on MutatorFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := v.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
//
type MutatorAliases struct{}

// ID implements response to request for 'id' field.
func (_ MutatorAliases) ID(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'id'")
	}
	return ret, err
}

// Namespace implements response to request for 'namespace' field.
func (_ MutatorAliases) Namespace(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'namespace'")
	}
	return ret, err
}

// Name implements response to request for 'name' field.
func (_ MutatorAliases) Name(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'name'")
	}
	return ret, err
}

// Metadata implements response to request for 'metadata' field.
func (_ MutatorAliases) Metadata(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Command implements response to request for 'command' field.
func (_ MutatorAliases) Command(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'command'")
	}
	return ret, err
}

// Timeout implements response to request for 'timeout' field.
func (_ MutatorAliases) Timeout(p graphql.ResolveParams) (int, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := graphql1.Int.ParseValue(val).(int)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'timeout'")
	}
	return ret, err
}

// EnvVars implements response to request for 'envVars' field.
func (_ MutatorAliases) EnvVars(p graphql.ResolveParams) ([]string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.([]string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'envVars'")
	}
	return ret, err
}

// ToJSON implements response to request for 'toJSON' field.
func (_ MutatorAliases) ToJSON(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// MutatorType A Mutator is a mutator specification.
var MutatorType = graphql.NewType("Mutator", graphql.ObjectKind)

// RegisterMutator registers Mutator object type with given service.
func RegisterMutator(svc *graphql.Service, impl MutatorFieldResolvers) {
	svc.RegisterObject(_ObjectTypeMutatorDesc, impl)
}
func _ObjTypeMutatorIDHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(MutatorIDFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.ID(frp)
	}
}

func _ObjTypeMutatorNamespaceHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(MutatorNamespaceFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Namespace(frp)
	}
}

func _ObjTypeMutatorNameHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(MutatorNameFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Name(frp)
	}
}

func _ObjTypeMutatorMetadataHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(MutatorMetadataFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Metadata(frp)
	}
}

func _ObjTypeMutatorCommandHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(MutatorCommandFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Command(frp)
	}
}

func _ObjTypeMutatorTimeoutHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(MutatorTimeoutFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Timeout(frp)
	}
}

func _ObjTypeMutatorEnvVarsHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(MutatorEnvVarsFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.EnvVars(frp)
	}
}

func _ObjTypeMutatorToJSONHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(MutatorToJSONFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.ToJSON(frp)
	}
}

func _ObjectTypeMutatorConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "A Mutator is a mutator specification.",
		Fields: graphql1.Fields{
			"command": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Command is the command to be executed.",
				Name:              "command",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"envVars": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Env is a list of environment variables to use with command execution",
				Name:              "envVars",
				Type:              graphql1.NewList(graphql1.NewNonNull(graphql1.String)),
			},
			"id": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "The globally unique identifier of the record",
				Name:              "id",
				Type:              graphql1.NewNonNull(graphql1.ID),
			},
			"metadata": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "metadata contains name, namespace, labels and annotations of the record",
				Name:              "metadata",
				Type:              graphql1.NewNonNull(graphql.OutputType("ObjectMeta")),
			},
			"name": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Name is the unique identifier for a mutator.",
				Name:              "name",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"namespace": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Namespace in which this record resides",
				Name:              "namespace",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"timeout": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Timeout is the command execution timeout in seconds.",
				Name:              "timeout",
				Type:              graphql1.Int,
			},
			"toJSON": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "toJSON returns a REST API compatible representation of the resource. Handy for\nsharing snippets that can then be imported with `sensuctl create`.",
				Name:              "toJSON",
				Type:              graphql1.NewNonNull(graphql.OutputType("JSON")),
			},
		},
		Interfaces: []*graphql1.Interface{
			graphql.Interface("Node"),
			graphql.Interface("Namespaced"),
			graphql.Interface("Resource")},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see MutatorFieldResolvers.")
		},
		Name: "Mutator",
	}
}

// describe Mutator's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeMutatorDesc = graphql.ObjectDesc{
	Config: _ObjectTypeMutatorConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"command":   _ObjTypeMutatorCommandHandler,
		"envVars":   _ObjTypeMutatorEnvVarsHandler,
		"id":        _ObjTypeMutatorIDHandler,
		"metadata":  _ObjTypeMutatorMetadataHandler,
		"name":      _ObjTypeMutatorNameHandler,
		"namespace": _ObjTypeMutatorNamespaceHandler,
		"timeout":   _ObjTypeMutatorTimeoutHandler,
		"toJSON":    _ObjTypeMutatorToJSONHandler,
	},
}

// MutatorConnectionNodesFieldResolver implement to resolve requests for the MutatorConnection's nodes field.
type MutatorConnectionNodesFieldResolver interface {
	// Nodes implements response to request for nodes field.
	Nodes(p graphql.ResolveParams) (interface{}, error)
}

// MutatorConnectionPageInfoFieldResolver implement to resolve requests for the MutatorConnection's pageInfo field.
type MutatorConnectionPageInfoFieldResolver interface {
	// PageInfo implements response to request for pageInfo field.
	PageInfo(p graphql.ResolveParams) (interface{}, error)
}

//
// MutatorConnectionFieldResolvers represents a collection of methods whose products represent the
// response values of the 'MutatorConnection' type.
//
// == Example SDL
//
//   """
//   Dog's are not hooman.
//   """
//   type Dog implements Pet {
//     "name of this fine beast."
//     name:  String!
//
//     "breed of this silly animal; probably shibe."
//     breed: [Breed]
//   }
//
// == Example generated interface
//
//   // DogResolver ...
//   type DogFieldResolvers interface {
//     DogNameFieldResolver
//     DogBreedFieldResolver
//
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
// == Example implementation ...
//
//   // DogResolver implements DogFieldResolvers interface
//   type DogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *DogResolver) Name(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *DogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
//
type MutatorConnectionFieldResolvers interface {
	MutatorConnectionNodesFieldResolver
	MutatorConnectionPageInfoFieldResolver
}

// MutatorConnectionAliases implements all methods on MutatorConnectionFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := v.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
//
type MutatorConnectionAliases struct{}

// Nodes implements response to request for 'nodes' field.
func (_ MutatorConnectionAliases) Nodes(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// PageInfo implements response to request for 'pageInfo' field.
func (_ MutatorConnectionAliases) PageInfo(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// MutatorConnectionType A connection to a sequence of records.
var MutatorConnectionType = graphql.NewType("MutatorConnection", graphql.ObjectKind)

// RegisterMutatorConnection registers MutatorConnection object type with given service.
func RegisterMutatorConnection(svc *graphql.Service, impl MutatorConnectionFieldResolvers) {
	svc.RegisterObject(_ObjectTypeMutatorConnectionDesc, impl)
}
func _ObjTypeMutatorConnectionNodesHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(MutatorConnectionNodesFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Nodes(frp)
	}
}

func _ObjTypeMutatorConnectionPageInfoHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(MutatorConnectionPageInfoFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.PageInfo(frp)
	}
}

func _ObjectTypeMutatorConnectionConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "A connection to a sequence of records.",
		Fields: graphql1.Fields{
			"nodes": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "nodes",
				Type:              graphql1.NewNonNull(graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("Mutator")))),
			},
			"pageInfo": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "pageInfo",
				Type:              graphql1.NewNonNull(graphql.OutputType("OffsetPageInfo")),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see MutatorConnectionFieldResolvers.")
		},
		Name: "MutatorConnection",
	}
}

// describe MutatorConnection's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeMutatorConnectionDesc = graphql.ObjectDesc{
	Config: _ObjectTypeMutatorConnectionConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"nodes":    _ObjTypeMutatorConnectionNodesHandler,
		"pageInfo": _ObjTypeMutatorConnectionPageInfoHandler,
	},
}

// MutatorEdgeNodeFieldResolver implement to resolve requests for the MutatorEdge's node field.
type MutatorEdgeNodeFieldResolver interface {
	// Node implements response to request for node field.
	Node(p graphql.ResolveParams) (interface{}, error)
}

// MutatorEdgeCursorFieldResolver implement to resolve requests for the MutatorEdge's cursor field.
type MutatorEdgeCursorFieldResolver interface {
	// Cursor implements response to request for cursor field.
	Cursor(p graphql.ResolveParams) (string, error)
}

//
// MutatorEdgeFieldResolvers represents a collection of methods whose products represent the
// response values of the 'MutatorEdge' type.
//
// == Example SDL
//
//   """
//   Dog's are not hooman.
//   """
//   type Dog implements Pet {
//     "name of this fine beast."
//     name:  String!
//
//     "breed of this silly animal; probably shibe."
//     breed: [Breed]
//   }
//
// == Example generated interface
//
//   // DogResolver ...
//   type DogFieldResolvers interface {
//     DogNameFieldResolver
//     DogBreedFieldResolver
//
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
// == Example implementation ...
//
//   // DogResolver implements DogFieldResolvers interface
//   type DogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *DogResolver) Name(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *DogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
//
type MutatorEdgeFieldResolvers interface {
	MutatorEdgeNodeFieldResolver
	MutatorEdgeCursorFieldResolver
}

// MutatorEdgeAliases implements all methods on MutatorEdgeFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := v.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
//
type MutatorEdgeAliases struct{}

// Node implements response to request for 'node' field.
func (_ MutatorEdgeAliases) Node(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Cursor implements response to request for 'cursor' field.
func (_ MutatorEdgeAliases) Cursor(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'cursor'")
	}
	return ret, err
}

// MutatorEdgeType An edge in a connection.
var MutatorEdgeType = graphql.NewType("MutatorEdge", graphql.ObjectKind)

// RegisterMutatorEdge registers MutatorEdge object type with given service.
func RegisterMutatorEdge(svc *graphql.Service, impl MutatorEdgeFieldResolvers) {
	svc.RegisterObject(_ObjectTypeMutatorEdgeDesc, impl)
}
func _ObjTypeMutatorEdgeNodeHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(MutatorEdgeNodeFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Node(frp)
	}
}

func _ObjTypeMutatorEdgeCursorHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(MutatorEdgeCursorFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Cursor(frp)
	}
}

func _ObjectTypeMutatorEdgeConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "An edge in a connection.",
		Fields: graphql1.Fields{
			"cursor": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "cursor",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"node": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "node",
				Type:              graphql.OutputType("Mutator"),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see MutatorEdgeFieldResolvers.")
		},
		Name: "MutatorEdge",
	}
}

// describe MutatorEdge's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeMutatorEdgeDesc = graphql.ObjectDesc{
	Config: _ObjectTypeMutatorEdgeConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"cursor": _ObjTypeMutatorEdgeCursorHandler,
		"node":   _ObjTypeMutatorEdgeNodeHandler,
	},
}

// MutatorListOrder Describes ways in which a list of mutators can be ordered.
type MutatorListOrder string

// MutatorListOrders holds enum values
var MutatorListOrders = _EnumTypeMutatorListOrderValues{
	NAME:      "NAME",
	NAME_DESC: "NAME_DESC",
}

// MutatorListOrderType Describes ways in which a list of mutators can be ordered.
var MutatorListOrderType = graphql.NewType("MutatorListOrder", graphql.EnumKind)

// RegisterMutatorListOrder registers MutatorListOrder object type with given service.
func RegisterMutatorListOrder(svc *graphql.Service) {
	svc.RegisterEnum(_EnumTypeMutatorListOrderDesc)
}
func _EnumTypeMutatorListOrderConfigFn() graphql1.EnumConfig {
	return graphql1.EnumConfig{
		Description: "Describes ways in which a list of mutators can be ordered.",
		Name:        "MutatorListOrder",
		Values: graphql1.EnumValueConfigMap{
			"NAME": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "NAME",
			},
			"NAME_DESC": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "NAME_DESC",
			},
		},
	}
}

// describe MutatorListOrder's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _EnumTypeMutatorListOrderDesc = graphql.EnumDesc{Config: _EnumTypeMutatorListOrderConfigFn}

type _EnumTypeMutatorListOrderValues struct {
	// NAME - self descriptive
	NAME MutatorListOrder
	// NAME_DESC - self descriptive
	NAME_DESC MutatorListOrder
}
