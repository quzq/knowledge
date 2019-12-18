create table my_table (
  id integer default nextval('steps_id_seq'::regclass) not null
  , name character varying(255) not null
  , age integer not null
  , primary key (id)
);

[{"name": "Mike", "age": 18}, {"name":"Jhon", "age":45},{"name":"Bob", age: 23}]



https://stackoverflow.com/questions/33129526/loading-json-data-from-a-file-into-postgres

BEGIN;
-- let's create a temp table to bulk data into
create temporary table temp_json (values text) on commit drop;
copy temp_json from '/home/sample.json';

-- uncomment the line above to insert records into your table
-- insert into tbl_staging_eventlog1 ("EId", "Category", "Mac", "Path", "ID") 

select values->>'name' as name,
       values->>'age' as age
from   (
           select json_array_elements(replace(values,'\','\\')::json) as values 
           from   temp_json
       ) a; 
COMMIT;