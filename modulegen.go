package main

import (
    "os"
    "fmt"
    "path"
    "strings"
    "text/template"
)

type Package struct {
    Name        string
    Version     string
    Rootdir     string
} 

var module_template = `--- Metadata
-- Installed by: Matthew Hartley <Matthew.Hartley@jic.ac.uk>
-- From: linuxbrew {{.Name}} install
-- License: N/A
-- Test: 

whatis("Version: {{.Version}}")
whatis("Keywords: ")
whatis("Description: ")

-- Set the paths
prepend_path("PATH",                "{{.Rootdir}}/bin")
prepend_path("LD_LIBRARY_PATH",     "{{.Rootdir}}/lib")
prepend_path("LIBRARY_PATH",        "{{.Rootdir}}/lib")
prepend_path("CPATH",               "{{.Rootdir}}/include")
prepend_path("CPLUS_INCLUDE_PATH",  "{{.Rootdir}}/include")
prepend_path("MANPATH",             "{{.Rootdir}}/share/man")
prepend_path("PKG_CONFIG_PATH",     "{{.Rootdir}}/pkgconfig")
`

func find_package_details(root_path string) Package {

    cleanpath := path.Clean(root_path)

    s := strings.Split(cleanpath, "/")

    version := string(s[len(s)-1])
    name := string(s[len(s)-2])

    target := Package{name, version, cleanpath}   

    return target
}

func main() {

    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %v package_path\n", os.Args[0])
        os.Exit(2)
    }

    dirpath := os.Args[1]
    target := find_package_details(dirpath)

    tmpl, err := template.New("test").Parse(module_template)
    err = tmpl.Execute(os.Stdout, target)
    if err != nil { panic(err) }


}
