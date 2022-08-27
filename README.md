# mmapipc
This is a experiment to see how memory mapping files and inter-process communications behaves

### Notes
- `ipcs -m` list out all interprocess communications with shared memory segments
- `ipcs -mp` list out all interprocess communications with shared memory segments with PID
- `ipcrm -m <shmid>` remove specic shared memory segment
- `ipcrm -a` remove all shared memory segments (make sure that the programs that use these segments can deal with them being removed)
- `ps <PID>` info on process ID
