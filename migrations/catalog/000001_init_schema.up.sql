CREATE TABLE brands
(
    id    UUID PRIMARY KEY,
    title VARCHAR(200) NOT NULL
);

CREATE TABLE categories
(
    id    UUID PRIMARY KEY,
    title VARCHAR(200) NOT NULL
);

CREATE TABLE catalog_items
(
    id                UUID PRIMARY KEY,
    title             VARCHAR(200)     NOT NULL,
    short_description TEXT,
    full_description  TEXT,
    image_url         TEXT,
    brand_id          UUID REFERENCES brands (id),
    category_id       UUID REFERENCES categories (id),
    price             DOUBLE PRECISION NOT NULL DEFAULT 0
);

INSERT INTO brands (id, title)
VALUES ('b0000001-0000-0000-0000-000000000001',
        'Test brand');

INSERT INTO categories (id, title)
VALUES ('c0000001-0000-0000-0000-000000000001',
        'Test category');

INSERT INTO catalog_items
(id,
 title,
 short_description,
 full_description,
 image_url,
 brand_id,
 category_id,
 price)
VALUES ('d0000001-0000-0000-0000-000000000001',
        'Test item',
        'Test item 1',
        'Test item with test brand and test category',
        'https:link',
        'b0000001-0000-0000-0000-000000000001',
        'c0000001-0000-0000-0000-000000000001',
        '100');
