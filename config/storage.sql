BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS Segment
(
    id   serial primary key,
    slug varchar(255) not null
);
CREATE TABLE IF NOT EXISTS UsersSegments
(
    userId    integer,
    segmentId integer
);
COMMIT;