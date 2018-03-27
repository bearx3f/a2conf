package a2conf

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// FormatOption .
type FormatOption struct {
	Tab2Space bool
	TabSize   int
}

type walkingResource struct {
	FormatOption
	depth int
}

func (wr *walkingResource) GetIndent() string {
	tab := "\t"
	tabsize := 1
	if wr.Tab2Space {
		tab = " "
		tabsize = wr.TabSize
	}
	return strings.Repeat(tab, tabsize*wr.depth)
}

func (wr *walkingResource) StartBlock() string {
	indent := wr.GetIndent()
	wr.depth++
	return indent
}

func (wr *walkingResource) EndBlock() string {
	wr.depth--
	return wr.GetIndent()
}

/*
<VirtualHost *:80>
	<Location / >
	</Location>
</VirtualHost>
:#:
<VirtualHost *:80>
	<Location />
	</Location>
</VirtualHost>
*/

// Format .
func Format(in io.Reader, fmtOpt FormatOption) (string, error) {
	var indent string

	buf := new(bytes.Buffer)
	wkr := &walkingResource{fmtOpt, 0}

	sc := bufio.NewScanner(in)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if len(line) == 0 {
			buf.WriteByte('\n')
			continue
		}
		if line[0] == '<' {
			if line[1] == '/' {
				indent = wkr.EndBlock()
			} else {
				indent = wkr.StartBlock()
			}
		} else {
			indent = wkr.GetIndent()
		}
		buf.WriteString(
			fmt.Sprintf("%s%s\n",
				indent,
				line,
			),
		)
	}

	return buf.String(), nil
}
