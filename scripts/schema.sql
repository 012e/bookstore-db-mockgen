-- Bang cac nhan vien
create table employees
(
    id     integer primary key generated by default as identity,
    -- Ho va ten dem
    first_name      varchar(255) not null,
    -- Ten
    last_name       varchar(255) not null,
    -- Email
    email           varchar(255) not null,
    -- luong
    salary          money        not null check (salary > 0::money),
    profile_picture bytea        not null,
    -- Co phai quan ly khong
    is_manager      boolean      not null default false
);

-- Bang cac nha cung cap
create table providers
(
    id integer primary key generated by default as identity,
    -- Ten
    name        varchar(255) not null,
    -- Dia chi
    address     varchar(255) not null
);

-- Cac vat pham
create table items
(
    id     integer primary key generated by default as identity,
    name        varchar(255) not null,
    image       bytea        not null,
    -- So luong
    quantity    integer      not null check (quantity > 0),
    -- Thong tin chi tiet
    description text,                 -- text la nvarchar khong co gioi han do dai
    price       money        not null, -- Don gia
    provider_id integer references providers(id) on delete cascade
);

create table customers
(
    id     integer primary key generated by default as identity,
    first_name      varchar(255) not null,
    last_name       varchar(255) not null,
    phone_number    varchar(20) not null,
    email           varchar(255) not null
);

-- Hoa don
create table invoices
(
    id  integer primary key generated by default as identity,
    -- Tong tri gia hoa don
    total       money                               not null check (total > 0::money),
    -- Nhan vien xu ly hoa don
    employee_id integer                             not null references employees (id) on delete cascade,
    -- Thoi gian lap hoa don (khong can chen nua, thoi gian nay duoc tu dong chen)
    customer_id integer                             not null references customers(id) on delete cascade,
    created_at  timestamp default current_timestamp not null
);

-- Danh sach cac vat pham trong hoa don
create table invoices_items
(
    invoice_id integer not null references invoices (id) on delete cascade,
    item_id    integer not null references items (id) on delete cascade,
    quantity   integer                          not null check (quantity > 0),
    primary key (invoice_id, item_id)
);

create table tags
(
    id          integer primary key generated by default as identity,
    name        varchar(255) not null,
    description text
);

create table items_tags
(
    item_id  integer references items (id) on delete cascade,
    tag_id   integer references tags (id) on delete cascade,
    primary key (item_id, tag_id)
);
