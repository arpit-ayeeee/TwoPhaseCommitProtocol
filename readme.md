# 2-Phase Commit Protocol for Food Ordering System

This project is an implementation of a 2-phase commit (2PC) protocol for managing a food ordering process. It involves three main services that interact to coordinate food reservations, delivery agent assignments, and order confirmations in a reliable and consistent way.

## Project Structure

The project is divided into three main services:

1. **Food Service**: Handles food reservation and booking.
2. **Delivery Agent Service**: Manages the reservation and booking of delivery agents.
3. **Order Service**: Coordinates with the Food and Delivery Agent services to place an order using the 2-phase commit protocol.

## Overview of the 2-Phase Commit Protocol

The 2-Phase Commit Protocol is used to ensure consistency across distributed services in this food ordering system. It consists of two phases:

- **Prepare Phase**: The order service requests both the food and delivery services to reserve resources (food and delivery agent) for the order. If either reservation fails, the process is aborted.

- **Commit Phase**: If both reservations succeed, the order service proceeds to book the food and agent, finalizing the order.

If any service fails during these phases, the transaction is rolled back to maintain consistency.

## Prerequisites

- Go 1.16+
- MySQL Database
- Gin Framework for Go
- AWS SDK (optional for deployment)

Each service runs on separate ports as follows:
- Food Service: `localhost:8081`
- Delivery Agent Service: `localhost:8082`
- Order Service: initiates requests to these services

## Installation and Setup

1. **Clone the repository**:
   ```bash
   git clone <repository_url>
   cd repository_folder
   ```

2. **Database Setup**:
   - Create a MySQL database `zomato2pc`.
   - Adjust database credentials in each service file (`food.go` and `agents.go`) to match your environment.

3. **Install Dependencies**:
   - Use Go modules to handle dependencies:
     ```bash
     go mod tidy
     ```

## Service Endpoints

### Food Service (`food.go`)

- **`POST /food/reserve`**: Reserves a specific food item for an order.
- **`POST /food/book`**: Confirms the booking of the reserved food item for the order.

### Delivery Agent Service (`agents.go`)

- **`POST /delivery/agent/reserve`**: Reserves a delivery agent for the order.
- **`POST /delivery/agent/book`**: Confirms the booking of the reserved agent for the order.

### Order Service (`order.go`)

The order service initiates the 2-phase commit protocol:
1. Calls `POST /food/reserve` and `POST /delivery/agent/reserve`.
2. If both reservations succeed, calls `POST /food/book` and `POST /delivery/agent/book` to finalize the order.

## Running the Services

1. **Start the Food Service**:
   ```bash
   go run food.go
   ```
   - Runs on `localhost:8081`

2. **Start the Delivery Agent Service**:
   ```bash
   go run agents.go
   ```
   - Runs on `localhost:8082`

3. **Start the Order Service**:
   ```bash
   go run order.go
   ```
   - Executes multiple orders in parallel, simulating the ordering process.

## Example Usage

Once all services are running, the order service will initiate multiple food orders. Each order goes through the 2PC protocol to ensure consistency. Example output:

```
order_id = a1b2c3d4
order placed for Burger, delivery agent J

order_id = e5f6g7h8
order placed for Pizza, delivery agent S
```

If an order cannot be completed, it will output an error message indicating the reason for the failure.

## Notes

- **Concurrency**: The Order Service creates multiple orders in parallel using Go routines.
- **Error Handling**: If any service fails during the reservation or booking steps, the order is canceled.
- **Database Configuration**: Ensure the MySQL connection details are correctly configured in `dsn`.

