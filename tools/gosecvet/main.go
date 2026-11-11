// Command gosecvet runs the gosec rules through gosec's own go/analysis entry point.
//
// The gosec command itself is not used: from v2.23.0 on it imports AI vendor SDKs whose
// build does not fit the development PC's memory (brief F-00, section 3.3).
package main

import (
	"github.com/securego/gosec/v2/goanalysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(goanalysis.Analyzer) }
