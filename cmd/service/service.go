package service

import (
	"encoding/json"
	"github.com/aeramu/gocto/pkg/generator"
	"github.com/aeramu/gocto/pkg/schema"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

func Add(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatalln("method name required")
	}
	f, err := os.Open("schema.json")
	if err != nil {
		log.Fatalln(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}
	var s schema.Schema
	if err := json.Unmarshal(b, &s); err != nil {
		log.Fatalln(err)
	}
	g := generator.Generator{
		ModulePath:     s.ModulePath,
		EntitiesName:   s.Entity,
		MethodsName:    s.Service,
		InterfacesName: s.Adapter,
	}
	log.Println("Generating...")
	g.AddService(args[0])
	log.Println("Done.")
}
