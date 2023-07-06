package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

var commonInitialisms = []string{"ACL", "API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SQL", "SSH", "TCP", "TLS", "TTL", "UDP", "UI", "UID", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XMPP", "XSRF", "XSS"}
var commonInitialismsReplacer *strings.Replacer
var uncommonInitialismsReplacer *strings.Replacer

func generateStruct() {

	input := `view_company_project
#公司ID,公司中文名,公司英文名,公司删除状态,公司项目关系ID
view_company_id,view_company_cn,view_company_en,view_company_del_flag,view_pr_id,
#项目ID,项目中文名,项目英文名,部门ID,部门名,项目类型
view_project_id,view_project_cn,view_project_en,view_dept_id,view_dept_name,view_project_type,
#项目群号,消息群类型中文名,消息群类型英文名,开发人员QQ,项目删除状态
view_group_qq,view_group_type_cn,view_group_type_en,view_group_dev_qq,view_project_del_flag

view_assets
#资产ID,公网IP,内网IP,服务器用途ID,云商ID,云商英文名,云商中文名
view_asset_id,view_outer_ip,view_inner_ip,view_host_role_id,view_provider_id,view_provider_name_en,view_provider_name_cn,
#服务器硬件信息,SSH端口,初始化状态,清理状态,初始登录信息,状态变更备注信息,资产备注信息
view_hardware_info,view_ssh_port,view_init_type,view_clean_type,view_init_login_info,view_change_status_remark,view_remark,
#资产创建时间,资产更新时间,资产删除状态
view_asset_create_time,view_asset_update_time,view_asset_del_flag,
#公司项目关系ID,出机方公司ID,出机方公司中文名,出机方工资英文名
view_pr_id,view_asset_ownership_company_id,view_asset_ownership_company_cn,view_asset_ownership_company_en,
#出机方公司删除状态,公司项目关系删除状态,使用方公司ID,使用方公司中文名,使用方公司英文名
view_asset_ownership_company_deleted,view_server_affiliation_deleted,view_user_company_id,view_user_company_cn,view_user_company_en,
#使用方公司删除状态,使用方项目ID,使用方项目中文名,使用方项目英文名,使用方项目删除状态
view_user_company_deleted,view_user_project_id,view_user_project_cn,view_user_project_en,view_user_project_deleted
`
	reader := bufio.NewReader(bytes.NewBuffer([]byte(input)))
	structList := make([]string, 0)
	golangStructList := make([]string, 0)
	protoList := make([]string, 0)
	num := 1
	for {
		line, err := reader.ReadString('\n') //注意是字符

		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println("end:", line)
			}
			//fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}

		compile := regexp.MustCompile(`#|^$`)
		allString := compile.FindAllString(strings.TrimSpace(line), -1)
		if len(allString) == 0 {
			a := strings.Split(strings.TrimSpace(line), ",")

			if len(a) == 1 {
				c := fmt.Sprintf(`}
type ` + Marshal(strings.TrimSpace(line)) + ` struct {`)

				m := fmt.Sprintf(`}
message %s {
`, Marshal(strings.TrimSpace(line)))
				e := fmt.Sprintf(`}
%s struct {`, Marshal(strings.TrimSpace(line)))
				golangStructList = append(golangStructList, c)
				protoList = append(protoList, m)
				structList = append(structList, e)
				num = 1
			} else {
				for _, v := range a {
					if len(strings.TrimSpace(v)) != 0 {
						types := "sql.NullString"
						typesn := "string"
						mustCompile := regexp.MustCompile(`id|deleted|del_flag`)
						if mustCompile.MatchString(v) {
							types = "sql.NullInt64"
							typesn = "int64"
						}
						c := fmt.Sprintf(`%s %s %s`, Marshal(strings.TrimSpace(v)), types, "`db:\""+v+"\"`")
						golangStructList = append(golangStructList, c)
						m := fmt.Sprintf("%s %s = %d;", typesn, v, num)
						protoList = append(protoList, m)
						mm := strings.ToLower(Marshal(strings.TrimSpace(v))[:1]) + Marshal(strings.TrimSpace(v))[1:]
						e := fmt.Sprintf("%s %s `json:\"%s\"`", Marshal(strings.TrimSpace(v)), typesn, mm)
						structList = append(structList, e)
						num++
					}

				}
			}

		}

	}
	fmt.Println(strings.Join(golangStructList, "\n"))
	fmt.Println(strings.Join(protoList, "\n"))
	fmt.Println(strings.Join(structList, "\n"))
	fmt.Println("}")

}

func init() {
	var commonInitialismsForReplacer []string
	var uncommonInitialismsForReplacer []string
	for _, initialism := range commonInitialisms {
		commonInitialismsForReplacer = append(commonInitialismsForReplacer, initialism, strings.Title(strings.ToLower(initialism)))
		uncommonInitialismsForReplacer = append(uncommonInitialismsForReplacer, strings.Title(strings.ToLower(initialism)), initialism)
	}
	commonInitialismsReplacer = strings.NewReplacer(commonInitialismsForReplacer...)
	uncommonInitialismsReplacer = strings.NewReplacer(uncommonInitialismsForReplacer...)

}

func Marshal(name string) string {
	if name == "" {
		return ""
	}
	temp := strings.Split(name, "_")
	var s string
	for _, v := range temp {
		vv := []rune(v)
		if len(vv) > 0 {
			if bool(vv[0] >= 'a' && vv[0] <= 'z') { //首字母大写
				vv[0] -= 32
			}

			s += string(vv)
		}
	}

	//s = uncommonInitialismsReplacer.Replace(s)
	//smap.Set(name, s)
	return s
}
