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
    if page*50 > 0 {
        num = page*50 - 50
    }

    var stmt *sql.Stmt
    var err error
    if other_set["search_type"] == "title" {
        stmt, err = db.Prepare(tool.DB_change("SELECT doc_name, set_data FROM data_set WHERE doc_name collate nocase LIKE ? AND set_name = 'view_count' AND doc_rev = '' ORDER BY doc_name LIMIT ?, 50"))
        if err != nil {
            panic(err)
        }
    } else {
        stmt, err = db.Prepare(tool.DB_change("SELECT d.title, set_data from data d join data_set ds on d.title = ds.doc_name where d.data collate nocase Like ? and ds.set_name = 'view_count' and ds.doc_rev = ''  LIMIT ?, 50"))
        if err != nil {
            panic(err)
        }
    }
    defer stmt.Close()

    // 두 개의 값을 저장할 구조체 배열 선언
    type Result struct {
        Title   string `json:"title"`
        SetData string `json:"set_data"`
    }
    results := []Result{}

    rows, err := stmt.Query("%"+other_set["name"]+"%", num)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
        var title string
        var setData string

        err := rows.Scan(&title, &setData) // 두 개의 값을 읽음
        if err != nil {
            panic(err)
        }

        // 읽은 데이터를 구조체로 추가
        results = append(results, Result{Title: title, SetData: setData})
    }

    // JSON으로 변환
    json_data, _ := json.Marshal(results)
    return string(json_data)
}
