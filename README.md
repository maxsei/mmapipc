# mmapipc
This is a experiment to see how memory mapping files and inter-process communications behaves

### Notes
- `ipcs -m` list out all interprocess communications (ipc) with shared memory (shm) segments
- `ipcs -mp` list out all ipc with shm segments with PID
- `ipcrm -m <shmid>` remove specic shm segment
- `ipcrm -a` remove all shm segments (make sure that the programs that use these segments can deal with them being removed)
- `ipcs -m | awk '{if($6 == 0){print $2}}' | xargs -L 1 ipcrm -m` remove all shm segments that have no processes attached (in Makefile)
- `ps <PID>` info on process ID
