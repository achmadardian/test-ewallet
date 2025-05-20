--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4 (Debian 17.4-1.pgdg120+2)
-- Dumped by pg_dump version 17.4

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
-- Name: transaction_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.transaction_status AS ENUM (
    'PENDING',
    'SUCCESS',
    'FAILED'
);


ALTER TYPE public.transaction_status OWNER TO postgres;

--
-- Name: transaction_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.transaction_type AS ENUM (
    'DEBIT',
    'CREDIT'
);


ALTER TYPE public.transaction_type OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: payments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payments (
    id uuid NOT NULL,
    transaction_id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.payments OWNER TO postgres;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- Name: topups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.topups (
    id uuid NOT NULL,
    transaction_id uuid NOT NULL,
    amount bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.topups OWNER TO postgres;

--
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id uuid NOT NULL,
    status public.transaction_status NOT NULL,
    user_id uuid NOT NULL,
    transaction_type public.transaction_type NOT NULL,
    amount bigint NOT NULL,
    remarks character varying(50),
    balance_before bigint NOT NULL,
    balance_after bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- Name: transfers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transfers (
    id uuid NOT NULL,
    sender_transaction_id uuid NOT NULL,
    receiver_transaction_id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.transfers OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    first_name character varying(50) NOT NULL,
    last_name character varying(50) NOT NULL,
    phone_number character varying(20) NOT NULL,
    address text NOT NULL,
    pin character(60) NOT NULL,
    balance bigint DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Data for Name: payments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payments (id, transaction_id, created_at, updated_at, deleted_at) FROM stdin;
f19b95a9-8549-4bde-9932-4157f65ce2dd	2de42b3e-304b-4c00-b4ba-c6797c624350	2025-05-20 21:20:39.402987	2025-05-20 21:20:39.402987	\N
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schema_migrations (version, dirty) FROM stdin;
5	f
\.


--
-- Data for Name: topups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.topups (id, transaction_id, amount, created_at, updated_at, deleted_at) FROM stdin;
0cdc7404-1dfd-41ec-96df-31a9f28ded9a	48160786-9178-4b04-8eab-1a2c309bfb03	10000	2025-05-20 21:19:36.814742	2025-05-20 21:19:36.814742	\N
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transactions (id, status, user_id, transaction_type, amount, remarks, balance_before, balance_after, created_at, updated_at, deleted_at) FROM stdin;
48160786-9178-4b04-8eab-1a2c309bfb03	SUCCESS	62b9653f-6f63-4cb8-b18e-7a45a898fe9c	CREDIT	10000	\N	0	10000	2025-05-20 21:19:36.806606	2025-05-20 21:19:36.806606	\N
2de42b3e-304b-4c00-b4ba-c6797c624350	SUCCESS	62b9653f-6f63-4cb8-b18e-7a45a898fe9c	DEBIT	2000	beli cilok online	10000	8000	2025-05-20 21:20:39.399409	2025-05-20 21:20:39.399409	\N
6cf5faf3-15fd-4216-9b58-3bbd61392326	SUCCESS	62b9653f-6f63-4cb8-b18e-7a45a898fe9c	DEBIT	5000	hadiah buat kamu	8000	3000	2025-05-20 21:21:30.416972	2025-05-20 21:21:30.416972	\N
9c9f8a4b-b7ad-4501-9409-712297460f06	SUCCESS	3e35d9d4-0641-467d-8f96-aaf64cfaa971	CREDIT	5000	hadiah buat kamu	0	5000	2025-05-20 21:21:30.417676	2025-05-20 21:21:30.417676	\N
\.


--
-- Data for Name: transfers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transfers (id, sender_transaction_id, receiver_transaction_id, created_at, updated_at, deleted_at) FROM stdin;
99efb095-c254-4b7c-908b-1ae205b1b7db	6cf5faf3-15fd-4216-9b58-3bbd61392326	9c9f8a4b-b7ad-4501-9409-712297460f06	2025-05-20 21:21:30.418761	2025-05-20 21:21:30.418761	\N
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, first_name, last_name, phone_number, address, pin, balance, created_at, updated_at, deleted_at) FROM stdin;
62b9653f-6f63-4cb8-b18e-7a45a898fe9c	achmad	ardian	0881	jl. malaka gg. sakura	$2a$10$NnWJlVVEV59FJrtGzLzOK.7RrFCr2kUvidkSTJ/ztWNCUCdsMEV8m	3000	2025-05-20 21:17:01.222601	2025-05-20 21:19:18.379378	\N
3e35d9d4-0641-467d-8f96-aaf64cfaa971	fulan	bin fulan	0882	jl. mangga	$2a$10$OgPPxOfEqljzTLEH.Q0liOTh.zEa9j68tKQtmF7STh82OgpZgzyUO	5000	2025-05-20 21:21:10.897797	2025-05-20 21:21:10.897797	\N
\.


--
-- Name: payments payments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: topups topups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topups
    ADD CONSTRAINT topups_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: transfers transfers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transfers
    ADD CONSTRAINT transfers_pkey PRIMARY KEY (id);


--
-- Name: users users_phone_number_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_phone_number_key UNIQUE (phone_number);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: payments payments_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES public.transactions(id);


--
-- Name: topups topups_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topups
    ADD CONSTRAINT topups_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES public.transactions(id);


--
-- Name: transactions transactions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: transfers transfers_receiver_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transfers
    ADD CONSTRAINT transfers_receiver_transaction_id_fkey FOREIGN KEY (receiver_transaction_id) REFERENCES public.transactions(id);


--
-- Name: transfers transfers_sender_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transfers
    ADD CONSTRAINT transfers_sender_transaction_id_fkey FOREIGN KEY (sender_transaction_id) REFERENCES public.transactions(id);


--
-- PostgreSQL database dump complete
--

