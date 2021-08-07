CREATE TABLE public.notes (
                              id serial NOT NULL,
                              note varchar NULL,
                              pubtime timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              CONSTRAINT id_pk PRIMARY KEY (id)
);

CREATE TABLE public.messages (
                                 message_id serial NOT NULL,
                                 recipient varchar NOT NULL DEFAULT 'admin',
                                 sender varchar NOT NULL DEFAULT 'admin',
                                 message varchar NOT NULL,
                                 send_time timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 CONSTRAINT message_pk PRIMARY KEY (message_id)
);

INSERT INTO public.messages
(recipient, sender, message, send_time)
VALUES('admin'::character varying, 'admin'::character varying, '', CURRENT_TIMESTAMP);

INSERT INTO public.messages (message) VALUES('hello from bloms');
