# Challenge F

## How would you solve such a problem?

we could in the transaction that SELECT the product first to see if it's in stock or not then UPDATE it put a FOR UPDATE in the end of the SELECT statement this will make the transaction hold an exclusive lock on the row from the beginning (from the SELECT) instead of shared lock then upgrading it to exclusive lock when reaching the UPDATE statement, by this after we read the row no other transaction can read or write in the same row until we commit or rollback so the other transaction while be waiting for us to finish and when we finish it will read the row to find the stock equal zero.

## Would you marginalize/add a solution if you knew that the used database in the system is MongoDB?

we can solve it with optimistic concurrency control it won't prevent 2 nodes from approaching the same document at the same time, by we can know that some one else edited the document from the last time we read it, so we will add a version and with every read we get it, and when we write we increment it, so when the first node look at the stock it get the version and in_stock = 1 , so if the second node update the in_stock, the first node when update it will put a condition that do the update if the version still equal the old one. and if it fails we do the whole operation from the beginning. 