-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2017-10-20 18:22:42.299

-- tables
-- Table: Contactos
CREATE TABLE Contactos (
    Usuarios_1_id int  NOT NULL,
    Usuarios_2_id int  NOT NULL,
    CONSTRAINT Contactos_pk PRIMARY KEY (Usuarios_1_id,Usuarios_2_id)
);

-- Table: Estados
CREATE TABLE Estados (
    id serial  NOT NULL,
    estado varchar(64)  NOT NULL,
    CONSTRAINT Estados_pk PRIMARY KEY (id)
);

-- Table: Mensajes
CREATE TABLE Mensajes (
    id bigserial  NOT NULL,
    mensaje varchar(1024)  NOT NULL,
    fecha_creacion date  NOT NULL,
    visto smallint  NOT NULL,
    Reportes_id int8  NOT NULL,
    Usuarios_id int  NOT NULL,
    CONSTRAINT Mensajes_pk PRIMARY KEY (id)
);

-- Table: Reportes
CREATE TABLE Reportes (
    id bigserial  NOT NULL,
    fecha_creacion date  NOT NULL,
    estado boolean  NOT NULL,
    posicion_lat decimal  NOT NULL,
    posicion_long decimal  NOT NULL,
    Usuarios_id int  NOT NULL,
    CONSTRAINT Reportes_pk PRIMARY KEY (id)
);

-- Table: Reportes_Estados
CREATE TABLE Reportes_Estados (
    Reportes_id int8  NOT NULL,
    Estados_id int  NOT NULL,
    CONSTRAINT Reportes_Estados_pk PRIMARY KEY (Reportes_id,Estados_id)
);

-- Table: Usuarios
CREATE TABLE Usuarios (
    id serial  NOT NULL,
    mail varchar(128)  NOT NULL,
    pass varchar(256)  NOT NULL,
    nombre_primero varchar(128)  NOT NULL,
    nombre_segundo varchar(128)  NULL,
    apellido_paterno varchar(128)  NOT NULL,
    apellido_materno varchar(128)  NOT NULL,
    fecha_nacimiento date  NOT NULL,
    telefono varchar(20)  NOT NULL,
    fb_usuario varchar(128)  NOT NULL,
    CONSTRAINT Usuarios_pk PRIMARY KEY (id)
);

-- foreign keys
-- Reference: Contactos_Usuarios_1 (table: Contactos)
ALTER TABLE Contactos ADD CONSTRAINT Contactos_Usuarios_1
    FOREIGN KEY (Usuarios_2_id)
    REFERENCES Usuarios (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Contactos_Usuarios_2 (table: Contactos)
ALTER TABLE Contactos ADD CONSTRAINT Contactos_Usuarios_2
    FOREIGN KEY (Usuarios_1_id)
    REFERENCES Usuarios (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Mensajes_Reportes (table: Mensajes)
ALTER TABLE Mensajes ADD CONSTRAINT Mensajes_Reportes
    FOREIGN KEY (Reportes_id)
    REFERENCES Reportes (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Mensajes_Usuarios (table: Mensajes)
ALTER TABLE Mensajes ADD CONSTRAINT Mensajes_Usuarios
    FOREIGN KEY (Usuarios_id)
    REFERENCES Usuarios (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Reportes_Usuarios (table: Reportes)
ALTER TABLE Reportes ADD CONSTRAINT Reportes_Usuarios
    FOREIGN KEY (Usuarios_id)
    REFERENCES Usuarios (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Repotes_Estados_Estados (table: Reportes_Estados)
ALTER TABLE Reportes_Estados ADD CONSTRAINT Repotes_Estados_Estados
    FOREIGN KEY (Estados_id)
    REFERENCES Estados (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Repotes_Estados_Reportes (table: Reportes_Estados)
ALTER TABLE Reportes_Estados ADD CONSTRAINT Repotes_Estados_Reportes
    FOREIGN KEY (Reportes_id)
    REFERENCES Reportes (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- End of file.

