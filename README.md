# Test assignment

---

## Service for track items in every warehouse

### For start the service use `docker-compose up`
### When server start use `make m-up` for up migrations
### Now you can use service
### Request examples from client in `internal/app/client/client.go`

## API Description

* ### UploadProduct
    *  **input data:** `name, size, code, quantity, warehouseId`
    *  **output:** `responseMsg`

  ### request: 
  ![img.png](img/img.png)
  ### response: `{success}`
  ![img_3.png](img/img_3.png)

* ### UpdateProduct
  *  **input data:** `name, size, code, quantity, warehouseId`
  *  ***output:*** `reponseMsg` 

  ### request:
  ![img_1.png](img/img_1.png)
  ### response: `{success}`
  ![img_4.png](img/img_4.png)

* ### GetProduct
  *  **input data:** `productCode, warehouseId`
  *  **output:** `product info` 
  
  ### request:
  ![img_5.png](img/img_5.png)
  ### response: `{b668e4bf-adb6-4aec-89be-169a8302a364 T-Shirt M 73271291 45 ce6ba671-e63c-4bd9-99af-f2f0cd110498}`

* ### ListProduct
  *  **input data:** `warehouseId`
  *  **output:** `[]product info`

  ### request:
  ![img_6.png](img/img_6.png)
  ### response: `[{b668e4bf-adb6-4aec-89be-169a8302a364 T-Shirt M 73271291 45 ce6ba671-e63c-4bd9-99af-f2f0cd110498} {64800149-ccf7-402c-8e57-969ad67d0bd6 Pants M 73278291 31 ce6ba671-e63c-4bd9-99af-f2f0cd110498}]`

* ### DeleteProduct
  *  **input data:** `productCode, warehouseId`
  *  **output:** `responseMsg`
  
  ### request:
  ![img.png](img/img7.png)
  ### response: `{success}`
  ![img_1.png](img/img_8.png)

* ### ReserveProduct
  *  **input data:** `[]productCode, warehouseId, quantity`
  *  **output:** `responseMsg`
  
  ### request:
  ![img.png](img/img8.png)
  ### response: `{success}`
  ![img_1.png](img/img_9.png)
  

* ### ReleaseReserve
  *  **input data:** `[]productCode, warehouseId`
  *  **output:** `responseMsg`

  ### request:
  ![img_2.png](img/img_10.png)
  ### response: `{success}`
  ![img_3.png](img/img11.png)
