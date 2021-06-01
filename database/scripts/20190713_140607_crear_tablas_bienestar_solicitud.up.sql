
--
-- TOC entry 198 (class 1259 OID 53193)
-- Name: asignacion_atencion_medica_id_seq; Type: SEQUENCE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE SEQUENCE bienestar_solicitud.asignacion_atencion_medica_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


-- ALTER TABLE bienestar_solicitud.asignacion_atencion_medica_id_seq OWNER TO postgres;

--
-- TOC entry 199 (class 1259 OID 53195)
-- Name: asignacion_atencion_medica; Type: TABLE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE TABLE bienestar_solicitud.asignacion_atencion_medica (
    id integer DEFAULT nextval('bienestar_solicitud.asignacion_atencion_medica_id_seq'::regclass) NOT NULL,
    hoja_historia_id integer NOT NULL,
    asignacion_solicitud_id integer NOT NULL,
    fecha_creacion timestamp(6) with time zone DEFAULT now() NOT NULL,
    fecha_modificacion timestamp(6) with time zone DEFAULT now() NOT NULL
);


-- ALTER TABLE bienestar_solicitud.asignacion_atencion_medica OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 53201)
-- Name: asignacion_solicitud_id_seq; Type: SEQUENCE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE SEQUENCE bienestar_solicitud.asignacion_solicitud_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


-- ALTER TABLE bienestar_solicitud.asignacion_solicitud_id_seq OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 53203)
-- Name: asignacion_solicitud; Type: TABLE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE TABLE bienestar_solicitud.asignacion_solicitud (
    id integer DEFAULT nextval('bienestar_solicitud.asignacion_solicitud_id_seq'::regclass) NOT NULL,
    solicitud_servicio_id integer NOT NULL,
    fecha_asignacion timestamp with time zone NOT NULL,
    asesor_id integer NOT NULL,
    periodo_id integer NOT NULL,
    fecha_inicio_atencion timestamp with time zone NOT NULL,
    fecha_fin_atencion timestamp with time zone NOT NULL,
    observaciones text NOT NULL,
    fecha_creacion timestamp(6) with time zone DEFAULT now() NOT NULL,
    fecha_modificacion timestamp(6) with time zone DEFAULT now() NOT NULL
);


-- ALTER TABLE bienestar_solicitud.asignacion_solicitud OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 53212)
-- Name: dato_solicitud_id_seq; Type: SEQUENCE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE SEQUENCE bienestar_solicitud.dato_solicitud_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


-- ALTER TABLE bienestar_solicitud.dato_solicitud_id_seq OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 53214)
-- Name: dato_solicitud; Type: TABLE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE TABLE bienestar_solicitud.dato_solicitud (
    id integer DEFAULT nextval('bienestar_solicitud.dato_solicitud_id_seq'::regclass) NOT NULL,
    solicitud_servicio_id integer NOT NULL,
    tipo_atencion_id integer NOT NULL,
    observaciones text NOT NULL,
    asesor_preferencia_id integer NOT NULL,
    fecha_creacion timestamp(6) with time zone DEFAULT now() NOT NULL,
    fecha_modificacion timestamp(6) with time zone DEFAULT now() NOT NULL
);


-- ALTER TABLE bienestar_solicitud.dato_solicitud OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 53223)
-- Name: motivo_aval_id_seq; Type: SEQUENCE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE SEQUENCE bienestar_solicitud.motivo_aval_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


-- ALTER TABLE bienestar_solicitud.motivo_aval_id_seq OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 53225)
-- Name: motivo_aval; Type: TABLE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE TABLE bienestar_solicitud.motivo_aval (
    id integer DEFAULT nextval('bienestar_solicitud.motivo_aval_id_seq'::regclass) NOT NULL,
    nombre character varying(50) NOT NULL,
    descripcion character varying(250),
    codigo_abreviacion character varying(20),
    numero_orden numeric(5,2),
    activo boolean NOT NULL,
    fecha_creacion timestamp with time zone DEFAULT now() NOT NULL,
    fecha_modificacion timestamp with time zone DEFAULT now() NOT NULL
);


-- ALTER TABLE bienestar_solicitud.motivo_aval OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 53231)
-- Name: persona_id_seq; Type: SEQUENCE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE SEQUENCE bienestar_solicitud.persona_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


-- ALTER TABLE bienestar_solicitud.persona_id_seq OWNER TO postgres;

--
-- TOC entry 207 (class 1259 OID 53233)
-- Name: persona; Type: TABLE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE TABLE bienestar_solicitud.persona (
    id integer DEFAULT nextval('bienestar_solicitud.persona_id_seq'::regclass) NOT NULL,
    primer_nombre character varying(50) NOT NULL,
    segundo_nombre character varying(50),
    primer_apellido character varying(50) NOT NULL,
    segundo_apellido character varying(50),
    tipo_identificacion character varying(20),
    numero_identificacion character varying(50) NOT NULL,
    fecha_nacimiento date,
    genero character varying(20),
    telefono character varying(200) NOT NULL,
    correo character varying(200) NOT NULL,
    eps character varying(200) NOT NULL,
    facultad_id integer NOT NULL,
    dependencia_id integer NOT NULL,
    estamento character varying(200) NOT NULL,
    codigo_estudiantil character varying(200),
    estado_condor character varying(200)
);


