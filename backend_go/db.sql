--
-- PostgreSQL database dump
--

\restrict v8bfne2UvjUMrhUcW8mV8d2uw9WkPY7HKHUPTI8gcPO6AEe8cdupaoxb2XW7cOf

-- Dumped from database version 17.7
-- Dumped by pg_dump version 17.7

-- Started on 2026-01-02 19:02:34

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 3 (class 3079 OID 16400)
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- TOC entry 5040 (class 0 OID 0)
-- Dependencies: 3
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


--
-- TOC entry 2 (class 3079 OID 16389)
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- TOC entry 5041 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- TOC entry 903 (class 1247 OID 16438)
-- Name: client_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.client_status AS ENUM (
    'active',
    'revoked'
);


ALTER TYPE public.client_status OWNER TO postgres;

--
-- TOC entry 274 (class 1255 OID 16531)
-- Name: set_updated_at(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.set_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$;


ALTER FUNCTION public.set_updated_at() OWNER TO postgres;

SET default_tablespace = '';

--
-- TOC entry 221 (class 1259 OID 16457)
-- Name: api_logs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.api_logs (
                                 id bigint NOT NULL,
                                 client_id uuid NOT NULL,
                                 endpoint character varying(255) NOT NULL,
                                 ip_address inet NOT NULL,
                                 created_at timestamp without time zone DEFAULT now() NOT NULL
)
    PARTITION BY RANGE (created_at);


ALTER TABLE public.api_logs OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 16456)
-- Name: api_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.api_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.api_logs_id_seq OWNER TO postgres;

--
-- TOC entry 5042 (class 0 OID 0)
-- Dependencies: 220
-- Name: api_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.api_logs_id_seq OWNED BY public.api_logs.id;


SET default_table_access_method = heap;

--
-- TOC entry 223 (class 1259 OID 16481)
-- Name: api_logs_2025_01; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.api_logs_2025_01 (
                                         id bigint DEFAULT nextval('public.api_logs_id_seq'::regclass) NOT NULL,
                                         client_id uuid NOT NULL,
                                         endpoint character varying(255) NOT NULL,
                                         ip_address inet NOT NULL,
                                         created_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.api_logs_2025_01 OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 16493)
-- Name: api_logs_2025_02; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.api_logs_2025_02 (
                                         id bigint DEFAULT nextval('public.api_logs_id_seq'::regclass) NOT NULL,
                                         client_id uuid NOT NULL,
                                         endpoint character varying(255) NOT NULL,
                                         ip_address inet NOT NULL,
                                         created_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.api_logs_2025_02 OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 16469)
-- Name: api_logs_default; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.api_logs_default (
                                         id bigint DEFAULT nextval('public.api_logs_id_seq'::regclass) NOT NULL,
                                         client_id uuid NOT NULL,
                                         endpoint character varying(255) NOT NULL,
                                         ip_address inet NOT NULL,
                                         created_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.api_logs_default OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 16539)
-- Name: client_ip_whitelists; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.client_ip_whitelists (
                                             id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
                                             client_id uuid NOT NULL,
                                             ip_address inet NOT NULL,
                                             created_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.client_ip_whitelists OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16443)
-- Name: clients; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.clients (
                                id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
                                client_id character varying(64) NOT NULL,
                                name character varying(255) NOT NULL,
                                email bytea NOT NULL,
                                api_key_hash character varying(255) NOT NULL,
                                status public.client_status DEFAULT 'active'::public.client_status NOT NULL,
                                created_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.clients OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 16514)
-- Name: daily_usage; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.daily_usage (
                                    id bigint NOT NULL,
                                    client_id uuid NOT NULL,
                                    date date NOT NULL,
                                    total_requests bigint DEFAULT 0 NOT NULL,
                                    created_at timestamp without time zone DEFAULT now() NOT NULL,
                                    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.daily_usage OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 16513)
-- Name: daily_usage_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.daily_usage_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.daily_usage_id_seq OWNER TO postgres;

--
-- TOC entry 5043 (class 0 OID 0)
-- Dependencies: 225
-- Name: daily_usage_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.daily_usage_id_seq OWNED BY public.daily_usage.id;


--
-- TOC entry 4820 (class 0 OID 0)
-- Name: api_logs_2025_01; Type: TABLE ATTACH; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.api_logs ATTACH PARTITION public.api_logs_2025_01 FOR VALUES FROM ('2025-01-01 00:00:00') TO ('2025-02-01 00:00:00');


