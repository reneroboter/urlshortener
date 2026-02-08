# ToDos
- analyze execution time, memory usage and CPU usage
- remove check if url alread exists
  - hashing cost time
  - checking store cost time (save roundtrips)
  - it doesn't matter, if the same hash (no collision) is written again in memory/redis
  - “Write first, handle conflict”
  - “Optimistic write”