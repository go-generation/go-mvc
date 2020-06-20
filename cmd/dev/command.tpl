package gomvc

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/hoisie/mustache"
	"github.com/spf13/cobra"
)

// {{TitleName}} is the cli command that creates new {{Name}}
var {{TitleName}} = &cobra.Command{
	Use:   "{{Name}}",
	Short: "Generate {{Name}} files",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a name for your new {{Name}}")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		log.Infof("preparing to create a new {{Name}}: %s\n", name)
		/*
				ADD LOGIC HERE OR USE/MODIFY BELOW
		*/
		// example boilerplate below
		baseTemplateDir := "./templates"
		templateDir := fmt.Sprintf("%s/{{Name}}.tpl", baseTemplateDir)
		data := map[string]string{
			"Name":      name,
			"TitleName": strings.Title(name),
		}
		r := mustache.RenderFile(templateDir, data)
		destPath := fmt.Sprintf("%s/%s.go", ".", name)
		if err := CreateFileFromString(destPath, r); err != nil {
			panic(err)
		}
	},
}
