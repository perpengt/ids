# ids

Random ID library written in Go language.

# Installtion

```
go get https://github.com/perpengt/ids
```

# Usage

```go
// Generate the ID
var id = ids.GenerateID()

// Or create from byte slice
var id = ids.New([]byte{ ... })

// Validate id
var err = id.Valid()

// Stringify
var str = id.String()
var uri = id.URIString()

// You can also create key
var key = id.Key()

// Key can be used for index of map
var m map[ids.Key]string
```
