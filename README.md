
# epayments-restapi

## Task
Task: Implement a Rest API for a financial institution where it provides e-wallet services to its partners.  It has two types of e-wallet accounts: identified and non-identified.  The API can support multiple clients and only http, post methods with json as data format should be used.  Clients must be authenticated via the http parameter header X-UserId and X-Digest.  X-Digest is hmac-sha1, the hash sum of the request body.  There must be pre-recorded e-wallets, with different balances, and the maximum balance is TJS 10,000 for unidentified accounts and TJS 100,000 for identified accounts.  For data storage you can use according to your choice.  Service API methods: 1. Check if there is an e-wallet account.  2. Replenishment of the electronic wallet.  3. Get the total number and amounts of replenishment transactions for the current month.  4. Get the balance of the electronic wallet.



 
## Run Locally  
Clone the project  

~~~bash  
  git clone https://github.com/ProninIgor/epayments-restapi
~~~

Go to the project directory  


Start the server  

~~~bash  
./run.sh

~~~  

Request Examples
~~~bash
curl -X GET -H "X-UserId: user1" -H "X-Digest: f92a65136e9c7b7a31590f87b850c96f7c257dac" http://localhost:8000/wallet/exists
curl -X GET -H "X-UserId: user1" -H "X-Digest: f92a65136e9c7b7a31590f87b850c96f7c257dac" http://localhost:8000/wallet/balance
curl -X POST -H "X-UserId: user1" -H "X-Digest: f92a65136e9c7b7a31590f87b850c96f7c25ac" -d '{"amount": 100}' http://localhost:8000/wallet/replenish
curl -X GET -H "X-UserId: user1" -H "X-Digest: f92a65136e9c7b7a31590f87b850c96f7c257dac" http://localhost:8000/wallet/transactions
~~~