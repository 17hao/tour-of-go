#include <errno.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>
#include <time.h>
#include <unistd.h>

int main() {
  char *cgroup = "/sys/fs/cgroup/system.slice/naive-container.service";
  struct stat statbuf;
  if (stat(cgroup, &statbuf) != 0) {
    if (errno == ENOENT) {
      if (mkdir(cgroup, 0700) != 0) {
        perror("mkdir");
        exit(EXIT_FAILURE);
      }
    } else {
      perror("stat");
      exit(EXIT_FAILURE);
    }
  }

  char memory[128];
  strcpy(memory, cgroup);
  strcat(memory, "/memory.max");
  int fd;
  if ((fd = open(memory, O_RDWR)) < 0) {
    perror("open");
    exit(EXIT_FAILURE);
  }
  if (write(fd, "3M", 3) < 0) {
    perror("write");
    exit(EXIT_FAILURE);
  }

  char procs[128];
  strcpy(procs, cgroup);
  strcat(procs, "/cgroup.procs");
  if ((fd = open(procs, O_RDWR)) < 0) {
    perror("open");
    exit(EXIT_FAILURE);
  }
  pid_t pid = getpid();
  char pidStr[64];
  sprintf(pidStr, "%d", pid);
  if (write(fd, pidStr, 64) < 0) {
    perror("write");
    exit(EXIT_FAILURE);
  }

  int total = 0;
  for (int i = 1;; i++) {
    int len = 1024 * 1000;
    char *chars;
    chars = malloc(len * sizeof(char));
    if (chars == NULL) {
      perror("malloc");
      exit(EXIT_FAILURE);
    }
    memset(chars, 0, len);
    total += len;
    printf("[%d] allocate %d-bytes, total=%d\n", i, len, total);
    struct timespec ts;
    ts.tv_sec = 0;
    ts.tv_nsec = 500 * 1000 * 1000;
    nanosleep(&ts, &ts);
  }
}