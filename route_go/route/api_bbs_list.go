package route

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"opennamu/route/tool"
)

func bbs_list(db *sql.DB, db_set map[string]string) map[string]string {
	rows, err := db.Query(tool.DB_change(db_set, "select set_data, set_id from bbs_set where set_name = 'bbs_name'"))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	data_list := map[string]string{}

	for rows.Next() {
		var name string
		var id string

		err := rows.Scan(&name, &id)
		if err != nil {
			log.Fatal(err)
		}

		data_list[name] = id
	}

	return data_list
}

func Api_bbs_list(call_arg []string) {
	db_set := map[string]string{}
	json.Unmarshal([]byte(call_arg[0]), &db_set)

	db := tool.DB_connect(db_set)
	if db == nil {
		return
	}
	defer db.Close()

	data_list := bbs_list(db, db_set)

	if len(data_list) == 0 {
		fmt.Print("{}")
	} else {
		json_data, _ := json.Marshal(data_list)
		fmt.Print(string(json_data))
	}
}
