package client

import "github.com/DrusGalkin/libs/proto/generate"

func (a *Auth) IsAdmin(uid int) bool {
	ctx, cancel := a.getContext()
	defer cancel()

	var req = &generate.IsAdminRequest{
		Id: int64(uid),
	}

	res, err := a.client.IsAdmin(ctx, req)
	if err != nil {
		return false
	}
	return res.IsAdmin
}
