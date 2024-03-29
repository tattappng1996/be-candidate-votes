// Package v1 Code generated by swaggo/swag. DO NOT EDIT
package v1

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
        "/": {
            "get": {
                "description": "Just HealthCheck",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Generals"
                ],
                "summary": "HealthCheck",
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/candidate": {
            "put": {
                "description": "You can use this API to update/edit candidate in the system.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Private (Candidate)"
                ],
                "summary": "UpdateCandidate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer ${token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateCandidateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GeneralResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "You can use this API to create candidates in the system.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Private (Candidate)"
                ],
                "summary": "CreateCandidate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer ${token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateCandidateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/candidate/{id}": {
            "get": {
                "description": "You can use this API to get candidate by id in the system.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Private (Candidate)"
                ],
                "summary": "GetCandidate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer ${token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Candidate ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetCandidateResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "You can use this API to delete candidate in the system (Can use when candidate vote is 0).",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Private (Candidate)"
                ],
                "summary": "DeleteCandidate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer ${token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Candidate ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/candidates": {
            "delete": {
                "description": "You can use this API to clear all candidates and votes in the system.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Private (Candidate)"
                ],
                "summary": "ClearCandidates",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer ${token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/list-candidate": {
            "post": {
                "description": "You can use this API to list candidate in the system.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Private (Candidate)"
                ],
                "summary": "ListCandidate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer ${token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ListCandidateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ListCandidateResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/login": {
            "post": {
                "description": "Users need to use this API to log in and grant access token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public (User)"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RegisterUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/register": {
            "post": {
                "description": "You can use this API to create users in the system.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public (User)"
                ],
                "summary": "RegisterUser",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RegisterUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/report": {
            "get": {
                "description": "You can use this API to get report in the system.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Private (Bonus)"
                ],
                "summary": "ExportReport",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer ${token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ReportResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/vote": {
            "post": {
                "description": "You can use this API to vote for candidates in the system. (User can only have 1 vote)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Private (Vote)"
                ],
                "summary": "Vote",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer ${token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.VoteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/vote-status": {
            "put": {
                "description": "You can use this API to update/edit vote status in the system.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Private (Bonus)"
                ],
                "summary": "UpdateVoteStatus",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer ${token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.VoteStatus"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/votes": {
            "delete": {
                "description": "You can use this API to clear all votes in the system. (After clearing all votes user can vote again)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Private (Vote)"
                ],
                "summary": "ClearVotes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer ${token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GeneralResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CandidateResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "vote_count": {
                    "type": "integer"
                }
            }
        },
        "models.CreateCandidateRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 30,
                    "example": "karen (Max 30 Chars)"
                }
            }
        },
        "models.GeneralResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "$ref": "#/definitions/models.ResponseStatus"
                }
            }
        },
        "models.GetCandidateResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.CandidateResponse"
                },
                "status": {
                    "$ref": "#/definitions/models.ResponseStatus"
                }
            }
        },
        "models.ListCandidateRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "order_by_vote_count": {
                    "type": "boolean"
                },
                "page": {
                    "type": "integer"
                },
                "search_by_name": {
                    "type": "string"
                }
            }
        },
        "models.ListCandidateResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.ListCandidateResponseData"
                },
                "status": {
                    "$ref": "#/definitions/models.ResponseStatus"
                }
            }
        },
        "models.ListCandidateResponseData": {
            "type": "object",
            "properties": {
                "candidate_count": {
                    "type": "integer"
                },
                "candidates": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CandidateResponse"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "page_count": {
                    "type": "integer"
                }
            }
        },
        "models.LoginResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.LoginResponseData"
                },
                "status": {
                    "$ref": "#/definitions/models.ResponseStatus"
                }
            }
        },
        "models.LoginResponseData": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "models.RegisterUserRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "example": "john (Max 50 Chars)"
                }
            }
        },
        "models.ReportCandidateResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "vote_count": {
                    "type": "integer"
                }
            }
        },
        "models.ReportResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.ReportResponseData"
                },
                "status": {
                    "$ref": "#/definitions/models.ResponseStatus"
                }
            }
        },
        "models.ReportResponseData": {
            "type": "object",
            "properties": {
                "candidates": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ReportCandidateResponse"
                    }
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                },
                "vote_status": {
                    "$ref": "#/definitions/models.VoteStatus"
                }
            }
        },
        "models.ResponseStatus": {
            "type": "object",
            "properties": {
                "app_code": {
                    "type": "string"
                },
                "app_message": {
                    "type": "string"
                }
            }
        },
        "models.UpdateCandidateRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "maxLength": 30,
                    "example": "karen (Max 30 Chars)"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                },
                "voted_to": {
                    "type": "string"
                }
            }
        },
        "models.VoteRequest": {
            "type": "object",
            "required": [
                "candidate_id"
            ],
            "properties": {
                "candidate_id": {
                    "type": "integer"
                }
            }
        },
        "models.VoteStatus": {
            "type": "object",
            "properties": {
                "is_active": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
