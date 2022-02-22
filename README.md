# YS-GoLang-REST-API
A REST-API service that works as an in memory key-value store  

Requests:  
&nbsp; -GET  
&nbsp; &nbsp;Takes a key string and returns the value of the key.If the key is exist, responds 200;if not responds 404  
&nbsp; -PUT  
&nbsp; &nbsp;Takes a key-value string pair and updates the value of the given key.If the key is not exist, creates it  
&nbsp; -DELETE  
&nbsp; &nbsp;Removes all pairs  
&nbsp; -VIEW  
&nbsp; &nbsp;Sends all pairs  
  
Respond Codes:  
&nbsp; StatusInternalServerError  : 500  
&nbsp; StatusBadGateway           : 502  
&nbsp; StatusNotFound             : 404  
&nbsp; StatusNoContent            : 204  
&nbsp; StatusMethodNotAllowed     : 405  
&nbsp; StatusOK                   : 200  
  
Note:  
&nbsp; If you want the put request does not create new pair when the given key is not exist,you need to toggle block comment.  