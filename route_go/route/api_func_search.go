package route

import (
    "database/sql"
    "strconv"

    "opennamu/route/tool"

    jsoniter "github.com/json-iterator/go"
)

func Api_func_search(db *sql.DB, call_arg []string) string {
    var json = jsoniter.ConfigCompatibleWithStandardLibrary

    other_set := map[string]string{}
    json.Unmarshal([]byte(call_arg[0]), &other_set)

    page, _ := strconv.Atoi(other_set["num"])
    num := 0
    if page * 50 > 0 {
        num = page * 50 - 50
    }

    var stmt *sql.Stmt
    var err error
    if other_set["search_type"] == "title" {
        stmt, err = db.Prepare(tool.DB_change("select d.title from data d join data_set ds on d.title = ds.doc_name where d.title collate nocase like ? and ds.doc_rev='' and ds.set_name = 'view_count' order by CAST(ds.set_data AS UNSIGNED) DESC limit ?, 50"))
        if err != nil {
            panic(err)
        }
    } else {
        stmt, err = db.Prepare(tool.DB_change("select d.title from data d join data_set ds on d.title = ds.doc_name where d.data collate nocase like ? and ds.doc_rev='' and ds.set_name = 'view_count' order by CAST(ds.set_data AS UNSIGNED) DESC limit ?, 50"))
        if err != nil {
            panic(err)
        }
    }
    defer stmt.Close()

    title_list := []string{}

    rows, err := stmt.Query("%"+other_set["name"]+"%", num)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
        var title string

        err := rows.Scan(&title)
        if err != nil {
            panic(err)
        }

        title_list = append(title_list, title)
    }

    json_data, _ := json.Marshal(title_list)
    return string(json_data)
}
