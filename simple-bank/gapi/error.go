package gapi

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func fieldViolation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

func maybeInvalidArgumentError(violations []*errdetails.BadRequest_FieldViolation) error {
	if violations != nil {
		badRequest := &errdetails.BadRequest{FieldViolations: violations}
		statusInvalid := status.New(codes.InvalidArgument, "invalid parameters")
		statusDetails, err := statusInvalid.WithDetails(badRequest)
		if err != nil {
			// probably log here since this shouldn't really happen
			return statusInvalid.Err() // No details
		}
		return statusDetails.Err()
	}
	return nil
}
