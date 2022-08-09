USE NotebookStore;
CREATE TABLE `NotebookInventory` (
  `Printno` bigserial PRIMARY KEY,
  `Model` varchar(255),
  `Manufacturer` varchar(255),
  `Width` decimal,
  `Binding` varchar(255)
);

CREATE TABLE `Transactions` (
  `TID` bigserial PRIMARY KEY,
  `Method` ENUM ('CASH_ON_DELIVERY', 'CREDIT', 'BANK_TRANSFER', 'INTERNET_PURCHASE') NOT NULL,
  `Amount` decimal NOT NULL,
  `By` varchar(255),
  `To` varchar(255),
  `ProductID` bigserial
);

CREATE TABLE `Users` (
  `Username` varchar(255) PRIMARY KEY,
  `Password` varchar(255) NOT NULL,
  `Balance` decimal
);

CREATE TABLE `Orders` (
  `OID` bigserial PRIMARY KEY,
  `Username` varchar(255),
  `TID` bigserial COMMENT 'An order is a transaction only after payment',
  `Status` ENUM ('DELIVERED', 'TRANSIT', 'CANCELLED', 'UNCONFIRMED'),
  `DeliveryTime` datetime
);

CREATE INDEX `Users_index_0` ON `Users` (`Username`);

CREATE INDEX `Orders_index_1` ON `Orders` (`OID`);

CREATE INDEX `Orders_index_2` ON `Orders` (`OID`, `Username`);

CREATE TABLE `NotebookInventory_Transactions` (
  `NotebookInventory_Printno` bigserial NOT NULL,
  `Transactions_ProductID` bigserial NOT NULL,
  PRIMARY KEY (`NotebookInventory_Printno`, `Transactions_ProductID`)
);

ALTER TABLE `NotebookInventory_Transactions` ADD FOREIGN KEY (`NotebookInventory_Printno`) REFERENCES `NotebookInventory` (`Printno`);

ALTER TABLE `NotebookInventory_Transactions` ADD FOREIGN KEY (`Transactions_ProductID`) REFERENCES `Transactions` (`ProductID`);


CREATE TABLE `Users_Orders` (
  `Users_Username` varchar NOT NULL,
  `Orders_Username` varchar NOT NULL,
  PRIMARY KEY (`Users_Username`, `Orders_Username`)
);

ALTER TABLE `Users_Orders` ADD FOREIGN KEY (`Users_Username`) REFERENCES `Users` (`Username`);

ALTER TABLE `Users_Orders` ADD FOREIGN KEY (`Orders_Username`) REFERENCES `Orders` (`Username`);


ALTER TABLE `Orders` ADD FOREIGN KEY (`TID`) REFERENCES `Transactions` (`TID`);
