package main

import "fmt"

func main() {
	l := LAP{}
	l.addTrigger(PrismaEnum{})

	source := `
		enum Foo {
    		BAR
    		BAZ
		}

		// Hello
	`

	sourceFile := SourceFile{
		name:    "schema.prisma",
		content: source,
	}

	cgs := l.genAll(sourceFile)
	for _, cg := range cgs {
		fmt.Printf("cg.outputPath: %s", cg.outputPath)
		fmt.Printf("len(cg.code): %d", len(cg.code))
	}
}
