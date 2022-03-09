--
-- PostgreSQL database dump
--

-- Dumped from database version 12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 13.1

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

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
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
-- Name: notes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gotest
--

ALTER SEQUENCE public.notes_id_seq OWNED BY public.notes.id;


--
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
-- Name: project_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gotest
--

ALTER SEQUENCE public.project_id_seq OWNED BY public.project.id;


--
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
-- Name: project_state; Type: TABLE; Schema: public; Owner: gotest
--

CREATE TABLE public.project_state (
    id bigint NOT NULL,
    name character varying(150)
);


ALTER TABLE public.project_state OWNER TO gotest;

--
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
-- Name: project_state_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gotest
--

ALTER SEQUENCE public.project_state_id_seq OWNED BY public.project_state.id;


--
-- Name: project_user; Type: TABLE; Schema: public; Owner: gotest
--

CREATE TABLE public.project_user (
    project_id bigint NOT NULL,
    user_id bigint NOT NULL,
    active boolean
);


ALTER TABLE public.project_user OWNER TO gotest;

--
-- Name: task; Type: TABLE; Schema: public; Owner: gotest
--

CREATE TABLE public.task (
    id bigint NOT NULL,
    title character varying(150) NOT NULL,
    description text,
    project_id bigint NOT NULL,
    user_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone
);


ALTER TABLE public.task OWNER TO gotest;

--
-- Name: task_id_seq; Type: SEQUENCE; Schema: public; Owner: gotest
--

CREATE SEQUENCE public.task_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.task_id_seq OWNER TO gotest;

--
-- Name: task_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gotest
--

ALTER SEQUENCE public.task_id_seq OWNED BY public.task.id;


--
-- Name: task_state; Type: TABLE; Schema: public; Owner: gotest
--

CREATE TABLE public.task_state (
    id bigint NOT NULL,
    name character varying(150)
);


ALTER TABLE public.task_state OWNER TO gotest;

--
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
-- Name: task_state_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gotest
--

ALTER SEQUENCE public.task_state_id_seq OWNED BY public.task_state.id;


--
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
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name character varying(150) NOT NULL,
    username character varying(150) NOT NULL,
    password character varying(255) NOT NULL,
    email character varying(150) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
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
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: notes id; Type: DEFAULT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.notes ALTER COLUMN id SET DEFAULT nextval('public.notes_id_seq'::regclass);


--
-- Name: project id; Type: DEFAULT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.project ALTER COLUMN id SET DEFAULT nextval('public.project_id_seq'::regclass);


--
-- Name: project_state id; Type: DEFAULT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.project_state ALTER COLUMN id SET DEFAULT nextval('public.project_state_id_seq'::regclass);


--
-- Name: task id; Type: DEFAULT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.task ALTER COLUMN id SET DEFAULT nextval('public.task_id_seq'::regclass);


--
-- Name: task_state id; Type: DEFAULT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.task_state ALTER COLUMN id SET DEFAULT nextval('public.task_state_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: notes notes_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.notes
    ADD CONSTRAINT notes_pkey PRIMARY KEY (id);


--
-- Name: project project_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.project
    ADD CONSTRAINT project_pkey PRIMARY KEY (id);


--
-- Name: project_project_state project_project_state_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.project_project_state
    ADD CONSTRAINT project_project_state_pkey PRIMARY KEY (state_id, project_id);


--
-- Name: project_state project_state_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.project_state
    ADD CONSTRAINT project_state_pkey PRIMARY KEY (id);


--
-- Name: project_user project_user_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.project_user
    ADD CONSTRAINT project_user_pkey PRIMARY KEY (project_id, user_id);


--
-- Name: task task_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_pkey PRIMARY KEY (id);


--
-- Name: task_state task_state_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.task_state
    ADD CONSTRAINT task_state_pkey PRIMARY KEY (id);


--
-- Name: task_task_state task_task_state_pkey; Type: CONSTRAINT; Schema: public; Owner: gotest
--

ALTER TABLE ONLY public.task_task_state
    ADD CONSTRAINT task_task_state_pkey PRIMARY KEY (state_id, task_id);


--
-- Name: users username_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT username_unique UNIQUE (username);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

