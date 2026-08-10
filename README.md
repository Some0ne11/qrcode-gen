# QR Studio

A full-stack, modern web application for generating and scanning QR codes.

## Architecture

This project is separated into two distinct services to allow for better maintainability and scalability:

*   **`frontend/`**: The client-side interface built with Vue 3, Vite, and Tailwind CSS v4.
*   **`backend/`**: The REST API built with Go, utilizing the `go-chi` router.

### Project Structure

```text
qrcode-genapp/
├── backend/                  # Go REST API
│   ├── internal/
│   │   ├── api/
│   │   │   ├── handlers/     # Request handlers (generate, scan)
│   │   │   │   ├── generate.go
│   │   │   │   ├── helpers.go
│   │   │   │   └── scan.go
│   │   │   └── routes.go     # API route definitions
│   │   └── models/           # Data structures and JSON schemas
│   │       └── models.go
│   ├── .env                  # Environment variables (e.g., PORT)
│   ├── go.mod                # Go module dependencies
│   └── main.go               # Entry point for the Go server
│
└── frontend/                 # Vue 3 Client Application
    ├── public/               # Static assets
    ├── src/
    │   ├── assets/           # CSS and styles (Tailwind import)
    │   ├── components/       # Reusable Vue components
    │   │   ├── GenerateTab.vue # QR generation UI & logic
    │   │   └── ScanTab.vue     # QR scanner UI & logic
    │   ├── App.vue           # Main layout and tab controller
    │   └── main.js           # Vue application entry point
    ├── .env                  # Environment variables (e.g., API Base URL)
    ├── package.json          # Node dependencies
    └── vite.config.js        # Vite bundler configuration
```

---

## 🚀 Getting Started

To run this application locally, you will need to start both the Go backend and the Vue frontend development servers.

### 1. Backend (Go) Setup

The backend handles the core logic for encoding text into QR images and decoding uploaded QR images back into text.

**Requirements**: Go 1.20+

1.  Navigate to the backend directory:
    ```bash
    cd backend
    ```
2.  Install dependencies:
    ```bash
    go mod download
    ```
3.  Configure environment variables:
    *   Create a `.env` file in the `backend/` directory.
    *   Add the desired port (default is 8080 if not provided):
        ```env
        PORT=8080
        ```
4.  Run the server:
    ```bash
    go run main.go
    ```
    *The server will start at `http://localhost:8080`.*

### 2. Frontend (Vue) Setup

The frontend provides a sleek, dark-mode user interface with tabs for Generating and Scanning.

**Requirements**: Node.js and `pnpm`

1.  Navigate to the frontend directory:
    ```bash
    cd frontend
    ```
2.  Install dependencies:
    ```bash
    pnpm install
    ```
3.  Configure environment variables:
    *   Create a `.env` file in the `frontend/` directory.
    *   Specify the full URL of your Go backend so the frontend knows where to send API requests:
        ```env
        VITE_API_BASE_URL=http://localhost:8080
        ```
4.  Run the development server:
    ```bash
    pnpm dev
    ```
    *The application will be accessible at `http://localhost:5173` (or whichever port Vite assigns).*

---

## 📡 API Endpoints

The Go backend exposes two primary endpoints:

### `POST /api/generate`
Generates a QR code PNG image from the provided text.
*   **Request Body**: `{"text": "https://example.com"}`
*   **Response**: `image/png` binary data.

### `POST /api/scan`
Reads a QR code from an uploaded image.
*   **Request Body**: `multipart/form-data` containing an `image` file.
*   **Response**: `{"result": "decoded text here"}`
