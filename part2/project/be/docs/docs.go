// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/health": {
            "get": {
                "description": "Check API health",
                "produces": [
                    "application/json"
                ],
                "summary": "Check API health",
                "responses": {}
            }
        },
        "/todos": {
            "get": {
                "description": "Get ToDo's list",
                "produces": [
                    "application/json"
                ],
                "summary": "Get ToDo's list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ToDo"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Post ToDo",
                "produces": [
                    "application/json"
                ],
                "summary": "Post ToDo",
                "parameters": [
                    {
                        "description": "ToDo json payload",
                        "name": "ToDo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RawToDo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ToDo"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "description": "Get ToDo by id",
                "produces": [
                    "application/json"
                ],
                "summary": "Get ToDo by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ToDo Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ToDo"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete ToDo",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete ToDo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the ToDo to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ToDo"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update ToDo",
                "produces": [
                    "application/json"
                ],
                "summary": "Update ToDo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the ToDo to update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "ToDo updated body",
                        "name": "ToDo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RawToDo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ToDo"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.RawToDo": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "isDone": {
                    "type": "boolean"
                }
            }
        },
        "models.ToDo": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isDone": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
