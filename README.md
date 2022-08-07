# Little Brown Book Shop

## `Took time 22 hours with beginner knowledge and my endeavor.`

## Personal Information

Please provide all information in English.

|  |  |
| --- | --- |
| **⚠️ First Name:** | Chawalit |
| **⚠️ Last Name:** | Janinta |
| **⚠️ Email:** | chawalit.janinta@gmail.com |
| **⚠️ Phone Number:** | 0811002132 |

## Overview

You were assigned to implement a POS(Point of Sale) system for our client, who owns a shop name “Little Brown Book Shop.” They need a precise system, easy to use, and also an elegant user interface.

Their requirements are simple. Firstly, a staff stands at the POS, and the first page they see is the item-list page. If customers bring books from the shelf and come to checkout, the staff just has to add what item in the customer’s basket, And the system will calculate the total price and offer some discounts if it meets the criteria. Luckily, there is no tax/vat here. When all customer items were added, the system is ready to proceed to the next step, payment. At Little Brown, they only accept cash (paying by credit card is still on their mind). In this way, the staff has to fill the amount of cash into the system and let it process the change back to the customers. Don’t forget to show staff that everything is alright through the payment process. Here, you have to show customers a screen that their payment was successful (instead of using a paper receipt, save the world!). After everything is done, the system must be ready for the next order again.

As we mention about the discount, there is a special discount only for all Harry Potter book series (7 books);

* buy 2 unique series books discount 10% of those 2 books
* buy 3 unique series books discount 11% of those 3 books
* buy 4 unique series books discount 12% of those 4 books
* buy 5 unique series books discount 13% of those 5 books
* buy 6 unique series books discount 14% of those 6 books
* buy 7 unique series books discount 15% of those 7 books

You, as our best back-end developer, have to develop all of these API parts within 1 week. API spec has been provided by our API team (see document below).

So, we hope you give them everything essential for this task to get impressed from our client. And don’t forget, correctness is a MUST.

Also, we have prepared `docker-compose.yml` file for you, which means when you submit, we will use this file to run your application.

## API Spec

### API `GET` Books

**URI:** `/api/books`

**Response:**
```json
[{
  "id": 1,
  "title": "book1",
  "price": 200,
  "description": "description here"
  "rating": 1,
  "author": "J.K. Rolling",
  "images": {
    "jpeg": "https://vos.line-scdn.net/assignment/book-shop/jpeg/harry-1.jpeg",
    "avif": "https://vos.line-scdn.net/assignment/book-shop/avif/harry-1.avif",
    "webp": "https://vos.line-scdn.net/assignment/book-shop/webp/harry-1.webp"
  }
},
...
]
```

### API `GET` Search Books

**URI:** `/api/search?query=Harry%20Potter%20I`

**Response:**
```json
[{
    "title": "Brown Potter I",
    "price": 200,
    "rating": 5,
    "author": "Brown",
    "images": [...]
  },
  ...
]
```

### API `GET` Book Promotion

**URI:** `/api/promotions`

**Response:**
```json
[
  {
    "id": 9001,
    "type": "MULTI_HARRY",
    "name": "Buy multiple harry potter volume",
    "targetIds": [10001, 10002, 10003, 10004, 10005, 10006, 10007],
    "discounts": ["10.00", "15.00", "20.00", "25.00", "30.00", "70.00"]
  },
  {
    "id": 9002,
    "type": "3_FREE_1",
    "name": "Buy 3 free 1",
    "targetIds": [10008, 10009, 10010, 10011, 10012, 10013, 10014]
  }
]
```

### API `POST` Add to cart

**URI:** `/api/orders`

**Payload:**
```json
{
  "id": 1,
  "quantity": 1
}
```

**Response:**

Success Case
```json
{
  "status": "success",
}
```

Failed Case
```json
{
 "status": "failed",
  "message": "Please try again later"
}
```

### API `GET` Order Summary

**URI:** `/api/orders/summary`

**Notes:**
1. id is order id
2. It should return proper discount based on promotions

**Response:**
```json
{
  "id": 1,
  "products": [{
    "name": "Harry Potter I",
    "price": 200,
    "rating": 5,
    "author": "Brown",
    "coverImage": "https://www.line.me"
  },
  ...
  ],
  "discount": 40,
  "summary": 560
}
```

### API `POST` Order Checkout

**URI:** `/api/orders/checkout`

**Notes:** id is order id

**Payload:**
```json
{
  "id": 1
}
```

**Response:**

Success Case
```json
{
 "status": "success",
}
```

Failed Case
```json
{
 "status": "failed",
  "message": "Please try again later"
}
```

## How to start service
`docker compose up` to run backend service locally. The service is able to do hot reload. Happy Coding!
