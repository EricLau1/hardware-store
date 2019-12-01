drop database supertest;

create database supertest
default character set utf8
default collate utf8_general_ci;
use supertest;

create table if not exists categories(
id bigint primary key auto_increment,
description varchar(256) not null unique,
created_at timestamp default current_timestamp,
updated_at timestamp default current_timestamp
)
engine = InnoDB
default charset = utf8;

create table if not exists products(
id bigint primary key auto_increment,
name varchar(512) not null unique,
price decimal(10,2) default 0.0,
quantity int(10) unsigned default 0,
status char(1) default 0,
category_id bigint not null,
created_at timestamp default current_timestamp,
updated_at timestamp default current_timestamp,
constraint products_category_id foreign key(category_id) references categories(id)
on delete cascade on update cascade
)
engine = InnoDB
default charset = utf8;

insert into categories (description) values 
('Placa Mãe'),
('Processador'),
('Placa de Vídeo'),
('SSD'),
('Memória RAM');

insert into products (name, price, quantity, status, category_id) values
('Placa Mãe Colorful Battle-AX Z390AK Gaming V20, Chipset Z390, Intel LGA 1151, ATX, DDR4', 860.43, 50, 1, 1),
('Placa-Mãe Gigabyte GA-AB350M-DS3H V2, AMD AM4, mATX, DDR4', 339.90, 20, 1, 1),
('Placa-Mãe Asus Prime B450M Gaming/BR, AMD AM4, mATX, DDR4',419.90, 35, 1, 1),
('Placa-Mãe ASRock A320M-HD, AMD AM4, mATX, DDR4', 299.90, 0, 0, 1),
('Placa-Mãe Gigabyte X570 Aorus Xtreme, AM4, eATX, DDR4', 4799.90, 2000, 1, 1),
('Placa-mãe MSI MEG X570 Godlike, AMD AM4, E-ATX, DDR4', 4599.90, 250, 1, 1),
('Placa-Mãe Gigabyte Z390 Aorus Xtreme, Intel LGA 1151, eATX, DDR4 (Rev. 1.0)', 4325.90, 10, 1, 1),
('Placa-mãe Gigabyte X570 Aorus Elite, AMD AM4, ATX, DDR4', 1659.90, 0, 0, 1);

insert into products (name, price, quantity, status, category_id) values
('Processador Intel Core i5 9400F 2.90GHz (4.10GHz Turbo), 9ª Geração, 6-Core 6-Thread, LGA 1151, BX80684I59400F', 829.00, 80, 1,2),
('Processador AMD Ryzen 5 3600 3.6GHz (4.2GHz Turbo), 6-Core 12-Thread, Cooler Wraith Stealth, AM4, 100-100000031BOX, S/ Video', 979.00, 50, 1, 2),
('Processador AMD Ryzen 9 3900x 3.8ghz (4.6ghz Turbo), 12-core 24-thread, Wraith Prism RGB, AM4, S/ Video', 3299.00, 3, 1, 2),
('Processador AMD Athlon 240GE 3.5GHz, Dual Core 5MB AM4, YD240GC6FBBOX', 369.00, 1, 1, 2),
('Processador AMD Ryzen 5 2400G 3.6GHz (3.9GHz Turbo), 4-Core 8-Thread, Cooler Wraith Stealth, AM4, YD2400C5FBBOX', 749.90, 20, 1, 2),
('Processador Intel Core i7 9700 3.0GHz (4.70GHz Turbo), 9ª Geração, 8-Core 8-Thread, LGA 1151, BX80684I79700', 2174.13, 10, 1, 2),
('Processador Intel Core i3 8100 3.6GHz, 8ª Geração, 4-Core 4-Thread, LGA 1151, BX80684I38100', 689.00, 30, 1, 2),
('Processador Intel Pentium Gold G5400 3.7GHz 4MB BX80684G5400 8ª GERAÇÃO Coffee Lake LGA 1151', 347.13, 25, 1, 2),
('Processador Intel Core i7 8700K 3.70GHz (4.7GHz Turbo), 8ª Geração, 6-Core 12-Thread, LGA 1151, BX80684I78700K', 2295.93, 0, 0, 2);

