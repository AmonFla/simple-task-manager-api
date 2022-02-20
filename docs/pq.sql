--
-- PostgreSQL database dump
--

-- Dumped from database version 12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)

-- Started on 2022-02-20 19:41:22 -03

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 215 (class 1259 OID 17389)
-- Name: notes; Type: TABLE; Schema: public; Owner: gotest
--

CREATE TABLE public.notes (
    id bigint NOT NULL,
    comment text,
    created_at timestamp without time zone,
    user_id bigint,
    task_id bigint
);


ALTER TABLE public.notes OWNER TO gotest;

--
-- TOC entry 214 (class 1259 OID 17387)
-- Name: notes_id_seq; Type: SEQUENCE; Schema: public; Owner: gotest
--

CREATE SEQUENCE public.notes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.notes_id_seq OWNER TO gotest;

--
-- TOC entry 3091 (class 0 OID 0)
-- Dependencies: 214
-- Name: notes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gotest
--

ALTER SEQUENCE public.notes_id_seq OWNED BY public.notes.id;


--
-- TOC entry 207 (class 1259 OID 17331)
-- Name: project; Type: TABLE; Schema: public; Owner: gotest
--

CREATE TABLE public.project (
    id bigint NOT NULL,
    name character varying(150),
    description text,
    created_at timestamp without time zone DEFAULT now(),
    closed_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.project OWNER TO gotest;

--
-- TOC entry 206 (class 1259 OID 17329)
-- Name: project_id_seq; Type: SEQUENCE; Schema: public; Owner: gotest
--

CREATE SEQUENCE public.project_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.project_id_seq OWNER TO gotest;

--
-- TOC entry 3092 (class 0 OID 0)
-- Dependencies: 206
-- Name: project_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gotest
--

ALTER SEQUENCE public.project_id_seq OWNED BY public.project.id;


--
-- TOC entry 208 (class 1259 OID 17341)
-- Name: project_project_state; Type: TABLE; Schema: public; Owner: gotest
--

CREATE TABLE public.project_project_state (
    state_id bigint NOT NULL,
    project_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    user_id bigint
);


ALTER TABLE public.project_project_state OWNER TO gotest;

--
-- TOC entry 210 (class 1259 OID 17361)
-- Name: project_state; Type: TABLE; Schema: public; Owner: gotest
--

CREATE TABLE public.project_state (
    id bigint NOT NULL,
    name character varying(150)
);


ALTER TABLE public.project_state OWNER TO gotest;

--
-- TOC entry 209 (class 1259 OID 17359)
-- Name: project_state_id_seq; Type: SEQUENCE; Schema: public; Owner: gotest
--

CREATE SEQUENCE public.project_state_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.project_state_id_seq OWNER TO gotest;

--
-- TOC entry 3093 (class 0 OID 0)
-- Dependencies: 209
-- Name: project_state_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gotest
--

ALTER SEQUENCE public.project_state_id_seq OWNED BY public.project_state.id;


--
-- TOC entry 211 (class 1259 OID 17367)
-- Name: project_user; Type: TABLE; Schema: public; Owner: gotest
--

CREATE TABLE public.project_user (
    project_id bigint NOT NULL,
    user_id bigint NOT NULL,
    active boolean
);


ALTER TABLE public.project_user OWNER TO gotest;

--
-- TOC entry 213 (class 1259 OID 17378)
-- Name: task; Type: TABLE; Schema: public; Owner: gotest
--

CREATE TABLE public.task (
    id bigint NOT NULL,
    title character varying(150) NOT NULL,
    description text,
    project_id bigint NOT NULL,
    user_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone
);


ALTER TABLE public.task OWNER TO gotest;

--
-- TOC entry 205 (class 1259 OID 17311)
-- Name: task_state; Type: TABLE; Schema: public; Owner: gotest
--

CREATE TABLE public.task_state (
    id bigint NOT NULL,
    name character varying(150)
);


ALTER TABLE public.task_state OWNER TO gotest;

--
-- TOC entry 204 (class 1259 OID 17309)
-- Name: task_state_id_seq; Type: SEQUENCE; Schema: public; Owner: gotest
--

CREATE SEQUENCE public.task_state_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.task_state_id_seq OWNER TO gotest;

--
-- TOC entry 3094 (class 0 OID 0)
-- Dependencies: 204
-- Name: task_state_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gotest
--

ALTER SEQUENCE public.task_state_id_seq OWNED BY public.task_state.id;


--
-- TOC entry 212 (class 1259 OID 17372)
-- Name: task_task_state; Type: TABLE; Schema: public; Owner: gotest
--

CREATE TABLE public.task_task_state (
    state_id bigint NOT NULL,
    task_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    user_id bigint
);


ALTER TABLE public.task_task_state OWNER TO gotest;

--
-- TOC entry 203 (class 1259 OID 17297)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name character varying(150) NOT NULL,
    username character varying(150) NOT NULL,
    password character varying(255) NOT NULL,
    email character varying(150) NOT NULL,
    created_at time without time zone DEFAULT now(),
    updated_at time without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 17295)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 3095 (class 0 OID 0)
-- Dependencies: 202
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 2925 (class 2604 OID 17392)
-- Name: notes id; Type: DEFAULT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.notes ALTER COLUMN id SET DEFAULT nextval('public.notes_id_seq'::regclass);


--
-- TOC entry 2919 (class 2604 OID 17334)
-- Name: project id; Type: DEFAULT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.project ALTER COLUMN id SET DEFAULT nextval('public.project_id_seq'::regclass);


