// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/trajectory/machineTypeAdd": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Trajectory_轨迹属性控制器"
                ],
                "summary": "添加轨迹绑定机型",
                "parameters": [
                    {
                        "description": "机型请求体",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Machinetype"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/trajectory/machineTypeSearch": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Trajectory_轨迹属性控制器"
                ],
                "summary": "按照lastappeared_id,machinetype 查询机型",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "轨迹Id",
                        "name": "lastappeared_ids",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "机型",
                        "name": "machinetype",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/trajectory/machineTypeUpdate": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Trajectory_轨迹属性控制器"
                ],
                "summary": "更新机型",
                "parameters": [
                    {
                        "description": "机型请求体",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Machinetype"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/trajectory/objecttrajactory/{lastappeared_id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Trajectory_轨迹属性控制器"
                ],
                "summary": "按照lastappeared_id查询轨迹",
                "parameters": [
                    {
                        "type": "string",
                        "description": "轨迹Id",
                        "name": "lastappeared_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "entity.Machinetype": {
            "type": "object",
            "properties": {
                "lastappeared_id": {
                    "type": "integer"
                },
                "machinetype": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:2345",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "POSTGRESQL-SERVER-GOLANG",
	Description:      "使用golang语言,整合Swagger实现简单Postgres整合Gis服务",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}