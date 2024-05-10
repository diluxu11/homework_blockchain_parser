Due to time limitation, the system is not yet ready for deployment and still needs testing phase.

The following is the brief description of my system design:

1. Timer for Block Scanning:
By comparing the current scanned block number with the block number obtained from RPC, the system determines if there are any new transaction information. Once new transaction information is detected during the scanning process, it is updated in the local storage of the system. In the existing code, I use local memory to mimic the read/write operation of Redis.

2. Implementation of Parser:
This feature enables the ability to query transactions for subscribed addresses. The parser module parses the relevant data and provides functionality to retrieve transaction details for the subscribed addresses.
