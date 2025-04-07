package main

import "fmt"

func main() {
	e := NewGenEngine()
	e.addTriggers(PrismaEnum{})

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

	cgs := e.gen(sourceFile)
	for _, cg := range cgs {
		fmt.Printf("cg.outputPath: %s", cg.outputPath)
		fmt.Printf("len(cg.code): %d", len(cg.code))
	}
}
