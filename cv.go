package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type profile struct {
  Name string `yaml:"name"`
  URL string `yaml:"URL"`
  Producer string `yaml:"producer"`
  ProducerURL string `yaml:"producerURL"`
  Date string `yaml:"date"`
}

type achievement struct {
  Project string `yaml:"project"`
  Logo string `yaml:"logo"`
  AuxLogo string `yaml:"auxLogo"`
  STAR []string `yaml:"STAR"`
}

type tech struct {
  Tech string `yaml:"tech"`
  Desc string `yaml:"desc"`
}

type employment struct {
  Position string `yaml:"position"`
  Contract string `yaml:"contract"`
  Duration string `yaml:"duration"`
  Summary string `yaml:"summary"`
  Company string `yaml:"company"`
  URL string `yaml:"URL"`
  Location string `yaml:"location"`
  Description string `yaml:"description"`
  Industry string `yaml:"industry"`
  Logo string `yaml:"logo"`
  AuxLogo string `yaml:"auxLogo"`
  Responsibilities []string `yaml:"responsibilities"`
  Skills []string `yaml:"skills"`
  Technologies []tech `yaml:"technologies"`
}

type course struct {
  Title string `yaml:"title"`
  Duration string `yaml:"duration"`
}

type education struct {
  Courses []course `yaml:"courses"`
  Center string `yaml:"center"`
  URL string `yaml:"URL"`
  Location string `yaml:"location"`
  Logo string `yaml:"logo"`
}

type cv struct {
  Title string `yaml:"title"`
  Name string `yaml:"name"`
  Role string `yaml:"role"`
  Insudstries string `yaml:"insudstries"`
  Email string `yaml:"email"`
  Phone string `yaml:"phone"`
  Location string `yaml:"location"`
  URL string `yaml:"URL"`
  LinkedIn string `yaml:"linkedin"`
  Languages []string `yaml:"languages"`
  Selfie string `yaml:"selfie"`
  Summary string `yaml:"summary"`
  Profiles []profile `yaml:"profiles"`
  Capabilities []string `yaml:"capabilities"`
  Achievements []achievement `yaml:"achievements"`
  Employment []employment `yaml:"employment"`
  Education []education `yaml:"education"`
}

func exitOnError(err error) {
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func main() {
  file, err := os.Open("vlad.yaml")
  exitOnError(err)
  defer file.Close()
  var cv cv
  err = yaml.NewDecoder(file).Decode(&cv)
  exitOnError(err)
  fmt.Println(cv)
}
