#include <sys/shm.h>
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>

void panic(char* s) {
	fprintf(stderr, s);
	fprintf(stderr, "\n");
	exit(1);
}

 int main(int argc, char *argv[]) {
	key_t key = EEXIST;
	/*
	for (int i; key == EEXIST; i++) {
		key = ftok("/", i);
		if (key == -1)
			panic("ftok");
	}
	*/

	// Sometimes you might want the same memory key.
	key = ftok("./ipc/hello", 0xdeadbeef);
	if (key == -1)
		panic("ftok");

	size_t shmget_size = 4000;
	int shmid = shmget(key, shmget_size, 0644 | IPC_CREAT);
	if (shmid == -1) {
		panic("shmget");
	}
	printf("shmid = %d\n", shmid);

	void *data;
	data = shmat(shmid, NULL, 0);
	if (data == (void*)(-1)) {
		panic("shmat");
	}

	if (shmdt(data) == -1)
		panic("shmdt"); 
}
