# Test and Confirm Leboncoin 503 Issues

Run the test using 3 terminals:

- Start the API server:
```bash
make run-api
```

- Start Envoy:
```bash
start-envoy
```

- Run the tests (using oha with `--burst-delay 10s`):
```bash
make test
```