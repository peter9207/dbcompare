
create table warehouses (
  id serial PRIMARY KEY,
  name varchar
);


create table item_types (
  id serial PRIMARY KEY,
  name varchar
  );

create table items (
  id serial PRIMARY KEY,
  name varchar,
  item_type_id int,
  warehouse_id int,
  created_at timestamp,
  updated_at timestamp,
  FOREIGN KEY (item_type_id) references item_types(id),
  FOREIGN KEY (warehouse_id) references warehouses(id)
  );







