-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2017-08-11 21:55:06.859

-- tables
-- Table: abilities
CREATE TABLE abilities (
    id SERIAL  NOT NULL,
    ability varchar(256)  NOT NULL,
    CONSTRAINT abilities_pk PRIMARY KEY (id)
);

-- Table: answers
CREATE TABLE answers (
    id SERIAL  NOT NULL,
    "createAt" datetime  NOT NULL,
    answer text  NOT NULL,
    problem_id int  NOT NULL,
    user_id int  NOT NULL,
    CONSTRAINT answers_pk PRIMARY KEY (id)
);

-- Table: emergencies
CREATE TABLE emergencies (
    id SERIAL  NOT NULL,
    "createAt" datetime  NOT NULL,
    title varchar(100)  NOT NULL,
    place_latitude double precision  NOT NULL,
    place_longitude double precision  NOT NULL,
    place_radius double precision  NOT NULL,
    description text  NOT NULL,
    commune varchar(50)  NOT NULL,
    city varchar(50)  NOT NULL,
    region varchar(50)  NOT NULL,
    emergency_type_id int  NOT NULL,
    user_id int  NOT NULL,
    emergency_status_id int  NOT NULL,
    CONSTRAINT id PRIMARY KEY (id)
);

-- Table: emergencies_statuses
CREATE TABLE emergencies_statuses (
    id SERIAL  NOT NULL,
    status varchar(30)  NOT NULL,
    CONSTRAINT emergencies_statuses_pk PRIMARY KEY (id)
);

-- Table: emergencies_types
CREATE TABLE emergencies_types (
    id SERIAL  NOT NULL,
    type varchar(256)  NOT NULL,
    CONSTRAINT emergencies_types_pk PRIMARY KEY (id)
);

-- Table: files
CREATE TABLE files (
    id SERIAL  NOT NULL,
    file text  NOT NULL,
    mission_id int  NOT NULL,
    CONSTRAINT files_pk PRIMARY KEY (id)
);

-- Table: history_missions
CREATE TABLE history_missions (
    mission_id int  NOT NULL,
    volunteer_id int  NOT NULL,
    history_mission_state_id int  NOT NULL,
    CONSTRAINT history_missions_pk PRIMARY KEY (mission_id,volunteer_id)
);

-- Table: history_missions_states
CREATE TABLE history_missions_states (
    id SERIAL  NOT NULL,
    state varchar(30)  NOT NULL,
    CONSTRAINT history_missions_states_pk PRIMARY KEY (id)
);

-- Table: missions
CREATE TABLE missions (
    id SERIAL  NOT NULL,
    "createAt" datetime  NOT NULL,
    meeting_point_latitude double precision  NOT NULL,
    meeting_point_longitude double precision  NOT NULL,
    title varchar(100)  NOT NULL,
    description text  NOT NULL,
    meeting_point_address varchar(1024)  NOT NULL,
    start_date datetime  NOT NULL,
    finish_date datetime  NOT NULL,
    scheduled_start_date datetime  NOT NULL,
    scheduled_finish_date datetime  NOT NULL,
    assertiveness_text double precision  NOT NULL,
    emergency_id int  NOT NULL,
    user_id int  NOT NULL,
    mission_status_id int  NOT NULL,
    CONSTRAINT missions_pk PRIMARY KEY (id)
);

-- Table: missions_abilities
CREATE TABLE missions_abilities (
    mission_id int  NOT NULL,
    ability_id int  NOT NULL,
    CONSTRAINT missions_abilities_pk PRIMARY KEY (mission_id,ability_id)
);

-- Table: missions_statuses
CREATE TABLE missions_statuses (
    id SERIAL  NOT NULL,
    status varchar(40)  NOT NULL,
    CONSTRAINT missions_statuses_pk PRIMARY KEY (id)
);

-- Table: problems
CREATE TABLE problems (
    id SERIAL  NOT NULL,
    "createAt" datetime  NOT NULL,
    title varchar(100)  NOT NULL,
    description text  NOT NULL,
    status int  NOT NULL,
    assertiveness_text double precision  NOT NULL,
    mission_id int  NOT NULL,
    user_id int  NOT NULL,
    CONSTRAINT problems_pk PRIMARY KEY (id)
);

-- Table: users
CREATE TABLE users (
    id SERIAL  NOT NULL,
    first_name varchar(40)  NOT NULL,
    last_name varchar(40)  NOT NULL,
    birthday date  NOT NULL,
    password varchar(512)  NOT NULL,
    contact_phone_number int  NOT NULL,
    emergency_phone_number int  NOT NULL,
    life_insurance boolean  NOT NULL,
    enabled boolean  NOT NULL,
    user_type_id int  NOT NULL,
    email varchar(256)  NOT NULL,
    CONSTRAINT email UNIQUE (email) NOT DEFERRABLE  INITIALLY IMMEDIATE,
    CONSTRAINT users_pk PRIMARY KEY (id)
);

-- Table: users_types
CREATE TABLE users_types (
    id SERIAL  NOT NULL,
    type varchar(30)  NOT NULL,
    CONSTRAINT users_types_pk PRIMARY KEY (id)
);

