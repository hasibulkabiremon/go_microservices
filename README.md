# Go MicroService :RESTfull 
## Struct data to json convert:
### json.Marshal(--struct--) return []byte , err

    - get response in cmd : 1. Run main.go
    -                        2. curl localhost:9090 (from another console)
    - For better format: curl localhost:9090 | jq
    - Struct in go accept annotation like : Details: https://pkg.go.dev/encoding/json#Marshal
                                            - // Field appears in JSON as key "myName".
                                            -    Field int `json:"myName"`

                                            - // Field appears in JSON as key "myName" and
                                            - // the field is omitted from the object if - - its value is empty,
                                            - // as defined above.
                                            - Field int `json:"myName,omitempty"`

                                            - // Field appears in JSON as key "Field" (the - default), but
                                            - // the field is skipped if empty.
                                            - // Note the leading comma.
                                            - Field int `json:",omitempty"`

                                            - // Field is ignored by this package.
                                            - Field int `json:"-"`

                                            - // Field appears in JSON as key "-".
                                            - Field int `json:"-,"`
    - Use Json Encoder
