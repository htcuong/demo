---- drop ----
DROP TABLE IF EXISTS `buys`;
DROP TABLE IF EXISTS `wagers`;

---- create ----
create table IF not exists `wagers`
(
    id INT auto_increment NOT NULL,
	total_wager_value INT NOT NULL,
	odds INT NOT NULL,
	selling_percentage INT NOT NULL,
	selling_price DOUBLE NOT NULL,
	current_selling_price DOUBLE NOT NULL,
	percentage_sold INT NULL,
	amount_sold INT NULL,
    placed_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

create table IF not exists `buys`
(
    id INT auto_increment NOT NULL,
	wager_id INT NOT NULL,
	buying_price DOUBLE NOT NULL,
	bought_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (wager_id) REFERENCES wagers(id)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

---- stored-procedure ----

DROP PROCEDURE IF EXISTS `buy_wager`;

DELIMITER $$
CREATE PROCEDURE buy_wager (
	IN $wager_id INT,
	IN $buying_price DOUBLE,
	IN $bought_at TIMESTAMP
)
BEGIN
	START TRANSACTION;

	INSERT INTO buys (wager_id,buying_price,bought_at) 
	SELECT $wager_id,$buying_price,$bought_at
	FROM wagers
	WHERE 
		wagers.current_selling_price > $buying_price
		AND id = $wager_id;
	
	IF FOUND_ROWS() > 0 THEN
		UPDATE wagers SET
			current_selling_price = (selling_price - COALESCE(amount_sold,0) - $buying_price), 
			percentage_sold = ((COALESCE(amount_sold,0) + $buying_price)*100/selling_price), 
			amount_sold = (COALESCE(amount_sold,0) + $buying_price)
		WHERE id = $wager_id;
		COMMIT;
		SELECT * FROM buys WHERE id = LAST_INSERT_ID();
	ELSE
		ROLLBACK;
		SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'Can not insert data into buys table';
	END IF;
END$$
DELIMITER ;