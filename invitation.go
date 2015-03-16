package main
import (
"log"
"os"
"text/template"
)
func main() {

const response = `
Dear {{.Honorific}}{{.Name}},
{{if .Attended}}
It was a pleasure to see you at the fundraiser.{{else}}
We missed you at the fundraiser.{{end}}
{{if .Donated}}
Thanks for the generous donation.
Reminder for the coming event:{{else}}
Reminder for the coming event:{{end}}
{{range .Events}}{{.}}
{{end}}
Warm Regards,
Akashdeep
`

type Person struct {
Honorific, Name string
Attended, Donated bool
Events []string
}
var people = []Person{
{"Mrs.", "Wilson", true, true, []string{"RedCross DATE:11/14", "WHO seminar DATE: 12/3", "Somilia Relief Fund  DATE:1/22", "UNO conference  DATE:3/3", "Disaster relief  DATE:4/2 "}},
{"Mr.", "Singh", true, false, []string{"WHO seminar  DATE:12/3", "Somilia Relief Fund  DATE:1/22", "UNO conference  DATE:3/3"}},
{"Ms.", "Preeti", false, false, []string{"WHO seminar  DATE:12/3"}},
}

t := template.Must(template.New("response").Parse(response))

for _, r := range people {
err := t.Execute(os.Stdout, r)
if err != nil {
log.Println("executing template:", err)
}
}
}
