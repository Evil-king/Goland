package pdb

import "github.com/gogf/gf/database/gdb"

func GetDBInstance(name ...string) (gdb.DB, error) {
	if len(name) == 0 {
		return gdb.Instance(gdb.DefaultGroupName)
	}
	return gdb.Instance(name...)
}
