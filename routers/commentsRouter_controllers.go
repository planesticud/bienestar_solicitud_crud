package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionAtencionMedicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionAtencionMedicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionAtencionMedicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionAtencionMedicaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionAtencionMedicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionAtencionMedicaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionAtencionMedicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionAtencionMedicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionAtencionMedicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionAtencionMedicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:AsignacionSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:DatoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:DatoSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:DatoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:DatoSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:DatoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:DatoSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:DatoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:DatoSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:DatoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:DatoSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:MotivoAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:MotivoAvalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:MotivoAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:MotivoAvalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:MotivoAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:MotivoAvalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:MotivoAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:MotivoAvalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:MotivoAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:MotivoAvalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SolicitudServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SolicitudServicioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SolicitudServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SolicitudServicioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SolicitudServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SolicitudServicioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SolicitudServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SolicitudServicioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SolicitudServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SolicitudServicioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SoporteAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SoporteAvalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SoporteAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SoporteAvalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SoporteAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SoporteAvalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SoporteAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SoporteAvalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SoporteAvalController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SoporteAvalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SubtipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SubtipoServicioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SubtipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SubtipoServicioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SubtipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SubtipoServicioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SubtipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SubtipoServicioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SubtipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:SubtipoServicioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoAtencionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoAtencionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoAtencionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoAtencionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoAtencionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoAtencionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoAtencionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoAtencionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoAtencionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoAtencionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoServicioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoServicioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoServicioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoServicioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoServicioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/bienestar_solicitud_crud/controllers:TipoServicioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
