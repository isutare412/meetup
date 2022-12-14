// Package api GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package api

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
        "/api/v1/users": {
            "post": {
                "description": "Create an user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "description": "Request to create an user",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.createUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.createUserResp"
                        }
                    },
                    "default": {
                        "description": "Somethings got wrong",
                        "schema": {
                            "$ref": "#/definitions/http.errorResp"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{userId}": {
            "get": {
                "description": "Get an user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "number",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.getUserResp"
                        }
                    },
                    "default": {
                        "description": "Somethings got wrong",
                        "schema": {
                            "$ref": "#/definitions/http.errorResp"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "number",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.deleteUserResp"
                        }
                    },
                    "default": {
                        "description": "Somethings got wrong",
                        "schema": {
                            "$ref": "#/definitions/http.errorResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.createUserReq": {
            "type": "object",
            "properties": {
                "nickname": {
                    "type": "string",
                    "example": "redshore"
                }
            }
        },
        "http.createUserResp": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string",
                    "example": "2022-08-20T18:54:53.965295+09:00"
                },
                "id": {
                    "type": "integer",
                    "example": 412
                },
                "nickname": {
                    "type": "string",
                    "example": "redshore"
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2022-08-20T18:54:53.965295+09:00"
                }
            }
        },
        "http.deleteUserResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 412
                }
            }
        },
        "http.errorResp": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        },
        "http.getUserResp": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string",
                    "example": "2022-08-20T18:54:53.965295+09:00"
                },
                "id": {
                    "type": "integer",
                    "example": 412
                },
                "nickname": {
                    "type": "string",
                    "example": "redshore"
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2022-08-20T18:54:53.965295+09:00"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Meetup Gateway",
	Description:      "API gateway for Meetup project.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
