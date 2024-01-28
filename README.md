# dirWatcher

#Schema Diagram

+------------------------+
|   FileMonitoring       |
+------------------------+
| Id           : int     |
| RunTimestamp : time    |
| FileName     : string  |
| MagicString  : int     |
| NewFileAdded : bool    |
| FileDeleted  : bool    |
+------------------------+

#Execution instructions
1.Install latest Go 
2.Install latest postgresql
3.Import the code in separate directory
4.Make sure you have necessary go packages 
5.Finally hit "go run ." from the current working directory

#API Details

API 1: Run Background Task
Route: /api/run-background-task
Method: GET
Description: Initiates a background task to run a specific task.
o/p:"Background task executed successfully"


API 2: Fetch File Monitoring Details
Route: /api/file-monitoring/details
Method: POST
RequestPayload:{"file_name": "main.go"}
o/p:`[
    {
        "run_timestamp": "0001-01-01T00:00:00Z",
        "file_name": "main.go"
    },
    {
        "run_timestamp": "0001-01-01T00:00:00Z",
        "file_name": "main.go"
    }
]`

API 3: Fetch All File Names
Route: /api/file-monitoring/details/fetch-all
Method: POST
RequestPayload:{"fetch_all": true}
o/p: `[
    ".env",
    "Migrations.go",
    "cron.go",
    "database.go",
    "fileMonitoringDetails.go",
    "go.mod",
    "go.sum",
    "main.go",
    "model_file_monitoring.go",
    "server.go",
    ".env",
    "Migrations.go",
    "cron.go",
    "database.go",
    "fileMonitoringDetails.go",
    "go.mod",
    "go.sum",
    "main.go",
    "model_file_monitoring.go",
    "server.go"
]`
