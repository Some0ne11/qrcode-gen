# QR Studio

A full-stack, modern web application for generating and scanning QR codes.

## ✨ Features

*   **Generate 5 Formats**: Easily generate QR codes for Links (URLs), Plain Text, Contacts (vCards), Phone Numbers, and WiFi Networks.
*   **Scan & Decode**: Upload any QR code image (PNG/JPEG) to instantly decode and copy its contents.
*   **Dynamic Theme**: Flawless Light and Dark mode toggle that remembers your preference.
*   **Strict Validation**: Robust frontend and backend validation ensures your QR codes are always generated correctly.

## 🏗️ Architecture

This project is separated into two distinct services to allow for better maintainability and scalability:

*   🌐 **`frontend/`**: The client-side interface built with ![Vue.js](https://img.shields.io/badge/Vue.js-35495E?style=flat&logo=vuedotjs&logoColor=4FC08D) ![Vite](https://img.shields.io/badge/Vite-B73BFE?style=flat&logo=vite&logoColor=FFD62E) ![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-38B2AC?style=flat&logo=tailwind-css&logoColor=white)
*   ⚙️ **`backend/`**: The REST API built with ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white) utilizing the `go-chi` router.

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
    │   │   ├── forms/        # Dynamic form components (URL, Text, WiFi, etc.)
    │   │   ├── ui/           # Reusable UI elements (Modals, etc.)
    │   │   ├── GenerateTab.vue # QR generation UI & logic
    │   │   └── ScanTab.vue     # QR scanner UI & logic
    │   ├── App.vue           # Main layout and tab controller
    │   └── main.js           # Vue application entry point
    ├── .env                  # Environment variables (e.g., API Base URL)
    ├── package.json          # Node dependencies
    └── vite.config.js        # Vite bundler configuration
```

---

## 🚀 Local Development

To run this application locally on your machine, you need to start both the Go backend and the Vue frontend development servers.

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
    *   Specify the local URL of your Go backend so the frontend knows where to send API requests:
        ```env
        VITE_API_BASE_URL=http://localhost:8080
        ```
4.  Run the development server:
    ```bash
    pnpm dev
    ```
    *The application will be accessible at `http://localhost:5173` (or whichever port Vite assigns).*

---

## ☁️ Cloud Deployment

For production, it is highly recommended to decouple the frontend and backend using standard cloud hosting providers.

### 1. Deploying the Backend (Go)
The Go backend can be deployed as a standard web service on platforms like **Render**, **Railway**, or **Heroku**.

1.  Connect your Git repository to your chosen platform.
2.  Set the root directory of the service to `backend/`.
3.  The platform should automatically detect it as a Go application and build it.
4.  Once deployed, copy the live URL of your backend service (e.g., `https://qr-api.onrender.com`).
5.  **Security Tip**: Update the `cors` settings in `backend/main.go` to explicitly allow traffic from your future frontend domain, rather than `*`.

### 2. Deploying the Frontend (Vue)
The Vue frontend is a statically generated site that is perfect for platforms like **Vercel**, **Netlify**, or **Cloudflare Pages**.

1.  Connect your Git repository to your chosen platform.
2.  Set the root directory of the site to `frontend/`.
3.  Set the Build Command to `pnpm build` (or `npm run build`) and the Output Directory to `dist`.
4.  Add an Environment Variable on the hosting platform:
    *   **Key**: `VITE_API_BASE_URL`
    *   **Value**: The live URL of your deployed Go backend (e.g., `https://qr-api.onrender.com`).
5.  Deploy the site! The frontend will now securely communicate with your live backend.

---

## 📡 API Endpoints

The Go backend exposes two primary endpoints:

### `POST /api/generate`
Generates a QR code PNG image from the provided text.
*   **Request Body**: `{"type": "url", "text": "https://example.com"}` (Supported types: `url`, `text`, `contact`, `phone`, `wifi`)
*   **Response**: `image/png` binary data.

### `POST /api/scan`
Reads a QR code from an uploaded image.
*   **Request Body**: `multipart/form-data` containing an `image` file.
*   **Response**: `{"result": "decoded text here"}`
