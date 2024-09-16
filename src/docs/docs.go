// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mike Zinyoni",
            "email": "mzinyoni7"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1": {
            "get": {
                "description": "Gets Index",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Index"
                ],
                "summary": "Index",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/rooms/check-in": {
            "post": {
                "description": "Enter room",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Check In",
                "parameters": [
                    {
                        "description": "CheckIn JSON",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.CheckInBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RoomUserModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error400Response"
                        }
                    }
                }
            }
        },
        "/api/v1/rooms/extend": {
            "patch": {
                "description": "Update users stay in room",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Extend",
                "parameters": [
                    {
                        "description": "CheckIn JSON",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.ExtendBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RoomUserModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error400Response"
                        }
                    }
                }
            }
        },
        "/api/v1/rooms/rooms-users": {
            "get": {
                "description": "Get rooms and users in them",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Rooms Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.RoomUsers"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.RoomUserModel": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "expires": {
                    "description": "Room    RoomModel ` + "`" + `json:\"-\"` + "`" + `",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "roomID": {
                    "description": "User    UserModel ` + "`" + `json:\"-\"` + "`" + `",
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "responses.Error400Response": {
            "type": "object",
            "properties": {
                "error": {}
            }
        },
        "structs.CheckInBody": {
            "type": "object",
            "required": [
                "roomUid",
                "userUid"
            ],
            "properties": {
                "roomUid": {
                    "type": "string"
                },
                "userUid": {
                    "type": "string"
                }
            }
        },
        "structs.ExtendBody": {
            "type": "object",
            "required": [
                "expires",
                "userRoomID"
            ],
            "properties": {
                "expires": {
                    "type": "string"
                },
                "userRoomID": {
                    "type": "integer"
                }
            }
        },
        "structs.RoomUsers": {
            "type": "object",
            "required": [
                "roomID",
                "roomUid",
                "users"
            ],
            "properties": {
                "roomID": {
                    "type": "integer"
                },
                "roomUid": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/structs.UserInRoom"
                    }
                }
            }
        },
        "structs.UserInRoom": {
            "type": "object",
            "required": [
                "created_at",
                "expires",
                "id",
                "userID",
                "userName",
                "userUid"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "expires": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "userID": {
                    "type": "integer"
                },
                "userName": {
                    "type": "string"
                },
                "userUid": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "W2W-GO",
	Description:      "W2W-GO",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}