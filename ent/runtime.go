// Code generated by ent, DO NOT EDIT.

package ent

import (
	"threads/ent/message"
	"threads/ent/schema"
	"threads/ent/thread"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescText is the schema descriptor for text field.
	messageDescText := messageFields[1].Descriptor()
	// message.TextValidator is a validator for the "text" field. It is called by the builders before save.
	message.TextValidator = messageDescText.Validators[0].(func(string) error)
	// messageDescCreated is the schema descriptor for created field.
	messageDescCreated := messageFields[2].Descriptor()
	// message.DefaultCreated holds the default value on creation for the created field.
	message.DefaultCreated = messageDescCreated.Default.(func() time.Time)
	// messageDescID is the schema descriptor for id field.
	messageDescID := messageFields[0].Descriptor()
	// message.DefaultID holds the default value on creation for the id field.
	message.DefaultID = messageDescID.Default.(func() string)
	threadFields := schema.Thread{}.Fields()
	_ = threadFields
	// threadDescName is the schema descriptor for name field.
	threadDescName := threadFields[1].Descriptor()
	// thread.NameValidator is a validator for the "name" field. It is called by the builders before save.
	thread.NameValidator = threadDescName.Validators[0].(func(string) error)
	// threadDescID is the schema descriptor for id field.
	threadDescID := threadFields[0].Descriptor()
	// thread.DefaultID holds the default value on creation for the id field.
	thread.DefaultID = threadDescID.Default.(func() string)
}
