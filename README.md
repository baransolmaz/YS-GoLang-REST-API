# YS-GoLang-REST-API

A REST-API service that works as an in memory key-value store  
[API Doc Link](https://documenter.getpostman.com/view/19679607/UVknsb2t){:target="_blank"}  
Requests:  
&emsp; -GET  
&emsp;  &emsp; Takes a key string and returns the value of the key.If the key is exist, responds 200;if not responds 404  
&emsp; -PUT  
&emsp;  &emsp; Takes a key-value string pair and updates the value of the given key.If the key is not exist, creates it  
&emsp; -DELETE  
&emsp;  &emsp; Removes all pairs  
&emsp; -VIEW  
&emsp;  &emsp; Sends all pairs  
  
Respond Codes:  
&emsp; 500 :&emsp;StatusInternalServerError  
&emsp; 502 :&emsp;StatusBadGateway  
&emsp; 404 :&emsp;StatusNotFound  
&emsp; 204 :&emsp;StatusNoContent  
&emsp; 405 :&emsp;StatusMethodNotAllowed  
&emsp; 200 :&emsp;StatusOK  
  
Note:  
&emsp; If you want the put request does not create new pair when the given key is not exist,you need to toggle block comment.  
