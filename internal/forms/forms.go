package forms

import (
	"net/http"
	"net/url"
)

//Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

//Valid returns true if the form is valid else false.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}), //curly braces because the string is empty
	}
}

//Has check if form is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field can't be blank!")
		return false
	}
	return true
}
