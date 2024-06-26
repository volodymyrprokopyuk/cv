package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"text/template"

	"github.com/yuin/goldmark"
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

var reCleanPhone = regexp.MustCompile(`[^\d\+]`)
var reRemovePar = regexp.MustCompile(`<p>|</p>`)

var tplFuncs = template.FuncMap{
  "phone": func(phone string) string {
    return reCleanPhone.ReplaceAllString(phone, "")
  },
  "md": func(md string) string {
    var htmlBuf bytes.Buffer
    err := goldmark.Convert([]byte(md), &htmlBuf)
    if err != nil {
      panic(err)
    }
    return reRemovePar.ReplaceAllString(htmlBuf.String(), "")
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
  _, err = tpl.ParseFiles("cv2.html")
  if err != nil {
    return err
  }
  w, err := os.Create("index2.html")
  if err != nil {
    return err
  }
  defer w.Close()
  err = tpl.ExecuteTemplate(w, "cv2.html", cv)
  if err != nil {
    return err
  }
  twCmd := exec.Command(
    "bunx", "tailwindcss", "--input", "cv2.css", "--output", "tw.css",
  )
  twCmd.Stdout = os.Stdout
  twCmd.Stderr = os.Stderr
  err = twCmd.Run()
  if err != nil {
    return err
  }
  // minCmd := exec.Command("minify-html", "index2.html", "tw.css")
  // minCmd.Stdout = os.Stdout
  // minCmd.Stderr = os.Stderr
  // return minCmd.Run()
  return nil
}

func main() {
  err := render()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