-- Table: volunteers
CREATE TABLE volunteers (
    id SERIAL  NOT NULL,
    user_id int  NOT NULL,
    token varchar(256) NOT NULL,
    volunteer_status_id int  NOT NULL,
    CONSTRAINT volunteers_pk PRIMARY KEY (id)
);

-- Table: volunteers_abilities
CREATE TABLE volunteers_abilities (
    volunteer_id int  NOT NULL,
    ability_id int  NOT NULL,
    CONSTRAINT volunteers_abilities_pk PRIMARY KEY (volunteer_id,ability_id)
);

-- Table: volunteers_statuses
CREATE TABLE volunteers_statuses (
    id SERIAL  NOT NULL,
    status varchar(30)  NOT NULL,
    CONSTRAINT volunteers_statuses_pk PRIMARY KEY (id)
);

-- foreign keys
-- Reference: answers_problems (table: answers)
ALTER TABLE answers ADD CONSTRAINT answers_problems
    FOREIGN KEY (problem_id)
    REFERENCES problems (id) ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: answers_users (table: answers)
ALTER TABLE answers ADD CONSTRAINT answers_users
    FOREIGN KEY (user_id)
    REFERENCES users (id)  ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: emergencies_emergencies_statuses (table: emergencies)
ALTER TABLE emergencies ADD CONSTRAINT emergencies_emergencies_statuses
    FOREIGN KEY (emergency_status_id)
    REFERENCES emergencies_statuses (id) ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: emergencies_emergencies_types (table: emergencies)
ALTER TABLE emergencies ADD CONSTRAINT emergencies_emergencies_types
    FOREIGN KEY (emergency_type_id)
    REFERENCES emergencies_types (id)  ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: emergencies_users (table: emergencies)
ALTER TABLE emergencies ADD CONSTRAINT emergencies_users
    FOREIGN KEY (user_id)
    REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE 
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: files_missions (table: files)
ALTER TABLE files ADD CONSTRAINT files_missions
    FOREIGN KEY (mission_id)
    REFERENCES missions (id)  ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: history_missions_history_missions_state (table: history_missions)
ALTER TABLE history_missions ADD CONSTRAINT history_missions_history_missions_state
    FOREIGN KEY (history_mission_state_id)
    REFERENCES history_missions_states (id)  ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: history_missions_missions (table: history_missions)
ALTER TABLE history_missions ADD CONSTRAINT history_missions_missions
    FOREIGN KEY (mission_id)
    REFERENCES missions (id)  ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: history_missions_volunteers (table: history_missions)
ALTER TABLE history_missions ADD CONSTRAINT history_missions_volunteers
    FOREIGN KEY (volunteer_id)
    REFERENCES volunteers (id) ON UPDATE CASCADE ON DELETE CASCADE 
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: missions_abilities_abilities (table: missions_abilities)
ALTER TABLE missions_abilities ADD CONSTRAINT missions_abilities_abilities
    FOREIGN KEY (ability_id)
    REFERENCES abilities (id)  ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: missions_abilities_missions (table: missions_abilities)
ALTER TABLE missions_abilities ADD CONSTRAINT missions_abilities_missions
    FOREIGN KEY (mission_id)
    REFERENCES missions (id)  ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: missions_emergencies (table: missions)
ALTER TABLE missions ADD CONSTRAINT missions_emergencies
    FOREIGN KEY (emergency_id)
    REFERENCES emergencies (id)  ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: missions_missions_statuses (table: missions)
ALTER TABLE missions ADD CONSTRAINT missions_missions_statuses
    FOREIGN KEY (mission_status_id)
    REFERENCES missions_statuses (id)  ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: missions_users (table: missions)
ALTER TABLE missions ADD CONSTRAINT missions_users
    FOREIGN KEY (user_id)
    REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE 
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: problems_missions (table: problems)
ALTER TABLE problems ADD CONSTRAINT problems_missions
    FOREIGN KEY (mission_id)
    REFERENCES missions (id)  ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: problems_users (table: problems)
ALTER TABLE problems ADD CONSTRAINT problems_users
    FOREIGN KEY (user_id)
    REFERENCES users (id)  ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: users_users_type (table: users)
ALTER TABLE users ADD CONSTRAINT users_users_type
    FOREIGN KEY (user_type_id)
    REFERENCES users_types (id) ON UPDATE CASCADE ON DELETE CASCADE 
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: volunteers_abilities_abilities (table: volunteers_abilities)
ALTER TABLE volunteers_abilities ADD CONSTRAINT volunteers_abilities_abilities
    FOREIGN KEY (ability_id)
    REFERENCES abilities (id)  ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: volunteers_abilities_volunteers (table: volunteers_abilities)
ALTER TABLE volunteers_abilities ADD CONSTRAINT volunteers_abilities_volunteers
    FOREIGN KEY (volunteer_id)
    REFERENCES volunteers (id)  ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: volunteers_users (table: volunteers)
ALTER TABLE volunteers ADD CONSTRAINT volunteers_users
    FOREIGN KEY (user_id)
    REFERENCES users (id)  ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: volunteers_volunteers_statuses (table: volunteers)
ALTER TABLE volunteers ADD CONSTRAINT volunteers_volunteers_statuses
    FOREIGN KEY (volunteer_status_id)
    REFERENCES volunteers_statuses (id)  ON UPDATE CASCADE ON DELETE CASCADE
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- End of file.

