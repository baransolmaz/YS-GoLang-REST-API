# YS-GoLang-REST-API
A REST-API service that works as an in memory key-value store

Requests:
-GET
    Takes a key string and returns the value of the key.If the key is exist, responds 200;if not responds 404 
-PUT
    Takes a key-value string pair and updates the value of the given key.If the key is not exist, creates it
-DELETE
    Removes all pairs
-VIEW
    Sends all pairs

Respond Codes:
    -StatusInternalServerError  : 500
    -StatusBadGateway           : 502
    -StatusNotFound             : 404
    -StatusNoContent            : 204
    -StatusMethodNotAllowed     : 405
    -StatusOK                   : 200

Note:
    If you want the put request does not create new pair when the given key is not exist,you need to toggle block comment. 