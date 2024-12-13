package config

import (
	"database/sql"
	"fmt"
	"log"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
)

func NewEnforcer(db *sql.DB) (*casbin.Enforcer, error) {
	adapter, err := sqladapter.NewAdapter(db, "postgres", "rules")
	if err != nil {
		return nil, fmt.Errorf("casbin: failed to initialize adapter - %v", err)
	}
	enforce, err := casbin.NewEnforcer("./casbin.conf", adapter)
	if err != nil {
		return nil, fmt.Errorf("casbin: failed to create enforcer - %v", err)
	}

	// policies := [][]string{
	// 	{"SUPERADMIN", "/proto.TaskService/List"},
	// 	{"SUPERADMIN", "/proto.TaskService/Delete"},
	// 	{"ADMIN", "/proto.TaskService/Delete"},
	// }

	// _, err = infrastructure.SqlxDB.Exec("delete from rules where p_type = 'p' ")
	// if err != nil {
	// 	fmt.Printf("casbin: failed to delete policies - %v \n", err)
	// 	os.Exit(1)
	// }
	// _, err = enforce.AddPoliciesEx(policies)
	// if err != nil {
	// 	fmt.Printf("casbin: failed to add policies - %v \n", err)
	// 	os.Exit(1)
	// }
	enforce.LoadPolicy()

	log.Println("casbin: permission established")
	return enforce, nil
}
