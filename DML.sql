INSERT INTO "public"."mst_customer" ("id", "name", "phone", "active_member", "join_date", "gender") VALUES ('C001', 'Sari Ayunda', '081234567999', true, '2020-02-23', 'F');
INSERT INTO "public"."mst_customer" ("id", "name", "phone", "active_member", "join_date", "gender") VALUES ('C002', 'Bagas Putra', '082345778996', true, '2021-02-01', 'M');
INSERT INTO "public"."mst_customer" ("id", "name", "phone", "active_member", "join_date", "gender") VALUES ('C003', 'Nikola Tesla', '081234789555', true, '2022-12-12', 'M');
INSERT INTO "public"."mst_customer" ("id", "name", "phone", "active_member", "join_date", "gender") VALUES ('C004', 'Albert Einstein', '081234789010', true, '2023-11-11', 'M');
INSERT INTO "public"."mst_customer" ("id", "name", "phone", "active_member", "join_date", "gender") VALUES ('C005', 'Eminem Yeah', '081234555777', true, '2021-02-01', 'M');

INSERT INTO "public"."mst_service" ("id", "service_name", "satuan", "price") VALUES ('S002', 'Laundry Bedcover', 'Buah', 50000);
INSERT INTO "public"."mst_service" ("id", "service_name", "satuan", "price") VALUES ('S003', 'Laundry Boneka', 'Buah', 25000);
INSERT INTO "public"."mst_service" ("id", "service_name", "satuan", "price") VALUES ('S001', 'Cuci + Setrika', 'KG', 8000);

INSERT INTO "public"."orders" ("id", "date_received", "date_finished", "customer_id", "receiver") VALUES ('O001', '2022-08-18', '2022-08-20', 'C001', 'Mirna');

INSERT INTO "public"."order_detail" ("id", "order_id", "service_id", "quantity") VALUES ('D001', 'O001', 'S001', 6);


