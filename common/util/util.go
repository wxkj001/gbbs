package util

import (
	"bbs/config"
	"crypto/md5"
	"fmt"
)

type Util struct {
	Config config.Config
}

func (u *Util) NewPassword(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s+u.Config.Pass+"bbs")))
}
func (u *Util) PassVerify(p1, p2 string) bool {
	return fmt.Sprintf("%x", md5.Sum([]byte(p1+u.Config.Pass+"bbs"))) == p2
}
