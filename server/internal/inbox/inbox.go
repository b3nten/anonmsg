package inbox

import (
	"context"
	"crypto/rand"
	"database/sql"
	"net/http"

	"benton.codes/anonmsg/internal/core"
	"benton.codes/anonmsg/internal/database"
	"benton.codes/anonmsg/internal/messages"
	"github.com/danielgtaylor/huma/v2"
)

type GenericMessageResponse struct {
	Body string `json:"message"`
}

func newGenericMessageResponse(message string) *GenericMessageResponse {
	return &GenericMessageResponse{Body: message}
}

func Register(parent huma.API, core *core.Context) {
	api := huma.NewGroup(parent, "/inbox")

	api.UseSimpleModifier(func(op *huma.Operation) {
		op.Tags = []string{"Inbox"}
	})

	type CreateInboxResponse struct {
		Body struct {
			PrivateKey string `json:"private_key"`
		} `json:"body"`
	}
	huma.Register(api,
		huma.Operation{
			OperationID: "create-inbox",
			Method:      http.MethodPost,
			Path:        "/",
			Summary:     "Create a new inbox",
			Description: "Create a new inbox, returning the private key",
		},
		func(ctx context.Context, in *struct{}) (*CreateInboxResponse, error) {
			// private keys start with 0, public keys start with 1
			privateKey := "0" + rand.Text()
			_, err := core.Queries.CreateInbox(ctx, privateKey)
			if err != nil {
				return nil, huma.Error500InternalServerError("Internal error, could not create inbox")
			}
			return &CreateInboxResponse{
				Body: struct {
					PrivateKey string `json:"private_key"`
				}{
					PrivateKey: privateKey,
				},
			}, nil
		})

	type GetInboxRequest struct {
		PrivateKey string `path:"private_key"`
	}
	type GetInboxResponse struct {
		Body struct {
			PrivateKey   string `json:"private_key"`
			PublicKey    string `json:"public_key"`
			Active       bool   `json:"active"`
			MessageCount int64  `json:"message_count"`
		}
	}
	huma.Register(api,
		huma.Operation{
			OperationID: "get-inbox",
			Method:      http.MethodGet,
			Path:        "/{private_key}",
			Summary:     "Get an inbox",
		},
		func(ctx context.Context, in *GetInboxRequest) (*GetInboxResponse, error) {
			inbox, err := core.Queries.GetInbox(ctx, in.PrivateKey)
			if err != nil {
				return nil, huma.Error500InternalServerError("Could not retreive inbox")
			}
			return &GetInboxResponse{Body: struct {
				PrivateKey   string `json:"private_key"`
				PublicKey    string `json:"public_key"`
				Active       bool   `json:"active"`
				MessageCount int64  `json:"message_count"`
			}{
				PrivateKey:   inbox.PrivateKey,
				PublicKey:    inbox.PublicKey.String,
				Active:       inbox.Active,
				MessageCount: inbox.MessageCount,
			}}, nil
		})

	type DeleteInboxRequest struct {
		PrivateKey string `path:"private_key"`
	}
	huma.Register(api,
		huma.Operation{
			OperationID: "delete-inbox",
			Method:      http.MethodDelete,
			Path:        "/{private_key}",
			Summary:     "Delete an inbox",
		},
		func(ctx context.Context, in *DeleteInboxRequest) (*GenericMessageResponse, error) {
			err := core.Queries.DeleteInbox(ctx, in.PrivateKey)
			if err != nil {
				return nil, huma.Error500InternalServerError("Could not delete inbox")
			}
			return newGenericMessageResponse("Inbox deleted for " + in.PrivateKey), nil
		})

	type SetActiveStatusRequest struct {
		PrivateKey string `path:"private_key"`
		Status     bool   `query:"status"`
	}
	type SetActiveStatusResponse struct {
		Body struct {
			Active bool `json:"active"`
		}
	}
	huma.Register(api,
		huma.Operation{
			OperationID: "set-inbox-active-status",
			Method:      http.MethodPost,
			Path:        "/{private_key}/active",
			Summary:     "Set active status of an inbox",
		},
		func(ctx context.Context, in *SetActiveStatusRequest) (*SetActiveStatusResponse, error) {
			_, err := core.Queries.SetInboxActive(ctx, database.SetInboxActiveParams{
				PrivateKey: in.PrivateKey,
				Active:     in.Status,
			})
			if err != nil {
				return nil, huma.Error500InternalServerError("Could not set active status")
			}
			return &SetActiveStatusResponse{
				Body: struct {
					Active bool `json:"active"`
				}{
					Active: in.Status,
				},
			}, nil
		})

	type GetMessagesRequest struct {
		PrivateKey string `path:"private_key"`
	}
	type GetMessagesResponse struct {
		Body struct {
			Messages []messages.ClientMessage `json:"messages"`
		}
	}
	huma.Register(api,
		huma.Operation{
			OperationID: "get-inbox-messages",
			Method:      http.MethodGet,
			Path:        "/{private_key}/messages",
			Summary:     "Get messages for an inbox",
		},
		func(ctx context.Context, in *GetMessagesRequest) (*GetMessagesResponse, error) {
			messagesFromDb, err := core.Queries.GetMessagesByPrivateKey(ctx, in.PrivateKey)

			if err != nil {
				return nil, huma.Error500InternalServerError("Could not retreive messages")
			}

			// if len == 0, the inbox doesn't exist
			if len(messagesFromDb) == 0 {
				return nil, huma.Error404NotFound("Inbox not found for private key")
			}

			// if the timestamp is zero (null), the inbox exists but is empty
			if messagesFromDb[0].CreatedAt.Time.IsZero() {
				return &GetMessagesResponse{Body: struct {
					Messages []messages.ClientMessage `json:"messages"`
				}{
					Messages: []messages.ClientMessage{},
				}}, nil
			}

			finalMessages := []messages.ClientMessage{}
			for _, message := range messagesFromDb {
				finalMessages = append(finalMessages, messages.ClientMessage{
					ID:        message.ID.Int64,
					Content:   message.MsgContent.String,
					CreatedAt: message.CreatedAt.Time.Unix(),
				})
			}

			return &GetMessagesResponse{Body: struct {
				Messages []messages.ClientMessage `json:"messages"`
			}{
				Messages: finalMessages,
			}}, nil
		})

	type SetPublicKeyRequest struct {
		PrivateKey string `path:"private_key"`
	}
	type SetPublicKeyResponse struct {
		Body struct {
			PublicKey string `json:"public_key"`
		}
	}
	huma.Register(api,
		huma.Operation{
			OperationID: "set-inbox-public-key",
			Method:      http.MethodPost,
			Path:        "/{private_key}/set-public-key",
			Summary:     "Add a public key to an inbox",
		},
		func(ctx context.Context, in *SetPublicKeyRequest) (*SetPublicKeyResponse, error) {
			publicKey := "1" + rand.Text()
			_, err := core.Queries.AddPublicKeyToInbox(ctx, database.AddPublicKeyToInboxParams{
				PrivateKey: in.PrivateKey,
				PublicKey: sql.NullString{
					String: publicKey,
					Valid:  true,
				},
			})
			if err != nil {
				return nil, huma.Error500InternalServerError("Could not set public key")
			}
			return &SetPublicKeyResponse{Body: struct {
				PublicKey string `json:"public_key"`
			}{
				PublicKey: publicKey,
			}}, nil
		})

	type RemovePublicKeyRequest struct {
		PrivateKey string `path:"private_key"`
	}
	huma.Register(api,
		huma.Operation{
			OperationID: "remove-inbox-public-key",
			Method:      http.MethodPost,
			Path:        "/{private_key}/remove-public-key",
			Summary:     "Remove a public key from an inbox",
		},
		func(ctx context.Context, in *RemovePublicKeyRequest) (*GenericMessageResponse, error) {
			err := core.Queries.RemovePublicKeyFromInbox(ctx, in.PrivateKey)
			if err != nil {
				return nil, huma.Error500InternalServerError("Could not remove public key")
			}
			return newGenericMessageResponse("Public key removed successfully"), nil
		})

	type DeleteMessageRequest struct {
		PrivateKey string `path:"private_key"`
		Body       struct {
			MessageID int64 `json:"message_id"`
		}
	}
	huma.Register(api,
		huma.Operation{
			OperationID: "delete-message",
			Method:      http.MethodDelete,
			Path:        "/{private_key}/message",
			Summary:     "delete a message from an inbox",
		},
		func(ctx context.Context, in *DeleteMessageRequest) (*GenericMessageResponse, error) {
			err := core.Queries.DeleteMessage(ctx, database.DeleteMessageParams{
				PrivateKey: in.PrivateKey,
				ID:         in.Body.MessageID,
			})
			if err != nil {
				return nil, huma.Error500InternalServerError("Could not delete message")
			}
			return newGenericMessageResponse("Message deleted successfully"), nil
		})
}
