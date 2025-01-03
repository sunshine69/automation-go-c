package lib

import (
	"fmt"

	"github.com/Jemmic/go-pcre2"
)

// go-pcre2 implementation of FindStringSubmatch
func Pcre2FindStringSubmatch[T string | *pcre2.Regexp](pattern T, input string) ([]string, error) {
	resultB, err := Pcre2FindSubmatch(pattern, []byte(input))
	if err != nil {
		return nil, err
	}
	var output []string

	for _, gb := range resultB {
		gs := string(gb)
		output = append(output, gs)
	}
	return output, nil
}

// go-pcre2 implementation of FindStringSubmatch
func Pcre2FindSubmatch[T string | *pcre2.Regexp](pattern T, input []byte) ([][]byte, error) {
	var re *pcre2.Regexp
	anyPattern := any(pattern)
	if _ptn, ok := anyPattern.(*pcre2.Regexp); ok {
		re = _ptn
	} else {
		var err error
		if _ptn, ok := anyPattern.(string); ok {
			re, err = pcre2.Compile(_ptn, 0)
			if err != nil {
				return nil, fmt.Errorf("failed to compile pattern: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to assert pattern is string %v", pattern)
		}
	}
	defer re.Free()

	// Match the input string
	m := re.Matcher(input, 0)
	if !m.Matches() {
		return nil, nil
	}

	// Extract matched groups
	var groups [][]byte
	// Always append the full match (group 0) first
	groups = append(groups, m.Group(0))

	// Extract other capturing groups
	for i := 1; i <= m.Groups(); i++ {
		group := m.Group(i)
		groups = append(groups, group)
	}
	return groups, nil
}

// go-pcre2 implementation of FindAllStringSubmatch. We need to compile the ptn just before calling this func or at least the same scope otherwise we hit error
// panic: Matcher.Init: uninitialized. If pass using string it does not have the issue
func Pcre2FindAllStringSubmatch[T string | *pcre2.Regexp](pattern T, input string) ([][]string, error) {
	return Pcre2FindAllSubmatch(pattern, []byte(input))
}

// go-pcre2 implementation of FindAllSubmatch. We need to compile the ptn just before calling this func or at least the same scope otherwise we hit error
// panic: Matcher.Init: uninitialized. If pass using string it does not have the issue
func Pcre2FindAllSubmatch[T string | *pcre2.Regexp](pattern T, input []byte) ([][]string, error) {
	var re *pcre2.Regexp
	anyPattern := any(pattern)
	if _ptn, ok := anyPattern.(*pcre2.Regexp); ok {
		re = _ptn
	} else {
		var err error
		if _ptn, ok := anyPattern.(string); ok {
			re, err = pcre2.Compile(_ptn, 0)
			if err != nil {
				return nil, fmt.Errorf("failed to compile pattern: %w", err)
			}
			defer re.Free()
		} else {
			return nil, fmt.Errorf("failed to assert pattern is string %v", pattern)
		}
	}

	var allMatches [][]string
	offset := 0

	for {
		// Match the string from the current offset
		m := re.Matcher(input[offset:], 0)
		if !m.Matches() {
			break
		}

		// Extract matched groups
		var groups []string
		// Always append the full match (group 0) first
		groups = append(groups, m.GroupString(0))

		// Extract other capturing groups
		for i := 1; i <= m.Groups(); i++ {
			groups = append(groups, m.GroupString(i))
		}

		// Use the first group (full match) to update the offset correctly
		groupIndices := m.GroupIndices(0) // Get indices for the full match (group 0)
		offset += groupIndices[1]         // Shift offset to the end of the match

		// Store the match result
		allMatches = append(allMatches, groups)

		// Stop if we've processed the entire string
		if offset >= len(input) {
			break
		}
	}

	return allMatches, nil
}
