package client

import (
	"fmt"
	"github.com/DrusGalkin/libs/proto/generate"
	"skins/internal/domain/dto"
)

func (a *Auth) CheckAuth(sess_id string) *dto.ResponseUser {
	ctx, cancel := a.getContext()
	defer cancel()

	auth, err := a.client.CheckAuth(ctx, &generate.AuthRequest{SessId: sess_id})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &dto.ResponseUser{
		ID:    int(auth.Id),
		Name:  auth.Name,
		Email: auth.Email,
		Auth:  auth.Auth,
		Pay:   false,
		Admin: false,
	}
}
