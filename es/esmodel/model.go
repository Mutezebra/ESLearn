package esmodel

type ESModels interface {
	Mapping() string
	Index() string
}

type User struct {
	ID       uint     `json:"id"`   // integer
	Name     string   `json:"name"` // text
	Tag      []string `json:"tag"`  // keyword
	Age      int      `json:"age"`
	CreateAt string   `json:"create_at"` // date
}

func (*User) Index() string {
	return "user_index"
}

func (*User) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "id": {
        "type": "integer"
      },
      "name": {
        "analyzer": "ik_smart",
        "type": "text"
      },
      "age": {
	    "type": "integer"
      },
      "tag": {
        "type": "keyword"
      },
      "create_at": {
        "type": "date",
        "null_value": null,
        "format": ["yyyy-MM-dd HH:mm:ss"]
      }
    }
  }
}
`
}
