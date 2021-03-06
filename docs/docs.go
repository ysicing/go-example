// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "ysicing",
            "url": "http://github.com/ysicing",
            "email": "i@ysicing.me"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/apis/example.dev/v1beta/db/total": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "示例API"
                ],
                "summary": "查看DB大小",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/err500": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "默认"
                ],
                "summary": "errpage",
                "responses": {
                    "500": {}
                }
            }
        },
        "/errpanic": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "默认"
                ],
                "summary": "errpanic",
                "responses": {
                    "500": {}
                }
            }
        },
        "/gentoken": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "默认"
                ],
                "summary": "生成测试Token",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/health.User"
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/health": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "默认"
                ],
                "summary": "health",
                "responses": {
                    "200": {}
                }
            }
        },
        "/version": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "默认"
                ],
                "summary": "version",
                "responses": {
                    "200": {}
                }
            }
        }
    },
    "definitions": {
        "health.User": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                },
                "userrole": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.2",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Go Example API",
	Description: "This is a sample server Petstore server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
