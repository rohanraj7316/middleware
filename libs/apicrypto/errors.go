package apicrypto

import (
	"net/http"

	"github.com/rohanraj7316/middleware/libs/response"
)

var DecryptKeySuccess = func(sType, clientId, secretKey, publicKey string) response.BodyStruct {
	fSucc := map[string]response.BodyStruct{
		"Success": {
			StatusCode: http.StatusOK,
			Data: map[string]string{
				"clientId":  clientId,
				"secretKey": secretKey,
			},
		},
		"SuccessfullyCreatedPublicKey": {
			StatusCode: http.StatusBadRequest,
			Message:    "Private Key Not Found",
			Data: map[string]string{
				"publicKey": publicKey,
			},
		},
	}

	return fSucc[sType]
}

var DecryptKeyErr = func(errType string, err error) response.BodyStruct {
	errs := map[string]response.BodyStruct{
		"InvalidClient": {
			StatusCode: http.StatusUnauthorized,
			Err:        "unable to find `clientId`",
		},
		"FailedToFetchClientCache": {
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		},
		"FailedToGenerateKeyPairs": {
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		},
		"FailedToUpdatePublicKeyCache": {
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		},
		"FailedToUpdatePrivateKeyCache": {
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		},
		"FailedToDecryptKey": {
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		},
	}

	fErr := errs[errType]
	if err != nil {
		fErr.Err = err
	}

	return fErr
}

var DecryptPayloadSuccess = func(payload string) response.BodyStruct {
	fSucc := response.BodyStruct{
		StatusCode: http.StatusOK,
		Data:       payload,
	}

	return fSucc
}

var DecryptPayloadErr = func(errType string, err error) response.BodyStruct {
	errs := map[string]response.BodyStruct{
		"FailedToDecryptPayload": {
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		},
	}

	fErr := errs[errType]
	if err != nil {
		fErr.Err = err
	}

	return fErr
}

var EncryptPayloadSuccess = func(encryptedPayload string) response.BodyStruct {
	fSucc := response.BodyStruct{
		StatusCode: http.StatusOK,
		Data: map[string]string{
			"payload": encryptedPayload,
		},
	}

	return fSucc
}

var EncryptPayloadErr = func(errType string, err error) response.BodyStruct {
	errs := map[string]response.BodyStruct{
		"FailedToEncryptPayload": {
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		},
	}

	fErr := errs[errType]
	if err != nil {
		fErr.Err = err
	}

	return fErr
}

var HandshakeSuccess = func(publicKey string) response.BodyStruct {
	fSucc := response.BodyStruct{
		StatusCode: http.StatusOK,
		Message:    "Handshake Successful",
		Data: map[string]string{
			"publicKey": publicKey,
		},
	}

	return fSucc
}

var HandshakeErr = func(errType string, err error) response.BodyStruct {
	errs := map[string]response.BodyStruct{
		"InvalidClient": {
			StatusCode: http.StatusUnauthorized,
			Err:        "unable to find `clientId`",
		},
		"FailedToFetchClientCache": {
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		},
		"FailedToGenerateKeyPairs": {
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		},
		"FailedToUpdatePublicKeyCache": {
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		},
		"FailedToUpdatePrivateKeyCache": {
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		},
	}

	fErr := errs[errType]
	if err != nil {
		fErr.Err = err
	}

	return fErr
}
