
# YS-GoLang-REST-API

A REST-API service that works as an in memory key-value store

  

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

&emsp; StatusInternalServerError &emsp; : 500

&emsp; StatusBadGateway &emsp; &emsp; &emsp; &emsp;: 502

&emsp; StatusNotFound &emsp; &emsp; &emsp;&emsp;&emsp;&emsp;: 404

&emsp; StatusNoContent &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;: 204

&emsp; StatusMethodNotAllowed &emsp;: 405

&emsp; StatusOK &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp; &emsp; &emsp;: 200

Note:

&emsp; If you want the put request does not create new pair when the given key is not exist,you need to toggle block comment.