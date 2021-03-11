package templates

var protoTmpl = `syntax = "proto3";
//go:generate generator . Proto

package protos;

option go_package = "{{.AppName}}/service";

service {{.Name}}Service {
    rpc Create{{.Name}}({{.Name}}) returns ({{.Name}}Response) {}
}

message {{.Name}} {
    string public_id = 1;
    {{range $i, $f := .Fields}} {{if $f.ProtoField}} 
    {{$f.DataType}} {{$f.ProtoField}} = {{inc $i}}; {{end}}
	{{end}}
}

message {{.Name}}Response {
    bool created = 1;
    {{.Name}} entity = 2;
}
`