insert into products (name, price, quantity, status, category_id) values
('Placa de Video Gigabyte GeForce RTX 2080 Ti Aorus, 11GB GDDR6, 352Bit, GV-N208TAORUS-11GC', 6459.00, 39, 1, 3),
('Placa de Vídeo Galax GeForce GTX 1660 (1-Click OC) Dual, 6GB GDDR5, 192Bit, 60SRH7DSY91C', 1079.00, 5, 1, 3),
('Placa de Vídeo Galax Geforce RTX 2060 Ex White Dual, 6GB GDDR6, 192Bit, 26NRL7HPY3EW', 1692.15, 11, 1, 3),
('Placa de Vídeo Asus GeForce RTX 2060 Rog Strix Gaming Oc Edition, 6GB GDDR6, 192Bit, ROG-STRIX-RTX2060-O6G-GAMING', 1965.33, 0, 0, 3),
('Placa de Vídeo Galax GeForce GTX 1050 Ti EXOC Dual, 4GB GDDR5, 128Bit, 50IQH8DVN6EC', 649.00, 0, 0, 3),
('Placa de Vídeo Galax GeForce RTX 2070 Super EX Gamer Black, 8GB GDDR6, 256Bit, 27ISL6MDW0BG', 3018.03, 500, 1, 3),
('Placa de Vídeo MSI AMD Radeon RX 570 Armor 4G OC, GDDR5', 649.90, 220, 1, 3),
('Placa de Vídeo Gigabyte AMD Radeon RX 590 Gaming 8G, GDDR5 - GV-RX590GAMING-8GD', 919.90, 75, 1, 3),
('Placa de Vídeo XFX AMD Radeon RX 580, 8GB, GDDR5 - RX-580P828D6', 982.24, 130, 1, 3),
('Placa de Vídeo PCYes AMD Radeon HD 6570, 2GB, DDR3 - PJ657012802D3', 249.90, 0, 0, 3),
('Placa de Vídeo Gigabyte AMD Radeon RX 5700 8G, GDDR6 - GV-R57-8GD-B', 1829.90, 90, 1, 3);

insert into products (name, price, quantity, status, category_id) values
('SSD Adata SU650 120GB, Sata III, Leitura 520MBs e Gravação 450MBs, ASU650SS-120GT-R', 119.0, 180, 1, 4),
('SSD Kingston UV500 120GB, M.2 2280, Leitura 520MBs e Gravação 320MBs, SUV500M8/120G', 179.0, 240, 1, 4),
('SSD WD Black SN750 250GB, M.2 2280, Leitura 3100MBs e Gravação 1600MBs, WDS250G3X0C', 439.0, 300, 1, 4),
('SSD WD Blue 1TB, M.2 2280, Leitura 560MBs e Gravação 530MBs, WDS100T2B0B', 799.90, 50, 1, 4),
('SSD WD Green 1TB, Sata III, Leitura 545MB/S e Gravação 430MB/s, WDS100T2G0A', 639.0, 120, 1, 4),
('SSD WD Blue 250GB, Sata III, Leitura 550MBs e Gravação 525MBs, WDS250G2B0A', 309.20, 0, 0, 4),
('SSD Corsair MP600 2TB, M.2 2280, Leitura 4.950MBs e Gravação 4.250MBs, CSSD-F2000GBMP600', 2565.63, 70, 1, 4);

insert into products (name, price, quantity, status, category_id) values
('Memória DDR4 Corsair Vengeance RGB Pro, 16GB (2x8GB) 3600MHz, CMW16GX4M2D3600C18', 695.13, 10, 1, 5),
('Memória DDR4 Corsair Vengeance RGB Pro, 16GB (2x8GB) 3200MHz, White, CMW16GX4M2C3200C16W', 616.83, 200, 1, 5),
('Memória DDR4 Geil EVO X II RGB SYNC 16GB (2x8GB) 3600mhz, GAEXSY416GB3600C18BDC', 608.13, 600, 1, 5),
('Memória DDR4 G.Skill Trident Z RGB AMD, 16GB (2x8GB) 3600MHz, F4-3600C18D-16GTZRX', 782.13, 400, 1, 5),
('Memória DDR4 Kingston HyperX Fury RGB, 8GB 3200MHz, Black, HX432C16FB3A/8', 338.43, 900, 1, 5),
('Memória DDR3 Corsair Vengeance, SDP, 4GB 1600MHz, CMZ4GX3M1A1600C9', 199.88, 0, 0, 5),
('Memória DDR4 G.Skill Ripjaws V, 8GB 2800MHz, F4-2800C17S-8GVR', 249.00, 0, 0, 5),
('Memória DDR3 Kingston HyperX Fury, 8GB 1866MHz, Blue, HX318C10F/8', 1080.10, 3, 1, 5);