CREATE TABLE "public"."order_detail" ( 
  "id" VARCHAR(100) NOT NULL,
  "order_id" VARCHAR(100) NOT NULL,
  "service_id" VARCHAR(100) NOT NULL,
  "quantity" INTEGER NULL,
  CONSTRAINT "order_detail_pkey" PRIMARY KEY ("id")
);

CREATE TABLE "public"."orders" ( 
  "id" VARCHAR(100) NOT NULL,
  "date_received" DATE NOT NULL,
  "date_finished" DATE NOT NULL,
  "customer_id" VARCHAR(100) NOT NULL,
  "receiver" VARCHAR(250) NOT NULL,
  CONSTRAINT "order_pkey" PRIMARY KEY ("id")
);

CREATE TABLE "public"."mst_customer" ( 
  "id" VARCHAR(100) NOT NULL,
  "name" VARCHAR(100) NOT NULL,
  "phone" VARCHAR(15) NULL,
  "active_member" BOOLEAN NOT NULL DEFAULT true ,
  "join_date" DATE NOT NULL,
  "gender" VARCHAR(1) NOT NULL,
  CONSTRAINT "mst_customer_pkey" PRIMARY KEY ("id")
);

CREATE TABLE "public"."mst_service" ( 
  "id" VARCHAR(100) NOT NULL,
  "service_name" VARCHAR(50) NULL,
  "satuan" VARCHAR(20) NULL,
  "price" INTEGER NULL,
  CONSTRAINT "mst_service_pkey" PRIMARY KEY ("id")
);

ALTER TABLE "public"."order_detail" ADD CONSTRAINT "fk_order_detail_service_id" FOREIGN KEY ("service_id") REFERENCES "public"."mst_service" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."order_detail" ADD CONSTRAINT "fk_order_detail_order_id" FOREIGN KEY ("order_id") REFERENCES "public"."orders" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."orders" ADD CONSTRAINT "fk_order_customer_id" FOREIGN KEY ("customer_id") REFERENCES "public"."mst_customer" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
