--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.2

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
-- Name: article_favorites; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.article_favorites (
    id uuid NOT NULL,
    article_id uuid,
    user_id uuid
);


ALTER TABLE public.article_favorites OWNER TO postgres;

--
-- Name: article_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.article_tags (
    id uuid NOT NULL,
    article_id uuid,
    tag_id uuid
);


ALTER TABLE public.article_tags OWNER TO postgres;

--
-- Name: articles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.articles (
    id uuid NOT NULL,
    title character varying(255),
    slug character varying(255),
    description character varying(255),
    body text,
    user_id uuid,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.articles OWNER TO postgres;

--
-- Name: comments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.comments (
    id uuid NOT NULL,
    body text,
    article_id uuid,
    user_id uuid,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.comments OWNER TO postgres;

--
-- Name: follows; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.follows (
    id uuid NOT NULL,
    user_id uuid,
    follow_id uuid
);


ALTER TABLE public.follows OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tags (
    id uuid NOT NULL,
    name character varying(255)
);


ALTER TABLE public.tags OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: article_favorites article_favorites_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.article_favorites
    ADD CONSTRAINT article_favorites_pkey PRIMARY KEY (id);


--
-- Name: article_tags article_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.article_tags
    ADD CONSTRAINT article_tags_pkey PRIMARY KEY (id);


--
-- Name: articles articles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.articles
    ADD CONSTRAINT articles_pkey PRIMARY KEY (id);


--
-- Name: comments comments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (id);


--
-- Name: follows follows_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.follows
    ADD CONSTRAINT follows_pkey PRIMARY KEY (id);


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: article_favorites_article_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX article_favorites_article_id_idx ON public.article_favorites USING btree (article_id);


--
-- Name: article_favorites_article_id_user_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX article_favorites_article_id_user_id_idx ON public.article_favorites USING btree (article_id, user_id);


--
-- Name: article_tags_article_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX article_tags_article_id_idx ON public.article_tags USING btree (article_id);


--
-- Name: articles_slug_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX articles_slug_idx ON public.articles USING btree (slug);


--
-- Name: comments_article_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX comments_article_id_idx ON public.comments USING btree (article_id);


--
-- Name: comments_user_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX comments_user_id_idx ON public.comments USING btree (user_id);


--
-- Name: follows_user_id_follow_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX follows_user_id_follow_id_idx ON public.follows USING btree (user_id, follow_id);


--
-- Name: follows_user_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX follows_user_id_idx ON public.follows USING btree (user_id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: tags_name_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX tags_name_idx ON public.tags USING btree (name);


--
-- Name: users_email_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX users_email_idx ON public.users USING btree (email);


--
-- Name: article_tags article_tags_articles_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.article_tags
    ADD CONSTRAINT article_tags_articles_id_fk FOREIGN KEY (article_id) REFERENCES public.articles(id);


--
-- Name: article_tags article_tags_tags_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.article_tags
    ADD CONSTRAINT article_tags_tags_id_fk FOREIGN KEY (tag_id) REFERENCES public.tags(id);


--
-- Name: articles articles_users_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.articles
    ADD CONSTRAINT articles_users_id_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

