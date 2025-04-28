# Nave a la Deriva REST Service

This repository contains a REST API implemented in Go for the AltScore Makers Challenge "Nave a la Deriva". The challenge simulates a stranded spaceship in the year 2315, requiring an API to send distress signals to a repair robot via HTTP requests. The API implements three endpoints (`GET /status`, `GET /repair-bay`, `POST /teapot`) with specific JSON, HTML, and HTTP status code requirements.

## Project Overview

The API fulfills the following requirements:
- **GET /status**: Returns a JSON object with a `"damaged_system"` field (e.g., `{"damaged_system": "engines"}`), representing the spaceship's damaged system.
- **GET /repair-bay**: Returns an HTML page with a `<div class="anchor-point">` containing a code (e.g., `ENG-04` for `engines`), based on a predefined mapping.
- **POST /teapot**: Returns HTTP status code 418 ("I'm a teapot").

## Requirements

- **Go**: Version 1.22.2 or higher.
- **Operating System**: Any OS supported by Go (Linux, macOS, Windows).

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Asac2142/go-nave-deriva.git
   cd go-nave-deriva
   ```

## Running Locally

1. **Start the server**:
   Run the API on the default port (`:3030`):
   ```bash
   go run cmd/api/main.go 
   ```

2. **Test the endpoints**:
   Use cURL or Postman to verify the endpoints:
   ```bash
   curl http://localhost:3030/status
   # Expected: {"damaged_system":"engines"}
   curl http://localhost:3030/repair-bay
   # Expected: HTML with <div class="anchor-point">ENG-04</div>
   curl -X POST http://localhost:3030/teapot
   # Expected: HTTP 418 (I'm a teapot)
   ```

3. **Stop the server**:
   Press `Ctrl+C` to terminate the server.
