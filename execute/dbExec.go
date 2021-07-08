package execute

import (
	"database/sql"
	"generate/fetcher"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type dbTem struct {
	FileName     string
	TableName    string
	TableComent  string
	ArrFieldsMap string
	ArrTypesMap  string
	CreateTime   string
}

type Table struct {
	FieldName string
	TableName string
	Field     string
	Type      string
	Comment   string
}

func GetTemplate(conf fetcher.DBConfig) []dbTem {
	TypeMap := map[string]string{
		"int":     " Hk_Service_Db::TYPE_INT,\n",
		"tinyint": " Hk_Service_Db::TYPE_INT,\n",
		"bigint":  " Hk_Service_Db::TYPE_INT,\n",
		"varchar": " Hk_Service_Db::TYPE_STR,\n",
		"text":    " Hk_Service_Db::TYPE_STR,\n",
	}

	connect := conf.Uname + ":" + conf.Passwd + "@tcp(" + conf.Url + ":" + conf.Port + ")/" + conf.DbName + "?charset=" + conf.CharSet + "&parseTime=True&loc=Local"
	db, err := sql.Open(conf.DbType, connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var tems []dbTem
	for _, tableName := range conf.Table {
		var tem dbTem
		tem.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		tem.FileName = tableName[3:]
		tem.TableName = tableName

		fieldsMap := ""
		typesMap := ""
		execSql := "SELECT COLUMN_NAME, DATA_TYPE,COLUMN_COMMENT FROM information_schema.columns WHERE table_schema='" + conf.DbName + "' and table_name='" + tableName + "';"
		rows, err := db.Query(execSql)
		if err != nil {
			panic(err)
		}

		execSql = "select table_name,table_comment from information_schema.tables where table_schema = '" + conf.DbName + "' and table_name ='" + tableName + "';"
		tableInfos, err := db.Query(execSql)

		var name string
		tableInfos.Next()
		err = tableInfos.Scan(&name, &tem.TableComent)

		var tables []Table
		cnt := 0

		for rows.Next() {
			var table Table
			err = rows.Scan(&table.Field, &table.Type, &table.Comment)
			table.FieldName = Marshal(table.Field)
			tables = append(tables, table)
			if cnt < len(table.FieldName) {
				cnt = len(table.FieldName)
			}
		}
		for _, table := range tables {
			space := ""
			for i := 0; i < cnt-len(table.FieldName)+1; i++ {
				space += " "
			}
			fieldsMap += "            '" + table.FieldName + "'" + space + "=> '" + table.Field + "',  //" + table.Comment + "\n"
			typesMap += "            '" + table.FieldName + "'" + space + "=>" + TypeMap[table.Type]
		}
		tem.ArrFieldsMap = fieldsMap
		tem.ArrTypesMap = typesMap
		tems = append(tems, tem)
	}
	return tems
}

func Marshal(name string) string {
	if name == "" {
		return ""
	}

	temp := strings.Split(name, "_")
	var s string
	for k, v := range temp {
		vv := []rune(v)
		if len(vv) > 0 {
			if bool(vv[0] >= 'a' && vv[0] <= 'z' && k > 0) { //首字母大写
				vv[0] -= 32
			}
			s += string(vv)
		}
	}
	return s
}
