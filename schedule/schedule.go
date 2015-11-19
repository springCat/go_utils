package schedule

import (
	"fmt"
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	"html/template"
)

type Jobs func()

func (self Jobs) Run() {
	self()
}

var Me schedule

type schedule int

func (s schedule) Start() schedule {
	jobrunner.Start()
	return s
}

func (s schedule) Monitor(e *gin.Engine,path,username,password string) schedule {
	r := e.RouterGroup
	m := gin.Accounts{
		username:password,
	}
	r.Use(gin.BasicAuth(m))
	r.GET(path, JobHtml)
	return s
}

func (s schedule) Schedule(rate string, f interface{}) schedule {
	jobrunner.Schedule("@every 10s", Jobs(f.(func())))
	return s
}

func JobHtml(c *gin.Context) {
	Template := template.New("Status")
	s, err := Template.Parse(status)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	s.Execute(c.Writer,jobrunner.StatusPage())
}
//copy from Status.html
var status = `
<html>
	<head>
		<style>
body {
	margin: 30px 0 0 0;
  font-size: 11px;
  font-family: sans-serif;
  color: #345;

}
h1 {
	font-size: 24px;
	text-align: center;
	padding: 10 0 30px;
}
table {
	/*max-width: 80%;*/
	margin: 0 auto;
  border-collapse: collapse;
  border: none;
}
table td, table th {
	min-width: 25px;
	width: auto;
  padding: 15px 20px;
  border: none;
}


table tr:nth-child(odd) {
  background-color: #f0f0f0;
}
table tr:nth-child(1) {
  background-color: #345;
  color: white;
}
th {
  text-align: left;
}

		</style>
	</head>
	<body>

<h1>JobRunner Status Report</h1>

<table>
	<tr><th>ID</th><th>Name</th><th>Status</th><th>Last run</th><th>Next run</th><th>Latency</th></tr>
{{range .}}

	<tr>
		<td>{{.Id}}</td>
		<td>{{.JobRunner.Name}}</td>
		<td>{{.JobRunner.Status}}</td>
		<td>{{if not .Prev.IsZero}}{{.Prev.Format "2006-01-02 15:04:05"}}{{end}}</td>
		<td>{{if not .Next.IsZero}}{{.Next.Format "2006-01-02 15:04:05"}}{{end}}</td>
		<td>{{.JobRunner.Latency}}</td>
	</tr>
{{end}}
</table>
`
