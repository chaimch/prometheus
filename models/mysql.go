package models

import (
	"container/list"
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	_SQL_DB *sql.DB
	_ERR    error
)

type Group struct {
	// 组 ID
	GroupID int
	// 组名称
	Name string
	// 组昵称
	NickName string
}

type RuleItem struct {
	// 规则所属组名称
	Name string
	// 类别
	Fn string
	// 规则计算间隔
	Interval int
	// 告警名称
	Alert string
	// 记录名称
	Record string
	// 规则表达式
	Expr string
	// 持续时间
	For string
	// 规则维度信息
	Labels map[string]string
	// 规则描述信息
	Annotations map[string]string
}

func Initialization(db_url string) {
	_SQL_DB, _ERR = sql.Open("mysql", db_url)
	if _ERR != nil {
		log.Fatalf("Open database error: %s\n", _ERR)
		return
	}
	_ERR = _SQL_DB.Ping()
	if _ERR != nil {
		log.Fatal(_ERR)
		return
	}
}

func QueryGroup(datasource string) (map[int]Group, error) {
	groups := make(map[int]Group)
	groupid := 1
	groups[groupid] = Group{
		GroupID:  groupid,
		Name:     "sre",
		NickName: "运维组",
	}
	return groups, nil
}

func QueryRuleString(datasource string) (*list.List, error) {
	var (
		rule_labels, rule_annotations string
	)

	l := list.New()
	// 查询规则表
	rows, err := _SQL_DB.Query("select scene, `name`, `interval`, expr, `for`, labels, annotations, group_id from rule where datasource = ?;", datasource)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			scene   string
			item    RuleItem
			groupid int
		)
		item.Labels = make(map[string]string)
		item.Annotations = make(map[string]string)
		err := rows.Scan(&scene, &item.Alert, &item.Interval, &item.Expr, &item.For, &rule_labels, &rule_annotations, &groupid)
		if err != nil {
			log.Fatal(err)
		}

		if scene == "record" {
			item.Record = item.Alert
			item.Alert = ""
		}

		groups, err := QueryGroup(datasource)
		if err != nil {
			log.Fatal(err)
		}

		groupName := groups[groupid].Name
		if groupName == "" {
			groupName = "sre"
		}
		item.Name = groupName

		item.Fn = datasource

		//  label 数据格式转换
		if rule_labels != "" {
			labels := strings.Split(rule_labels, ",")
			labeln := len(labels)
			for i := 0; i < labeln; i++ {
				pars := strings.Split(labels[i], "=")
				plen := len(pars)
				for j := 0; j < plen; j += 2 {
					item.Labels[pars[j]] = pars[j+1]
				}
			}
		}
		item.Labels["business_group"] = groupName

		// annotations 数据格式转换
		if rule_annotations != "" {
			annotations := strings.Split(rule_annotations, ",")
			annlen := len(annotations)
			for k := 0; k < annlen; k++ {
				pars := strings.Split(annotations[k], "=")
				plen := len(pars)
				for j := 0; j < plen; j += 2 {
					item.Annotations[pars[j]] = pars[j+1]
				}
			}
		}

		l.PushBack(item)
	}
	return l, err
}
