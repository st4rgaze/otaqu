## TASK
Buatlah API menggunakan Golang (GoFiber Framework) dengan ketentuan sebagai berikut :

1. Buatlah script untuk mendapatkan list Hotel dari link di bawah ini
http://115.85.80.33/test-scrapping/avail.html
2. Kemudian Parsing pada hasil scrapping di atas dan insert ke dalam table "hotel".
3. Buatlah satu endpoint untuk menampilkan list hotel tersebut dengan format response JSON.
4. Gunakan script sql di bawah ini untuk membuat table "hotel".

CREATE TABLE `hotel` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(100) NOT NULL COLLATE 'latin1_swedish_ci',
	`address` VARCHAR(250) NOT NULL COLLATE 'latin1_swedish_ci',
	`image_url` VARCHAR(500) NOT NULL COLLATE 'latin1_swedish_ci',
	`star_rating` INT(10) UNSIGNED NOT NULL,
	`price` INT(10) UNSIGNED NOT NULL,
	`created_at` DATETIME NULL DEFAULT NULL,
	`updated_at` DATETIME NULL DEFAULT NULL,
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='latin1_swedish_ci'
ENGINE=InnoDB
;


## How to run the app

1. copy the .env file
```
cp .env.example .env
```
2. Setup your database connection etc in .env file
3. Create database on your postgresql based on what you set on .env
4. you can run by:
- run the project by command line "go run main.go"
- or build the package by command line "go build" then run by command line "./otaqu"