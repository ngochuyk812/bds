package global

import "github.com/go-playground/validator/v10"

var (
	Validate     = validator.New()
	PoliciesPath = make(map[string][]string)
)

func init() {
	initalPoliciesPath()
}

func initalPoliciesPath() {
	PoliciesPath["/greet.v1.GreetService/Greet"] = []string{"user"}
}
