package security

import (
	"strings"

	"github.com/incwadi-warehouse/monorepo-go/search-api/user"
	"github.com/incwadi-warehouse/monorepo-go/search-api/util"
)

func HasRole(token, role string) bool {
	s := strings.Split(token, " ")

	if len(s) != 2 {
		return false
	}

	u, err := user.GetUser(s[1])
	if err != nil {
		return false
	}

	return util.Contains(role, u.Roles)
}
