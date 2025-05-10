package strutil_test

import (
	"testing"

	"github.com/gomooth/utils/strutil"
	"github.com/stretchr/testify/assert"
)

func TestSnake(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple", "HelloWorld", "hello_world"},
		{"multiple caps", "HTMLElement", "html_element"},
		{"existing underscore", "hello_world", "hello_world"},
		{"mixed case", "XMLHttpRequest", "xml_http_request"},
		{"all caps", "HELLOWORLD", "helloworld"},
		{"empty", "", ""},
		{"single word", "hello", "hello"},
		{"numbers", "UserID42", "user_id42"},
		{"unicode", "ПриветМир", "привет_мир"},
		{"acronym", "JSONData", "json_data"},
		{"consecutive caps", "HTMLElement", "html_element"},
		{"complex", "MyHTTPRequest", "my_http_request"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, strutil.Snake(tt.input))
		})
	}
}

func TestCamel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple", "hello_world", "HelloWorld"},
		{"multiple underscores", "this_is_a_test", "ThisIsATest"},
		{"no underscores", "hello", "Hello"},
		{"mixed case", "XML_http_request", "XMLHttpRequest"},
		{"empty", "", ""},
		{"single underscore", "a_b", "AB"},
		{"numbers", "user_id_42", "UserId42"},
		{"unicode", "привет_мир", "ПриветМир"},
		{"unicode with numbers", "user_ид_42", "UserИд42"},
		{"mixed unicode", "имя_пользователя", "ИмяПользователя"},
		{"leading underscore", "_hello", "Hello"},
		{"trailing underscore", "hello_", "Hello"},
		{"consecutive underscores", "hello__world", "HelloWorld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, strutil.Camel(tt.input))
		})
	}
}

func BenchmarkSnake(b *testing.B) {
	input := "ThisIsALongStringToTestPerformanceOfTheSnakeFunction"
	for i := 0; i < b.N; i++ {
		strutil.Snake(input)
	}
}

func BenchmarkCamel(b *testing.B) {
	input := "this_is_a_long_string_to_test_performance_of_the_camel_function"
	for i := 0; i < b.N; i++ {
		strutil.Camel(input)
	}
}
