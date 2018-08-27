package cluster

import (
	"flag"
	"fmt"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/kubernauts/tk8/internal/templates"
	"github.com/spf13/viper"
)

type Config struct {
	ClusterName string
	SSHName     string
}

func namer(name string) Config {
	return Config{
		ClusterName: name,
		SSHName:     name,
	}
}

func generateName() string {
	var (
		words     = flag.Int("words", 2, "The number of words in generated name")
		separator = flag.String("separator", "", "The separator between words in the name")
	)
	flag.Parse()
	generatedName := petname.Generate(*words, *separator)
	return generatedName
}

func CreateConfig() {
	generatedName := generateName()
	fmt.Printf("\nNo default config was provided. Generating one for you...\n")
	parseTemplate(templates.Config, "./config.yaml", namer(generatedName))
	ReadViperConfigFile("config")
	region := viper.GetString("aws.aws_default_region")
	fmt.Printf("\nCluster Name:\t%s\nSSH Key name:\t%s\nAWS Region:\t%s\n", generatedName, generatedName, region)
	CreateSSHKey(generatedName, region)
}
