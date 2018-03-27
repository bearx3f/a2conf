package a2conf

import (
	"bufio"
	"io"
	"strings"
)

/*
<VirtualHost *:80>
	ServerName example.com
	ServerAlias www.example.com

	<Location "/">

	</Location>
</VirtualHost>
*/

// Parse .
func Parse(in io.Reader) (tree *Tree, err error) {
	tree = new(Tree)
	sc := bufio.NewScanner(in)
	for sc.Scan() {
		lineText := strings.TrimSpace(sc.Text())
		if len(lineText) == 0 {
			continue
		}
	}
	return
}

/*
<Location / >ProxyReverse /</Location>
:
<Location / >
	ProxyReverse /
</Location>
*/
func parse(in bufio.Scanner, tree *Tree) error {
	return nil
}
