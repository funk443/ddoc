// Package ddoc provides functionalities to construct a [Document] from a file.
package ddoc

import (
	"bufio"
	"fmt"
	log "github.com/funk443/simplog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// State represents the internal state for the parser.
type State int

const (
	StateTitleLine = State(iota)
	StateMetadata
	StateContent
	StateHeading1
	StateHeading2
	StateHeading3
	StateCodeblock
	StateParagraph
	StateReference
)

var StateNames = map[State]string{
	StateTitleLine: "Title line",
	StateMetadata:  "Metadata",
	StateContent:   "Content",
	StateHeading1:  "Heading 1",
	StateHeading2:  "Heading 2",
	StateHeading3:  "Heading 3",
	StateCodeblock: "Codeblock",
	StateParagraph: "Paragraph",
	StateReference: "Reference",
}

func (state State) String() string {
	return StateNames[state]
}

type ContentType int

const (
	ContentH1 = ContentType(iota)
	ContentH2
	ContentH3
	ContentParagraph
	ContentReference
	ContentCodeblock
)

var ContentTypeNames = map[ContentType]string{
	ContentH1:        "Heading 1",
	ContentH2:        "Heading 2",
	ContentH3:        "Heading 3",
	ContentParagraph: "Paragraph",
	ContentReference: "Reference",
	ContentCodeblock: "Codeblock",
}

func (t ContentType) String() string {
	return ContentTypeNames[t]
}

// Content is the representation of a block of data inside a DDoc file.
type Content struct {
	Type   ContentType
	RefNum int
	Text   string
}

// Document is the representation for a DDoc file.
type Document struct {
	Filename  string
	Title     string
	Metadatas map[string]string
	Contents  []Content
}

// RawRegexps contains regular expressions used for parsing, it can be compiled
// using [CompileRegexps].
var RawRegexps = map[State]string{
	StateTitleLine: `^\*([^*]+)\*\s+(\S.*)$`,
	StateMetadata:  `^\s+([^: ][^:]+): (.+)$`,
	StateHeading1:  `^={2,}$`,
	StateHeading2:  `^-{2,}$`,
	StateHeading3:  `^\.{2,}$`,
	StateCodeblock: `^ {4}(.*)$`,
	StateReference: `^\[(\d+)\]: (.*)$`,
}

// FromFile parses the file at the path filename, and returns a [Document]
// representing the file.
func FromFile(filename string) (Document, error) {
	doc := Document{
		Metadatas: map[string]string{},
		Contents:  []Content{},
	}

	regexps, err := CompileRegexps()
	if err != nil {
		errr := fmt.Errorf("Failed to initialize regexps: %w", err)
		return doc, errr
	}

	inputFile, err := os.Open(filename)
	if err != nil {
		errr := fmt.Errorf("Failed to open file `%s`: %w", filename, err)
		return doc, errr
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	lineCount := 0
	state := StateTitleLine
	storage := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		switch state {
		case StateTitleLine:
			matches := regexps[StateTitleLine].FindStringSubmatch(line)
			if matches == nil {
				err := fmt.Errorf("Failed to match title line")
				return doc, err
			}

			doc.Filename = matches[1]
			doc.Title = matches[2]
			state = StateMetadata

		case StateMetadata:
			if len(line) == 0 {
				state = StateContent
				continue
			}

			matches := regexps[StateMetadata].FindStringSubmatch(line)
			if matches == nil {
				err := fmt.Errorf("Failed to match metadata")
				return doc, err
			}
			doc.Metadatas[matches[1]] = matches[2]

		case StateContent:
			storage = []string{}
			switch {
			case regexps[StateHeading1].FindStringIndex(line) != nil:
				state = StateHeading1

			case regexps[StateHeading2].FindStringIndex(line) != nil:
				state = StateHeading2

			case regexps[StateHeading3].FindStringIndex(line) != nil:
				state = StateHeading3

			case regexps[StateCodeblock].FindStringSubmatch(line) != nil:
				storage = append(
					storage,
					regexps[StateCodeblock].FindStringSubmatch(line)[1],
				)
				state = StateCodeblock

			case regexps[StateReference].FindStringSubmatch(line) != nil:
				matches := regexps[StateReference].FindStringSubmatch(line)
				content := Content{
					Type: ContentReference,
				}
				refNum, err := strconv.Atoi(matches[1])
				if err != nil {
					errr := fmt.Errorf(
						"Failed to parse reference number: %w",
						err,
					)
					log.E("%v.", errr)
					content.Type = ContentParagraph
					content.Text = strings.TrimSpace(line)
				} else {
					content.RefNum = refNum
					content.Text = matches[2]
				}

				doc.Contents = append(doc.Contents, content)

			case len(line) == 0:
				continue

			default:
				storage = append(storage, strings.TrimSpace(line))
				state = StateParagraph
			}

		case StateHeading1:
			doc.Contents = append(doc.Contents, Content{
				Type: ContentH1,
				Text: strings.TrimSpace(line),
			})
			state = StateContent

		case StateHeading2:
			doc.Contents = append(doc.Contents, Content{
				Type: ContentH2,
				Text: strings.TrimSpace(line),
			})
			state = StateContent

		case StateHeading3:
			doc.Contents = append(doc.Contents, Content{
				Type: ContentH3,
				Text: strings.TrimSpace(line),
			})
			state = StateContent

		case StateCodeblock:
			matches := regexps[StateCodeblock].FindStringSubmatch(line)
			if matches == nil {
				state = StateContent
				doc.Contents = append(doc.Contents, Content{
					Type: ContentCodeblock,
					Text: strings.Join(storage, "\n"),
				})
				continue
			}

			storage = append(storage, matches[1])

		case StateParagraph:
			if len(line) == 0 {
				state = StateContent
				doc.Contents = append(doc.Contents, Content{
					Type: ContentParagraph,
					Text: strings.Join(storage, " "),
				})
			}

			storage = append(storage, strings.TrimSpace(line))

		default:
			err := fmt.Errorf("State %s not implemented", state)
			return doc, err
		}
	}

	return doc, nil
}

// CompileRegexps returns a map of pointers to [regexp.Regexp] compiled from
// [RawRegexps]. It will return an error if [regexp.Compile] returns one.
func CompileRegexps() (map[State]*regexp.Regexp, error) {
	regexps := map[State]*regexp.Regexp{}

	for state, s := range RawRegexps {
		re, err := regexp.Compile(s)
		if err != nil {
			errr := fmt.Errorf(
				"Failed to compile regexp for %s: %w",
				state,
				err,
			)
			return regexps, errr
		}

		regexps[state] = re
	}

	return regexps, nil
}
