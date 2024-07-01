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
## APIs
API documentation available in [Postman](https://www.postman.com/shivam-g10/workspace/kodingkorp-public-apis/collection/8417084-3f699e65-579a-46b6-8472-e6875bb47d03?action=share&creator=8417084)

## Project Structure

The directory structure of the project is as follows:
```sh
.
├── src 
│   ├── app
│   │   ├── handlers # server APIs
│   │   │   ├── auth.go
│   │   │   └── kyc.go
│   │   ├── kyc_activities # temporal activities
│   │   │   └── send_kyc_notification.go
│   │   ├── kyc_workflows # temporal workflows
│   │   │   ├── kyc_workflow.go
│   │   │   └── request_kyc.go
│   │   ├── models # data models
│   │   │   ├── kyc_request.go
│   │   │   └── user.go
│   │   └── shared.go # shared constants
│   ├── server # server main
│   │   └── main.go
│   └── worker # worker main
│       └── main.go
├── LICENSE
├── README.md
├── go.mod
└── go.sum
```
## Flowchart
![alt text](./docs/KYC%20PoC.png "Title")
## Additional Information

For more information on how Temporal works and its capabilities, please refer to the [Temporal Documentation](https://docs.temporal.io/docs/go-overview).

If you encounter any issues or have questions, feel free to reach out to the project maintainers.

---

By following these steps, you should have a working instance of the Temporal KYC System PoC running locally. Enjoy exploring Temporal's capabilities!

---

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
