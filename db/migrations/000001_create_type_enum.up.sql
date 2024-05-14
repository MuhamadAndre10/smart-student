BEGIN;

-- Create type Enum Gender
CREATE TYPE gender_type AS ENUM (
    'LAKI-LAKI',
    'PEREMPUAN');


COMMIT;