package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/flily/monkey-lang/compiler"
	"github.com/flily/monkey-lang/object"
)

type InteractiveScanner struct {
	In     io.Reader
	Out    io.Writer
	Prompt string

	Interactive bool

	scanner   *bufio.Scanner
	files     []string
	fileIndex int
}

func NewInteractiveScanner(in io.Reader, out io.Writer, prompt string) *InteractiveScanner {
	s := &InteractiveScanner{
		In:      in,
		Out:     out,
		Prompt:  prompt,
		scanner: bufio.NewScanner(in),
	}

	return s
}

func (s *InteractiveScanner) readFile() (string, error) {
	if s.fileIndex < 0 || s.fileIndex >= len(s.files) {
		return "", io.EOF
	}

	path := s.files[s.fileIndex]
	fd, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer fd.Close()
	content, err := ioutil.ReadAll(fd)
	if err != nil {
		return "", err
	}

	s.fileIndex += 1
	return string(content), nil
}

func (s *InteractiveScanner) Scan() (string, bool, error) {
	if s.fileIndex < len(s.files) {
		line, err := s.readFile()
		if err != nil {
			return "", false, err
		}

		return line, true, nil
	}

	if len(s.files) > 0 && !s.Interactive {
		return "", false, nil
	}

	_, err := io.WriteString(s.Out, s.Prompt)
	if err != nil {
		return "", false, err
	}

	if s.scanner.Scan() {
		return s.scanner.Text(), true, nil
	}

	return "", false, s.scanner.Err()
}

func (s *InteractiveScanner) WriteString(str string) {
	_, _ = io.WriteString(s.Out, str)
}

func (s *InteractiveScanner) Printf(format string, args ...interface{}) {
	s.WriteString(fmt.Sprintf(format, args...))
}

func (s *InteractiveScanner) WriteLines(lines []string) {
	for _, line := range lines {
		s.WriteString(line + "\n")
	}
}

func (s *InteractiveScanner) AddFiles(files []string) {
	s.files = append(s.files, files...)
}

func (s *InteractiveScanner) DumpBytecode(bytecode *compiler.Bytecode) {
	s.WriteLines([]string{
		"-------- CODE --------",
		bytecode.Instructions.String(),
		"-------- DATA --------",
	})
	for i, data := range bytecode.Constants {
		switch data := data.(type) {
		case *object.CompiledFunction:
			s.Printf("FN %d:\n%s", i, data.Instructions.String())

		default:
			s.Printf("OBJECT %d: %s\n", i, data.Inspect())
		}
	}
}
