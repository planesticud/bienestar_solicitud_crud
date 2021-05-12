// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/planesticud/bienestar_solicitud_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/persona",
			beego.NSInclude(
				&controllers.PersonaController{},
			),
		),

		beego.NSNamespace("/asignacion_atencion_medica",
			beego.NSInclude(
				&controllers.AsignacionAtencionMedicaController{},
			),
		),

		beego.NSNamespace("/motivo_aval",
			beego.NSInclude(
				&controllers.MotivoAvalController{},
			),
		),

		beego.NSNamespace("/asignacion_solicitud",
			beego.NSInclude(
				&controllers.AsignacionSolicitudController{},
			),
		),

		beego.NSNamespace("/soporte_aval",
			beego.NSInclude(
				&controllers.SoporteAvalController{},
			),
		),

		beego.NSNamespace("/solicitud_servicio",
			beego.NSInclude(
				&controllers.SolicitudServicioController{},
			),
		),

		beego.NSNamespace("/tipo_servicio",
			beego.NSInclude(
				&controllers.TipoServicioController{},
			),
		),

		beego.NSNamespace("/subtipo_servicio",
			beego.NSInclude(
				&controllers.SubtipoServicioController{},
			),
		),

		beego.NSNamespace("/tipo_atencion",
			beego.NSInclude(
				&controllers.TipoAtencionController{},
			),
		),

		beego.NSNamespace("/dato_solicitud",
			beego.NSInclude(
				&controllers.DatoSolicitudController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
