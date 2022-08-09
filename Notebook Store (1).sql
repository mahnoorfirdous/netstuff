CREATE TABLE [NotebookInventory] (
  [Printno] bigserial PRIMARY KEY,
  [Model] nvarchar(255),
  [Manufacturer] nvarchar(255),
  [Width] decimal,
  [Binding] nvarchar(255)
)
GO

CREATE TABLE [Transactions] (
  [TID] bigserial PRIMARY KEY,
  [Method] nvarchar(255) NOT NULL CHECK ([Method] IN ('CASH_ON_DELIVERY', 'CREDIT', 'BANK_TRANSFER', 'INTERNET_PURCHASE')) NOT NULL,
  [Amount] decimal NOT NULL,
  [By] nvarchar(255),
  [To] nvarchar(255),
  [ProductID] bigserial
)
GO

CREATE TABLE [Users] (
  [Username] nvarchar(255) PRIMARY KEY,
  [Password] nvarchar(255) NOT NULL,
  [Balance] decimal
)
GO

CREATE TABLE [Orders] (
  [OID] bigserial PRIMARY KEY,
  [Username] nvarchar(255),
  [TID] bigserial,
  [Status] nvarchar(255) NOT NULL CHECK ([Status] IN ('DELIVERED', 'TRANSIT', 'CANCELLED', 'UNCONFIRMED')),
  [DeliveryTime] datetime
)
GO

CREATE INDEX [Users_index_0] ON [Users] ("Username")
GO

CREATE INDEX [Orders_index_1] ON [Orders] ("OID")
GO

CREATE INDEX [Orders_index_2] ON [Orders] ("OID", "Username")
GO

EXEC sp_addextendedproperty
@name = N'Column_Description',
@value = 'An order is a transaction only after payment',
@level0type = N'Schema', @level0name = 'dbo',
@level1type = N'Table',  @level1name = 'Orders',
@level2type = N'Column', @level2name = 'TID';
GO

CREATE TABLE [NotebookInventory_Transactions] (
  [NotebookInventory_Printno] bigserial NOT NULL,
  [Transactions_ProductID] bigserial NOT NULL,
  PRIMARY KEY ([NotebookInventory_Printno], [Transactions_ProductID])
);
GO

ALTER TABLE [NotebookInventory_Transactions] ADD FOREIGN KEY ([NotebookInventory_Printno]) REFERENCES [NotebookInventory] ([Printno]);
GO

ALTER TABLE [NotebookInventory_Transactions] ADD FOREIGN KEY ([Transactions_ProductID]) REFERENCES [Transactions] ([ProductID]);
GO


CREATE TABLE [Users_Orders] (
  [Users_Username] varchar NOT NULL,
  [Orders_Username] varchar NOT NULL,
  PRIMARY KEY ([Users_Username], [Orders_Username])
);
GO

ALTER TABLE [Users_Orders] ADD FOREIGN KEY ([Users_Username]) REFERENCES [Users] ([Username]);
GO

ALTER TABLE [Users_Orders] ADD FOREIGN KEY ([Orders_Username]) REFERENCES [Orders] ([Username]);
GO


ALTER TABLE [Orders] ADD FOREIGN KEY ([TID]) REFERENCES [Transactions] ([TID])
GO
