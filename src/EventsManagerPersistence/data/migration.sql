CREATE EXTENSION "uuid-ossp";


CREATE TABLE country (
    country_id uuid DEFAULT uuid_generate_v4(),
    code VARCHAR(20) NOT NULL,
    name VARCHAR(50) NOT NULL,
    primary key (country_id)
);

CREATE TABLE region (
    region_id uuid DEFAULT uuid_generate_v4(),
    code VARCHAR(20) NOT NULL,
    name VARCHAR(50) NOT NULL,  
    country_id uuid not null,
    primary key (region_id),
    constraint fk_region_country
     foreign key (country_id) 
     references country (country_id)
);


CREATE TABLE sensor (
    sensor_id uuid DEFAULT uuid_generate_v4(),
    code VARCHAR(20) NOT NULL,
    name VARCHAR(50) NOT NULL,  
    region_id uuid not null,
    primary key (sensor_id),
    constraint fk_sensor_region
     foreign key (region_id) 
     references region (region_id)
);


CREATE TABLE event_state (
    event_state_id uuid DEFAULT uuid_generate_v4(),
    code VARCHAR(20) NOT NULL,
    name VARCHAR(50) NOT NULL,
    primary key (event_state_id)
);


CREATE TABLE event (
    event_id uuid DEFAULT uuid_generate_v4(),
    sensor_id uuid not null,
    event_state_id uuid not null,
    value varchar(100),
    timestamp timestamptz not null,
    
    primary key (event_id), 
    constraint fk_event_sensor
     foreign key (sensor_id) 
     references sensor (sensor_id),
    
    constraint fk_event_event_state
     foreign key (event_state_id) 
     references event_state (event_state_id)
);