--
-- TOC entry 4821 (class 0 OID 0)
-- Name: api_logs_2025_02; Type: TABLE ATTACH; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.api_logs ATTACH PARTITION public.api_logs_2025_02 FOR VALUES FROM ('2025-02-01 00:00:00') TO ('2025-03-01 00:00:00');


--
-- TOC entry 4819 (class 0 OID 0)
-- Name: api_logs_default; Type: TABLE ATTACH; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.api_logs ATTACH PARTITION public.api_logs_default DEFAULT;


--
-- TOC entry 4825 (class 2604 OID 16460)
-- Name: api_logs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.api_logs ALTER COLUMN id SET DEFAULT nextval('public.api_logs_id_seq'::regclass);


--
-- TOC entry 4833 (class 2604 OID 16517)
-- Name: daily_usage id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.daily_usage ALTER COLUMN id SET DEFAULT nextval('public.daily_usage_id_seq'::regclass);


--
-- TOC entry 5030 (class 0 OID 16481)
-- Dependencies: 223
-- Data for Name: api_logs_2025_01; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.api_logs_2025_01 (id, client_id, endpoint, ip_address, created_at) FROM stdin;
\.


--
-- TOC entry 5031 (class 0 OID 16493)
-- Dependencies: 224
-- Data for Name: api_logs_2025_02; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.api_logs_2025_02 (id, client_id, endpoint, ip_address, created_at) FROM stdin;
\.


--
-- TOC entry 5029 (class 0 OID 16469)
-- Dependencies: 222
-- Data for Name: api_logs_default; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.api_logs_default (id, client_id, endpoint, ip_address, created_at) FROM stdin;
\.


--
-- TOC entry 5034 (class 0 OID 16539)
-- Dependencies: 227
-- Data for Name: client_ip_whitelists; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.client_ip_whitelists (id, client_id, ip_address, created_at) FROM stdin;
c1c7d529-0aca-4a6d-9f58-afb0fd636eed	9c8fd397-e5fc-4a78-af5d-8139e3dbafc3	127.0.0.1	2026-01-02 11:08:25.675
82df00dc-286b-41ca-875f-8c6f8db664e0	9c8fd397-e5fc-4a78-af5d-8139e3dbafc3	192.168.1.10	2026-01-02 11:08:25.686
\.


--
-- TOC entry 5027 (class 0 OID 16443)
-- Dependencies: 219
-- Data for Name: clients; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.clients (id, client_id, name, email, api_key_hash, status, created_at) FROM stdin;
9c8fd397-e5fc-4a78-af5d-8139e3dbafc3	client_001	Acme Corp	\\x61646d696e4061636d652e636f6d3a7333637233546b33794e33784d33643135	ab5ceab122a1e817847be3bdd806e6dce7f49555a873de537665f5efd0df6550	active	2026-01-01 19:56:31.896143
49f3aefb-fcff-436d-8e57-13984c71f228	client_002	Test Corp	\\x61646d696e40746573742e636f6d3a7333637233546b33794e33784d33643135	2dcf8c17a1f5fd46d41c30e82d1c35798ed6efaeec88c707b0e0e702eee3ecb6	active	2026-01-01 21:56:39.791567
7be0291d-6d35-4351-8814-c708ba797393	client_003	Test	\\x61646d696e4074657374332e636f6d3a7333637233546b33794e33784d33643135	091c5f556c745298b257a0381a3af17598b023424c25a6c9691849f44fc9361f	active	2026-01-01 21:58:35.408624
1607bb4d-74cc-414f-b0f2-2cbe9c30d954	client_004	Test	\\x61646d696e4074657374332e636f6d3a7333637233546b33794e33784d33643135	8780c38ceb8ce074e6eabb76f109b463512a8189ce96b408cb485892671a6b97	active	2026-01-02 04:26:33.132
b714e2ed-dbaf-4604-a007-d3b2f13fce58	client_005	Test	\\x61646d696e4074657374332e636f6d3a7333637233546b33794e33784d33643135	5fa7e4eeb2e3ee59a8be770757d1c4634f53ff88ee355be967c97c9ded7a9a2b	active	2026-01-02 04:26:55.899
\.


--
-- TOC entry 5033 (class 0 OID 16514)
-- Dependencies: 226
-- Data for Name: daily_usage; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.daily_usage (id, client_id, date, total_requests, created_at, updated_at) FROM stdin;
1	9c8fd397-e5fc-4a78-af5d-8139e3dbafc3	2026-01-01	1004	2026-01-02 00:21:55.469658	2026-01-02 17:26:57.515418
2	9c8fd397-e5fc-4a78-af5d-8139e3dbafc3	2026-01-02	22	2026-01-02 00:21:55.501397	2026-01-02 17:26:57.543835
\.


