package main

import (
	// "bytes"
	"fmt"
	"html/template"
	"os"
	"regexp"

	// "github.com/yuin/goldmark"
	// "github.com/yuin/goldmark/parser"
	// "github.com/yuin/goldmark/renderer/html"
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
  Industries string `yaml:"industries"`
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

var tplFuncs = template.FuncMap{
  "phone": func(phone string) string {
    re := regexp.MustCompile(`[^\d\+]`)
    return re.ReplaceAllString(phone, "")
  },
  "md": func(md string) string {
    return md
    // var htmlBuf bytes.Buffer
    // gmark := goldmark.New(
    //   // goldmark.WithParserOptions(parser.WithInlineParsers(bs ...util.PrioritizedValue))
    //   goldmark.WithRenderer(r renderer.Renderer)
    //   goldmark.WithRendererOptions(html.WithUnsafe())
    // )
    // err := gmark.Convert([]byte(md), &htmlBuf)
    // if err != nil {
    //   panic(err)
    // }
    // return htmlBuf.String()
  },
}

func render() error {
  file, err := os.Open("cv.yaml")
  if err != nil {
    return err
  }
  defer file.Close()
  var cv cv
  err = yaml.NewDecoder(file).Decode(&cv)
  if err != nil {
    return err
  }
  tpl := template.New("cv")
  tpl.Funcs(tplFuncs)
  _, err = tpl.ParseFiles("cv.html")
  if err != nil {
    return err
  }
  w, err := os.Create("index2.html")
  if err != nil {
    return err
  }
  defer w.Close()
  return tpl.ExecuteTemplate(w, "cv.html", cv)
}

func main() {
  err := render()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  // var htmlBuf bytes.Buffer
  // err := goldmark.Convert([]byte("one **two**"), &htmlBuf)
  // if err != nil {
  //   fmt.Println(err)
  //   return
  // }
  // fmt.Println(htmlBuf.String())
}
