import { readFile, writeFile } from "fs/promises"
import { load } from "js-yaml"
import { marked } from "marked"
import njk from "nunjucks"

const data = load(await readFile("vlad.yaml"))
const env = njk.configure(".", { autoescape: false })
env.addFilter("md", marked.parseInline)
const doc = env.render("template.njk", { cv: data })
await writeFile("index.html", doc)
