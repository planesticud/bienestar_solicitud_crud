package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type AsignacionSolicitud struct {
	Id                  int                `orm:"column(id);pk;auto"`
	SolicitudServicioId *SolicitudServicio `orm:"column(solicitud_servicio_id);rel(fk)"`
	FechaAsignacion     time.Time          `orm:"column(fecha_asignacion);type(timestamp with time zone)"`
	AsesorId            int                `orm:"column(asesor_id)"`
	PeriodoId           int                `orm:"column(periodo_id)"`
	FechaInicioAtencion time.Time          `orm:"column(fecha_inicio_atencion);type(timestamp with time zone)"`
	FechaFinAtencion    time.Time          `orm:"column(fecha_fin_atencion);type(timestamp with time zone)"`
	Observaciones       string             `orm:"column(observaciones)"`
	FechaCreacion       string             `orm:"column(fecha_creacion);null"`
	FechaModificacion   string             `orm:"column(fecha_modificacion);null"`
}

func (t *AsignacionSolicitud) TableName() string {
	return "asignacion_solicitud"
}

func init() {
	orm.RegisterModel(new(AsignacionSolicitud))
}

// AddAsignacionSolicitud insert a new AsignacionSolicitud into database and returns
// last inserted Id on success.
func AddAsignacionSolicitud(m *AsignacionSolicitud) (id int64, err error) {
	m.FechaCreacion = time_bogota.TiempoBogotaFormato()
	m.FechaModificacion = time_bogota.TiempoBogotaFormato()
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAsignacionSolicitudById retrieves AsignacionSolicitud by Id. Returns error if
// Id doesn't exist
func GetAsignacionSolicitudById(id int) (v *AsignacionSolicitud, err error) {
	o := orm.NewOrm()
	v = &AsignacionSolicitud{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAsignacionSolicitud retrieves all AsignacionSolicitud matches certain condition. Returns empty list if
// no records exist
func GetAllAsignacionSolicitud(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AsignacionSolicitud))
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

	var l []AsignacionSolicitud
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

// UpdateAsignacionSolicitud updates AsignacionSolicitud by Id and returns error if
// the record to be updated doesn't exist
func UpdateAsignacionSolicitudById(m *AsignacionSolicitud) (err error) {
	o := orm.NewOrm()
	v := AsignacionSolicitud{Id: m.Id}
	m.FechaModificacion = time_bogota.TiempoBogotaFormato()
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "SolicitudServicioId", "FechaAsignacion", "AsesorId", "PeriodoId", "FechaInicioAtencion", "FechaFinAtencion", "Observaciones", "FechaModificacion"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAsignacionSolicitud deletes AsignacionSolicitud by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAsignacionSolicitud(id int) (err error) {
	o := orm.NewOrm()
	v := AsignacionSolicitud{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AsignacionSolicitud{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