--
-- TOC entry 5044 (class 0 OID 0)
-- Dependencies: 220
-- Name: api_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.api_logs_id_seq', 1, false);


--
-- TOC entry 5045 (class 0 OID 0)
-- Dependencies: 225
-- Name: daily_usage_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.daily_usage_id_seq', 302, true);


--
-- TOC entry 4845 (class 2606 OID 16463)
-- Name: api_logs api_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.api_logs
    ADD CONSTRAINT api_logs_pkey PRIMARY KEY (id, created_at);


--
-- TOC entry 4855 (class 2606 OID 16487)
-- Name: api_logs_2025_01 api_logs_2025_01_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.api_logs_2025_01
    ADD CONSTRAINT api_logs_2025_01_pkey PRIMARY KEY (id, created_at);


--
-- TOC entry 4859 (class 2606 OID 16499)
-- Name: api_logs_2025_02 api_logs_2025_02_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.api_logs_2025_02
    ADD CONSTRAINT api_logs_2025_02_pkey PRIMARY KEY (id, created_at);


--
-- TOC entry 4851 (class 2606 OID 16475)
-- Name: api_logs_default api_logs_default_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.api_logs_default
    ADD CONSTRAINT api_logs_default_pkey PRIMARY KEY (id, created_at);


--
-- TOC entry 4866 (class 2606 OID 16547)
-- Name: client_ip_whitelists client_ip_whitelists_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.client_ip_whitelists
    ADD CONSTRAINT client_ip_whitelists_pkey PRIMARY KEY (id);


--
-- TOC entry 4840 (class 2606 OID 16454)
-- Name: clients clients_client_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.clients
    ADD CONSTRAINT clients_client_id_key UNIQUE (client_id);


--
-- TOC entry 4842 (class 2606 OID 16452)
-- Name: clients clients_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.clients
    ADD CONSTRAINT clients_pkey PRIMARY KEY (id);


--
-- TOC entry 4861 (class 2606 OID 16522)
-- Name: daily_usage daily_usage_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.daily_usage
    ADD CONSTRAINT daily_usage_pkey PRIMARY KEY (id);


--
-- TOC entry 4868 (class 2606 OID 16549)
-- Name: client_ip_whitelists uq_client_ip; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.client_ip_whitelists
    ADD CONSTRAINT uq_client_ip UNIQUE (client_id, ip_address);


--
-- TOC entry 4864 (class 2606 OID 16524)
-- Name: daily_usage uq_daily_usage_client_date; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.daily_usage
    ADD CONSTRAINT uq_daily_usage_client_date UNIQUE (client_id, date);


--
-- TOC entry 4846 (class 1259 OID 16505)
-- Name: idx_api_logs_client_time; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_api_logs_client_time ON ONLY public.api_logs USING btree (client_id, created_at);


--
-- TOC entry 4852 (class 1259 OID 16506)
-- Name: api_logs_2025_01_client_id_created_at_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX api_logs_2025_01_client_id_created_at_idx ON public.api_logs_2025_01 USING btree (client_id, created_at);


--
-- TOC entry 4847 (class 1259 OID 16509)
-- Name: idx_api_logs_time; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_api_logs_time ON ONLY public.api_logs USING btree (created_at);


--
-- TOC entry 4853 (class 1259 OID 16510)
-- Name: api_logs_2025_01_created_at_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX api_logs_2025_01_created_at_idx ON public.api_logs_2025_01 USING btree (created_at);


--
-- TOC entry 4856 (class 1259 OID 16507)
-- Name: api_logs_2025_02_client_id_created_at_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX api_logs_2025_02_client_id_created_at_idx ON public.api_logs_2025_02 USING btree (client_id, created_at);


--
-- TOC entry 4857 (class 1259 OID 16511)
-- Name: api_logs_2025_02_created_at_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX api_logs_2025_02_created_at_idx ON public.api_logs_2025_02 USING btree (created_at);


--
-- TOC entry 4848 (class 1259 OID 16508)
-- Name: api_logs_default_client_id_created_at_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX api_logs_default_client_id_created_at_idx ON public.api_logs_default USING btree (client_id, created_at);


--
-- TOC entry 4849 (class 1259 OID 16512)
-- Name: api_logs_default_created_at_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX api_logs_default_created_at_idx ON public.api_logs_default USING btree (created_at);


