INSERT INTO users_types  (type)
values
('Administrador'),
('Coordinador'),
('Voluntario');

INSERT INTO users (first_name,last_name,birthday,password,email,contact_phone_number,emergency_phone_number,life_insurance,enabled,user_type_id) 
values
('Carolina','Bonacic','1990-10-12','12345','carolinab@usach.cl',12345678,12345678,false,true,2);

INSERT INTO volunteers_statuses (status)
values
('Inactivo'),
('Activo');

INSERT INTO abilities (ability)
values
('Fuerza'),
('Primeros auxilios'),
('Construccion'),
('Electricidad');

INSERT INTO emergencies_statuses (status)
values
('Iniciada'),
('Archivada');

INSERT INTO emergencies_types (type)
values
('Terremoto'),
('Inundacion'),
('Incendio'),
('Aluvion');

INSERT INTO missions_statuses (status)
values
('Iniciada'),
('Finalizada'),
('Archivada');

INSERT INTO history_missions_states (state)
values
('Invitado'),
('Acepto'),
('Inicio'),
('Termino'),
('Rechazo');