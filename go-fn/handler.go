package function

import (
	"fmt"

	"gopkg.in/loremipsum.v1"
)

// Handle a serverless request
func Handle(req []byte) string {
	text := loremipsum.New().Words(10)

	return fmt.Sprintf("Request body: %s\nResponse: %s", string(req), string(text))
}
