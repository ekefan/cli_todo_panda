# cli_todo_panda:
A minimalist command line todo application

## Features:
- Create a task
- Read(get) a tasks
- Delete a task
- Receive help for using the application


## How it work in brief
panda writes to, and reads from a json file which acts a storage for the tasks.
The application reads commands from the cli and processes supported commands.
If an invalid entry is made the application returns a help information.
Now that's minimalistic.

### Supported Commands
| Commands | Use    |
| :---| :---                                                            |
| add | needs two args, the task as quoted and the priority (h or l, n) |
| task | need no arg, its is used to show the list of incompleted tasks |
| complete | need one arg - the id, and is used to complete a task[the id is the serial number of the tasks] |
| help | need no arg, it displays the help information on request and in situations where an supported command is inputed |