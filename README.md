# Temporal KYC System PoC

This is a Proof of Concept (PoC) to evaluate the potential of Temporal for systems requiring multiple types of user interactions.

## Development Setup

To set up the local Temporal development environment, follow the instructions at the [Temporal Getting Started Guide](https://learn.temporal.io/getting_started/go/dev_environment/).

No additional setup is required beyond this.

## Running the Application

Open two terminal windows and execute the following commands:

1. **Start the Worker:**
    ```sh
    go run src/worker/main.go
    ```

2. **Start the Server:**
    ```sh
    go run src/server/main.go
    ```

## Additional Information

For more information on how Temporal works and its capabilities, please refer to the [Temporal Documentation](https://docs.temporal.io/docs/go-overview).

If you encounter any issues or have questions, feel free to reach out to the project maintainers.

---

By following these steps, you should have a working instance of the Temporal KYC System PoC running locally. Enjoy exploring Temporal's capabilities!

---

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
