# ToDos
- analyze execution time, memory usage and CPU usage
- add normalization layer to handle following edge cases:
    - "https://example.com"
    - "https://EXAMPLE.COM"
    - "https://example.com/"
    - "  https://example.com"
    -  "https://example.com  "
    -  "https://example.com\n"
- add two-layer cache (memory, json file)