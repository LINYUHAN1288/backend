package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)


func init() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	orm.RegisterModel(new(Profile))
}

type Profile struct {
	Truename   string    `orm:"column(truename);size(32)" description:"鐪熷疄濮撳悕"`
	Age        uint8     `orm:"column(age)" description:"骞撮緞"`
	Sex        string    `orm:"column(sex)" description:"鎬у埆"`
	Birthday   time.Time `orm:"column(birthday);type(date)" description:"鐢熸棩"`
	Nickname   string    `orm:"column(nickname);size(32)" description:"鏄电О"`
	Id         int       `orm:"column(userid);pk" description:"鐢ㄦ埛鐨処D"`
	Createtime uint      `orm:"column(createtime)" description:"鍒涘缓鏃堕棿"`
}

func (t *Profile) TableName() string {
	return "profile"
}



// AddProfile insert a new Profile into database and returns
// last inserted Id on success.
func AddProfile(m *Profile) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProfileById retrieves Profile by Id. Returns error if
// Id doesn't exist
func GetProfileById(id int) (v *Profile, err error) {
	o := orm.NewOrm()
	v = &Profile{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProfile retrieves all Profile matches certain condition. Returns empty list if
// no records exist
func GetAllProfile(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Profile))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Profile
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateProfile updates Profile by Id and returns error if
// the record to be updated doesn't exist
func UpdateProfileById(m *Profile) (err error) {
	o := orm.NewOrm()
	v := Profile{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProfile deletes Profile by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProfile(id int) (err error) {
	o := orm.NewOrm()
	v := Profile{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Profile{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
