package generate

import (
	"encoding/json"
	"gocto/pkg/generator"
	"gocto/pkg/schema"

	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	f, _ := os.Open("schema.json")
	b, _ := ioutil.ReadAll(f)
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
	os.Mkdir("entity", 0755)
	os.Mkdir("service", 0755)
	os.Mkdir("service/api", 0755)
	os.Mkdir("repository", 0755)
	os.Mkdir("handler", 0755)
	os.Mkdir("cmd", 0755)
	g.Generate()
	log.Println("Done.")
}