--
-- TOC entry 2922 (class 2604 OID 17364)
-- Name: project_state id; Type: DEFAULT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.project_state ALTER COLUMN id SET DEFAULT nextval('public.project_state_id_seq'::regclass);


--
-- TOC entry 2918 (class 2604 OID 17314)
-- Name: task_state id; Type: DEFAULT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.task_state ALTER COLUMN id SET DEFAULT nextval('public.task_state_id_seq'::regclass);


--
-- TOC entry 2916 (class 2604 OID 17300)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3085 (class 0 OID 17389)
-- Dependencies: 215
-- Data for Name: notes; Type: TABLE DATA; Schema: public; Owner: gotest
--

COPY public.notes (id, comment, created_at, user_id, task_id) FROM stdin;
\.


--
-- TOC entry 3077 (class 0 OID 17331)
-- Dependencies: 207
-- Data for Name: project; Type: TABLE DATA; Schema: public; Owner: gotest
--

COPY public.project (id, name, description, created_at, closed_at, updated_at) FROM stdin;
\.


--
-- TOC entry 3078 (class 0 OID 17341)
-- Dependencies: 208
-- Data for Name: project_project_state; Type: TABLE DATA; Schema: public; Owner: gotest
--

COPY public.project_project_state (state_id, project_id, created_at, user_id) FROM stdin;
\.


--
-- TOC entry 3080 (class 0 OID 17361)
-- Dependencies: 210
-- Data for Name: project_state; Type: TABLE DATA; Schema: public; Owner: gotest
--

COPY public.project_state (id, name) FROM stdin;
\.


--
-- TOC entry 3081 (class 0 OID 17367)
-- Dependencies: 211
-- Data for Name: project_user; Type: TABLE DATA; Schema: public; Owner: gotest
--

COPY public.project_user (project_id, user_id, active) FROM stdin;
\.


--
-- TOC entry 3083 (class 0 OID 17378)
-- Dependencies: 213
-- Data for Name: task; Type: TABLE DATA; Schema: public; Owner: gotest
--

COPY public.task (id, title, description, project_id, user_id, created_at, updated_at) FROM stdin;
\.


--
-- TOC entry 3075 (class 0 OID 17311)
-- Dependencies: 205
-- Data for Name: task_state; Type: TABLE DATA; Schema: public; Owner: gotest
--

COPY public.task_state (id, name) FROM stdin;
\.


--
-- TOC entry 3082 (class 0 OID 17372)
-- Dependencies: 212
-- Data for Name: task_task_state; Type: TABLE DATA; Schema: public; Owner: gotest
--

COPY public.task_task_state (state_id, task_id, created_at, user_id) FROM stdin;
\.


--
-- TOC entry 3073 (class 0 OID 17297)
-- Dependencies: 203
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, username, password, email, created_at, updated_at) FROM stdin;
\.


--
-- TOC entry 3096 (class 0 OID 0)
-- Dependencies: 214
-- Name: notes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gotest
--

SELECT pg_catalog.setval('public.notes_id_seq', 1, false);


--
-- TOC entry 3097 (class 0 OID 0)
-- Dependencies: 206
-- Name: project_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gotest
--

SELECT pg_catalog.setval('public.project_id_seq', 1, false);


--
-- TOC entry 3098 (class 0 OID 0)
-- Dependencies: 209
-- Name: project_state_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gotest
--

SELECT pg_catalog.setval('public.project_state_id_seq', 1, false);


--
-- TOC entry 3099 (class 0 OID 0)
-- Dependencies: 204
-- Name: task_state_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gotest
--

SELECT pg_catalog.setval('public.task_state_id_seq', 2, true);


--
-- TOC entry 3100 (class 0 OID 0)
-- Dependencies: 202
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- TOC entry 2945 (class 2606 OID 17397)
-- Name: notes notes_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.notes
    ADD CONSTRAINT notes_pkey PRIMARY KEY (id);


--
-- TOC entry 2933 (class 2606 OID 17340)
-- Name: project project_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.project
    ADD CONSTRAINT project_pkey PRIMARY KEY (id);


--
-- TOC entry 2935 (class 2606 OID 17358)
-- Name: project_project_state project_project_state_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.project_project_state
    ADD CONSTRAINT project_project_state_pkey PRIMARY KEY (state_id, project_id);


--
-- TOC entry 2937 (class 2606 OID 17366)
-- Name: project_state project_state_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.project_state
    ADD CONSTRAINT project_state_pkey PRIMARY KEY (id);


--
-- TOC entry 2939 (class 2606 OID 17371)
-- Name: project_user project_user_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.project_user
    ADD CONSTRAINT project_user_pkey PRIMARY KEY (project_id, user_id);


--
-- TOC entry 2943 (class 2606 OID 17386)
-- Name: task task_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_pkey PRIMARY KEY (id);


--
-- TOC entry 2931 (class 2606 OID 17316)
-- Name: task_state task_state_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.task_state
    ADD CONSTRAINT task_state_pkey PRIMARY KEY (id);


--
-- TOC entry 2941 (class 2606 OID 17377)
-- Name: task_task_state task_task_state_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.task_task_state
    ADD CONSTRAINT task_task_state_pkey PRIMARY KEY (state_id, task_id);


--
-- TOC entry 2927 (class 2606 OID 17308)
-- Name: users username_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT username_unique UNIQUE (username);


--
-- TOC entry 2929 (class 2606 OID 17306)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


-- Completed on 2022-02-20 19:41:22 -03

--
-- PostgreSQL database dump complete
--

