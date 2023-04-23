package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --storage gremlin --idtype string --template ./templates/create.tmpl --template ./templates/query.tmpl --template ./templates/decode.tmpl ./schema
