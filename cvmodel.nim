import std/[os, streams]
import pkg/[yaml/serialization, markdown, regex, nimja/parser]

type
  Education = object
    courses: seq[tuple[title, duration: string]]
    center: string
    url: string
    location: string
    logo: string

  Employment = object
    position: string
    contract: string
    duration: string
    summary: string
    company: string
    url: string
    location: string
    description: string
    industry: string
    logo: string
    auxlogo: string
    # responsibilities: seq[string]
    # skills: seq[string]
    # technologies: seq[string]

  CV = object
    title: string
    name: string
    role: string
    email: string
    phone: string
    url: string
    location: string
    summary: string
    selfie: string
    languages: seq[string]
    keyCapabilities: seq[string]
    educationList: seq[Education]
    employmentList: seq[Employment]

proc load(file: string, cv: var CV) =
  let s = newFileStream(file, fmRead)
  defer: s.close
  s.load(cv)

proc md(doc: string): string =
  markdown(doc).replace(re"^<p>|</p>\n$", "")

proc render(cv: CV): string =
  compileTemplateFile(getScriptDir() / "cvtemplate.nwt")

proc render(cv: CV, file: string) =
  file.writeFile(render(cv))

var cv: CV
"cvvlad.yaml".load(cv)
cv.render("index.html")
