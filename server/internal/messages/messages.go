package messages

import (
	"context"
	"database/sql"
	"net/http"

	"benton.codes/anonmsg/internal/core"
	"benton.codes/anonmsg/internal/database"
	"github.com/danielgtaylor/huma/v2"
)

type createMessageRequest struct {
	Key     string `path:"key" doc:"inbox private or public key"`
	RawBody []byte `contentType:"text/plain"`
}

type createMessageResponse struct {
	Body struct {
		Message ClientMessage `json:"message"`
	}
}

func Register(parent huma.API, core *core.Context) {
	api := huma.NewGroup(parent, "/send")

	api.UseSimpleModifier(func(op *huma.Operation) {
		op.Tags = []string{"Messages"}
	})

	//**************************************************
	// Send Message
	//**************************************************

	huma.Register(api,
		huma.Operation{
			OperationID: "send-message",
			Method:      http.MethodPost,
			Path:        "/{key}",
			Summary:     "Send a message",
		},
		func(ctx context.Context, in *createMessageRequest) (*createMessageResponse, error) {
			var error error
			var id int64
			var createdAt int64
			var content string
			// branch on public/private key
			// ID's starting with 0 are private, 1 are public
			if string(in.Key[0]) == "0" {
				msg, err := core.Queries.CreatePrivateMessage(ctx, database.CreatePrivateMessageParams{
					PrivateKey: in.Key,
					MsgContent: string(in.RawBody),
				})
				if err != nil {
					error = huma.Error400BadRequest("the private key does not exist")
				} else {
					id = msg.ID
					content = string(in.RawBody)
					createdAt = msg.CreatedAt.Unix()
				}
			} else if string(in.Key[0]) == "1" {
				msg, err := core.Queries.CreatePublicMessage(ctx, database.CreatePublicMessageParams{
					PublicKey: sql.NullString{
						String: in.Key,
						Valid:  true,
					},
					MsgContent: string(in.RawBody),
				})
				if err != nil {
					error = huma.Error400BadRequest("the public key does not exist")
				} else {
					id = msg.ID
					content = string(in.RawBody)
					createdAt = msg.CreatedAt.Unix()
				}
			} else {
				error = huma.Error400BadRequest("Invalid ID format")
			}
			// handle error
			if error != nil {
				return nil, error
			}
			// fin
			return &createMessageResponse{
				Body: struct {
					Message ClientMessage `json:"message"`
				}{
					Message: ClientMessage{
						ID:        id,
						Content:   content,
						CreatedAt: createdAt,
					},
				},
			}, nil
		})
}
