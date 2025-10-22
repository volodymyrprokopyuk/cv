package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"text/template"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
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
  Technologies []technology `yaml:"technologies"`
}

type technology struct {
  Tech string `yaml:"tech"`
  Desc string `yaml:"desc"`
  Logo string `yaml:"logo"`
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
  Highlights []string `yaml:"highlights"`
  Email string `yaml:"email"`
  Phone string `yaml:"phone"`
  Location string `yaml:"location"`
  URL string `yaml:"URL"`
  LinkedIn string `yaml:"linkedin"`
  Languages []string `yaml:"languages"`
  Industries []string `yaml:"industries"`
  Selfie string `yaml:"selfie"`
  Summary string `yaml:"summary"`
  Profiles []profile `yaml:"profiles"`
  Capabilities []string `yaml:"capabilities"`
  Achievements []achievement `yaml:"achievements"`
  Employment []employment `yaml:"employment"`
  Technologies []technology `yaml:"technologies"`
  Education []education `yaml:"education"`
  Portfolio bool
}

var reCleanPhone = regexp.MustCompile(`[^\d\+]`)
var reRemovePar = regexp.MustCompile(`<p>|</p>`)

var tplFuncs = template.FuncMap{
  "phone": func(phone string) string {
    return reCleanPhone.ReplaceAllString(phone, "")
  },
  "md": func(md string) string {
    gm := goldmark.New(goldmark.WithRendererOptions(html.WithUnsafe()))
    var htmlBuf bytes.Buffer
    err := gm.Convert([]byte(md), &htmlBuf)
    if err != nil {
      panic(err)
    }
    return reRemovePar.ReplaceAllString(htmlBuf.String(), "")
  },
}

func render() error {
  // decode YAML
  file, err := os.Open("cv.yaml")
  if err != nil {
    return err
  }
  defer func() {
    _ = file.Close()
  }()
  var cv cv
  err = yaml.NewDecoder(file).Decode(&cv)
  if err != nil {
    return err
  }
  // read template
  tpl := template.New("cv")
  tpl.Funcs(tplFuncs)
  _, err = tpl.ParseFiles("cv.html")
  if err != nil {
    return err
  }
  // write index.html
  w, err := os.Create("index.html")
  if err != nil {
    return err
  }
  defer func() {
    _ = file.Close()
  }()
  err = tpl.ExecuteTemplate(w, "cv.html", cv)
  if err != nil {
    return err
  }
  // write portfolio/index.html
  cv.Portfolio = true
  cv.Title = "Volodymyr Portfolio"
  w, err = os.Create("portfolio/index.html")
  if err != nil {
    return err
  }
  defer func() {
    _ = file.Close()
  }()
  err = tpl.ExecuteTemplate(w, "cv.html", cv)
  if err != nil {
    return err
  }
  // generate CSS
  twCmd := exec.Command(
    "bunx", "@tailwindcss/cli", "--input", "cv.css", "--output", "tw.css", "--minify",
  )
  twCmd.Stdout = os.Stdout
  twCmd.Stderr = os.Stderr
  return twCmd.Run()
}

func main() {
  err := render()
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
