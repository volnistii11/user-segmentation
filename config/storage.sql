BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS Segment
(
    id   serial primary key,
    slug varchar(255) not null unique
);
CREATE TABLE IF NOT EXISTS UsersSegments
(
    userId    integer,
    segmentId integer not null
);
COMMIT;