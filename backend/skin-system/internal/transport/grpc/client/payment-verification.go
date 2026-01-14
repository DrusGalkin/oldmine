package client

import "libs/proto/generate"

func (a *Auth) PaymentVerification(uid int) bool {
	ctx, cancel := a.getContext()
	defer cancel()

	var req = &generate.PaymentVerificationRequest{Id: int64(uid)}

	ver, err := a.client.PaymentVerification(ctx, req)
	if err != nil {
		return false
	}

	return ver.Pay
}
