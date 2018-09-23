package validate

import (
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"
)

// FormData obtained from the request body or url query parameters or user custom setting.
type FormData struct {
	// Form holds any basic key-value string data
	// This includes all fields from urlencoded form,
	// and the form fields only (not files) from a multipart form
	Form url.Values
	// Files holds files from a multipart form only.
	// For any other type of request, it will always
	// be empty. Files only supports one file per key,
	// since this is by far the most common use. If you
	// need to have more than one file per key, parse the
	// files manually using r.MultipartForm.File.
	Files map[string]*multipart.FileHeader
	// jsonBodies holds the original body of the request.
	// Only available for json requests.
	jsonBodies []byte
}

/*************************************************************
 * Form data operate
 *************************************************************/

// Create a Validation from data
func (d *FormData) Create(scene ...string) *Validation {
	return d.Validation(scene...)
}

// Validation create from data
func (d *FormData) Validation(scene ...string) *Validation {
	return NewValidation(d, scene...)
}

// Add adds the value to key. It appends to any existing values associated with key.
func (d *FormData) Add(key string, value string) {
	d.Form.Add(key, value)
}

// AddValues to Data.Form
func (d *FormData) AddValues(values url.Values) {
	for key, vals := range values {
		for _, val := range vals {
			d.Form.Add(key, val)
		}
	}
}

// AddFiles adds the multipart form files to data
func (d *FormData) AddFiles(files map[string][]*multipart.FileHeader) {
	for key, files := range files {
		if len(files) != 0 {
			d.AddFile(key, files[0])
		}
	}
}

// AddFile adds the multipart form file to data with the given key.
func (d *FormData) AddFile(key string, file *multipart.FileHeader) {
	d.Files[key] = file
}

// Del deletes the values associated with key.
func (d *FormData) Del(key string) {
	d.Form.Del(key)
}

// DelFile deletes the file associated with key (if any).
// If there is no file associated with key, it does nothing.
func (d *FormData) DelFile(key string) {
	delete(d.Files, key)
}

// Encode encodes the values into “URL encoded” form ("bar=baz&foo=quux") sorted by key.
// Any files in d will be ignored because there is no direct way to convert a file to a
// URL encoded value.
func (d *FormData) EncodeForm() string {
	return d.Form.Encode()
}

// Set sets the key to value. It replaces any existing values.
func (d *FormData) Set(field string, val interface{}) error {
	d.Form.Set(field, val.(string))
	return nil
}

func (d FormData) Get(key string) (interface{}, bool) {
	return d.Form.Get(key), true
}

// Trimmed gets the first value by key, and apply strings.TrimSpace
func (d FormData) Trimmed(key string) string {
	return strings.TrimSpace(d.Form.Get(key))
}

// GetFile returns the multipart form file associated with key, if any, as a *multipart.FileHeader.
// If there is no file associated with key, it returns nil. If you just want the body of the
// file, use GetFileBytes.
func (d FormData) GetFile(key string) *multipart.FileHeader {
	return d.Files[key]
}

// Has key in the Data
func (d FormData) Has(key string) bool {
	if vs, ok := d.Form[key]; ok && len(vs) > 0 {
		return true
	}

	if _, ok := d.Files[key]; ok {
		return true
	}

	return false
}

// HasField returns true iff data.Form[key] exists. When parsing a request body, the key
// is considered to be in existence if it was provided in the request body, even if its value
// is empty.
func (d FormData) HasField(key string) bool {
	_, found := d.Form[key]
	return found
}

// HasFile returns true iff data.Files[key] exists. When parsing a request body, the key
// is considered to be in existence if it was provided in the request body, even if the file
// is empty.
func (d FormData) HasFile(key string) bool {
	_, found := d.Files[key]
	return found
}

// GetInt returns the first element in data[key] converted to an int.
func (d FormData) Int(key string) int {
	if !d.HasField(key) || len(d.Form[key]) == 0 {
		return 0
	}

	str, _ := d.Get(key)
	if result, err := strconv.Atoi(str.(string)); err != nil {
		panic(err)
	} else {
		return result
	}
}

// MustInt64 returns the first element in data[key] converted to an int64.
func (d FormData) MustInt64(key string) int64 {
	if !d.HasField(key) || len(d.Form[key]) == 0 {
		return 0
	}

	val, _ := strconv.ParseInt(d.Trimmed(key), 10, 64)
	return val
}

// GetFloat returns the first element in data[key] converted to a float.
func (d FormData) Float(key string) float64 {
	if !d.HasField(key) || len(d.Form[key]) == 0 {
		return 0.0
	}

	val, _ := d.Get(key)
	result, err := strconv.ParseFloat(val.(string), 64)
	if err != nil {
		panic(err)
	}

	return result
}

// GetBool returns the first element in data[key] converted to a bool.
func (d FormData) Bool(key string) bool {
	if !d.HasField(key) || len(d.Form[key]) == 0 {
		return false
	}

	val, _ := d.Get(key)
	if result, err := strconv.ParseBool(val.(string)); err != nil {
		panic(err)
	} else {
		return result
	}
}

// GetFileBytes returns the body of the file associated with key. If there is no
// file associated with key, it returns nil (not an error). It may return an error if
// there was a problem reading the file. If you need to know whether or not the file
// exists (i.e. whether it was provided in the request), use the FileExists method.
func (d FormData) GetFileBytes(key string) ([]byte, error) {
	fileHeader, found := d.Files[key]
	if !found {
		return nil, nil
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(file)
}

// GetStringsSplit returns the first element in data[key] split into a slice delimited by delim.
func (d FormData) GetStringsSplit(key string, delim string) []string {
	if !d.HasField(key) || len(d.Form[key]) == 0 {
		return nil
	}

	return strings.Split(d.Form[key][0], delim)
}

// BindJSON binds v to the json data in the request body. It calls json.Unmarshal and
// sets the value of v.
func (d FormData) BindJSON(v interface{}) error {
	if len(d.jsonBodies) == 0 {
		return nil
	}

	return Unmarshal(d.jsonBodies, v)
}

// MapTo the v
func (d FormData) MapTo(v interface{}) error {
	return d.BindJSON(v)
}
