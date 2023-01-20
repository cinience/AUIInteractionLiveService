package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/iancoleman/strcase"
)

func IsChangeKeyStyle(r *http.Request) bool {
	if r.Header.Get("x-key-style") == "LowerCamel" {
		return true
	}
	return false
}

func DefaultRequester(r *http.Request) {
	if !IsChangeKeyStyle(r) {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		var in map[string]interface{}
		err = json.Unmarshal(body, &in)
		if err == nil {
			in["user_id"] = "admin"
			out := DeepCopy(in, func(k string) string {
				val := strcase.ToSnake(k)
				return val
			})
			newBody, err := json.Marshal(out)
			if err == nil {
				log.Println(string(newBody))
				r.Body = io.NopCloser(bytes.NewReader(newBody))
			}
		}
	}
}

func DefaultResponder(w http.ResponseWriter, r *http.Request, v interface{}) {
	if IsChangeKeyStyle(r) {
		//tv := reflect.ValueOf(v)
		//log.Printf("kind is %v \n", tv.Kind())
		data, _ := json.Marshal(v)

		var m interface{}
		_ = json.Unmarshal(data, &m)

		v = DeepCopy(m, func(k string) string {
			val := strcase.ToLowerCamel(k)
			return val
		})
	}
	render.DefaultResponder(w, r, v)
}