-- ALTER TABLE bienestar_solicitud.persona OWNER TO postgres;

--
-- TOC entry 208 (class 1259 OID 53240)
-- Name: solicitud_servicio_id_seq; Type: SEQUENCE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE SEQUENCE bienestar_solicitud.solicitud_servicio_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


-- ALTER TABLE bienestar_solicitud.solicitud_servicio_id_seq OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 53242)
-- Name: solicitud_servicio; Type: TABLE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE TABLE bienestar_solicitud.solicitud_servicio (
    id integer DEFAULT nextval('bienestar_solicitud.solicitud_servicio_id_seq'::regclass) NOT NULL,
    persona_id integer NOT NULL,
    subtipo_servicio_id integer,
    fecha_solicitud timestamp with time zone NOT NULL,
    acepta_terminos boolean NOT NULL,
    datos_persona_solicitud json NOT NULL,
    fecha_creacion timestamp(6) with time zone DEFAULT now() NOT NULL,
    fecha_modificacion timestamp(6) with time zone DEFAULT now() NOT NULL
);


-- ALTER TABLE bienestar_solicitud.solicitud_servicio OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 53251)
-- Name: soporte_aval_id_seq; Type: SEQUENCE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE SEQUENCE bienestar_solicitud.soporte_aval_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


-- ALTER TABLE bienestar_solicitud.soporte_aval_id_seq OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 53253)
-- Name: soporte_aval; Type: TABLE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE TABLE bienestar_solicitud.soporte_aval (
    id integer DEFAULT nextval('bienestar_solicitud.soporte_aval_id_seq'::regclass) NOT NULL,
    solicitud_servicio_id integer NOT NULL,
    documento_id integer NOT NULL,
    motivo_aval_id integer NOT NULL,
    fecha_creacion timestamp(6) with time zone DEFAULT now() NOT NULL,
    fecha_modificacion timestamp(6) with time zone DEFAULT now() NOT NULL
);


-- ALTER TABLE bienestar_solicitud.soporte_aval OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 53259)
-- Name: subtipo_servicio_id_seq; Type: SEQUENCE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE SEQUENCE bienestar_solicitud.subtipo_servicio_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


-- ALTER TABLE bienestar_solicitud.subtipo_servicio_id_seq OWNER TO postgres;

--
-- TOC entry 213 (class 1259 OID 53261)
-- Name: subtipo_servicio; Type: TABLE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE TABLE bienestar_solicitud.subtipo_servicio (
    id integer DEFAULT nextval('bienestar_solicitud.subtipo_servicio_id_seq'::regclass) NOT NULL,
    nombre character varying(50) NOT NULL,
    descripcion character varying(250),
    codigo_abreviacion character varying(20),
    numero_orden numeric(5,2),
    activo boolean NOT NULL,
    fecha_creacion timestamp with time zone DEFAULT now() NOT NULL,
    fecha_modificacion timestamp with time zone DEFAULT now() NOT NULL,
    tipo_servicio_id integer
);


-- ALTER TABLE bienestar_solicitud.subtipo_servicio OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 53267)
-- Name: tipo_atencion_id_seq; Type: SEQUENCE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE SEQUENCE bienestar_solicitud.tipo_atencion_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


-- ALTER TABLE bienestar_solicitud.tipo_atencion_id_seq OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 53269)
-- Name: tipo_atencion; Type: TABLE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE TABLE bienestar_solicitud.tipo_atencion (
    id integer DEFAULT nextval('bienestar_solicitud.tipo_atencion_id_seq'::regclass) NOT NULL,
    nombre character varying(50) NOT NULL,
    descripcion character varying(250),
    codigo_abreviacion character varying(20),
    numero_orden numeric(5,2),
    activo boolean NOT NULL,
    fecha_creacion timestamp with time zone DEFAULT now() NOT NULL,
    fecha_modificacion timestamp with time zone DEFAULT now() NOT NULL
);


-- ALTER TABLE bienestar_solicitud.tipo_atencion OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 53275)
-- Name: tipo_servicio_id_seq; Type: SEQUENCE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE SEQUENCE bienestar_solicitud.tipo_servicio_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


-- ALTER TABLE bienestar_solicitud.tipo_servicio_id_seq OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 53277)
-- Name: tipo_servicio; Type: TABLE; Schema: bienestar_solicitud; Owner: postgres
--

CREATE TABLE bienestar_solicitud.tipo_servicio (
    id integer DEFAULT nextval('bienestar_solicitud.tipo_servicio_id_seq'::regclass) NOT NULL,
    nombre character varying(50) NOT NULL,
    descripcion character varying(250),
    codigo_abreviacion character varying(20),
    numero_orden numeric(5,2),
    activo boolean NOT NULL,
    fecha_creacion timestamp with time zone DEFAULT now() NOT NULL,
    fecha_modificacion timestamp with time zone DEFAULT now() NOT NULL
);


-- ALTER TABLE bienestar_solicitud.tipo_servicio OWNER TO postgres;

--
-- TOC entry 3148 (class 2604 OID 53594)
-- Name: tipo_antecedente id; Type: DEFAULT; Schema: bienestar_medico; Owner: postgres
--

-- ALTER TABLE ONLY bienestar_medico.tipo_antecedente ALTER COLUMN id SET DEFAULT nextval('bienestar_medico.tipo_antecedente_id_seq'::regclass);
