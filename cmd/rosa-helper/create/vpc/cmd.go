package vpc

import (
	"fmt"

	"github.com/spf13/cobra"
)

var args struct {
	region       string
	name         string
	cidr         string
	tags         string
	findExisting bool
}
var Cmd = &cobra.Command{
	Use:   "vpc",
	Short: "Create vpc",
	Long:  "Create vpc.",
	Example: `  # Create a vpc named "myvpc"
  rosa-helper create vpc --name=myvpc --region us-east-2`,

	Run: run,
}

func init() {
	flags := Cmd.Flags()
	flags.SortFlags = false
	flags.StringVarP(
		&args.name,
		"name",
		"n",
		"",
		"Name of the vpc",
	)
	flags.StringVarP(
		&args.region,
		"region",
		"",
		"",
		"Region of the vpc",
	)
	flags.StringVarP(
		&args.cidr,
		"cidr",
		"",
		"",
		"cidr of the vpc",
	)
	flags.StringVarP(
		&args.tags,
		"tags",
		"",
		"",
		"tags of the vpc, fmt tagName:tagValue,tagName2:tagValue2",
	)
	flags.BoolVarP(
		&args.findExisting,
		"find-existing",
		"",
		false,
		"Find the vpc with same name from current region. if not exsiting, create a new one",
	)
}
func run(cmd *cobra.Command, _ []string) {
	fmt.Println("creating VPC ...")
}
