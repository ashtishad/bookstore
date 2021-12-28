package handlers

import (
	"github.com/ashtishad/bookstore/lib"
	"strconv"
)

func getUserId(userIdParam string) (int64, lib.RestErr) {
	id, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		return 0, lib.NewBadRequestError("user id should be a number")
	}
	return id, nil
}
