package db

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	eq  = "$eq"
	ne  = "$ne"
	gt  = "$gt"
	gte = "$gte"
	lt  = "$lt"
	lte = "$lte"
	in  = "$in"
	nin = "$nin"

	and = "$and"
	or  = "$or"
	nor = "$nor"

	regx        = "$regex"
	regxOptions = "$options"

	exists     = "$exists"
	_type      = "$type"
	mod        = "$mod"
	where      = "$where"
	all        = "$all"
	size       = "$size"
	jsonSchema = "$jsonSchema"
)

type operation struct {
	operator        string
	field           any
	value           any
	isFieldRequired bool
}

type filter struct {
	query      bson.D
	operations []operation

	isBuild bool
}
type kyte struct {
	source     any
	fields     map[any]string
	fieldNames []string
	err        error
	checkField bool
}

func newKyte(source any, checkField bool) *kyte {
	if source == nil {
		return &kyte{}
	}
	kyte := &kyte{fields: make(map[any]string), checkField: checkField}
	kyte.setSourceAndPrepareFields(source)
	return kyte
}

// FilterExpr represents a filter expression that can be converted to BSON
type FilterExpr interface {
	ToBSON() (bson.D, error)
}

// Equal creates a filter expression that checks if a field equals a value
func Equal[T any](field T, value any) FilterExpr {
	return &comparisonExpr{
		fieldExtractor: extractField(field),
		operator:       eq,
		value:          value,
	}
}

// NotEqual creates a filter expression that checks if a field does not equal a value
func NotEqual[T any](field T, value any) FilterExpr {
	return &comparisonExpr{
		fieldExtractor: extractField(field),
		operator:       ne,
		value:          value,
	}
}

// GreaterThan creates a filter expression that checks if a field is greater than a value
func GreaterThan[T any](field T, value any) FilterExpr {
	return &comparisonExpr{
		fieldExtractor: extractField(field),
		operator:       gt,
		value:          value,
	}
}

// GreaterThanOrEqual creates a filter expression that checks if a field is greater than or equal to a value
func GreaterThanOrEqual[T any](field T, value any) FilterExpr {
	return &comparisonExpr{
		fieldExtractor: extractField(field),
		operator:       gte,
		value:          value,
	}
}

// LessThan creates a filter expression that checks if a field is less than a value
func LessThan[T any](field T, value any) FilterExpr {
	return &comparisonExpr{
		fieldExtractor: extractField(field),
		operator:       lt,
		value:          value,
	}
}

// LessThanOrEqual creates a filter expression that checks if a field is less than or equal to a value
func LessThanOrEqual[T any](field T, value any) FilterExpr {
	return &comparisonExpr{
		fieldExtractor: extractField(field),
		operator:       lte,
		value:          value,
	}
}

// In creates a filter expression that checks if a field is in a list of values
func In[T any](field T, values ...any) FilterExpr {
	return &comparisonExpr{
		fieldExtractor: extractField(field),
		operator:       in,
		value:          values,
	}
}

// NotIn creates a filter expression that checks if a field is not in a list of values
func NotIn[T any](field T, values ...any) FilterExpr {
	return &comparisonExpr{
		fieldExtractor: extractField(field),
		operator:       nin,
		value:          values,
	}
}

// And creates a filter expression that combines multiple expressions with logical AND
func And(expressions ...FilterExpr) FilterExpr {
	return &logicalExpr{
		operator:    and,
		expressions: expressions,
	}
}

// Or creates a filter expression that combines multiple expressions with logical OR
func Or(expressions ...FilterExpr) FilterExpr {
	return &logicalExpr{
		operator:    or,
		expressions: expressions,
	}
}

// Nor creates a filter expression that combines multiple expressions with logical NOR
func Nor(expressions ...FilterExpr) FilterExpr {
	return &logicalExpr{
		operator:    nor,
		expressions: expressions,
	}
}

// Exists creates a filter expression that checks if a field exists
func Exists[T any](field T, e bool) FilterExpr {
	return &comparisonExpr{
		fieldExtractor: extractField(field),
		operator:       exists,
		value:          e,
	}
}

// Where creates a filter expression with a JavaScript expression
func Where(js string) FilterExpr {
	return &whereExpr{
		js: js,
	}
}

// All creates a filter expression that checks if an array field contains all specified values
func All[T any](field T, values ...any) FilterExpr {
	return &comparisonExpr{
		fieldExtractor: extractField(field),
		operator:       all,
		value:          values,
	}
}

// Size creates a filter expression that checks the size of an array field
func Size[T any](field T, s int) FilterExpr {
	return &comparisonExpr{
		fieldExtractor: extractField(field),
		operator:       size,
		value:          s,
	}
}

// JSONSchema creates a filter expression with a JSON schema
func JSONSchema(schema bson.M) FilterExpr {
	return &jsonSchemaExpr{
		schema: schema,
	}
}

// FilterFrom creates a new filter from a function that returns a FilterExpr
func FilterFrom[T any](exprFunc func(model T) FilterExpr) *filter {
	var model T
	expr := exprFunc(model)

	bsonQuery, err := expr.ToBSON()

	f := &filter{
		kyte:  newKyte(reflect.TypeOf(model), true),
		query: bsonQuery,
	}

	if err != nil {
		f.kyte.setError(err)
	}

	// Apply global filters
	globalMutex.RLock()
	for _, globalFilter := range globalFilters {
		if globalQuery, err := globalFilter.Build(); err == nil {
			f.query = append(f.query, globalQuery...)
		}
	}
	globalMutex.RUnlock()

	return f
}

// Helper types for the implementation

type fieldExtractor func() (string, error)

func extractField(field any) fieldExtractor {
	return func() (string, error) {
		// This is a simplified version - in a real implementation,
		// you would use reflection to determine the field name from struct tags
		// For now, we assume the field is a string
		if fieldName, ok := field.(string); ok {
			return fieldName, nil
		}

		// In a real implementation, we would analyze the struct field
		// and extract the BSON field name from struct tags
		return "", ErrInvalidFieldType
	}
}

type comparisonExpr struct {
	fieldExtractor fieldExtractor
	operator       string
	value          any
}

func (e *comparisonExpr) ToBSON() (bson.D, error) {
	field, err := e.fieldExtractor()
	if err != nil {
		return nil, err
	}

	// Handle special case for regex
	if e.operator == regx {
		if regexMap, ok := e.value.(bson.M); ok {
			return bson.D{{Key: field, Value: regexMap}}, nil
		}
	}

	return bson.D{{Key: field, Value: bson.M{e.operator: e.value}}}, nil
}

type logicalExpr struct {
	operator    string
	expressions []FilterExpr
}

func (e *logicalExpr) ToBSON() (bson.D, error) {
	exprs := bson.A{}

	for _, expr := range e.expressions {
		bsonExpr, err := expr.ToBSON()
		if err != nil {
			return nil, err
		}

		for _, item := range bsonExpr {
			exprs = append(exprs, bson.M{item.Key: item.Value})
		}
	}

	return bson.D{{Key: e.operator, Value: exprs}}, nil
}

type whereExpr struct {
	js string
}

func (e *whereExpr) ToBSON() (bson.D, error) {
	return bson.D{{Key: where, Value: e.js}}, nil
}

type jsonSchemaExpr struct {
	schema bson.M
}

func (e *jsonSchemaExpr) ToBSON() (bson.D, error) {
	return bson.D{{Key: jsonSchema, Value: e.schema}}, nil
}

// ErrInvalidFieldType is returned when a field is not of the expected type
var ErrInvalidFieldType = error("invalid field type")
