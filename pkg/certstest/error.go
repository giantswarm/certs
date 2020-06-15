package certstest

import "github.com/giantswarm/microerror"

var notFoundError = &microerror.Error{
	Kind: "notFoundError",
}

func IsNotFound(err error) bool {
	return microerror.Cause(err) == notFoundError
}