--
-- TOC entry 4843 (class 1259 OID 16455)
-- Name: idx_clients_api_key_hash; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_clients_api_key_hash ON public.clients USING btree (api_key_hash);


--
-- TOC entry 4862 (class 1259 OID 16530)
-- Name: idx_daily_usage_date; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_daily_usage_date ON public.daily_usage USING btree (date);


--
-- TOC entry 4872 (class 0 OID 0)
-- Name: api_logs_2025_01_client_id_created_at_idx; Type: INDEX ATTACH; Schema: public; Owner: postgres
--

ALTER INDEX public.idx_api_logs_client_time ATTACH PARTITION public.api_logs_2025_01_client_id_created_at_idx;


--
-- TOC entry 4873 (class 0 OID 0)
-- Name: api_logs_2025_01_created_at_idx; Type: INDEX ATTACH; Schema: public; Owner: postgres
--

ALTER INDEX public.idx_api_logs_time ATTACH PARTITION public.api_logs_2025_01_created_at_idx;


--
-- TOC entry 4874 (class 0 OID 0)
-- Name: api_logs_2025_01_pkey; Type: INDEX ATTACH; Schema: public; Owner: postgres
--

ALTER INDEX public.api_logs_pkey ATTACH PARTITION public.api_logs_2025_01_pkey;


--
-- TOC entry 4875 (class 0 OID 0)
-- Name: api_logs_2025_02_client_id_created_at_idx; Type: INDEX ATTACH; Schema: public; Owner: postgres
--

ALTER INDEX public.idx_api_logs_client_time ATTACH PARTITION public.api_logs_2025_02_client_id_created_at_idx;


--
-- TOC entry 4876 (class 0 OID 0)
-- Name: api_logs_2025_02_created_at_idx; Type: INDEX ATTACH; Schema: public; Owner: postgres
--

ALTER INDEX public.idx_api_logs_time ATTACH PARTITION public.api_logs_2025_02_created_at_idx;


--
-- TOC entry 4877 (class 0 OID 0)
-- Name: api_logs_2025_02_pkey; Type: INDEX ATTACH; Schema: public; Owner: postgres
--

ALTER INDEX public.api_logs_pkey ATTACH PARTITION public.api_logs_2025_02_pkey;


--
-- TOC entry 4869 (class 0 OID 0)
-- Name: api_logs_default_client_id_created_at_idx; Type: INDEX ATTACH; Schema: public; Owner: postgres
--

ALTER INDEX public.idx_api_logs_client_time ATTACH PARTITION public.api_logs_default_client_id_created_at_idx;


--
-- TOC entry 4870 (class 0 OID 0)
-- Name: api_logs_default_created_at_idx; Type: INDEX ATTACH; Schema: public; Owner: postgres
--

ALTER INDEX public.idx_api_logs_time ATTACH PARTITION public.api_logs_default_created_at_idx;


--
-- TOC entry 4871 (class 0 OID 0)
-- Name: api_logs_default_pkey; Type: INDEX ATTACH; Schema: public; Owner: postgres
--

ALTER INDEX public.api_logs_pkey ATTACH PARTITION public.api_logs_default_pkey;


--
-- TOC entry 4881 (class 2620 OID 16532)
-- Name: daily_usage trg_daily_usage_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER trg_daily_usage_updated_at BEFORE UPDATE ON public.daily_usage FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- TOC entry 4878 (class 2606 OID 16464)
-- Name: api_logs fk_api_logs_client; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE public.api_logs
    ADD CONSTRAINT fk_api_logs_client FOREIGN KEY (client_id) REFERENCES public.clients(id);


--
-- TOC entry 4879 (class 2606 OID 16525)
-- Name: daily_usage fk_daily_usage_client; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.daily_usage
    ADD CONSTRAINT fk_daily_usage_client FOREIGN KEY (client_id) REFERENCES public.clients(id);


--
-- TOC entry 4880 (class 2606 OID 16550)
-- Name: client_ip_whitelists fk_ip_whitelist_client; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.client_ip_whitelists
    ADD CONSTRAINT fk_ip_whitelist_client FOREIGN KEY (client_id) REFERENCES public.clients(id);


-- Completed on 2026-01-02 19:02:34

--
-- PostgreSQL database dump complete
--

\unrestrict v8bfne2UvjUMrhUcW8mV8d2uw9WkPY7HKHUPTI8gcPO6AEe8cdupaoxb2XW7cOf

