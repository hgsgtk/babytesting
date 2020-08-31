// TODO exec this main function by another main
package main

import (
	"os"
	"reflect"
	"testing"

	"github.com/hgsgtk/babytesting/internal/testdeps"
	// ref: https://github.com/golang/go/blob/bb998747d6c5213e3a366936c482e149dce62720/src/cmd/go/internal/load/test.go#L616
	// Todo import path
	//{{if .ImportTest}}
	//{{if .NeedTest}}_test{{else}}_{{end}} {{.Package.ImportPath | printf "%q"}}
	//{{end}}
	//{{if .ImportXtest}}
	//{{if .NeedXtest}}_xtest{{else}}_{{end}} {{.Package.ImportPath | printf "%s_test" | printf "%q"}}
	//{{end}}
	//{{if .Cover}}
	//{{range $i, $p := .Cover.Vars}}
	//_cover{{$i}} {{$p.Package.ImportPath | printf "%q"}}
	//{{end}}
	//{{end}}
)

var tests = []testing.InternalTest{
	// Todo gather tests
	//{{range .Tests}}
	//{"{{.Name}}", {{.Package}}.{{.Name}}},
	//{{end}}
}

var benchmarks = []testing.InternalBenchmark{
	// Todo gather benchmarks
	//{{range .Benchmarks}}
	//{"{{.Name}}", {{.Package}}.{{.Name}}},
	//{{end}}
}

var examples = []testing.InternalExample{
	// Todo gather examples
	//{{range .Examples}}
	//{"{{.Name}}", {{.Package}}.{{.Name}}, {{.Output | printf "%q"}}, {{.Unordered}}},
	//{{end}}
}

func init() {
	// Todo inportpath
	//testdeps.ImportPath = {{.ImportPath | printf "%q"}}
}

// Todo coverage report
//{{if .Cover}}
//
//// Only updated by init functions, so no need for atomicity.
//var (
//	coverCounters = make(map[string][]uint32)
//	coverBlocks = make(map[string][]testing.CoverBlock)
//)
//
//func init() {
//	{{range $i, $p := .Cover.Vars}}
//	{{range $file, $cover := $p.Vars}}
//	coverRegisterFile({{printf "%q" $cover.File}}, _cover{{$i}}.{{$cover.Var}}.Count[:], _cover{{$i}}.{{$cover.Var}}.Pos[:], _cover{{$i}}.{{$cover.Var}}.NumStmt[:])
//	{{end}}
//	{{end}}
//}
//
//func coverRegisterFile(fileName string, counter []uint32, pos []uint32, numStmts []uint16) {
//	if 3*len(counter) != len(pos) || len(counter) != len(numStmts) {
//		panic("coverage: mismatched sizes")
//	}
//	if coverCounters[fileName] != nil {
//		// Already registered.
//		return
//	}
//	coverCounters[fileName] = counter
//	block := make([]testing.CoverBlock, len(counter))
//	for i := range counter {
//		block[i] = testing.CoverBlock{
//			Line0: pos[3*i+0],
//			Col0: uint16(pos[3*i+2]),
//			Line1: pos[3*i+1],
//			Col1: uint16(pos[3*i+2]>>16),
//			Stmts: numStmts[i],
//		}
//	}
//	coverBlocks[fileName] = block
//}
//{{end}}

func main() {
	// TODO coverage report
	//{{if .Cover}}
	//testing.RegisterCover(testing.Cover{
	//	Mode: {{printf "%q" .Cover.Mode}},
	//	Counters: coverCounters,
	//	Blocks: coverBlocks,
	//	CoveredPackages: {{printf "%q" .Covered}},
	//})
	//{{end}}
	m := testing.MainStart(testdeps.TestDeps{}, tests, benchmarks, examples)
	//{{with .TestMain}}
	//{{.Package}}.{{.Name}}(m)
	// Fixme only support TestMain
	os.Exit(int(reflect.ValueOf(m).Elem().FieldByName("exitCode").Int()))
	//{{else}}
	//os.Exit(m.Run())
	//{{end}}
}
