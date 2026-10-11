package testutil

import (
	"fmt"
	"strings"
)

// Sample markers occupy a whole line. The first input marker is optional,
// matching samples copied without the heading for the first example.
func parseSampleCases(rawText string) ([][2]string, error) {
	rawText = strings.TrimSpace(rawText)
	if rawText == "" {
		return nil, fmt.Errorf("rawText is empty")
	}
	var cases [][2]string
	var input, output []string
	started, inOutput := false, false
	appendCase := func() {
		cases = append(cases, [2]string{strings.Join(input, "\n"), strings.Join(output, "\n")})
		input, output = nil, nil
	}
	for i, line := range strings.Split(rawText, "\n") {
		switch strings.TrimSpace(line) {
		case "input", "inputCopy":
			if started {
				if !inOutput {
					return nil, fmt.Errorf("line %d: case %d is missing an output marker", i+1, len(cases)+1)
				}
				appendCase()
			}
			started, inOutput = true, false
		case "output", "outputCopy":
			if !started || inOutput {
				return nil, fmt.Errorf("line %d: unexpected output marker", i+1)
			}
			inOutput = true
		default:
			started = true
			if inOutput {
				output = append(output, line)
			} else {
				input = append(input, line)
			}
		}
	}
	if !inOutput {
		return nil, fmt.Errorf("case %d is missing an output marker", len(cases)+1)
	}
	appendCase()
	return cases, nil
}
