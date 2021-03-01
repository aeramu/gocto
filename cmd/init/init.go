package init

import (
	"encoding/json"
	"fmt"
	"github.com/aeramu/gocto/pkg/schema"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func Run(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatalln("module path for go module needed")
	}
	f, err := os.Create("schema.json")
	if err != nil {
		log.Fatalln("can't create schema.json file")
	}
	s := schema.Schema{
		ModulePath: args[0],
		Entity:     []string{"Foo", "Bar"},
		Service:    []string{"Foo", "Bar"},
		Adapter:    []string{"FooRepository", "BarClient"},
	}
	b, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		log.Fatalln("failed marshal schema to json")
	}
	if _, err := f.Write(b); err != nil {
		log.Fatalln("failed write schema to file")
	}
	f, err = os.Create("Makefile")
	if err != nil {
		log.Fatalln("failed create makefile")
	}
	b = []byte(fmt.Sprintf(
		"generate:\n" +
			"\tgo mod init %s\n" +
			"\tgocto generate\n" +
			"\tmockery --all\n" +
			"\tgo mod tidy\n" +
			"mock:\n" +
			"\tmockery --all\n" +
			"test:\n" +
			"\tgo test ./...\n",
			args[0]))
	if _, err := f.Write(b); err != nil {
		log.Fatalln("failed write makefile")
	}
}
