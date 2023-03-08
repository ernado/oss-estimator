package lang

import "strings"

// All returns the list of all counted languages.
func All() []string {
	return []string{
		"Go",
		"C",
		"Go Template",
		"BASH",
		"Python",
		"C++",
		"Java",
		"Swift",
		"Ruby",
		"JavaScript",
		"TypeScript",
		"Haskell",
		"Dart",
		"PHP",
		"C#",
		"Kotlin",
		"Cython",
		"Rust",
		"Scala",
		"Objective C++",
		"Objective C",
		"Protocol Buffers",
		"Bazel",
		"Erlang",
		"Lua",
		"Perl",
		"BASH",
	}
}

var _languages = func() map[string]struct{} {
	list := All()
	m := make(map[string]struct{}, len(list))
	for _, v := range list {
		m[strings.ToLower(v)] = struct{}{}
	}
	return m
}()

// In returns true if the given language is in the list.
func In(k string) bool {
	_, ok := _languages[strings.ToLower(k)]
	return ok
}
